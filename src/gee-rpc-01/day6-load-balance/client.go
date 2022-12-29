package geerpc6

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"geerpc6/codec"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Call struct {
	Seq           uint64
	ServiceMethod string
	Args          interface{}
	Reply         interface{}
	Error         error
	// 当调用结束时， 会调用call.done 通知调用方
	Done chan *Call
}

func (call *Call) done() {
	call.Done <- call
}

type Client struct {
	cc  codec.Codec
	opt *Option
	//sending 是一个互斥锁，和服务端类似
	sending sync.Mutex
	header  codec.Header
	mu      sync.Mutex
	// 给发送的请求编号
	seq uint64
	// 存储未处理完的请求
	pending map[uint64]*Call
	//
	closing  bool
	shutdown bool
}

type clientResult struct {
	client *Client
	err    error
}

type newClientFunc func(conn net.Conn, opt *Option) (client *Client, err error)

var _ io.Closer = (*Client)(nil)

var ErrShutdown = errors.New("connection is shut down")

func (client *Client) Close() error {
	client.mu.Lock()
	defer client.mu.Unlock()
	if client.closing {
		return ErrShutdown
	}
	client.closing = true
	return client.cc.Close()
}

func (client *Client) IsAvailable() bool {
	client.mu.Lock()
	defer client.mu.Unlock()
	return !client.shutdown && !client.closing
}

// 将参数call添加到client.pending 中，并更新client.seq
func (client *Client) registerCall(call *Call) (uint64, error) {
	client.mu.Lock()
	defer client.mu.Unlock()

	if client.closing || client.shutdown {
		return 0, ErrShutdown
	}

	call.Seq = client.seq
	client.pending[call.Seq] = call
	client.seq++
	return call.Seq, nil
}

// 从 client.pending 中移除对应的 call，并返回。
func (client *Client) removeCall(seq uint64) *Call {
	client.mu.Lock()
	defer client.mu.Unlock()
	call := client.pending[seq]
	delete(client.pending, seq)
	return call
}

// 将 shutdown 设置为 true，且将错误信息通知所有 pending 状态的 call。
func (client *Client) terminalCalls(err error) {
	client.sending.Lock()
	defer client.sending.Unlock()

	client.mu.Lock()
	defer client.mu.Unlock()

	client.shutdown = true
	for _, call := range client.pending {
		call.Error = err
		call.done()
	}
}

func (client *Client) receive() {
	var err error
	for err == nil {
		var h codec.Header
		if err = client.cc.ReadHeader(&h); err != nil {
			break
		}

		call := client.removeCall(h.Seq)
		switch {
		case call == nil:
			//
			err = client.cc.ReaderBody(nil)
		case h.Error != "":
			call.Error = fmt.Errorf(h.Error)
			err = client.cc.ReaderBody(nil)
			call.done()
		default:
			err = client.cc.ReaderBody(call.Reply)
			if err != nil {
				call.Error = errors.New("reading body " + err.Error())
			}
			call.done()
		}
	}
	client.terminalCalls(err)
}

func NewClient(conn net.Conn, opt *Option) (*Client, error) {
	f := codec.NewCodecFuncMap[opt.CodecType]
	if f == nil {
		err := fmt.Errorf("invalid coec type %s ", opt.CodecType)
		log.Println("rpc client: codec error:", err)
		return nil, err
	}

	if err := json.NewEncoder(conn).Encode(opt); err != nil {
		log.Println("rpc client: options error:", err)
		_ = conn.Close()
		return nil, err
	}
	return newClientCodec(f(conn), opt), nil
}

func newClientCodec(cc codec.Codec, opt *Option) *Client {
	client := &Client{
		seq:     1,
		cc:      cc,
		opt:     opt,
		pending: make(map[uint64]*Call),
	}
	go client.receive()
	return client
}

func parseOptions(opts ...*Option) (*Option, error) {
	if len(opts) == 0 || opts[0] == nil {
		return DefaultOption, nil
	}

	if len(opts) != 1 {
		return nil, errors.New("number of options is more than 1")
	}
	opt := opts[0]
	opt.MagicNumber = DefaultOption.MagicNumber
	if opt.CodecType == "" {
		opt.CodecType = DefaultOption.CodecType
	}
	return opt, nil
}

// Dial 便于用户传入服务端地址，创建 Client 实例。为了简化用户调用，通过 ...*Option 将 Option 实现为可选参数。
func Dial(network, address string, opts ...*Option) (client *Client, err error) {
	return dialTimeOut(NewClient, network, address, opts...)
}

func dialTimeOut(f newClientFunc, network, address string, opts ...*Option) (client *Client, err error) {
	opt, err := parseOptions(opts...)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialTimeout(network, address, opt.ConnectionTimeout)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			_ = conn.Close()
		}
	}()

	// 使用子协程执行 NewClient
	ch := make(chan clientResult)
	go func() {
		client, err := f(conn, opt)
		ch <- clientResult{client: client, err: err}
	}()

	if opt.ConnectionTimeout == 0 {
		result := <-ch
		return result.client, result.err
	}
	select {
	//如果 time.After() 信道先接收到消息，则说明 NewClient 执行超时，返回错误。
	case <-time.After(opt.ConnectionTimeout):
		return nil, fmt.Errorf("rpc client: connect timeout: expect winthin %s", opt.ConnectionTimeout)
	case result := <-ch:
		return result.client, result.err
	}

}

func (client *Client) send(call *Call) {
	client.sending.Lock()
	defer client.sending.Unlock()

	seq, err := client.registerCall(call)
	if err != nil {
		call.Error = err
		call.done()
		return
	}

	client.header.ServiceMethod = call.ServiceMethod
	client.header.Seq = seq
	client.header.Error = ""

	if err := client.cc.Write(&client.header, call.Args); err != nil {
		call := client.removeCall(seq)
		if call != nil {
			call.Error = err
			call.done()
		}
	}
}

// Go 和 Call 是客户端暴露给用户的两个 RPC 服务调用接口，Go 是一个异步接口，返回 call 实例。
func (client *Client) Go(serviceMethod string, args, reply interface{}, done chan *Call) *Call {
	if done == nil {
		done = make(chan *Call, 10)
	} else if cap(done) == 0 {
		log.Panic("rpc client: done channel is unbuffered")
	}

	call := &Call{
		ServiceMethod: serviceMethod,
		Args:          args,
		Reply:         reply,
		Done:          done,
	}
	client.send(call)
	return call
}

// Call 是对 Go 的封装，阻塞 call.Done，等待响应返回，是一个同步接口。
func (client *Client) Call(ctx context.Context, serviceMethod string, args, reply interface{}) error {
	call := client.Go(serviceMethod, args, reply, make(chan *Call, 1))
	select {
	case <-ctx.Done():
		client.removeCall(call.Seq)
		return errors.New("rpc client: call failed: " + ctx.Err().Error())
	case call := <-call.Done:
		return call.Error
	}
}

// NewHTTPClient new a Client instance via HTTP as transport protocol
func NewHTTPClient(conn net.Conn, opt *Option) (*Client, error) {
	_, _ = io.WriteString(conn, fmt.Sprintf("CONNECT %s HTTP/1.0\n\n", defaultRPCPath))

	// Require successful HTTP response
	// before switching to RPC protocol.
	resp, err := http.ReadResponse(bufio.NewReader(conn), &http.Request{Method: "CONNECT"})

	if err == nil && resp.Status == connected {
		return NewClient(conn, opt)
	}
	if err == nil {
		err = errors.New("unexpected HTTP response:" + resp.Status)
	}
	return nil, err
}

func DialHTTP(network, address string, opts ...*Option) (*Client, error) {
	return dialTimeOut(NewHTTPClient, network, address, opts...)
}

// XDial calls different functions to connect to a rpc server
func XDial(rpcAddr string, opts ...*Option) (*Client, error) {
	parts := strings.Split(rpcAddr, "@")
	if len(parts) != 2 {
		return nil, fmt.Errorf("rpc client err: wrong format '%s',expect protocol@addr", rpcAddr)
	}

	protocol, addr := parts[0], parts[1]
	switch protocol {
	case "http":
		return DialHTTP("tcp", addr, opts...)
	default:
		return Dial(protocol, addr, opts...)
	}

}

package study

import "errors"

// 要求 *study 去实现 Study，若 Study 接口被更改或未全部实现时，在编译时就会报错。
var _ Study = (*study)(nil)

type Study interface {
	Listen(msg string) string
	Speak(msg string) string
	Read(msg string) string
	Write(msg string) string
}

// 之所以定义为私有的结构体，是因为不想在其他地方被使用，比如后面将 Name 改成 UserName 只需要在本包内修改即可。
type study struct {
	Name string
}

func (s *study) Listen(msg string) string {
	return s.Name + "听" + msg
}

func (s *study) Speak(msg string) string {
	return s.Name + "说" + msg
}

func (s *study) Read(msg string) string {
	return s.Name + "读" + msg
}

func (s *study) Write(msg string) string {
	return s.Name + "写" + msg
}

func New(name string) (Study, error) {
	if name == "" {
		return nil, errors.New("name required")
	}
	return &study{
		Name: name,
	}, nil
}

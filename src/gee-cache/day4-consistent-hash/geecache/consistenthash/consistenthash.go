package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32

type Map struct {
	hash Hash
	//虚拟节点倍数 replicas
	replicas int
	keys     []int
	// 虚拟节点与真实节点的映射表 hashMap
	hashMap map[int]string
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		hash:     fn,
		replicas: replicas,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// Add 实现添加真实节点/机器的 Add() 方法。
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}

// Get 计算 key 的哈希值
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}
	//计算 key 的哈希值
	hash := int(m.hash([]byte(key)))

	// 二分查找从排好序的里面找到第一个
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	//通过 hashMap 映射得到真实的节点。
	return m.hashMap[m.keys[idx%len(m.keys)]]
}

package lru

import "container/list"

type Cache struct {
	// 允许使用的最大内存
	maxBytes int64
	// 当前已使用的内存
	nbytes int64
	// 双向链表
	ll *list.List
	//键是字符串，值是双向链表中对应节点的指针。
	cache map[string]*list.Element
	//是某条记录被移除时的回调函数
	OnEvicted func(key string, value Value)
}

// 键值对 entry 是双向链表节点的数据类型，
type entry struct {
	key   string
	value Value
}

type Value interface {
	len() int
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     map[string]*list.Element{},
		OnEvicted: onEvicted,
	}
}

func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele) //链表中的节点 ele 移动到队尾
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.len()) - int64(kv.value.len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

func (c *Cache) RemoveOldest() {
	// 取到队首节点，从链表中删除
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		//从字典中 c.cache 删除该节点的映射关系
		delete(c.cache, kv.key)
		// 更新内存
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.len())
		//执行回调函数
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

func (c *Cache) len() int {
	return c.ll.Len()
}

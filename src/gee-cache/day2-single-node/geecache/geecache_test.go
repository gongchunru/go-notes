package geecache

import (
	"fmt"
	"log"
	"reflect"
	"testing"
	"time"
)

var set = make(map[int]bool, 0)

func PrintOne(num int) {
	if _, exist := set[num]; !exist {
		fmt.Println(num)
	}
	set[num] = true
}

func TestForLoop(t *testing.T) {
	for i := 0; i < 100; i++ {
		PrintOne(100)
	}
	time.Sleep(time.Second)
}

var db = map[string]string{
	"Tom":  "123",
	"Jack": "234",
	"Sam":  "756",
}

func TestGetter(t *testing.T) {
	var f Getter = GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})

	expect := []byte("key")
	if v, _ := f.Get("key"); !reflect.DeepEqual(v, expect) {
		t.Fatal("callback failed")
	}
}

func TestGet(t *testing.T) {
	loadCounts := make(map[string]int, len(db))
	gee := NewGroup("scores", 2<<10, GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				if _, ok := loadCounts[key]; !ok {
					loadCounts[key] = 0
				}
				loadCounts[key]++
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	//for k, v := range db {
	//	if view, err := gee.Get(k); err != nil || view.String() != v {
	//		t.Fatal("failed to get value of Tom")
	//	}
	//	if _, err := gee.Get(k); err != nil || loadCounts[k] > 1 {
	//		t.Fatalf("cache %s miss", k)
	//	}
	//}
	//
	//if view, err := gee.Get("unknown"); err == nil {
	//	t.Fatalf("the value of unknow should be empty , but %s got", view)
	//}

	for k, v := range db {
		if view, err := gee.Get(k); err != nil || view.String() != v {
			t.Fatal("failed to get value of Tom")
		}
		if _, err := gee.Get(k); err != nil || loadCounts[k] > 1 {
			t.Fatalf("cache %s miss", k)
		}
	}

	if view, err := gee.Get("unknown"); err == nil {
		t.Fatalf("the value of unknow should be empty, but %s got", view)
	}
}

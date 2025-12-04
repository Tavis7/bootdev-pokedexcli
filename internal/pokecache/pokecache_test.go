package pokecache

import (
	"testing"
	"time"
	"slices"
	"sync"
)

type testCaseItem struct {
	key string
	val []byte

	get bool
	ok bool

	delay time.Duration
}

type testCase struct {
	cache_interval time.Duration
	items []testCaseItem
}

func milliseconds(s float32) time.Duration {
	return time.Duration(float32(time.Millisecond) * s)
}

func TestCache(t *testing.T) {
	cases := []testCase{
		{
			cache_interval: milliseconds(1),
			items: []testCaseItem{
				testCaseItem{
					key: "first",
					val: []byte("hello"),
					get: false,
					ok: false,
					delay: milliseconds(0),
				},
				testCaseItem{
					key: "first",
					val: []byte("hello"),
					get: true,
					ok: true,
					delay: milliseconds(0),
				},
				testCaseItem{
					key: "first",
					val: []byte(""),
					get: true,
					ok: false,
					delay: milliseconds(3),
				},
			},
		},
		{
			cache_interval: milliseconds(1),
			items: []testCaseItem{
				testCaseItem{
					key: "first",
					val: []byte("hello"),
					get: false,
					ok: false,
					delay: milliseconds(0),
				},
				testCaseItem{
					key: "first",
					val: []byte("hello"),
					get: true,
					ok: true,
					delay: milliseconds(0),
				},
				testCaseItem{
					key: "first",
					val: []byte("world"),
					get: false,
					ok: false,
					delay: milliseconds(0),
				},
				testCaseItem{
					key: "first",
					val: []byte("world"),
					get: true,
					ok: true,
					delay: milliseconds(0),
				},
				testCaseItem{
					key: "first",
					val: []byte(""),
					get: true,
					ok: false,
					delay: milliseconds(3),
				},
			},
		},
	}

	mu := sync.Mutex{}
	ch := make(chan any)
	for i, c := range cases {
		go testItems(i, c, mu, t, ch)
	}

	for _ = range cases {
		<- ch
	}
}

func testItems(set int, c testCase, mu sync.Mutex, t *testing.T, ch chan any) {
	cache := NewCache(c.cache_interval)
	for i, item := range c.items {
		time.Sleep(item.delay)
		if item.get {
			val, ok := cache.Get(item.key)

			if (!slices.Equal(val, item.val) || ok != item.ok) {
				mu.Lock()
				t.Errorf("%v:%v: Expected '%v = %v (present:%v)' but got '%v' (present:%v)",
				set, i,
				item.key,
				item.val, item.ok, val, ok)
				mu.Unlock()
			}
		} else {
			cache.Add(item.key, item.val)
		}
	}
	ch <- struct{}{}
}

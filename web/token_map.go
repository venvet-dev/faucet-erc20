package main

import (
	"sync"
)

type tokensMap struct {
	m  map[string]bool
	rw *sync.RWMutex
}

func (t *tokensMap) exists(v string) bool {
	t.rw.RLock()
	_, ex := t.m[v]
	t.rw.RUnlock()
	return ex
}

func (t *tokensMap) add(v string) {
	t.rw.Lock()
	t.m[v] = false
	t.rw.Unlock()
}

func (t *tokensMap) retrievePending() []string {
	t.rw.RLock()

	entries := make([]string, 0, len(t.m))
	for addr, ok := range t.m {
		if ok {
			continue
		}
		entries = append(entries, addr)
	}

	t.rw.RUnlock()
	return entries
}

func (t *tokensMap) markAsDone(k string) {
	t.rw.Lock()
	t.m[k] = true
	t.rw.Unlock()
}

package main

import (
	"sync"
)

type tokensMap struct {
	m  map[string]bool
	rw *sync.RWMutex
}

func (t *tokensMap) retrievePending() []string {
	t.rw.RLock()

	entries := make([]string, 0, len(t.m))
	for addr := range t.m {
		entries = append(entries, addr)
	}

	t.rw.RUnlock()
	return entries
}

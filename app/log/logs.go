package log

import (
	"github.com/econron/dbgo/app/files"
	"sync"
)

var lastSavedLSN int = -1

type SafeCounter struct {
	mu sync.Mutex
}

func Flush(lsn int) {
	if lsn >= lastSavedLSN {
		flush()
	}
}

func flush(directory string) {
	files.FileMgr(directory)
}

func iterator() {

}

func append(c SafeCounter) {
	c.mu.Lock()

	c.mu.Unlock()
}

func appendNewBlock() {

}

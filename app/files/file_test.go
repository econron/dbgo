package files

import (
	"sync"
	"testing"
)

func TestActionReadDoesMove(t *testing.T) {
	blkId := BlockId{blknum: 1, filename: "test.txt"}
	sc := safeCounter{mu: sync.Mutex{}}
	Read(10, sc, blkId)
}

package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type safeCounter struct {
	mu sync.Mutex
}

type BlockId struct {
	blknum   int
	filename string
}

func FileMgr(directory string) {
	if !exists(directory) {
		os.MkdirAll(directory, os.ModePerm)
	}
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		os.Remove(file.Name())
	}
}

func exists(directory string) bool {
	if _, err := os.Stat(directory); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func Read(bksize int, c safeCounter, blk BlockId) {
	c.mu.Lock()
	// ここにファイル読み出し処理を書く
	file, error := os.Open(blk.filename)
	// bytebufferに読み出したデータを格納する
	if error != nil {
		panic(error)
	}
	var bl int = bksize * blk.blknum
	seek, err := file.Seek(int64(bl), 0)
	if err != nil {
		return
	}
	fmt.Println(seek)
	// ファイル内のblocksize *
	defer file.Close()
	defer c.mu.Unlock()
}

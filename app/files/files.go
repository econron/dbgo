package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

type safeCounter struct {
	mu sync.Mutex
}

type BlockId struct {
	blknum   int
	filename string
}

type Page struct {
	buffer []byte
}

func SetStringByte(target string, blkId BlockId) []byte {
	var logicalBlock []byte
	//・確保した領域に文字列（abcdefghijklm）を設定する
	stringBytes := strings.NewReader(target)
	var len int = int(stringBytes.Size())
	for i := 0; i < len; i++ {
		byte, error := stringBytes.ReadByte()
		if error != nil {
			panic(error)
		}
		logicalBlock = append(logicalBlock, byte)
	}
	return logicalBlock
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

func Read(bksize int, c safeCounter, blk BlockId, p Page) {
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
	// seekメソッドで変客された値をoffsetとしてファイルに書き込む
	num, error := file.ReadAt(p.buffer, seek)
	if error != nil {
		panic(error)
	}
	fmt.Println(num)
	defer file.Close()
	defer c.mu.Unlock()
}

func Write(bksize int, c safeCounter, blk BlockId, p Page) {
	c.mu.Lock()
	// ここにファイル読み出し処理を書く
	file, error := os.Create(blk.filename)
	// bytebufferに読み出したデータを格納する
	if error != nil {
		panic(error)
	}
	var bl int = bksize * blk.blknum
	seek, err := file.Seek(int64(bl), 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(seek)
	// seekメソッドで変客された値をoffsetとしてファイルに書き込む
	num, error := file.WriteAt(p.buffer, seek)
	if error != nil {
		panic(error)
	}
	fmt.Println(num)
	defer file.Close()
	defer c.mu.Unlock()
}

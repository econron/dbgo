package files

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestWriteContentToFile(t *testing.T) {
	// ブロックサイズ（400)、バッファーサイズ(8)、ディレクトリ(filetest)を指定する
	//var blocksize int = 400
	//var buffersize int = 8
	//var filename string = 'test.txt'
	f, err := os.Create("test.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	//// ファイル名（testfile）をつけてblknumだけブロックを指定する -3
	//type BlockId struc
	//	//	panic(err)
	//	//}t {
	//	filename string
	//	blknum   int
	//}
	//var blockId BlockId = BlockId{filename: filename, blknum: 2}
	//・ヒープメモリ上にブロックサイズ分の領域を確保する
	// goだとヒープメモリ上に意図的に乗せるのはコンパイラーの判断になる
	// ブロックサイズ分だけ確保するとは・・・？
	// →　論理的なブロックサイズだけ
	// 論理的なブロックサイズだけ確保する処理はどうする・・・？
	var logicalBlock []byte
	//・確保した領域に文字列（abcdefghijklm）を設定する
	//var target string = "abcdefghijklm"
	var target = strings.NewReader("abcdefghijklm")
	var len int = int(target.Size())
	for i := 0; i < len; i++ {
		byte, error := target.ReadByte()
		if error != nil {
			panic(error)
		}
		logicalBlock = append(logicalBlock, byte)
	}
	//fmt.Println(result)
	fmt.Println(logicalBlock)
	f.Write(logicalBlock)

	//・確保した領域に文字列（abcdefghijklm）を設定する
	//・char型1文字について最大の保存領域を確保する。
	//　・その保存領域数値分 * 文字数　を取得する　ー①
	//　・ブロックサイズに①を追加する　-②
	//・②の末尾箇所に「345」を数値としてセットする
	//・ブロック数とブロックサイズをかけた数分だけ3のtestfile内をseekする
	//・そこにbytebufferを全て詰め込む
}

func TestFileMangerWrite(t *testing.T) {
	blkId := BlockId{blknum: 60, filename: "test.txt"}
	stringBytes := SetStringByte("abcdefgh", blkId)
	page := Page{buffer: stringBytes}
	Write(100, safeCounter{}, blkId, page)
}

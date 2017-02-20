// 56writingfiles
/*
	创建新的新建，并写入内容
*/
package xiaowenbasic

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func WritingFilesMain() {
	fmt.Println("WritingFile Main")

	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("tmp/dat", d1, 0644)
	checkWrite(err)

	f, err := os.Create("tmp/dat1")
	checkWrite(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}

func checkWrite(e error) {
	if e != nil {
		panic(e)
	}
}

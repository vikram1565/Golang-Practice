package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

type ReaderInterface struct {
	Data []byte
}

func (ri ReaderInterface) Read(t []byte) (n int, err error) {
	// if ri.Data == nil {
	// 	log.Print("data is empty")
	// 	return n, err
	// }
	n = copy(t, ri.Data)
	return n, io.EOF
}
func main() {
	// TODO: use new reader to input
	ri := ReaderInterface{Data: []byte("vikram")}
	fmt.Println(ri)
	t := make([]byte, 6)
	n, e := ri.Read(t)
	fmt.Println(n, e, string(t))
	fmt.Println("ioutil reader implementation--------")
	ba, e := ioutil.ReadAll(ri)
	fmt.Println("---", string(ba))
}

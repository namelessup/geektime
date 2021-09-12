package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"
)

/*
	根据PPT里的内容(goim_tcp.jpg)来看，数据包分为六部分
	1. package length       4byte
	2. header length        2byte
	3. protocol version     2byte
	4. operation            4byte
	5. sequence id          4byte
	6. body                 packLen-headerLen
*/

var (
	defaultProtocol  = []byte{0x00, 0x01}
	defaultOperation = []byte{0xff, 0xff}
)

// PackMessage 装包
func PackMessage(body []byte) []byte {
	ans := make([]byte, 6)
	// 设置package length
	binary.BigEndian.PutUint32(ans[:4], uint32(len(body)+12))
	fmt.Println("source package length =", uint32(len(body)+12))
	// 设置header length (header length + protocol length + operation length + sequence id length)
	binary.BigEndian.PutUint16(ans[4:6], 12)
	fmt.Println("source header length =", 12)

	fmt.Println("source protocol length =", defaultProtocol)
	fmt.Println("source operation length =", defaultOperation)

	seq := sequence()
	fmt.Println("source sequence id length =", seq)

	fmt.Println("source body =", body)

	buffer := bytes.NewBuffer(ans)
	buffer.Write(defaultProtocol)
	buffer.Write(defaultOperation)
	buffer.Write(seq)
	buffer.Write(body)

	return buffer.Bytes()
}

func UnpackMessage(data []byte) {
	packLen := binary.BigEndian.Uint32(data[:4])
	headerLen := binary.BigEndian.Uint16(data[4:6])
	bodyLen := packLen - uint32(headerLen)
	protocol := data[6:8]
	operation := data[8:10]
	seq := data[10:14]
	body := data[14 : 14+bodyLen]
	fmt.Println("unpack package length =", packLen)
	fmt.Println("unpack header length=", headerLen)
	fmt.Println("unpack protocol=", protocol)
	fmt.Println("unpack operation=", operation)
	fmt.Println("unpack sequence id=", seq)
	fmt.Println("unpack body=", body)
}

// 注意：非线程安全
var source = rand.NewSource(time.Now().Unix())

func sequence() []byte {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs[0:4], uint32(source.Int63()&(1<<32-1)))
	return bs
}

func main() {
	// source package length = 23
	// source header length = 12
	// source protocol length = [0 1]
	// source operation length = [255 255]
	// source sequence id length = [125 149 80 78]
	// source body = [104 101 108 108 111 32 119 111 114 108 100]
	// -----------------------
	// unpack package length = 23
	// unpack header length= 12
	// unpack protocol= [0 1]
	// unpack operation= [255 255]
	// unpack sequence id= [125 149 80 78]
	// unpack body= [104 101 108 108 111 32 119 111 114 108 100]
	body := PackMessage(bytes.NewBufferString("hello world").Bytes())
	UnpackMessage(body)
}

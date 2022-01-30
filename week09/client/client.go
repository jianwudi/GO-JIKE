package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

func Enpack(message []byte) []byte {
	buf := append(Int32ToBytes(len(message)+16), Int16ToBytes(16)...)
	buf = append(buf, Int16ToBytes(1)...)
	buf = append(buf, Int32ToBytes(2)...)
	buf = append(buf, Int32ToBytes(3)...)
	buf = append(buf, message...)

	return buf
}

//整数转换成字节
func Int32ToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//整数转换成字节
func Int16ToBytes(n int) []byte {
	x := int16(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:7373")
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	defer conn.Close()
	data := "helloworld"
	conn.Write(Enpack([]byte(data)))

}

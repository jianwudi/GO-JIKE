package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

const (
	PackageLengthBytes   = 4
	HeaderLengthBytes    = 2
	ProtocolVersionBytes = 2
	OperationBytes       = 4
	SequenceIDBytes      = 4

	HeaderLength = PackageLengthBytes + HeaderLengthBytes + ProtocolVersionBytes + OperationBytes + SequenceIDBytes

	//Body
)

func Depack(buffer []byte) []byte {
	length := len(buffer)

	if length < HeaderLength {
		return nil
	}
	fmt.Println(buffer)
	packageLength := ByteToInt(buffer[:PackageLengthBytes])
	if length < packageLength {
		return nil
	}

	site := PackageLengthBytes
	headerLength := ByteToInt16(buffer[site : site+HeaderLengthBytes])
	site += HeaderLengthBytes

	protocolVersion := ByteToInt16(buffer[site : site+ProtocolVersionBytes])
	site += ProtocolVersionBytes

	operation := ByteToInt(buffer[site : site+OperationBytes])
	site += OperationBytes

	sequenceID := ByteToInt(buffer[site : site+SequenceIDBytes])
	site += SequenceIDBytes

	fmt.Printf("packageLength: %d, headerLength: %d , protocolVersion: %d, operation: %d, sequenceID: %d \n", packageLength, headerLength, protocolVersion, operation, sequenceID)

	data := buffer[HeaderLength:packageLength]

	return data
}

func ByteToInt(n []byte) int {
	bytesbuffer := bytes.NewBuffer(n)
	var x int32
	binary.Read(bytesbuffer, binary.BigEndian, &x)

	return int(x)
}

func ByteToInt16(n []byte) int {
	bytesbuffer := bytes.NewBuffer(n)
	var x int16
	binary.Read(bytesbuffer, binary.BigEndian, &x)

	return int(x)
}
func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	reader := bufio.NewReader(conn)

	var buf [128]byte
	n, err := reader.Read(buf[:]) // 读取数据
	if err != nil {
		fmt.Println("read from client failed, err:", err)
		return
	}
	data := Depack(buf[:n])
	fmt.Printf("%s\n", data)

}
func main() {
	netListen, err := net.Listen("tcp", "localhost:7373")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	defer netListen.Close()

	for {
		conn, err := netListen.Accept()
		if err != nil {

			continue
		}
		go process(conn)
	}
}

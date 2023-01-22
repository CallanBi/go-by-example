package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		panic(err)
	}
	for {
		client, err := server.Accept()
		if err != nil {
			log.Printf("Accept failed %v", err)
			continue
		}
		go process(client)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	// 带缓冲流
	reader := bufio.NewReader(conn)
	for {
		// 一个一个字节读
		b, err := reader.ReadByte()
		// 底层：读一个字节时会把剩下的都读完，不会低效
		if err != nil {
			break
		}
		// conn.Write 写入一个类型为 byte 的 slice
		_, err = conn.Write([]byte{b})
		if err != nil {
			break
		}
	}
}

// 用 nc 127.0.0.1 1080 建立一个 TCP 链接

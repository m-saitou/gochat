package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Print("Hi, What are you name ? ")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()

	name := stdin.Text()

	fmt.Println("hello", name)

	// TCP サーバを立ち上げる
	listen, _ := net.Listen("tcp", "127.0.0.1:8888")
	fmt.Println("Listen 127.0.0.1:8888")

	// クライアントからの通信を待って、 Connection を張る
	conn, _ := listen.Accept()

	// クライアントから送信されたデータを読み込む
	// n はデータ長（バイト）
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)

	// buf の n バイト目までを文字列化して出力
	fmt.Printf("[Message]\n%s", string(buf[:n]))

	// Connection を閉じる
	conn.Close()
}

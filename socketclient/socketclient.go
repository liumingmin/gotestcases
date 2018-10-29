
package main

import (
	"time"
	"net"
	"fmt"
	"bufio"
)

func main(){
	startclient()

	for{
		time.Sleep(time.Hour)
	}
}

func startclient(){
	var conn net.Conn
	conn, err := net.Dial("tcp", "127.0.0.1:9999")

	if err!=nil{
		fmt.Println(err)
	}

	conn.SetReadDeadline(time.Now().Add(time.Second*20))
	go read(conn)
}

func  read(conn net.Conn) {
	defer func() {
		conn.Close()

		fmt.Println("socket read timeout ")
	}()

	reader := bufio.NewReader(conn)
	scanner := bufio.NewScanner(reader)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		//fmt.Println(data)
		if !atEOF {
			if len(data) == 0{
				fmt.Println("recv [] data")
			}

			if len(data) > 2 {
				//length := int16(0)
				//binary.Read(bytes.NewReader(data[0:2]), binary.LittleEndian, &length)
				//if int(length) <= len(data) {
				//	return int(length), data[:int(length)], nil
				//}else {
				//	return 0,nil,nil
				//}
			}
		}
		return
	})

	for scanner.Scan() {
		fmt.Println(scanner.Bytes())
	}
}

package main

import (
	"os"
	"fmt"
	"net"
	"bufio"
	"encoding/binary"
	"bytes"
	"time"
)

func main() {
	startServer()
}

func startServer(){
	listen, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {

		fmt.Println("Socket listen %d failed,", err)
		os.Exit(1)
	}
	defer listen.Close()
	fmt.Println("Begin listen url: ")

	for {
		//logutil.Info(context.Background(), "listen :")
		conn, err := listen.Accept()
		if err != nil {
			//logutil.Error(context.Background(), "Err :%v", err)
			continue
		}
		go readFromGameChatServer(conn)
		go writeToGameChatServer(conn)
	}
}

func  writeToGameChatServer(conn net.Conn) {
	ticker := time.NewTicker(time.Second*10)
	defer func() {
		//logutil.Info(context.Background(), "Finish writing to game. id: %d, ptr: %p", c.CenterID, c)
		ticker.Stop()
		conn.Close()
	}()
	for {
		select {

		case <-ticker.C:
			if err := conn.SetWriteDeadline(time.Now().Add(time.Second*20)); err != nil {
				//todo: if conn is closed, should trigger notification to app msgr server
				//if c.IsAgentStoped {
				//	ticker.Stop()
				//}
				//logutil.Error(context.Background(), "Set write deadline to client failed. id：%v, error: %v", c.CenterID, err)
				time.Sleep(time.Second)
				//continue
				return //返回,跳出Select循环
			}
			//todo: send ping message
			//logutil.Debug(context.Background(), "Send ping message , to do :")
		}
	}
}
func  readFromGameChatServer(conn net.Conn) {
	defer func() {
		//logutil.Info(context.Background(), "WriteToServer finished [%d]", a.CenterID)
		conn.Close()
		//GameChatHub.UnregisterAgent(a)
	}()

	reader := bufio.NewReader(conn)
	scanner := bufio.NewScanner(reader)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		fmt.Println(data)
		if !atEOF {
			if len(data) > 2 {
				length := int16(0)
				binary.Read(bytes.NewReader(data[0:2]), binary.LittleEndian, &length)
				if int(length) <= len(data) {
					//return int(length), data[:int(length)+2], nil//不需要加2
					return int(length), data[:int(length)], nil
				}else {
					return 0,nil,nil
				}
			}
		}
		return
	})
	for scanner.Scan() {
		fmt.Println(scanner.Bytes())

	}
}
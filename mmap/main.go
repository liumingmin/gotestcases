package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"time"

	"golang.org/x/exp/mmap"
)

func main() {
	mmapCase()
	ioCase()
}

func mmapCase() {
	start := time.Now().UnixNano()

	defer func() {
		end := time.Now().UnixNano()
		fmt.Printf("mmap: %v\n", time.Duration(end-start))
	}()

	// 使用 mmap 包打开文件
	reader, err := mmap.Open("E:\\10040719_com.tencent.tmgp.dnf_h3153744_101.0.2.0_gqbJdf.apk")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer reader.Close()

	// 获取文件大小
	fileSize := reader.Len()

	sr := io.NewSectionReader(reader, 0, int64(fileSize))

	hasher := sha256.New()

	buffer := make([]byte, 512*1024)
	io.CopyBuffer(hasher, sr, buffer)

	hashVal := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	fmt.Println(hashVal)
}

func ioCase() {
	start := time.Now().UnixNano()

	defer func() {
		end := time.Now().UnixNano()
		fmt.Printf("io: %v \n", time.Duration(end-start))
	}()

	f, err := os.Open("E:\\10040719_com.tencent.tmgp.dnf_h3153744_101.0.2.0_gqbJdf.apk")
	if err != nil {
		return
	}

	defer f.Close()
	hasher := sha256.New()

	buffer := make([]byte, 512*1024)
	io.CopyBuffer(hasher, f, buffer)

	hashVal := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	fmt.Println(hashVal)
}

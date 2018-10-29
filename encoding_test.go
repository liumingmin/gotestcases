/*
 * Copyright (c) 2018. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package main

import (
	"testing"
	"golang.org/x/text/transform"
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"context"
)

func TryUTF8ToGBK(ctx context.Context, s string)string{
	str,err := UTF82GBK([]byte(s))
	if err != nil{
		return s
	}

	return string(str)
}

func TryUTF8ToGBK2(ctx context.Context,s string)string{
	return string(ConvertUTF8BytesTOGBKBytes([]byte(s)))
}

func UTF82GBK(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func BenchmarkPkgIn(b *testing.B) {
	for i:=0 ; i  < 100000 ; i++ {
		TryUTF8ToGBK2(nil,"我是测试UTF-8转GBK的性能测试")
	}
}
func BenchmarkXText(b *testing.B) {
	for i:=0 ; i  < 100000 ; i++ {
		TryUTF8ToGBK(nil,"我是测试UTF-8转GBK的性能测试")
	}
}




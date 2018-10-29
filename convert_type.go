package main

import (
	"encoding/binary"
	"strconv"
	//"github.com/kimiazhu/ginweb/util/log"
	//"context"
	"gopkg.in/iconv.v1"
	"fmt"
)

func Uint64ToUint32(in uint64)(uint32){
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, in)
	out := binary.BigEndian.Uint32(b)
	return out
}

func StringToInt64(in string)(int64){
	out,err := strconv.ParseInt(in,10,64)
	if err != nil {
		//logutil.Info(context.Background(),"convert error :%v",err)
	}
	return out
}
func StringToInt32(in string)(int32){
	out,err := strconv.ParseInt(in,10,32)
	if err != nil {
		//logutil.Info(context.Background(),"convert error :%v",err)
	}
	return int32(out)
}
func StringToUInt64(in string)(uint64){
	out,err := strconv.ParseUint(in,10,64)
	if err != nil {
		//logutil.Info(context.Background(),"convert error :%v",err)
	}
	return out
}
func StringToUInt32(in string)(uint32){
	tmpOut,err := strconv.ParseUint(in,10,32)
	if err != nil {
		//logutil.Info(context.Background(),"convert error :%v",err)
	}
	out := uint32(tmpOut)
	return out
}

func Int32ToUInt16(in int32)(uint16){
	inStr := fmt.Sprintf("%v",in)
	tmpOut,err := strconv.ParseInt(inStr,10,16)
	if err != nil {
		//logutil.Info(context.Background(),"convert error :%v",err)
	}
	out := uint16(tmpOut)
	return out
}


func ConvertUTF8BytesFromGBKBytes(gbkBs []byte) []byte {
	code_convertor, err := iconv.Open("utf-8","gb18030" ) // convert from gbk to utf8
	defer code_convertor.Close()
 	var str string
	if err == nil {
		str = code_convertor.ConvString(string(gbkBs))
	}else{
		//logutil.Error(context.Background(), "DecoderByteFromGBK err:%v",err)
	}
	return []byte(str)
}
func ConvertToUTF8StrFromGBKBytes(gbkBs []byte)string {
	result:=""

	if gbkBs!=nil && len(gbkBs)>0{
		code_convertor, err := iconv.Open("utf-8","gb18030" ) // convert from gbk to utf8
		defer code_convertor.Close()
		if err == nil {
			result = code_convertor.ConvString(string(gbkBs))
		}else{
			//logutil.Error(context.Background(), "DecoderByteFromGBK err:%v",err)
		}
	}
	return result
}

func ConvertUTF8BytesTOGBKBytes1(utf8Bs []byte)([]byte) {
	//code_convertor, err := iconv.Open("utf-8", "gb2312") // convert from gbk to utf8
	code_convertor, err := iconv.Open("gb18030","utf-8" ) // convert from utf8 to gbk
	defer code_convertor.Close()
	var str string
	if err == nil {
		str = code_convertor.ConvString(string(utf8Bs))
		return []byte(str)
	}
	return nil
}
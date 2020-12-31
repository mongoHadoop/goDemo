package socker_stick

import (
"bufio"
"bytes"
"encoding/binary"
"fmt"
)

//Encode 将消息编码
func Encode(message string) ([]byte,error) {
	length := int32(len(message))
	var pkg = new(bytes.Buffer)
	//写入消息头
	err := binary.Write(pkg,binary.LittleEndian,length)
	if err != nil {
		fmt.Println("写入消息头失败",err)
		return nil,err
	}
	//写入消息实体
	err = binary.Write(pkg,binary.LittleEndian,[]byte(message))
	if err != nil {
		fmt.Println("写入消息实体失败",err)
		return nil,err
	}
	return pkg.Bytes(),nil
}

//Decode解码消息
func Decode(reader *bufio.Reader) (string,error) {
	//读取信息长度
	lengthByte,_ := reader.Peek(4)
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff,binary.LittleEndian,&length)
	if err != nil {
		return "",err
	}
	//BuffRead 返回缓冲区现有的可读的字节数
	if int32(reader.Buffered()) < length+4 {
		return "",err
	}
	pack := make([]byte,int(4+length))
	_,err = reader.Read(pack)
	if err != nil {
		return "",err
	}
	return string(pack[4:]),nil
}
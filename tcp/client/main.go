package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

func Encode(message string) ([]byte, error) {
	// 读取消息的长度，转换成int32类型（占4个字节）
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

/**
出现粘包问题，“粘包”可发生在发送端也可发生在接收端：
*/
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, How are you!`
		data, err := Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err：", err)
			return
		}
		conn.Write(data)
	}
}

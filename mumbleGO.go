package mumbleGO

import (
	"C"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

func CreateBufC(x uint64) *bytes.Buffer {
	buf := new(bytes.Buffer)
	var data = []interface{}{
		C.int(0),
		C.longlong(x),
	}
	for _, v := range data {
		err := binary.Write(buf, binary.BigEndian, v)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	return buf
}

func CreateBuf(x uint64) *bytes.Buffer {
	buf := new(bytes.Buffer)
	var data = []interface{}{
		uint32(0),
		uint64(x),
	}
	for _, v := range data {
		err := binary.Write(buf, binary.BigEndian, v)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	return buf
}

func Connect2MumbleRecievePingPaket(identity uint64, server string, port string) []byte {

	var recv []byte = make([]byte, 24)

	addr, err := net.ResolveUDPAddr("udp", ConcatServerPort(server, port))
	if err != nil {
		fmt.Println("ResolveUDPAddr")
		fmt.Println(err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("DialUDP")
		fmt.Println(err)
	}

	_, err = conn.Write(CreateBuf(identity).Bytes())
	if err != nil {
		fmt.Println("Write")
		fmt.Println(err)
	}

	_, _, err = conn.ReadFromUDP(recv)
	if err != nil {
		fmt.Println("ReadFromUDP")
		fmt.Println(err)
	}
	return recv
}

func Byte2String(recv []byte) uint64 {
	var t1 T1
	buf := bytes.NewReader(recv)
	err := binary.Read(buf, binary.BigEndian, &t1)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	return t1.F2
}

func ConcatServerPort(server string, port string) string {
	var buffer bytes.Buffer
	_, err := buffer.WriteString(server)
	if err != nil {
		fmt.Println("couldn't write to buffer", err)
	}
	_, err = buffer.WriteString(":")
	if err != nil {
		fmt.Println("couldn't write to buffer", err)
	}
	_, err = buffer.WriteString(port)
	if err != nil {
		fmt.Println("couldn't write to buffer", err)
	}
	return buffer.String()
}

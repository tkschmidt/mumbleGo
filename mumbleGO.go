package mumbleGO

import (
	"C"
	"bytes"
	"encoding/binary"
	"fmt"
)
import (
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

func Connect2Mumble() []byte {

	var recv []byte = make([]byte, 24)

	addr, err := net.ResolveUDPAddr("udp", "95.143.172.148:61030")
	if err != nil {
		fmt.Println("ResolveUDPAddr")
		fmt.Println(err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("DialUDP")
		fmt.Println(err)
	}

	_, err = conn.Write(CreateBuf(154011).Bytes())
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

type T1 struct {
	// //F1 [5]byte
	// F1 [4]byte
	// F2 [8]byte
	// F3 [4]byte
	// F4 [4]byte
	// F5 [4]byte

	F1 uint32
	F2 uint64
	F3 uint32
	F4 uint32
	F5 uint32
}

type T2 struct {
	// //F1 [5]byte
	// F1 [4]byte
	// F2 [8]byte
	// F3 [4]byte
	// F4 [4]byte
	// F5 [4]byte
	F1 uint32
	F2 uint64
}

func Byte2String(recv []byte) {
	//b := CreateBuf(154011).Bytes()
	var t1 T1

	buf := bytes.NewReader(recv)
	err := binary.Read(buf, binary.BigEndian, &t1)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Println(t1)
}

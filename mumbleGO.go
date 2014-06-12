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

func createBuf(x uint64) *bytes.Buffer {
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

func Connect2Mumble() {
	var recv []byte

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

	_, err = conn.Write([]uint8{0x00, 0x02, 0x74, 0x66, 0x74, 0x70, 0x2e,
		0x67, 0x6f, 0x00, 0x6e, 0x65, 0x74, 0x61, 0x73, 0x63, 0x69, 0x69, 0x00})
	if err != nil {
		fmt.Println("Write")
		fmt.Println(err)
	}

	_, _, err = conn.ReadFromUDP(recv)
	if err != nil {
		fmt.Println("ReadFromUDP")
		fmt.Println(err)
	}
}

package mumbleGO

import (
	"C"
	"bytes"
	"encoding/binary"
	"fmt"
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

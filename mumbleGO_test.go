package mumbleGO

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestCreateBufC(t *testing.T) {
	var time uint64 = 239179
	fmt.Println()
	if x := hex.EncodeToString(CreateBufC(time).Bytes()); x != "00000000000000000003a64b" {
		t.Errorf("Bytes are not correct")
	}
}

func TestCreateBuf(t *testing.T) {
	// uint32 = int
	// uint64 = long
	var time uint64 = 239179
	result := hex.EncodeToString(CreateBuf(time).Bytes())
	fmt.Println(result)
	if x := result; x != "00000000000000000003a64b" {
		t.Errorf("Bytes are not correct")
	}
}

// func TestConnect2Mumble(t *testing.T) {

// 	recv := Connect2Mumble()
// }

func TestByte2String(t *testing.T) {
	recv := Connect2Mumble()
	Byte2String(recv)
}

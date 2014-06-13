package mumbleGO

import (
	"encoding/hex"
	"testing"
)

func TestCreateBufC(t *testing.T) {
	var time uint64 = 239179
	if x := hex.EncodeToString(CreateBufC(time).Bytes()); x != "00000000000000000003a64b" {
		t.Errorf("Bytes are not correct")
	}
}

func TestCreateBuf(t *testing.T) {
	// uint32 = int
	// uint64 = long
	var time uint64 = 239179
	result := hex.EncodeToString(CreateBuf(time).Bytes())
	if x := result; x != "00000000000000000003a64b" {
		t.Errorf("Bytes are not correct")
	}
}

func TestGettingIdentityBack(t *testing.T) {
	var ident uint64 = 154011
	recv := Connect2MumbleRecievePingPaket(ident, "95.143.172.148", "61030")

	if x := Byte2String(recv); x != 154011 {
		t.Errorf("ident doesn't come back")
	}
}

func TestConcatServerPort(t *testing.T) {
	var server string = "95.143.172.148"
	var port string = "61030"
	if x := ConcatServerPort(server, port); x != "95.143.172.148:61030" {
		t.Errorf("concatination doesn't work")
	}
}

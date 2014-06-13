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
	recv := Connect2MumbleRecievePingPaket(ident)

	if x := Byte2String(recv); x != 154011 {
		t.Errorf("ident doesn't come back")
	}
}

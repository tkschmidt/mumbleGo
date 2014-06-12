package mumbleGO

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestBufCreation(t *testing.T) {
	var time uint64 = 239179
	fmt.Println()
	if x := hex.EncodeToString(createBuf(time).Bytes()); x != "00000000000000000003a64b" {
		t.Errorf("Bytes are not correct")
	}
}

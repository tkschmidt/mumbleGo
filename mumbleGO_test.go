package mumbleGO

import "testing"
import "fmt"
import "bytes"

func TestBufCreation(t *testing.T) {
	value := bytes.NewBufferString("00000000000000000003a6K")
	//var value = "\x00\x00\x00\x00\x00\x00\x00\x00\x00\x03\xa6K"
	var time uint64 = 239179
	if x := createBuf(time); x != value {
		t.Errorf("Bytes are not correct")

	}
	fmt.Printf("% x", value)
	fmt.Println("")
	fmt.Printf("% x", createBuf(time))
	fmt.Println("")

}

// func TestSqrt(t *testing.T) {
// 	const in, out = 4, 2
// 	if x := Sqrt(in); x != out {
// 		t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
// 	}
// }

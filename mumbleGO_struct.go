package mumbleGO

type version struct {
	A1 [1]byte
	A2 [1]byte
	A3 [1]byte
	A4 [1]byte
}
type T1 struct {
	F1 version
	F2 uint64
	F3 uint32
	F4 uint32
	F5 uint32
}

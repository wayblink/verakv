package kv

type Modify struct {
	Data interface{}
}

type PUT struct {
	Key   []byte
	Value []byte
}

type DEL struct {
	Key []byte
}

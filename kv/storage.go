package kv

type Storage interface {
	Start() error
	Stop() error
	Write(batch []Modify) error
	Reader() (StorageReader, error)
}

type StorageReader interface {
	Get(key []byte) ([]byte, error)
}

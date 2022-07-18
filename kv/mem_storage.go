package kv

import "sync"

var _ Storage = (*MemStorage)(nil)

// MemStorage is in-memory kv storage, will vanish after restart, just for test
type MemStorage struct {
	Map sync.Map
}

func NewMemStorage() *MemStorage {
	return &MemStorage{Map: sync.Map{}}
}

func (m *MemStorage) Start() error {
	return nil
}

func (m *MemStorage) Stop() error {
	return nil
}

func (m *MemStorage) Write(batch []Modify) error {
	for _, modify := range batch {
		switch data := modify.Data.(type) {
		case PUT:
			m.Map.Store(string(data.Key), data.Value)
		case DEL:
			m.Map.Delete(string(data.Key))
		}
	}
	return nil
}

func (m *MemStorage) Reader() (StorageReader, error) {
	return &memReader{m, 0}, nil
}

// memReader is a StorageReader which reads from a MemStorage.
type memReader struct {
	inner     *MemStorage
	iterCount int
}

func (m memReader) Get(key []byte) ([]byte, error) {
	value, ok := m.inner.Map.Load(string(key))
	if ok {
		return value.([]byte), nil
	} else {
		return nil, nil
	}
}

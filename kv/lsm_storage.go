package kv

import "github.com/wayblink/verakv/kv/lsm"

var _ Storage = (*LsmStorage)(nil)

// LsmStorage is LSM-based storage engine
// 我还没想清楚storage enigne和LSM要怎么组织，先分成两层便于解藕
type LsmStorage struct {
	lsm lsm.LSM
}

func (l LsmStorage) Start() error {
	//TODO implement me
	panic("implement me")
}

func (l LsmStorage) Stop() error {
	//TODO implement me
	panic("implement me")
}

func (l LsmStorage) Write(batch []Modify) error {
	//TODO implement me
	panic("implement me")
}

func (l LsmStorage) Reader() (StorageReader, error) {
	//TODO implement me
	panic("implement me")
}

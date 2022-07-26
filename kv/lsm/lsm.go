package lsm

type LSM struct {
	memTable  *memTable
	immTables []*memTable
	levels    *levelManager
}

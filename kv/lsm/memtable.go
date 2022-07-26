package lsm

import "github.com/wayblink/verakv/kv/util"

type memTable struct {
	skipList util.SkipList
}

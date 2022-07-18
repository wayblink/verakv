package kv

import (
	"fmt"
	"testing"
)

func TestDb(t *testing.T) {
	db := Open()
	defer db.Close()
	db.Set(&Entry{
		Key:   ([]byte)("key1"),
		Value: ([]byte)("value1"),
	})

	entry, _ := db.Get(([]byte)("key1"))
	fmt.Println(string(entry.Value))

	db.Del(([]byte)("key1"))

	entry, _ = db.Get(([]byte)("key1"))
	fmt.Println(string(entry.Value))

}

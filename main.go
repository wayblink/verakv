package main

import (
	"fmt"
	"github.com/wayblink/verakv/kv"
)

func main() {
	db := kv.Open()
	defer db.Close()
	db.Set(&kv.Entry{
		Key:   ([]byte)("key1"),
		Value: ([]byte)("value1"),
	})

	entry, _ := db.Get(([]byte)("key1"))
	fmt.Println(string(entry.Value))

	db.Del(([]byte)("key1"))

	entry, _ = db.Get(([]byte)("key1"))
	fmt.Println(string(entry.Value))
}

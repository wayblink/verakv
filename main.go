package main

import (
	"fmt"
	"github.com/wayblink/verakv/kv"
)

func main() {

	storage := kv.NewMemStorage()
	reader, _ := storage.Reader()

	puts := []kv.Modify{
		{Data: kv.PUT{Key: ([]byte)("key1"), Value: ([]byte)("value1")}},
	}
	storage.Write(puts)

	value, _ := reader.Get(([]byte)("key1"))

	fmt.Println(string(value))
}

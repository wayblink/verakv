package kv

type (
	API interface {
		Set(data *Entry) error
		Get(key []byte) (*Entry, error)
		Del(key []byte) error
		//NewIterator(opt *iterator.Options) iterator.Iterator
		//Info() *Stats
		//Open() error
		Close() error
	}

	DB struct {
		storage Storage
	}
)

func Open() *DB {
	storage := NewMemStorage()
	storage.Start()
	db := &DB{storage: storage}
	return db
}

func (db *DB) Close() {
	db.storage.Stop()
}

func (db *DB) Del(key []byte) error {
	return db.storage.Write([]Modify{
		{Data: DEL{Key: key}},
	})
}

func (db *DB) Set(data *Entry) error {
	return db.storage.Write([]Modify{
		{Data: PUT{Key: data.Key, Value: data.Value}},
	})
}

func (db *DB) Get(key []byte) (*Entry, error) {
	reader, err := db.storage.Reader()
	if err != nil {
		return nil, err
	}
	value, err := reader.Get(([]byte)("key1"))
	if err != nil {
		return nil, err
	}
	return &Entry{Key: key, Value: value}, nil
}

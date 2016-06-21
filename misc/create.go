package misc

import (
	"fmt"

	"github.com/tecbot/gorocksdb"
)

// CreateRocksDb just creates a simple rocksdb store
func CreateRocksDb() error {

	opts := gorocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(true)
	db, err := gorocksdb.OpenDb(opts, "/Users/aayushkumar/Downloads/testDb")

	if err != nil {
		return err
	}
	fmt.Printf("RocksDb store is created with name %s\n", db.Name())
	return nil
}

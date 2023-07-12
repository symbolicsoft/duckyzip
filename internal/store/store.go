/* SPDX-FileCopyrightText: © 2023 Nadim Kobeissi <nadim@symbolic.software>
 * SPDX-License-Identifier: GPL-3.0-only */

package store

import (
	"log"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

var Database = func() *leveldb.DB {
	db, err := leveldb.OpenFile("ducky.db", nil)
	if err != nil {
		log.Fatal("could not open database")
	}
	return db
}()

func Put(key string, value string) error {
	return Database.Put([]byte(key), []byte(value), &opt.WriteOptions{
		Sync: true,
	})
}

func Get(key string) (string, error) {
	value, err := Database.Get([]byte(key), &opt.ReadOptions{})
	return string(value), err
}

func Has(key string) (bool, error) {
	return Database.Has([]byte(key), &opt.ReadOptions{})
}

func CloseDatabase() {
	err := Database.Close()
	if err != nil {
		log.Fatal("could not close database")
	}
}

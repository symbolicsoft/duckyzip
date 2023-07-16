/* SPDX-FileCopyrightText: © 2023 Nadim Kobeissi <nadim@symbolic.software>
 * SPDX-License-Identifier: GPL-3.0-only */

package store

import (
	"encoding/json"
	"errors"
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

type DBEntry struct {
	Payload   string
	VRFValue0 string
	VRFProof0 string
	VRFValue1 string
	VRFProof1 string
}

func PutDBEntry(shortID string, dbEntry DBEntry) error {
	has, err := Has(shortID)
	if has {
		return errors.New("refusing to overwrite entry")
	}
	if err != nil {
		return err
	}
	dbEntryString, err := json.Marshal(&dbEntry)
	if err != nil {
		return err
	}
	return Database.Put([]byte(shortID), []byte(dbEntryString), &opt.WriteOptions{
		Sync: true,
	})
}

func GetDBEntry(shortID string) (DBEntry, error) {
	dbEntry := DBEntry{}
	dbEntryBytes, err := Database.Get([]byte(shortID), &opt.ReadOptions{})
	if err != nil {
		return DBEntry{}, err
	}
	err = json.Unmarshal(dbEntryBytes, &dbEntry)
	return dbEntry, err
}

func Has(shortID string) (bool, error) {
	return Database.Has([]byte(shortID), &opt.ReadOptions{})
}

func CloseDatabase() {
	err := Database.Close()
	if err != nil {
		log.Fatal("could not close database")
	}
}

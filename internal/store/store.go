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

type URLEntry struct {
	LongURL   string
	VRFValue0 string
	VRFProof0 string
	VRFValue1 string
	VRFProof1 string
}

func PutURLEntry(shortURL string, urlEntry URLEntry) error {
	has, err := Has(shortURL)
	if has {
		return errors.New("refusing to overwrite entry")
	}
	if err != nil {
		return err
	}
	urlEntryString, err := json.Marshal(&urlEntry)
	if err != nil {
		return err
	}
	return Database.Put([]byte(shortURL), []byte(urlEntryString), &opt.WriteOptions{
		Sync: true,
	})
}

func GetURLEntry(shortURL string) (URLEntry, error) {
	urlEntry := URLEntry{}
	urlEntryBytes, err := Database.Get([]byte(shortURL), &opt.ReadOptions{})
	if err != nil {
		return URLEntry{}, err
	}
	err = json.Unmarshal(urlEntryBytes, &urlEntry)
	return urlEntry, err
}

func Has(shortURL string) (bool, error) {
	return Database.Has([]byte(shortURL), &opt.ReadOptions{})
}

func CloseDatabase() {
	err := Database.Close()
	if err != nil {
		log.Fatal("could not close database")
	}
}

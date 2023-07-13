/* SPDX-FileCopyrightText: © 2023 Nadim Kobeissi <nadim@symbolic.software>
 * SPDX-License-Identifier: GPL-3.0-only */

package store

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"

	"ducky.zip/m/v2/internal/vrf"
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

func PutURLEntry(shortURL string, longURL string) error {
	has, err := Has(shortURL)
	if has {
		return errors.New("refusing to overwrite entry")
	}
	if err != nil {
		return err
	}
	vrfValue0, vrfProof0 := vrf.GenShortURLProof(shortURL)
	vrfValue1, vrfProof1 := vrf.GenLongURLProof(longURL)
	urlEntry := URLEntry{
		LongURL:   longURL,
		VRFValue0: hex.EncodeToString(vrfValue0),
		VRFProof0: hex.EncodeToString(vrfProof0),
		VRFValue1: hex.EncodeToString(vrfValue1),
		VRFProof1: hex.EncodeToString(vrfProof1),
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

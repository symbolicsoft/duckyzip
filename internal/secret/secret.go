/* SPDX-FileCopyrightText: © 2023 Nadim Kobeissi <nadim@symbolic.software>
 * SPDX-License-Identifier: GPL-3.0-only */

package secret

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type StoreLoad struct {
	VRFK0 string
	VRFK1 string
	EthSK string
}

type StoreStruct struct {
	VRFK0 []byte
	VRFK1 []byte
	EthSK []byte
}

var Store = func() StoreStruct {
	loader := StoreLoad{}
	store := StoreStruct{}
	data, err := os.ReadFile(filepath.Join("internal", "secret", "assets", "secrets.json"))
	if err != nil {
		log.Fatal("could not find secrets")
	}
	err = json.Unmarshal(data, &loader)
	if err != nil {
		log.Fatal("could not decode secrets")
	}
	if len(loader.VRFK0) != 128 {
		if err != nil {
			log.Fatal("incorrect vrfk0 length")
		}
	}
	if len(loader.VRFK1) != 128 {
		if err != nil {
			log.Fatal("incorrect vrfk1 length")
		}
	}
	if len(loader.EthSK) != 64 {
		if err != nil {
			log.Fatal("incorrect ethsk length")
		}
	}
	store.VRFK0, err = hex.DecodeString(loader.VRFK0)
	if err != nil {
		log.Fatal("could not decode vrfk0")
	}
	store.VRFK1, err = hex.DecodeString(loader.VRFK1)
	if err != nil {
		log.Fatal("could not decode vrfk1")
	}
	store.EthSK, err = hex.DecodeString(loader.EthSK)
	if err != nil {
		log.Fatal("could not decode ethsk")
	}
	return store
}()

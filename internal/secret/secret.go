/* SPDX-FileCopyrightText: © 2023 Nadim Kobeissi <nadim@symbolic.software>
 * SPDX-License-Identifier: GPL-3.0-only */

package secret

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/coniks-sys/coniks-go/crypto/vrf"
)

type StoreLoad struct {
	VRFK0 string
	VRFK1 string
	EthSK string
}

type StoreStruct struct {
	VRFK0  []byte
	VRFK1  []byte
	VRFPK0 []byte
	VRFPK1 []byte
	EthSK  []byte
}

var Store = func() StoreStruct {
	loader := StoreLoad{}
	store := StoreStruct{}
	ok := false
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
	store.VRFPK0, ok = vrf.PrivateKey(store.VRFK0).Public()
	if !ok {
		log.Fatal("could not calculate vrfk0 public key")
	}
	store.VRFK1, err = hex.DecodeString(loader.VRFK1)
	if err != nil {
		log.Fatal("could not decode vrfk1")
	}
	store.VRFPK1, ok = vrf.PrivateKey(store.VRFK1).Public()
	if !ok {
		log.Fatal("could not calculate vrfk1 public key")
	}
	store.EthSK, err = hex.DecodeString(loader.EthSK)
	if err != nil {
		log.Fatal("could not decode ethsk")
	}
	return store
}()

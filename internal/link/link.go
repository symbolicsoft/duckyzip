/* SPDX-FileCopyrightText: © 2023 Nadim Kobeissi <nadim@symbolic.software>
 * SPDX-License-Identifier: GPL-3.0-only */

package link

import (
	"encoding/hex"
	"errors"
	"strings"

	"ducky.zip/m/v2/internal/contract"
	"ducky.zip/m/v2/internal/sanitize"
	"ducky.zip/m/v2/internal/store"
	"ducky.zip/m/v2/internal/vrf"
)

func GenShortID(payload string) (string, store.DBEntry, error) {
	var shortID string
	var err error
	alreadyInDatabase := true
	length := 8
	for alreadyInDatabase {
		shortID = RandomString(8)
		alreadyInDatabase, err = store.Has(shortID)
		length++
		if err != nil {
			return shortID, store.DBEntry{}, err
		}
	}
	if !sanitize.CheckShortID(shortID) {
		return shortID, store.DBEntry{}, errors.New("somehow generated invalid short id")
	}
	vrfValue0, vrfProof0 := vrf.GenShortIDProof(shortID)
	vrfValue1, vrfProof1 := vrf.GenPayloadProof(payload)
	dbEntry := store.DBEntry{
		Payload:   payload,
		VRFValue0: hex.EncodeToString(vrfValue0),
		VRFProof0: hex.EncodeToString(vrfProof0),
		VRFValue1: hex.EncodeToString(vrfValue1),
		VRFProof1: hex.EncodeToString(vrfProof1),
	}
	err = store.PutDBEntry(shortID, dbEntry)
	if err != nil {
		return shortID, store.DBEntry{}, err
	}
	err = contract.Write(
		strings.Join([]string{
			dbEntry.VRFValue0,
			dbEntry.VRFProof0,
		}, ""),
		strings.Join([]string{
			dbEntry.VRFValue1,
			dbEntry.VRFProof1,
		}, ""),
	)
	return shortID, dbEntry, err
}

func GetPayload(shortID string) (store.DBEntry, error) {
	if !sanitize.CheckShortID(shortID) {
		return store.DBEntry{}, errors.New("invalid short id")
	}
	dbEntry, err := store.GetDBEntry(shortID)
	if err != nil {
		return store.DBEntry{}, err
	}
	if !sanitize.CheckPayload(dbEntry.Payload) {
		return store.DBEntry{}, errors.New("somehow stored invalid long id")
	}
	return dbEntry, err
}

/* SPDX-FileCopyrightText: © 2023 Nadim Kobeissi <nadim@symbolic.software>
 * SPDX-License-Identifier: GPL-3.0-only */

package shorten

import (
	"encoding/hex"
	"errors"
	"strings"

	"ducky.zip/m/v2/internal/contract"
	"ducky.zip/m/v2/internal/sanitize"
	"ducky.zip/m/v2/internal/store"
	"ducky.zip/m/v2/internal/vrf"
)

func GenShortURL(longURL string) (string, store.URLEntry, error) {
	var shortURL string
	var err error
	alreadyInDatabase := true
	length := 8
	for alreadyInDatabase {
		shortURL = RandomString(8)
		alreadyInDatabase, err = store.Has(shortURL)
		length++
		if err != nil {
			return shortURL, store.URLEntry{}, err
		}
	}
	if !sanitize.CheckShortURL(shortURL) {
		return shortURL, store.URLEntry{}, errors.New("somehow generated invalid short url")
	}
	vrfValue0, vrfProof0 := vrf.GenShortURLProof(shortURL)
	vrfValue1, vrfProof1 := vrf.GenLongURLProof(longURL)
	urlEntry := store.URLEntry{
		LongURL:   longURL,
		VRFValue0: hex.EncodeToString(vrfValue0),
		VRFProof0: hex.EncodeToString(vrfProof0),
		VRFValue1: hex.EncodeToString(vrfValue1),
		VRFProof1: hex.EncodeToString(vrfProof1),
	}
	err = store.PutURLEntry(shortURL, urlEntry)
	if err != nil {
		return shortURL, store.URLEntry{}, err
	}
	err = contract.Write(
		strings.Join([]string{
			urlEntry.VRFValue0,
			urlEntry.VRFProof0,
		}, ""),
		strings.Join([]string{
			urlEntry.VRFValue1,
			urlEntry.VRFProof1,
		}, ""),
	)
	return shortURL, urlEntry, err
}

func GetLongURL(shortURL string) (store.URLEntry, error) {
	if !sanitize.CheckShortURL(shortURL) {
		return store.URLEntry{}, errors.New("invalid short url")
	}
	urlEntry, err := store.GetURLEntry(shortURL)
	if err != nil {
		return store.URLEntry{}, err
	}
	if !sanitize.CheckLongURL(urlEntry.LongURL) {
		return store.URLEntry{}, errors.New("somehow stored invalid long url")
	}
	return urlEntry, err
}

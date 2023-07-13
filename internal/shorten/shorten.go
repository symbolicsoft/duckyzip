/* SPDX-FileCopyrightText: © 2023 Nadim Kobeissi <nadim@symbolic.software>
 * SPDX-License-Identifier: GPL-3.0-only */

package shorten

import (
	"errors"
	"fmt"

	"ducky.zip/m/v2/internal/sanitize"
	"ducky.zip/m/v2/internal/store"
)

func GenShortURL(longURL string) (string, error) {
	var shortURL string
	var err error
	alreadyInDatabase := true
	for alreadyInDatabase {
		shortURL = RandomString(8)
		alreadyInDatabase, err = store.Has(shortURL)
		if err != nil {
			return shortURL, err
		}
	}
	if !sanitize.CheckShortURL(shortURL) {
		return shortURL, errors.New("somehow generated invalid short url")
	}
	err = store.PutURLEntry(shortURL, longURL)
	return shortURL, err
}

func GetLongURL(shortURL string) (string, error) {
	if !sanitize.CheckShortURL(shortURL) {
		return "", errors.New("invalid short url")
	}
	urlEntry, err := store.GetURLEntry(shortURL)
	if err != nil {
		return "", err
	}
	if !sanitize.CheckLongURL(urlEntry.LongURL) {
		return "", errors.New("somehow stored invalid long url")
	}
	fmt.Println(urlEntry)
	return urlEntry.LongURL, err
}

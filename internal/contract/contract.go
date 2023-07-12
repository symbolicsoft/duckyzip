/* SPDX-FileCopyrightText: © 2023 Nadim Kobeissi <nadim@symbolic.software>
 * SPDX-License-Identifier: GPL-3.0-only */

package contract

import (
	"crypto/ecdsa"
	"log"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/crypto"
)

var ethPrivKey = func() *ecdsa.PrivateKey {
	keyStringBytes, err := os.ReadFile(filepath.Join("secret", "ethPrivKey.txt"))
	if err != nil {
		log.Fatal("could not read ethereum private key")
	}
	privateKey, err := crypto.HexToECDSA(string(keyStringBytes))
	if err != nil {
		log.Fatal("could not parse ethereum private key")
	}
	return privateKey
}()

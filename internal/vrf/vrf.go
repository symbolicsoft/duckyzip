/* SPDX-FileCopyrightText: © 2023 Nadim Kobeissi <nadim@symbolic.software>
 * SPDX-License-Identifier: GPL-3.0-only */

package vrf

import (
	"ducky.zip/m/v2/internal/secret"
	"github.com/coniks-sys/coniks-go/crypto/vrf"
)

func GenShortURLProof(shortURL string) ([]byte, []byte) {
	vrfValue, vrfProof := vrf.PrivateKey(secret.Store.VRFK0).Prove([]byte(shortURL))
	return vrfValue, vrfProof
}

func GenLongURLProof(longURL string) ([]byte, []byte) {
	vrfValue, vrfProof := vrf.PrivateKey(secret.Store.VRFK1).Prove([]byte(longURL))
	return vrfValue, vrfProof
}

func VerifyShortURLProof(shortURL string, vrfValue []byte, vrfProof []byte) bool {
	return vrf.PublicKey(secret.Store.VRFPK0).Verify([]byte(shortURL), vrfValue, vrfProof)
}

func VerifyLongURLProof(longURL string, vrfValue []byte, vrfProof []byte) bool {
	return vrf.PublicKey(secret.Store.VRFPK1).Verify([]byte(longURL), vrfValue, vrfProof)
}

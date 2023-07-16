/* SPDX-FileCopyrightText: © 2023 Nadim Kobeissi <nadim@symbolic.software>
 * SPDX-License-Identifier: GPL-3.0-only */

package vrf

import (
	"ducky.zip/m/v2/internal/secret"
	"github.com/coniks-sys/coniks-go/crypto/vrf"
)

func GenShortIDProof(shortID string) ([]byte, []byte) {
	vrfValue, vrfProof := vrf.PrivateKey(secret.Store.VRFK0).Prove([]byte(shortID))
	return vrfValue, vrfProof
}

func GenPayloadProof(payload string) ([]byte, []byte) {
	vrfValue, vrfProof := vrf.PrivateKey(secret.Store.VRFK1).Prove([]byte(payload))
	return vrfValue, vrfProof
}

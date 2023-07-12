/* SPDX-FileCopyrightText: © 2023 Nadim Kobeissi <nadim@symbolic.software>
 * SPDX-License-Identifier: GPL-3.0-only */

package shorten

import (
	"crypto/rand"
	"log"
)

func RandomString(n int) string {
	const base58 = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal("csprng failure")
	}
	for i := range b {
		b[i] = base58[b[i]%byte(len(base58))]
	}
	return string(b)
}

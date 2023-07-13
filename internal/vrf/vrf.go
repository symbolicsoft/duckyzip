/* SPDX-FileCopyrightText: © 2023 Nadim Kobeissi <nadim@symbolic.software>
 * SPDX-License-Identifier: GPL-3.0-only */

package vrf

import (
	"log"

	"github.com/coniks-sys/coniks-go/crypto/vrf"
)

func Test() {
	_, err := vrf.GenerateKey(nil)
	if err != nil {
		log.Fatal(err)
	}
}

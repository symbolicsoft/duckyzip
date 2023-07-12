/* SPDX-FileCopyrightText: © 2023 Nadim Kobeissi <nadim@symbolic.software>
 * SPDX-License-Identifier: GPL-3.0-only */

package vrf

import (
	"fmt"
	"log"

	"github.com/coniks-sys/coniks-go/crypto/vrf"
)

func test() {
	privateKey, err := vrf.GenerateKey(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(privateKey)
}

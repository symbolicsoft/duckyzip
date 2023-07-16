/* SPDX-FileCopyrightText: © 2023 Nadim Kobeissi <nadim@symbolic.software>
 * SPDX-License-Identifier: GPL-3.0-only */

package main

import (
	"encoding/hex"
	"syscall/js"

	"github.com/coniks-sys/coniks-go/crypto/vrf"
)

var VRFPK0, VRFPK1 = func() ([]byte, []byte) {
	vrfpk0, _ := hex.DecodeString("e77f139963840a1ec719133895e4a8687dca78d439566e35a2d0e4e05ef46f09")
	vrfpk1, _ := hex.DecodeString("9a37be5b9a35c442f093600695df32957af6f16f698a06caab5887fc7472975c")
	return vrfpk0, vrfpk1
}()

func VerifyShortIDProof(this js.Value, args []js.Value) any {
	shortID := []byte(args[0].String())
	vrfValue, err1 := hex.DecodeString(args[1].String())
	vrfProof, err2 := hex.DecodeString(args[2].String())
	if err1 != nil || err2 != nil {
		return false
	}
	return vrf.PublicKey(VRFPK0).Verify(shortID, vrfValue, vrfProof)
}

func VerifyPayloadProof(this js.Value, args []js.Value) any {
	shortID := []byte(args[0].String())
	vrfValue, err1 := hex.DecodeString(args[1].String())
	vrfProof, err2 := hex.DecodeString(args[2].String())
	if err1 != nil || err2 != nil {
		return false
	}
	return vrf.PublicKey(VRFPK1).Verify(shortID, vrfValue, vrfProof)
}

func main() {
	done := make(chan struct{}, 0)
	js.Global().Set("VerifyShortIDProof", js.FuncOf(VerifyShortIDProof))
	js.Global().Set("VerifyPayloadProof", js.FuncOf(VerifyPayloadProof))
	<-done
}

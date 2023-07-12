/* SPDX-FileCopyrightText: © 2023 Nadim Kobeissi <nadim@symbolic.software>
 * SPDX-License-Identifier: GPL-3.0-only */

package main

import (
	"fmt"
	"os"
	"os/signal"

	"ducky.zip/m/v2/internal/store"
	"ducky.zip/m/v2/internal/web"
)

func main() {
	fmt.Println("DuckyZip")
	handleSigInterrupt()
	web.StartServer()
}

func handleSigInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			fmt.Println("Closing database...")
			store.CloseDatabase()
			os.Exit(0)
		}
	}()
}

/* SPDX-FileCopyrightText: © 2023 Nadim Kobeissi <nadim@symbolic.software>
 * SPDX-License-Identifier: GPL-3.0-only */

package contract

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"strings"

	simpleStorage "ducky.zip/m/v2/internal/contract/simplestorage"
	"ducky.zip/m/v2/internal/secret"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const ContractAddress = "0x082Ff59678C0c5781f164c29C5A8f90008D5b1c0"
const InfuraAddress = "https://optimism-mainnet.infura.io/v3/"

func Read(key string) (string, error) {
	client, err := ethclient.Dial(strings.Join([]string{
		InfuraAddress,
		secret.Store.Infura,
	}, ""))
	if err != nil {
		return "", err
	}
	address := common.HexToAddress(ContractAddress)
	instance, err := simpleStorage.NewSimpleStorage(address, client)
	if err != nil {
		return "", err
	}
	return instance.GetValue(nil, key)
}

func Write(key string, value string) error {
	client, err := ethclient.Dial(strings.Join([]string{
		InfuraAddress,
		secret.Store.Infura,
	}, ""))
	if err != nil {
		return err
	}
	privateKey, err := crypto.ToECDSA(secret.Store.EthSK)
	if err != nil {
		return err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(10))
	if err != nil {
		return err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	address := common.HexToAddress(ContractAddress)
	instance, err := simpleStorage.NewSimpleStorage(address, client)
	if err != nil {
		return err
	}
	_, err = instance.AddToStore(auth, key, value)
	return err
}

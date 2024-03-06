package helpers

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"math/big"
	"net/http"
)

func GetEthClient(r *http.Request) (*bind.TransactOpts, *ethclient.Client, error) {
	client, err := ethclient.Dial(NetworkConfig(r).RpcUrl.String())
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to connect client to the RPC URL")
	}
	publicKeyECDSA, ok := NetworkConfig(r).PrivateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, nil, errors.New("failed to casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to to get account nonce")
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get gas price")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(NetworkConfig(r).PrivateKey, big.NewInt(NetworkConfig(r).ChainId))
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to create a transaction signer from a single private key")
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasPrice = gasPrice
	if NetworkConfig(r).GasLimit != 0 {
		auth.GasLimit = NetworkConfig(r).GasLimit
	}
	return auth, client, nil
}

func UpdateAuth(r *http.Request, client *ethclient.Client) (*bind.TransactOpts, error) {
	publicKeyECDSA, ok := NetworkConfig(r).PrivateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("failed to casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, errors.Wrap(err, "failed to to get account nonce")
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get gas price")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(NetworkConfig(r).PrivateKey, big.NewInt(NetworkConfig(r).ChainId))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create a transaction signer from a single private key")
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasPrice = gasPrice
	if NetworkConfig(r).GasLimit != 0 {
		auth.GasLimit = NetworkConfig(r).GasLimit
	}
	return auth, nil
}

package BTC

import (
	"encoding/hex"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"testing"
)

var client *rpcclient.Client

func init() {
	config := &rpcclient.ConnConfig{
		Host:         "https://go.getblock.io/3ebe61c87d824529b16f2fff7b42ecc7",
		HTTPPostMode: true,
		//DisableTLS:   true,
	}
	var err error
	client, err = rpcclient.New(config, nil)
	if err != nil {
		panic(err)
	}
}

func TestGeneratePrivateKey(t *testing.T) {
	privateKey, err := btcec.NewPrivateKey()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("private key: %x", privateKey.Serialize())
	// 6b15482c2d6cd81d914925343ff2bed1f8f6a1f52bc7414881ecf7640aadcc75

	publicKey := privateKey.PubKey()

	t.Logf("public key: %x", publicKey.SerializeCompressed())
	// 03b56c0e311edbe942d403a3af81bbef183d7de3e2510ca9b8b35a2071c5ef683d
}

func TestFromPrivateKeyByte(t *testing.T) {
	bytes, err := hex.DecodeString("6b15482c2d6cd81d914925343ff2bed1f8f6a1f52bc7414881ecf7640aadcc75")
	if err != nil {
		t.Error(err)
		return
	}
	privateKey, publicKey := btcec.PrivKeyFromBytes(bytes)
	t.Logf("private key: %x", privateKey.Serialize())
	t.Logf("public key: %x", publicKey.SerializeCompressed())
}

func TestAddress(t *testing.T) {
	bytes, err := hex.DecodeString("6b15482c2d6cd81d914925343ff2bed1f8f6a1f52bc7414881ecf7640aadcc75")
	if err != nil {
		t.Error(err)
		return
	}
	_, publicKey := btcec.PrivKeyFromBytes(bytes)
	publicKeyCompressed := publicKey.SerializeCompressed()

	net := &chaincfg.TestNet3Params

	addressPubKey, err := btcutil.NewAddressPubKey(publicKeyCompressed, net)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("address: %s", addressPubKey.EncodeAddress())

	addressPubKeyHash, err := btcutil.NewAddressPubKeyHash(addressPubKey.AddressPubKeyHash().ScriptAddress(), net)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("address: %s", addressPubKeyHash.EncodeAddress())

	script, err := txscript.PayToAddrScript(addressPubKey.AddressPubKeyHash())
	if err != nil {
		t.Error(err)
		return
	}

	addressScriptHash, err := btcutil.NewAddressScriptHash(script, net)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("address: %s", addressScriptHash.EncodeAddress())

	addressWitnessPubKeyHash, err := btcutil.NewAddressWitnessPubKeyHash(addressPubKeyHash.ScriptAddress(), net)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("address: %s", addressWitnessPubKeyHash.EncodeAddress())

	addressWitnessScriptHash, err := btcutil.NewAddressWitnessScriptHash(script, net)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("address: %s", addressWitnessScriptHash.EncodeAddress())
}

func TestGetBalance(t *testing.T) {
	balance, err := client.GetBalance("mmKfHs5ebWW1roeuRQLKe28YXwRLkhhSy2")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("balance: %d", balance)
}

func TestClient(t *testing.T) {
	message, err := client.RawRequest("getblockcount", nil)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("message: %s", message)
}

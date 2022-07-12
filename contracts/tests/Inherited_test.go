package tests

import (
	"fmt"
	"github.com/janetachain/janetachain/accounts/abi/bind"
	"github.com/janetachain/janetachain/accounts/abi/bind/backends"
	"github.com/janetachain/janetachain/common"
	"github.com/janetachain/janetachain/core"
	"github.com/janetachain/janetachain/crypto"
	"github.com/janetachain/janetachain/log"
	"math/big"
	"os"
	"testing"
)

var (
	mainKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	mainAddr   = crypto.PubkeyToAddress(mainKey.PublicKey)
)

func TestPriceFeed(t *testing.T) {
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.LvlTrace)
	log.Root().SetHandler(glogger)
	common.TIPJanotCancellationFee = big.NewInt(0)
	// init genesis
	contractBackend := backends.NewSimulatedBackend(core.GenesisAlloc{
		mainAddr: {Balance: big.NewInt(0).Mul(big.NewInt(10000000000000), big.NewInt(10000000000000))},
	})
	transactOpts := bind.NewKeyedTransactor(mainKey)
	// deploy payer swap SMC
	addr, contract, err := DeployMyInherited(transactOpts, contractBackend)
	if err != nil {
		t.Fatal("can't deploy smart contract: ", err)
	}
	fmt.Println("addr", addr.Hex())
	tx, err := contract.Foo()
	if err != nil {
		t.Fatal("can't run function Foo() in  smart contract: ", err)
	}
	fmt.Println("tx", tx)

}

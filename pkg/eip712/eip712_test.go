package eip712

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestBuildEIP712DomainSeparator(t *testing.T) {
	expectedAmoy := common.HexToHash("0x6b19e618edbcb3dff6263b5e2ff88e1bf35df68b801a7da3825c87e14d7ed8bc")
	name := crypto.Keccak256Hash([]byte("Kuest CTF Exchange"))
	version := crypto.Keccak256Hash([]byte("1"))
	chainId := big.NewInt(80002)
	address := common.HexToAddress("0x0000000000000000000000000000000000000000")

	actual, err := BuildEIP712DomainSeparator(name, version, chainId, address)
	assert.NoError(t, err)
	assert.NotEmpty(t, actual)
	assert.Equal(t, expectedAmoy.String(), actual.String())

	expectedPolygon := common.HexToHash("5387ff527161b318d38cceea7c9048368070f42aec88d28ca137a6e85adab706")
	chainId = big.NewInt(137)

	actual, err = BuildEIP712DomainSeparator(name, version, chainId, address)
	assert.NoError(t, err)
	assert.NotEmpty(t, actual)
	assert.NotEqual(t, expectedAmoy.String(), actual.String())
	assert.Equal(t, expectedPolygon.String(), actual.String())
}

func TestBuildEIP712DomainSeparatorNoContract(t *testing.T) {
	// Calculated in foundry
	expectedAmoy := common.HexToHash("0x99a2c264407aa721c2e5de6d3013e84caa68838ede505e61986d1e7e7c8b7a8b")
	chainId := big.NewInt(80002)

	name := crypto.Keccak256Hash([]byte("Kuest CTF Exchange"))
	version := crypto.Keccak256Hash([]byte("1"))

	actual, err := BuildEIP712DomainSeparatorNoContract(name, version, chainId)
	assert.NoError(t, err)
	assert.NotEmpty(t, actual)
	assert.Equal(t, expectedAmoy.String(), actual.String())

	// Calculated in foundry
	expectedPolygon := common.HexToHash("0a007ad5fc0a840753026d60eab9679d79a245a7ebc45a2ac5eea6f7fe2b2ab6")
	chainId = big.NewInt(137)

	actual, err = BuildEIP712DomainSeparatorNoContract(name, version, chainId)
	assert.NoError(t, err)
	assert.NotEmpty(t, actual)
	assert.NotEqual(t, expectedAmoy.String(), actual.String())
	assert.Equal(t, expectedPolygon.String(), actual.String())
}

func TestHashTypedDataV4(t *testing.T) {
	name := crypto.Keccak256Hash([]byte("Kuest CTF Exchange"))
	version := crypto.Keccak256Hash([]byte("1"))
	chainId := big.NewInt(80002)
	address := common.HexToAddress("0x0000000000000000000000000000000000000000")

	domainSeparator, err := BuildEIP712DomainSeparator(name, version, chainId, address)
	assert.NoError(t, err)
	assert.NotEmpty(t, domainSeparator)

	types := []abi.Type{
		Bytes32,
		String,
		Uint256,
	}
	values := []interface{}{
		crypto.Keccak256Hash([]byte("MockObj(string name, uint256 id)")),
		"test",
		big.NewInt(1),
	}

	dataHashBytes, err := HashTypedDataV4(domainSeparator, types, values)
	assert.NoError(t, err)
	assert.NotEmpty(t, dataHashBytes)

	expectedTypedDataHash := "0xf28dd0f4e2fa155c58c44b76b0b5569f07f6c6257f65f94d535df0c8679cc92d"
	assert.Equal(t, expectedTypedDataHash, dataHashBytes.String())
}

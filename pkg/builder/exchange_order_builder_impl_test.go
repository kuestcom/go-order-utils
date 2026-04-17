package builder

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/kuestcom/go-order-utils/pkg/model"
	"github.com/stretchr/testify/assert"
)

var (
	chainId = new(big.Int).SetInt64(80002)
	// publicly known private key
	privateKey, _ = crypto.ToECDSA(common.Hex2Bytes("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"))
	// private key address
	signerAddress = common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	salt = int64(479249096354)
)

func TestBuildOrder(t *testing.T) {
	// random salt
	builder := NewExchangeOrderBuilderImpl(chainId, nil)

	order, err := builder.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		Taker:       "0x0",
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	assert.True(t, order.Salt.Int64() > 0)
	assert.Equal(t, order.Maker, signerAddress)
	assert.Equal(t, order.Signer, signerAddress)
	assert.Equal(t, order.Taker, common.HexToAddress("0x0"))
	assert.Equal(t, order.TokenId.String(), "1234")
	assert.Equal(t, order.MakerAmount.String(), "100000000")
	assert.Equal(t, order.TakerAmount.String(), "50000000")
	assert.Equal(t, order.Side.String(), "0")
	assert.Equal(t, order.Expiration.String(), "0")
	assert.Equal(t, order.Nonce.String(), "0")
	assert.Equal(t, order.FeeRateBps.String(), "100")
	assert.Equal(t, order.SignatureType.String(), "0")

	// specific salt
	builder = NewExchangeOrderBuilderImpl(chainId, func() int64 { return salt })

	order, err = builder.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		Taker:       "0x1",
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	assert.Equal(t, order.Salt.Int64(), int64(salt))
	assert.Equal(t, order.Maker, signerAddress)
	assert.Equal(t, order.Signer, signerAddress)
	assert.Equal(t, order.Taker, common.HexToAddress("0x1"))
	assert.Equal(t, order.TokenId.String(), "1234")
	assert.Equal(t, order.MakerAmount.String(), "100000000")
	assert.Equal(t, order.TakerAmount.String(), "50000000")
	assert.Equal(t, order.Side.String(), "0")
	assert.Equal(t, order.Expiration.String(), "0")
	assert.Equal(t, order.Nonce.String(), "0")
	assert.Equal(t, order.FeeRateBps.String(), "100")
	assert.Equal(t, order.SignatureType.String(), "0")
}

func TestBuildOrderHash(t *testing.T) {
	// FEE
	// random salt
	builder := NewExchangeOrderBuilderImpl(chainId, nil)

	order, err := builder.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		Taker:       common.HexToAddress("0x0").Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	orderHash, err := builder.BuildOrderHash(order, model.CTFExchange)
	assert.NoError(t, err)
	assert.NotNil(t, orderHash)

	// specific salt
	builder = NewExchangeOrderBuilderImpl(chainId, func() int64 { return salt })

	order, err = builder.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		Taker:       common.HexToAddress("0x0").Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	orderHash, err = builder.BuildOrderHash(order, model.CTFExchange)
	assert.NoError(t, err)
	assert.NotNil(t, orderHash)

	expectedOrderHash := common.HexToHash("8dad71a6d00c68e46feffafee2e99694aa28d3b1669b0dc15481ef44aa19910e")
	assert.Equal(t, expectedOrderHash.String(), orderHash.String())

	// NegRisk
	// random salt
	builder = NewExchangeOrderBuilderImpl(chainId, nil)

	order, err = builder.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		Taker:       common.HexToAddress("0x0").Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	orderHash, err = builder.BuildOrderHash(order, model.NegRiskCTFExchange)
	assert.NoError(t, err)
	assert.NotNil(t, orderHash)

	// specific salt
	builder = NewExchangeOrderBuilderImpl(chainId, func() int64 { return salt })

	order, err = builder.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		Taker:       common.HexToAddress("0x0").Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	orderHash, err = builder.BuildOrderHash(order, model.NegRiskCTFExchange)
	assert.NoError(t, err)
	assert.NotNil(t, orderHash)

	expectedOrderHash = common.HexToHash("8b435cbda05f86756f3acd1deb3f57883e442a19453936131a6128887e437882")
	assert.Equal(t, expectedOrderHash.String(), orderHash.String())
}

func TestBuildOrderSignature(t *testing.T) {
	// FEE
	// random salt
	builder := NewExchangeOrderBuilderImpl(chainId, nil)

	order, err := builder.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		Taker:       common.HexToAddress("0x0").Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	orderHash, err := builder.BuildOrderHash(order, model.CTFExchange)
	assert.NoError(t, err)
	assert.NotNil(t, orderHash)

	orderSignature, err := builder.BuildOrderSignature(privateKey, orderHash)
	assert.NoError(t, err)
	assert.NotNil(t, orderSignature)

	// specific salt
	builder = NewExchangeOrderBuilderImpl(chainId, func() int64 { return salt })

	order, err = builder.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		Taker:       common.HexToAddress("0x0").Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	orderHash, err = builder.BuildOrderHash(order, model.CTFExchange)
	assert.NoError(t, err)
	assert.NotNil(t, orderHash)

	orderSignature, err = builder.BuildOrderSignature(privateKey, orderHash)
	assert.NoError(t, err)
	assert.NotNil(t, orderSignature)

	expectedSignature := "9cd249e296f7f84aa6bf832994c19196057a28cb5fe763f15f539168dc7b440e7b2673e8ae22d8247d078a6fd6322e39c5869bdbd7d9b6b1e074406f183060531c"
	assert.Equal(t, expectedSignature, common.Bytes2Hex(orderSignature))

	// NegRisk
	// random salt
	builder = NewExchangeOrderBuilderImpl(chainId, nil)

	order, err = builder.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		Taker:       common.HexToAddress("0x0").Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	orderHash, err = builder.BuildOrderHash(order, model.NegRiskCTFExchange)
	assert.NoError(t, err)
	assert.NotNil(t, orderHash)

	orderSignature, err = builder.BuildOrderSignature(privateKey, orderHash)
	assert.NoError(t, err)
	assert.NotNil(t, orderSignature)

	// specific salt
	builder = NewExchangeOrderBuilderImpl(chainId, func() int64 { return salt })

	order, err = builder.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		Taker:       common.HexToAddress("0x0").Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	orderHash, err = builder.BuildOrderHash(order, model.NegRiskCTFExchange)
	assert.NoError(t, err)
	assert.NotNil(t, orderHash)

	orderSignature, err = builder.BuildOrderSignature(privateKey, orderHash)
	assert.NoError(t, err)
	assert.NotNil(t, orderSignature)

	expectedSignature = "3af61c3c156c626fe594a06f5199673df8a1009091370a650cf6f99ac4120bdf2b2e8fd4eea2511e3ce44ec77e6607c5ddeda9e1dd8f16c13880b6064c25d28c1b"
	assert.Equal(t, expectedSignature, common.Bytes2Hex(orderSignature))
}

func TestBuildSignedOrder(t *testing.T) {
	// FEE
	// random salt
	builder := NewExchangeOrderBuilderImpl(chainId, nil)

	signedOrder, err := builder.BuildSignedOrder(privateKey, &model.OrderData{
		Maker:       signerAddress.Hex(),
		Taker:       common.HexToAddress("0x0").Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	}, model.CTFExchange)
	assert.NoError(t, err)
	assert.NotNil(t, signedOrder)

	assert.True(t, signedOrder.Salt.Int64() > 0)
	assert.Equal(t, signedOrder.Maker, signerAddress)
	assert.Equal(t, signedOrder.Signer, signerAddress)
	assert.Equal(t, signedOrder.TokenId.String(), "1234")
	assert.Equal(t, signedOrder.MakerAmount.String(), "100000000")
	assert.Equal(t, signedOrder.TakerAmount.String(), "50000000")
	assert.Equal(t, signedOrder.Side.String(), "0")
	assert.Equal(t, signedOrder.Expiration.String(), "0")
	assert.Equal(t, signedOrder.Nonce.String(), "0")
	assert.Equal(t, signedOrder.FeeRateBps.String(), "100")
	assert.Equal(t, signedOrder.SignatureType.String(), "0")
	assert.NotEmpty(t, signedOrder.Signature)
	assert.NotEmpty(t, hex.EncodeToString(signedOrder.Signature))

	// specific salt
	builder = NewExchangeOrderBuilderImpl(chainId, func() int64 { return salt })

	signedOrder, err = builder.BuildSignedOrder(privateKey, &model.OrderData{
		Maker:       signerAddress.Hex(),
		Taker:       common.HexToAddress("0x0").Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	}, model.CTFExchange)
	assert.NoError(t, err)
	assert.NotNil(t, signedOrder)

	assert.Equal(t, signedOrder.Salt.Int64(), salt)
	assert.Equal(t, signedOrder.Maker, signerAddress)
	assert.Equal(t, signedOrder.Signer, signerAddress)
	assert.Equal(t, signedOrder.TokenId.String(), "1234")
	assert.Equal(t, signedOrder.MakerAmount.String(), "100000000")
	assert.Equal(t, signedOrder.TakerAmount.String(), "50000000")
	assert.Equal(t, signedOrder.Side.String(), "0")
	assert.Equal(t, signedOrder.Expiration.String(), "0")
	assert.Equal(t, signedOrder.Nonce.String(), "0")
	assert.Equal(t, signedOrder.FeeRateBps.String(), "100")
	assert.Equal(t, signedOrder.SignatureType.String(), "0")
	assert.NotEmpty(t, hex.EncodeToString(signedOrder.Signature))

	expectedSignature := "9cd249e296f7f84aa6bf832994c19196057a28cb5fe763f15f539168dc7b440e7b2673e8ae22d8247d078a6fd6322e39c5869bdbd7d9b6b1e074406f183060531c"
	assert.Equal(t, expectedSignature, common.Bytes2Hex(signedOrder.Signature))

	// NegRisk
	// random salt
	builder = NewExchangeOrderBuilderImpl(chainId, nil)

	signedOrder, err = builder.BuildSignedOrder(privateKey, &model.OrderData{
		Maker:       signerAddress.Hex(),
		Taker:       common.HexToAddress("0x0").Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	}, model.NegRiskCTFExchange)
	assert.NoError(t, err)
	assert.NotNil(t, signedOrder)

	assert.True(t, signedOrder.Salt.Int64() > 0)
	assert.Equal(t, signedOrder.Maker, signerAddress)
	assert.Equal(t, signedOrder.Signer, signerAddress)
	assert.Equal(t, signedOrder.TokenId.String(), "1234")
	assert.Equal(t, signedOrder.MakerAmount.String(), "100000000")
	assert.Equal(t, signedOrder.TakerAmount.String(), "50000000")
	assert.Equal(t, signedOrder.Side.String(), "0")
	assert.Equal(t, signedOrder.Expiration.String(), "0")
	assert.Equal(t, signedOrder.Nonce.String(), "0")
	assert.Equal(t, signedOrder.FeeRateBps.String(), "100")
	assert.Equal(t, signedOrder.SignatureType.String(), "0")
	assert.NotEmpty(t, signedOrder.Signature)
	assert.NotEmpty(t, hex.EncodeToString(signedOrder.Signature))

	// specific salt
	builder = NewExchangeOrderBuilderImpl(chainId, func() int64 { return salt })

	signedOrder, err = builder.BuildSignedOrder(privateKey, &model.OrderData{
		Maker:       signerAddress.Hex(),
		Taker:       common.HexToAddress("0x0").Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	}, model.NegRiskCTFExchange)
	assert.NoError(t, err)
	assert.NotNil(t, signedOrder)

	assert.Equal(t, signedOrder.Salt.Int64(), salt)
	assert.Equal(t, signedOrder.Maker, signerAddress)
	assert.Equal(t, signedOrder.Signer, signerAddress)
	assert.Equal(t, signedOrder.TokenId.String(), "1234")
	assert.Equal(t, signedOrder.MakerAmount.String(), "100000000")
	assert.Equal(t, signedOrder.TakerAmount.String(), "50000000")
	assert.Equal(t, signedOrder.Side.String(), "0")
	assert.Equal(t, signedOrder.Expiration.String(), "0")
	assert.Equal(t, signedOrder.Nonce.String(), "0")
	assert.Equal(t, signedOrder.FeeRateBps.String(), "100")
	assert.Equal(t, signedOrder.SignatureType.String(), "0")
	assert.NotEmpty(t, hex.EncodeToString(signedOrder.Signature))

	expectedSignature = "3af61c3c156c626fe594a06f5199673df8a1009091370a650cf6f99ac4120bdf2b2e8fd4eea2511e3ce44ec77e6607c5ddeda9e1dd8f16c13880b6064c25d28c1b"
	assert.Equal(t, expectedSignature, common.Bytes2Hex(signedOrder.Signature))
}

func TestBuildSignedOrder2(t *testing.T) {
	builder := NewExchangeOrderBuilderImpl(chainId, nil)

	signedOrder, err := builder.BuildSignedOrder(privateKey, &model.OrderData{
		Maker:         "0xaFB8270A801862270FebB3763505b136491e557b",
		Signer:        "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
		Taker:         common.HexToAddress("0x0").Hex(),
		TokenId:       "100",
		MakerAmount:   "50000000",
		TakerAmount:   "100000000",
		Side:          model.BUY,
		FeeRateBps:    "100",
		Nonce:         "0",
		Expiration:    "0",
		SignatureType: model.KUEST_GNOSIS_SAFE,
	}, model.NegRiskCTFExchange)
	assert.NoError(t, err)
	assert.NotNil(t, signedOrder)

}

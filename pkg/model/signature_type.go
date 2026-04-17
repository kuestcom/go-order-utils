package model

type SignatureType = int

const (
	// ECDSA EIP712 signatures signed by EOAs
	EOA SignatureType = iota

	// EIP712 signatures signed by EOAs that own Kuest Proxy wallets
	KUEST_PROXY

	// EIP712 signatures signed by EOAs that own Kuest Gnosis safes
	KUEST_GNOSIS_SAFE
)

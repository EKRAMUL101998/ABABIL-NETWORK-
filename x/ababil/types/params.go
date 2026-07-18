package types

type Params struct {
	ChainName string
	CoinDenom string
}

func DefaultParams() Params {
	return Params{
		ChainName: "ABABIL Network",
		CoinDenom: "abl",
	}
}

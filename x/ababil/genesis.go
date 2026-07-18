package ababil

type GenesisState struct{}

func DefaultGenesis() *GenesisState {
	return &GenesisState{}
}

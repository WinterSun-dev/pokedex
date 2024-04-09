package state

type State struct {
	MapCurrent  string
	MapNext     string
	MapPrevious string
	Area        string
	Pokemon     string
}

type PokeDataInDex struct {
	Name           string
	BaseExperience int
	Height         int
	Weight         int
	Stats          []DexStat
	Types          []string
}
type DexStat struct {
	Name  string
	Value int
}

type PokeEntry struct {
	PokeData   PokeDataInDex
	IsCaptured bool
}

var PokeCatalog = make(map[string]PokeEntry)

func NewState() State {
	return State{
		MapCurrent:  "https://pokeapi.co/api/v2/location-area/",
		MapNext:     "https://pokeapi.co/api/v2/location-area/",
		MapPrevious: "https://pokeapi.co/api/v2/location-area/",
	}
}

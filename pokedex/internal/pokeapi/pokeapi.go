package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/WinterSun-dev/pokedex/internal/pokecache"
	"github.com/WinterSun-dev/pokedex/internal/state"
)

type apiResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Response `json:"results"`
}
type Response struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type apiArea struct {
	EncounterMethodRates []EncounterMethodRates `json:"encounter_method_rates"`
	GameIndex            int                    `json:"game_index"`
	ID                   int                    `json:"id"`
	Location             Response               `json:"location"`
	Name                 string                 `json:"name"`
	Names                []Names                `json:"names"`
	PokemonEncounters    []PokemonEncounters    `json:"pokemon_encounters"`
}
type VersionDetails1 struct {
	Rate    int      `json:"rate"`
	Version Response `json:"version"`
}
type EncounterMethodRates struct {
	EncounterMethod Response          `json:"encounter_method"`
	VersionDetails  []VersionDetails1 `json:"version_details"`
}

type Names struct {
	Language Response `json:"language"`
	Name     string   `json:"name"`
}
type EncounterDetails struct {
	Chance          int      `json:"chance"`
	ConditionValues []any    `json:"condition_values"`
	MaxLevel        int      `json:"max_level"`
	Method          Response `json:"method"`
	MinLevel        int      `json:"min_level"`
}
type VersionDetails2 struct {
	EncounterDetails []EncounterDetails `json:"encounter_details"`
	MaxChance        int                `json:"max_chance"`
	Version          Response           `json:"version"`
}
type PokemonEncounters struct {
	Pokemon        Response          `json:"pokemon"`
	VersionDetails []VersionDetails2 `json:"version_details"`
}

type ApiPokemon struct {
	ID                     int           `json:"id"`
	Name                   string        `json:"name"`
	BaseExperience         int           `json:"base_experience"`
	Height                 int           `json:"height"`
	IsDefault              bool          `json:"is_default"`
	Order                  int           `json:"order"`
	Weight                 int           `json:"weight"`
	Abilities              []Abilities   `json:"abilities"`
	Forms                  []Response    `json:"forms"`
	GameIndices            []GameIndices `json:"game_indices"`
	HeldItems              []HeldItems   `json:"held_items"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	Moves                  []Moves       `json:"moves"`
	Species                Response      `json:"species"`
	Sprites                Sprites       `json:"sprites"`
	Cries                  Cries         `json:"cries"`
	Stats                  []Stats       `json:"stats"`
	Types                  []Types       `json:"types"`
	PastTypes              []PastTypes   `json:"past_types"`
}
type Abilities struct {
	IsHidden bool     `json:"is_hidden"`
	Slot     int      `json:"slot"`
	Ability  Response `json:"ability"`
}
type GameIndices struct {
	GameIndex int      `json:"game_index"`
	Version   Response `json:"version"`
}
type VersionDetails struct {
	Rarity  int      `json:"rarity"`
	Version Response `json:"version"`
}
type HeldItems struct {
	Item           Response         `json:"item"`
	VersionDetails []VersionDetails `json:"version_details"`
}
type VersionGroupDetails struct {
	LevelLearnedAt  int      `json:"level_learned_at"`
	VersionGroup    Response `json:"version_group"`
	MoveLearnMethod Response `json:"move_learn_method"`
}
type Moves struct {
	Move                Response              `json:"move"`
	VersionGroupDetails []VersionGroupDetails `json:"version_group_details"`
}
type DreamWorld struct {
	FrontDefault string `json:"front_default"`
	FrontFemale  any    `json:"front_female"`
}
type Home struct {
	FrontDefault     string `json:"front_default"`
	FrontFemale      any    `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale any    `json:"front_shiny_female"`
}
type OfficialArtwork struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}
type Showdown struct {
	BackDefault      string `json:"back_default"`
	BackFemale       any    `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  any    `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      any    `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale any    `json:"front_shiny_female"`
}
type Other struct {
	DreamWorld      DreamWorld      `json:"dream_world"`
	Home            Home            `json:"home"`
	OfficialArtwork OfficialArtwork `json:"official-artwork"`
	Showdown        Showdown        `json:"showdown"`
}
type RedBlue struct {
	BackDefault  string `json:"back_default"`
	BackGray     string `json:"back_gray"`
	FrontDefault string `json:"front_default"`
	FrontGray    string `json:"front_gray"`
}
type Yellow struct {
	BackDefault  string `json:"back_default"`
	BackGray     string `json:"back_gray"`
	FrontDefault string `json:"front_default"`
	FrontGray    string `json:"front_gray"`
}
type GenerationI struct {
	RedBlue RedBlue `json:"red-blue"`
	Yellow  Yellow  `json:"yellow"`
}
type Crystal struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}
type Gold struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}
type Silver struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}
type GenerationIi struct {
	Crystal Crystal `json:"crystal"`
	Gold    Gold    `json:"gold"`
	Silver  Silver  `json:"silver"`
}
type Emerald struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}
type FireredLeafgreen struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}
type RubySapphire struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}
type GenerationIii struct {
	Emerald          Emerald          `json:"emerald"`
	FireredLeafgreen FireredLeafgreen `json:"firered-leafgreen"`
	RubySapphire     RubySapphire     `json:"ruby-sapphire"`
}
type DiamondPearl struct {
	BackDefault      string `json:"back_default"`
	BackFemale       any    `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  any    `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      any    `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale any    `json:"front_shiny_female"`
}
type HeartgoldSoulsilver struct {
	BackDefault      string `json:"back_default"`
	BackFemale       any    `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  any    `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      any    `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale any    `json:"front_shiny_female"`
}
type Platinum struct {
	BackDefault      string `json:"back_default"`
	BackFemale       any    `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  any    `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      any    `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale any    `json:"front_shiny_female"`
}
type GenerationIv struct {
	DiamondPearl        DiamondPearl        `json:"diamond-pearl"`
	HeartgoldSoulsilver HeartgoldSoulsilver `json:"heartgold-soulsilver"`
	Platinum            Platinum            `json:"platinum"`
}
type Animated struct {
	BackDefault      string `json:"back_default"`
	BackFemale       any    `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  any    `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      any    `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale any    `json:"front_shiny_female"`
}
type BlackWhite struct {
	Animated         Animated `json:"animated"`
	BackDefault      string   `json:"back_default"`
	BackFemale       any      `json:"back_female"`
	BackShiny        string   `json:"back_shiny"`
	BackShinyFemale  any      `json:"back_shiny_female"`
	FrontDefault     string   `json:"front_default"`
	FrontFemale      any      `json:"front_female"`
	FrontShiny       string   `json:"front_shiny"`
	FrontShinyFemale any      `json:"front_shiny_female"`
}
type GenerationV struct {
	BlackWhite BlackWhite `json:"black-white"`
}
type OmegarubyAlphasapphire struct {
	FrontDefault     string `json:"front_default"`
	FrontFemale      any    `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale any    `json:"front_shiny_female"`
}
type XY struct {
	FrontDefault     string `json:"front_default"`
	FrontFemale      any    `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale any    `json:"front_shiny_female"`
}
type GenerationVi struct {
	OmegarubyAlphasapphire OmegarubyAlphasapphire `json:"omegaruby-alphasapphire"`
	XY                     XY                     `json:"x-y"`
}
type Icons struct {
	FrontDefault string `json:"front_default"`
	FrontFemale  any    `json:"front_female"`
}
type UltraSunUltraMoon struct {
	FrontDefault     string `json:"front_default"`
	FrontFemale      any    `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale any    `json:"front_shiny_female"`
}
type GenerationVii struct {
	Icons             Icons             `json:"icons"`
	UltraSunUltraMoon UltraSunUltraMoon `json:"ultra-sun-ultra-moon"`
}
type GenerationViii struct {
	Icons Icons `json:"icons"`
}
type Versions struct {
	GenerationI    GenerationI    `json:"generation-i"`
	GenerationIi   GenerationIi   `json:"generation-ii"`
	GenerationIii  GenerationIii  `json:"generation-iii"`
	GenerationIv   GenerationIv   `json:"generation-iv"`
	GenerationV    GenerationV    `json:"generation-v"`
	GenerationVi   GenerationVi   `json:"generation-vi"`
	GenerationVii  GenerationVii  `json:"generation-vii"`
	GenerationViii GenerationViii `json:"generation-viii"`
}
type Sprites struct {
	BackDefault      string   `json:"back_default"`
	BackFemale       any      `json:"back_female"`
	BackShiny        string   `json:"back_shiny"`
	BackShinyFemale  any      `json:"back_shiny_female"`
	FrontDefault     string   `json:"front_default"`
	FrontFemale      any      `json:"front_female"`
	FrontShiny       string   `json:"front_shiny"`
	FrontShinyFemale any      `json:"front_shiny_female"`
	Other            Other    `json:"other"`
	Versions         Versions `json:"versions"`
}
type Cries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}
type Stats struct {
	BaseStat int      `json:"base_stat"`
	Effort   int      `json:"effort"`
	Stat     Response `json:"stat"`
}
type Types struct {
	Slot int      `json:"slot"`
	Type Response `json:"type"`
}
type PastTypes struct {
	Generation Response `json:"generation"`
	Types      []Types  `json:"types"`
}

func MapGet(c *state.State, prewious bool, cache *pokecache.Cache) ([]Response, error) {
	mapURL := c.MapNext

	if prewious {
		mapURL = c.MapPrevious
	}
	_, ok := cache.Get(mapURL)

	if !ok {
		result, err := http.Get(mapURL)

		if err != nil {
			fmt.Println("Err 01")
			return nil, err
		}
		body, err := io.ReadAll(result.Body)

		result.Body.Close()

		if result.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", result.StatusCode, body)
		}
		if err != nil {
			fmt.Println("Err 02")
			return nil, err

		}
		cache.Add(mapURL, body)
	}
	data, _ := cache.Get(mapURL)

	structuredResponse := apiResponse{}

	err := json.Unmarshal(data, &structuredResponse)
	if err != nil {
		return nil, err
	}
	c.MapCurrent = mapURL
	c.MapNext = structuredResponse.Next
	c.MapPrevious = structuredResponse.Previous
	return structuredResponse.Results, nil

}
func AreaGet(c *state.State, prewious bool, cache *pokecache.Cache, parameter string) ([]Response, error) {

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", parameter)

	_, ok := cache.Get(url)

	if !ok {
		result, err := http.Get(url)

		if err != nil {
			fmt.Println("Err 01")
			return nil, err
		}
		body, err := io.ReadAll(result.Body)

		result.Body.Close()

		if result.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", result.StatusCode, body)
		}
		if err != nil {
			fmt.Println("Err 02")
			return nil, err

		}
		cache.Add(url, body)
	}
	data, _ := cache.Get(url)

	structuredResponse := apiArea{}

	err := json.Unmarshal(data, &structuredResponse)
	if err != nil {
		return nil, err
	}

	ret := []Response{}
	for _, poke := range structuredResponse.PokemonEncounters {
		ret = append(ret, poke.Pokemon)
	}

	return ret, nil

}
func PokemonGet(c *state.State, prewious bool, cache *pokecache.Cache, parameter string) (state.PokeDataInDex, error) {

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", parameter)

	result, err := http.Get(url)

	if err != nil {
		fmt.Println("Err 01")
		return state.PokeDataInDex{}, err
	}
	body, err := io.ReadAll(result.Body)

	result.Body.Close()

	if result.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", result.StatusCode, body)
	}
	if err != nil {
		fmt.Println("Err 02")
		return state.PokeDataInDex{}, err

	}

	structuredResponse := ApiPokemon{}

	err = json.Unmarshal(body, &structuredResponse)
	if err != nil {
		return state.PokeDataInDex{}, err
	}

	var sts []state.DexStat
	for _, s := range structuredResponse.Stats {

		sts = append(sts, state.DexStat{
			Name:  s.Stat.Name,
			Value: s.BaseStat,
		})
	}
	var typs []string
	for _, n := range structuredResponse.Types {
		typs = append(typs, n.Type.Name)
	}

	ret := state.PokeDataInDex{
		Name:           structuredResponse.Name,
		BaseExperience: structuredResponse.BaseExperience,
		Height:         structuredResponse.Height,
		Weight:         structuredResponse.Weight,
		Stats:          sts,
		Types:          typs,
	}

	return ret, nil

}

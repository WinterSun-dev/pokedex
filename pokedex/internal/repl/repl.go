package repl

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"

	"github.com/WinterSun-dev/pokedex/internal/pokecache"
	"github.com/WinterSun-dev/pokedex/internal/state"

	"github.com/WinterSun-dev/pokedex/internal/pokeapi"
)

type clicommand struct {
	name        string
	description string
	callback    func(c *state.State, cmd map[string]clicommand, cache *pokecache.Cache, parameter string) ([]pokeapi.Response, error)
}

func commandHelp(c *state.State, cmd map[string]clicommand, cache *pokecache.Cache, parameter string) ([]pokeapi.Response, error) {
	for key := range cmd {
		fmt.Println(key, ":", cmd[key].description)
	}
	return nil, nil
}
func commandExit(c *state.State, cmd map[string]clicommand, cache *pokecache.Cache, parameter string) ([]pokeapi.Response, error) {
	return nil, errors.New("you have closed the pokedex")
}
func commandMap(c *state.State, cmd map[string]clicommand, cache *pokecache.Cache, parameter string) ([]pokeapi.Response, error) {
	fmt.Println(" --- Print 20 locations --- ")
	result, err := pokeapi.MapGet(c, false, cache)
	if err != nil {
		return nil, err
	}
	for _, r := range result {
		fmt.Println(r.Name)
	}

	return result, nil
}
func commandMapB(c *state.State, cmd map[string]clicommand, cache *pokecache.Cache, parameter string) ([]pokeapi.Response, error) {
	fmt.Println(" --- Print prewious 20 locations --- ")
	result, err := pokeapi.MapGet(c, true, cache)
	if err != nil {
		return nil, err
	}
	for _, r := range result {
		fmt.Println(r.Name)
	}

	return result, nil
}
func commandExplore(c *state.State, cmd map[string]clicommand, cache *pokecache.Cache, parameter string) ([]pokeapi.Response, error) {
	fmt.Println(" --- Print Pokemons in area --- ")
	result, err := pokeapi.AreaGet(c, false, cache, parameter)
	if err != nil {
		return nil, err
	}
	for _, r := range result {
		fmt.Println(r.Name)
	}

	return result, nil
}
func commandCatch(c *state.State, cmd map[string]clicommand, cache *pokecache.Cache, parameter string) ([]pokeapi.Response, error) {
	fmt.Println(" --- Print Catch sequense --- ")
	if parameter == "" {
		fmt.Println(" --- No Pokemon to catch --- ")
	}
	_, ok := state.PokeCatalog[parameter]
	if !ok {

		result, err := pokeapi.PokemonGet(c, false, cache, parameter)
		if err != nil {
			return nil, err
		}
		state.PokeCatalog[parameter] = state.PokeEntry{PokeData: result, IsCaptured: false}
	}
	requireLuck := (state.PokeCatalog[parameter].PokeData.BaseExperience / 100) + 2

	luck := rand.IntN(10)

	if requireLuck >= luck {
		fmt.Printf("You trough Ball towards %v\n", parameter)
		fmt.Printf("%v sreams in horror\n", parameter)
		fmt.Printf("%v is sedated and no longer resist\n", parameter)
		state.PokeCatalog[parameter] = state.PokeEntry{PokeData: state.PokeCatalog[parameter].PokeData, IsCaptured: true}
	} else {

		fmt.Printf("...Miss ... get gud %v requires %v\n", luck, requireLuck)
	}

	return nil, nil
}
func commandInspect(c *state.State, cmd map[string]clicommand, cache *pokecache.Cache, parameter string) ([]pokeapi.Response, error) {
	fmt.Println(" --- Print Pokemon information --- ")
	if parameter == "" {
		fmt.Println(" --- Select pokemon to inspect--- ")
	}
	v, ok := state.PokeCatalog[parameter]
	if !ok || !v.IsCaptured {

		fmt.Println("Capture pokemon first")
		return nil, nil
	}

	fmt.Printf("Name  : %v\n", v.PokeData.Name)
	fmt.Printf("Height: %v\n", v.PokeData.Height)
	fmt.Printf("Weight: %v\n", v.PokeData.Weight)
	fmt.Println("Stats :")

	for _, s := range v.PokeData.Stats {
		fmt.Printf("   - %v: %v\n", s.Name, s.Value)
	}
	fmt.Println("Types :")
	for _, t := range v.PokeData.Types {
		fmt.Printf("   - %v\n", t)
	}

	return nil, nil
}
func commandPokedex(c *state.State, cmd map[string]clicommand, cache *pokecache.Cache, parameter string) ([]pokeapi.Response, error) {
	fmt.Println(" --- Print captured Pokemons --- ")
	if len(state.PokeCatalog) == 0 {
		fmt.Println("Get some pokemons to oogle first")
		return nil, nil
	}
	for _, p := range state.PokeCatalog {

		fmt.Printf(" - %v\n", p.PokeData.Name)

	}

	return nil, nil
}

var comms = map[string]clicommand{
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"map": {
		name:        "map",
		description: "Show next 20 locations",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Show previous 20 locations",
		callback:    commandMapB,
	},
	"explore": {
		name:        "explore",
		description: "Show list of pokemons in given area: explore location",
		callback:    commandExplore,
	},
	"catch": {
		name:        "catch",
		description: "Try to catch pokemon",
		callback:    commandCatch,
	},
	"inspect": {
		name:        "inspect",
		description: "Show iformation of pokemons you have cathed",
		callback:    commandInspect,
	},
	"pokedex": {
		name:        "pokedex",
		description: "Show all pokemos captured",
		callback:    commandPokedex,
	},
}

func Repl(locka *state.State, cache *pokecache.Cache) {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("==> ")

	for scanner.Scan() {

		input := strings.Fields(scanner.Text())
		v, ok := comms[input[0]]
		if !ok {
			fmt.Println("Commad not found")
			fmt.Printf("--> ")
			continue
		}
		parameter := ""
		if len(input) > 1 {
			parameter = input[1]
		}

		_, err := v.callback(locka, comms, cache, parameter)
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Printf("--> ")
	}

	if err := scanner.Err(); err != nil {

		fmt.Fprintln(os.Stderr, "reading standard input:", err)

	}

}

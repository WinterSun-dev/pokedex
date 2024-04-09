package main

import (
	"fmt"
	"time"

	"github.com/WinterSun-dev/pokedex/internal/pokecache"
	"github.com/WinterSun-dev/pokedex/internal/repl"
	"github.com/WinterSun-dev/pokedex/internal/state"
)

func main() {

	fmt.Println("Poke GO Go >")
	state := state.NewState()
	//config := pokeapi.Newconfig()
	cache := pokecache.NewCache(5 * time.Second)

	repl.Repl(&state, cache)

}

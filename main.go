package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/Tavis7/bootdev-pokedexcli/internal/pokecache"
	"time"
	"math/rand"
)

const API_URL = "https://pokeapi.co/api/v2/"

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	next_url string
	prev_url string
	cache pokecache.Cache
	pokedex map[string]poke_Pokemon
}

var commands map[string]cliCommand

func main() {
	pokecache.Test()
	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name: 	 	 "explore",
			description: "List all the pokemon in a location",
			callback:    commandExplore,
		},
		"catch": {
			name: 	 	 "catch",
			description: "Try to catch a pokemon",
			callback:    commandCatch,
		},
		"pokedex": {
			name: 	 	 "pokedex",
			description: "List pokemon in pokedex",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	conf := config{
		prev_url: "",
		next_url: API_URL + "location-area/?offset=0&limit=20",
		cache: pokecache.NewCache(time.Minute * 5),
		pokedex: make(map[string]poke_Pokemon),
	}
	for {
		fmt.Print("Pokedex > ")
		valid := scanner.Scan()
		if valid {
			input := scanner.Text()
			command_line := cleanInput(input)
			if len(command_line) > 0 {
				command, ok := commands[command_line[0]]
				if ok {
					err := command.callback(&conf, command_line)
					if err != nil {
						fmt.Printf("Error: %v\n", err)
					}
				} else {
					fmt.Printf("Error: command not found: %v\n", command_line[0])
				}
			}
		} else {
			break
		}
	}
	fmt.Println()
}

// https://pokeapi.co/api/v2/location-area/{id or name}/

func commandExit(conf *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Failed to exit program")
}

func commandHelp(conf *config, args []string) error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("\n")
	fmt.Printf("Usage:\n")
	for _, command := range commands {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}

func fetchAndDecodeJson(url string, v any, cache pokecache.Cache) error {
	body,ok := cache.Get(url)
	if !ok {
		//fmt.Printf("Fetching %v from network\n", url)
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			return fmt.Errorf("Request returned status code %v (expected 200)",
				res.StatusCode)
		}

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		// todo: check Cache-Control header
		// todo: cache non-200 responses
		// https://developer.mozilla.org/en-US/docs/Glossary/Cacheable
		cache.Add(url, body)
	} else {
		fmt.Printf("Got %v from cache\n", url)
	}

	err := json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	return nil
}

func commandMap(conf *config, args []string) error {
	if conf.next_url == "" {
		fmt.Println("You're on the last page")
		return nil
	}

	resourceList := poke_NamedAPIResourceList{}

	err := fetchAndDecodeJson(conf.next_url, &resourceList, conf.cache)
	if err != nil {
		return fmt.Errorf("command failed: %w", err)
	}

	conf.next_url = resourceList.Next
	conf.prev_url = resourceList.Previous

	for _, resource := range resourceList.Results {
		fmt.Printf("%s\n", resource.Name)
	}

	return nil
}

func commandMapb(conf *config, args []string) error {
	if conf.prev_url == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	resourceList := poke_NamedAPIResourceList{}

	err := fetchAndDecodeJson(conf.prev_url, &resourceList, conf.cache)
	if err != nil {
		return fmt.Errorf("command failed: %w", err)
	}

	conf.next_url = resourceList.Next
	conf.prev_url = resourceList.Previous

	for _, resource := range resourceList.Results {
		fmt.Printf("%s\n", resource.Name)
	}

	return nil
}

func commandExplore(conf *config, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("Incorrect number of arguments, expecting '%v <area>'", args[0])
	}

	area_name := args[1]

	fmt.Printf("Exploring %v...\n", area_name)

	area := poke_LocationArea{}

	err := fetchAndDecodeJson(API_URL + "location-area/" + area_name, &area, conf.cache)
	if err != nil {
		return fmt.Errorf("Failed to load area '%v': %w", area_name, err)
	}

	if len(area.Pokemon_encounters) > 0 {
		fmt.Printf("Found Pokemon:\n")
		for _, encounter := range area.Pokemon_encounters {
			fmt.Printf(" - %v\n", encounter.Pokemon.Name)
		}
	} else {
		fmt.Printf("Didn't find any Pokemon in %v\n", area_name)
	}

	return nil
}

func commandCatch(conf *config, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("Incorrect number of arguments, expecting '%v <pokemon>'", args[0])
	}

	pokemon_name := args[1]

	pokemon_info := poke_Pokemon{}

	err := fetchAndDecodeJson(API_URL + "pokemon/" + pokemon_name, &pokemon_info, conf.cache)
	if err != nil {
		return fmt.Errorf("Couldn't find %v: %w", pokemon_name, err)
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon_info.Name)

	chance := (20.0 / float32(pokemon_info.Base_experience))
	/*
	fmt.Printf("Base experience: %v\n", pokemon_info.Base_experience)
	fmt.Printf("Odds: %v%%\n", chance * 100)
	*/

	if (rand.Float32() > chance) {
		fmt.Printf("%v escaped!\n", pokemon_info.Name)
	} else {
		fmt.Printf("%v was caught!\n", pokemon_info.Name)
		conf.pokedex[pokemon_info.Name] = pokemon_info
	}

	return nil
}

func printPokedexInfo(pokemon poke_Pokemon) {
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf(" - ID: %v\n", pokemon.Id)
	fmt.Printf(" - Height: %v\n", pokemon.Height)
	fmt.Printf(" - Weight: %v\n", pokemon.Weight)
}

func commandPokedex(conf *config, args []string) error {
	if len(args) == 2 {
		pokemon := args[1]
		info,ok := conf.pokedex[pokemon]
		if !ok {
			return fmt.Errorf("%v not found in pokedex", pokemon)
		}
		printPokedexInfo(info)
	} else if len(args) == 1 {
		for _,info := range conf.pokedex {
			printPokedexInfo(info)
		}
	} else {
		return fmt.Errorf(
			"Incorrect number of arguments, expecting '%v [<pokemon>]'",
			args[0])
	}

	return nil
}

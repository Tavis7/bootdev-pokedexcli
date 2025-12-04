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
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next_url string
	prev_url string
	cache pokecache.Cache
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
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	conf := config{
		prev_url: "",
		next_url: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		cache: pokecache.NewCache(time.Second * 5),
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
					err := command.callback(&conf)
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

func commandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Failed to exit program")
}

func commandHelp(conf *config) error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("\n")
	fmt.Printf("Usage:\n")
	for _, command := range commands {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}

type poke_NamedAPIResource struct {
	Name string `json:name`
	Url  string `json:url`
}

type poke_NamedAPIResourceList struct {
	Count    int                     `json:const`
	Previous string                  `json:previous`
	Next     string                  `json:next`
	Results  []poke_NamedAPIResource `json:results`
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
			return err
		}

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		cache.Add(url, body)
	} else {
		//fmt.Printf("Got %v from cache\n", url)
	}

	err := json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	return nil
}

func commandMap(conf *config) error {
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

func commandMapb(conf *config) error {
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

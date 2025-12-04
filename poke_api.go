package main

type poke_NamedAPIResource struct {
	Name string `json:name`
	Url  string `json:url`
}

type poke_NamedAPIResourceList struct {
	Count    int                     `json:count`
	Previous string                  `json:previous`
	Next     string                  `json:next`
	Results  []poke_NamedAPIResource `json:results`
}

type poke_LocationArea struct {
	Id                     int                        `json:id`
	Name                   string                     `json:name`
	Game_index             int                        `json:game_index`
	Encounter_method_rates []poke_EncounterMethodRate `json:encounter_method_rates`
	Location               poke_NamedAPIResource      `json:location`
	Names                  []poke_Name                `json:names`
	Pokemon_encounters     []poke_PokemonEncounter    `json:pokemon_encounters`
}

type poke_EncounterMethodRate struct {
	Encounter_method poke_NamedAPIResource        `json:encounter_method`
	Version_details  []poke_EncounterVersionDetails `json:version_details`
}

type poke_Name struct {
	Name     string                `json:name`
	Language poke_NamedAPIResource `json:language`
}

type poke_PokemonEncounter struct {
	Pokemon         poke_NamedAPIResource         `json:pokemon`
	Version_details []poke_VersionEncounterDetail `json:version_details`
}

type poke_EncounterVersionDetails struct {
	Rate    int                   `json:rate`
	Version poke_NamedAPIResource `json:version`
}

type poke_VersionEncounterDetail struct {
	Version           poke_NamedAPIResource `json:version`
	Max_chance        int                   `json:max_chance`
	Encounter_details []poke_Encounter      `json:encounter_details`
}

type poke_Encounter struct {
	Min_level        int                     `json:min_level`
	Max_level        int                     `json:max_level`
	Condition_values []poke_NamedAPIResource `json:condition_values`
	Chance           int                     `json:chance`
	Method           poke_NamedAPIResource   `json:method`
}


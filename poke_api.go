package main

type poke_Pokemon struct {
	Id                       int                       `json:id`
	Name                     string                    `json:name`
	Base_experience          int                       `json:base_experience`
	Height                   int                       `json:height` //decimeters
	Is_default               bool                      `json:is_default`
	Order                    int                       `json:order`
	Weight                   int                       `json:weight` //hectograms
	Abilities                []poke_PokemonAbility     `json:abilities`
	Forms                    []poke_NamedAPIResource   `json:forms`
	Game_indices             []poke_VersionGameIndex   `json:game_indices`
	Held_items               []poke_PokemonHeldItem    `json:held_items`
	Location_area_encounters string                    `json:location_area_encounters`
	Moves                    []poke_PokemonMove        `json:moves`
	Past_types               []poke_PokemonTypePast    `json:past_types`
	Past_abilities           []poke_PokemonAbilityPast `json:past_abilities`
	Sprites                  poke_PokemonSprites       `json:sprites`
	Cries                    poke_PokemonCries         `json:cries`
	Species                  poke_NamedAPIResource     `json:species`
	Stats                    []poke_PokemonStat        `json:stats`
	Types                    []poke_PokemonType        `json:types`
}

type poke_PokemonAbility struct {
	Is_hidden bool                  `json:is_hidden`
	Slot      int                   `json:slot`
	Ability   poke_NamedAPIResource `json:ability`
}

type poke_VersionGameIndex struct {
	Game_index int                   `json:game_index`
	Version    poke_NamedAPIResource `json:version`
}

type poke_PokemonHeldItem struct {
	Item            poke_NamedAPIResource         `json:item`
	Version_details []poke_PokemonHeldItemVersion `json:version_details`
}

type poke_PokemonMove struct {
	Move                  poke_NamedAPIResource   `json:move`
	Version_group_details []poke_PokemonMoveVersion `json:version_group_details`
}

type poke_PokemonTypePast struct {
	Generation poke_NamedAPIResource `json:generation`
	Types      []poke_PokemonType    `json:types`
}

type poke_PokemonAbilityPast struct {
	Generation poke_NamedAPIResource `json:generation`
	Abilities  []poke_PokemonAbility `json:abilities`
}

type poke_PokemonSprites struct {
	Front_default      string `json:front_default`
	Front_shiny        string `json:front_shiny`
	Front_female       string `json:front_female`
	Front_shiny_female string `json:front_shiny_female`

	Back_default      string `json:back_default`
	Back_shiny        string `json:back_shiny`
	Back_female       string `json:back_female`
	Back_shiny_female string `json:back_shiny_female`
}

type poke_PokemonCries struct {
	Latest string `json:latest`
	Legacy string `json:legacy`
}

type poke_PokemonStat struct {
	Stat      poke_NamedAPIResource `json:stat`
	Effort    int                   `json:effort`
	Base_stat int                   `json:base_stat`
}

type poke_PokemonType struct {
	Slot int                   `json:slot`
	Type poke_NamedAPIResource `json:type`
}

type poke_PokemonHeldItemVersion struct {
	Version poke_NamedAPIResource `json:version`
	Rarity  int                   `json:rarity`
}

type poke_PokemonMoveVersion struct {
	Move_learn_method poke_NamedAPIResource `json:move_learn_method`
	Version_group     poke_NamedAPIResource `json:version_group`
	Level_learned_at  int                   `json:level_learned_at`
	Order             int                   `json:order`
}

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


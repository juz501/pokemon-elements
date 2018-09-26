package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/normal", normal)
	http.HandleFunc("/fire", fire)
	http.HandleFunc("/water", water)
	http.HandleFunc("/electric", electric)
	http.HandleFunc("/grass", grass)
	http.HandleFunc("/ice", ice)
	http.HandleFunc("/fighting", fighting)
	http.HandleFunc("/poison", poison)
	http.HandleFunc("/ground", ground)
	http.HandleFunc("/flying", flying)
	http.HandleFunc("/psychic", psychic)
	http.HandleFunc("/bug", bug)
	http.HandleFunc("/rock", rock)
	http.HandleFunc("/ghost", ghost)
	http.HandleFunc("/dragon", dragon)
	http.HandleFunc("/dark", dark)
	http.HandleFunc("/steel", steel)
	http.HandleFunc("/fairy", fairy)
	port := ":" + os.Getenv("GO_SERVER_PORT")
	if port == ":" {
		port = ":18883"
	}
	addr := os.Getenv("GO_SERVER_ADDR")
	log.Fatal(http.ListenAndServe(addr+port, nil))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Pokémon Element Checker")
}

func createResult(attacks string, effectiveness string, pokemon string) string {
	if effectiveness == "yes" {
		return attacks + " attacks are super effective against " + pokemon + " Pokémon"
	} else if effectiveness == "no" {
		return attacks + " attacks are not effective against " + pokemon + " Pokémon"
	} else {
		return attacks + " attacks have no effect on " + pokemon + " Pokémon"
	}
}

func normal(w http.ResponseWriter, r *http.Request) {
	element := "Normal"
	fmt.Fprintln(w, createResult(element, "no", "Rock and Steel"))
	fmt.Fprintln(w, createResult(element, "", "Ghost"))
	fmt.Fprintln(w, createResult("Ghost", "", element))
	fmt.Fprintln(w, createResult("Fighting", "yes", element))
}

func fire(w http.ResponseWriter, r *http.Request) {
	element := "Fire"
	fmt.Fprintln(w, createResult(element, "yes", "Grass, Ice, Bug and Steel"))
	fmt.Fprintln(w, createResult(element, "no", "Fire, Water, Rock and Dragon"))
	fmt.Fprintln(w, createResult("Fire, Grass, Ice, Bug, Steel and Fairy", "no", element))
	fmt.Fprintln(w, createResult("Water, Ground and Rock", "yes", element))
}

func water(w http.ResponseWriter, r *http.Request) {
	element := "Water"
	fmt.Fprintln(w, createResult(element, "yes", "Fire, Ground and Rock"))
	fmt.Fprintln(w, createResult(element, "no", "Water, Grass and Dragon"))
	fmt.Fprintln(w, createResult("Fire, Water, Ice and Steel", "no", element))
	fmt.Fprintln(w, createResult("Electric and Grass", "yes", element))
}

func electric(w http.ResponseWriter, r *http.Request) {
	element := "Electric"
	fmt.Fprintln(w, createResult(element, "yes", "Water and Flying"))
	fmt.Fprintln(w, createResult(element, "no", "Electric, Grass and Dragon"))
	fmt.Fprintln(w, createResult(element, "", "Ground"))
	fmt.Fprintln(w, createResult("Electric, Flying and Steel", "no", element))
	fmt.Fprintln(w, createResult("Ground", "yes", element))
}

func grass(w http.ResponseWriter, r *http.Request) {
	element := "Grass"
	fmt.Fprintln(w, createResult(element, "yes", "Water, Ground and Rock"))
	fmt.Fprintln(w, createResult(element, "no", "Fire, Grass, Poison, Flying, Bug, Dragon and Steel"))
	fmt.Fprintln(w, createResult("Water, Electric, Grass and Ground", "no", element))
	fmt.Fprintln(w, createResult("Fire, Ice, Poison, Flying and Bug", "yes", element))
}

func ice(w http.ResponseWriter, r *http.Request) {
	element := "Ice"
	fmt.Fprintln(w, createResult(element, "yes", "Grass, Ground, Flying and Dragon"))
	fmt.Fprintln(w, createResult(element, "no", "Fire, Water, Ice and Steel"))
	fmt.Fprintln(w, createResult("Ice", "no", element))
	fmt.Fprintln(w, createResult("Fire, Fighting, Rock and Steel", "yes", element))
}

func fighting(w http.ResponseWriter, r *http.Request) {
	element := "Fighting"
	fmt.Fprintln(w, createResult(element, "yes", "Normal, Ice, Rock, Dark and Steel"))
	fmt.Fprintln(w, createResult(element, "no", "Poison, Flying, Psychic, Bug and Fairy"))
	fmt.Fprintln(w, createResult(element, "", "Ghost"))
	fmt.Fprintln(w, createResult("Bug, Rock and Dark", "no", element))
	fmt.Fprintln(w, createResult("Flying, Psychic and Fairy", "yes", element))
}

func poison(w http.ResponseWriter, r *http.Request) {
	element := "Poison"
	fmt.Fprintln(w, createResult(element, "yes", "Grass and Fairy"))
	fmt.Fprintln(w, createResult(element, "no", "Poison, Ground, Rock and Ghost"))
	fmt.Fprintln(w, createResult(element, "", "Steel"))
	fmt.Fprintln(w, createResult("Grass, Fighting, Poison, Bug and Fairy", "no", element))
	fmt.Fprintln(w, createResult("Ground and Psychic", "yes", element))
}

func ground(w http.ResponseWriter, r *http.Request) {
	element := "Ground"
	fmt.Fprintln(w, createResult(element, "yes", "Fire, Electric, Poison, Rock and Steel"))
	fmt.Fprintln(w, createResult(element, "no", "Grass and Bug"))
	fmt.Fprintln(w, createResult(element, "", "Flying"))
	fmt.Fprintln(w, createResult("Electric", "", element))
	fmt.Fprintln(w, createResult("Poison and Rock", "no", element))
	fmt.Fprintln(w, createResult("Water, Grass and Ice", "yes", element))
}

func flying(w http.ResponseWriter, r *http.Request) {
	element := "Flying"
	fmt.Fprintln(w, createResult(element, "yes", "Grass, Fighting and Bug"))
	fmt.Fprintln(w, createResult(element, "no", "Electric, Rock and Steel"))
	fmt.Fprintln(w, createResult("Ground", "", element))
	fmt.Fprintln(w, createResult("Grass, Fighting and Bug", "no", element))
	fmt.Fprintln(w, createResult("Electric, Ice and Rock", "yes", element))
}

func psychic(w http.ResponseWriter, r *http.Request) {
	element := "Psychic"
	fmt.Fprintln(w, createResult(element, "yes", "Fighting and Poison"))
	fmt.Fprintln(w, createResult(element, "no", "Psychic and Steel"))
	fmt.Fprintln(w, createResult(element, "", "Dark"))
	fmt.Fprintln(w, createResult("Fighting and Psychic", "no", element))
	fmt.Fprintln(w, createResult("Bug, Ghost and Dark", "yes", element))
}

func bug(w http.ResponseWriter, r *http.Request) {
	element := "Bug"
	fmt.Fprintln(w, createResult(element, "yes", "Grass, Psychic and Dark"))
	fmt.Fprintln(w, createResult(element, "no", "Fire, Fighting, Poison, Flying, Ghost, Steel and Fairy"))
	fmt.Fprintln(w, createResult("Grass, Fighting and Ground", "no", element))
	fmt.Fprintln(w, createResult("Fire, Flying and Rock", "yes", element))
}

func rock(w http.ResponseWriter, r *http.Request) {
	element := "Rock"
	fmt.Fprintln(w, createResult(element, "yes", "Fire, Ice, Flying and Bug"))
	fmt.Fprintln(w, createResult(element, "no", "Fighting, Ground and Steel"))
	fmt.Fprintln(w, createResult("Normal, Fire, Poison and Flying", "no", element))
	fmt.Fprintln(w, createResult("Water, Grass, Fighting, Ground and Steel", "yes", element))
}

func ghost(w http.ResponseWriter, r *http.Request) {
	element := "Ghost"
	fmt.Fprintln(w, createResult(element, "yes", "Psychic and Ghost"))
	fmt.Fprintln(w, createResult(element, "no", "Dark"))
	fmt.Fprintln(w, createResult(element, "", "Normal"))
	fmt.Fprintln(w, createResult("Normal and Fighting", "", element))
	fmt.Fprintln(w, createResult("Poison and Bug", "no", element))
	fmt.Fprintln(w, createResult("Ghost and Dark", "yes", element))
}

func dragon(w http.ResponseWriter, r *http.Request) {
	element := "Dragon"
	fmt.Fprintln(w, createResult(element, "yes", "Dragon"))
	fmt.Fprintln(w, createResult(element, "no", "Steel"))
	fmt.Fprintln(w, createResult(element, "", "Fairy"))
	fmt.Fprintln(w, createResult("Fire, Water, Electric and Grass", "no", element))
	fmt.Fprintln(w, createResult("Ice, Dragon and Fairy", "yes", element))
}

func dark(w http.ResponseWriter, r *http.Request) {
	element := "Dark"
	fmt.Fprintln(w, createResult(element, "yes", "Psychic and Ghost"))
	fmt.Fprintln(w, createResult(element, "no", "Fighting, Dark and Fairy"))
	fmt.Fprintln(w, createResult("Psychic", "", element))
	fmt.Fprintln(w, createResult("Ghost and Dark", "no", element))
	fmt.Fprintln(w, createResult("Fighting, Bug and Fairy", "yes", element))
}

func steel(w http.ResponseWriter, r *http.Request) {
	element := "Steel"
	fmt.Fprintln(w, createResult(element, "yes", "Ice, Rock and Fairy"))
	fmt.Fprintln(w, createResult(element, "no", "Fire, water, Electric and Steel"))
	fmt.Fprintln(w, createResult("Poison", "", element))
	fmt.Fprintln(w, createResult("Normal, Grass, Ice, Flying, Psychic, Bug, Rock, Dragon, Steel, Fairy", "no", element))
	fmt.Fprintln(w, createResult("Fire, Fighting and Ground", "yes", element))
}

func fairy(w http.ResponseWriter, r *http.Request) {
	element := "Fairy"
	fmt.Fprintln(w, createResult(element, "yes", "Fighting, Dragon and Steel"))
	fmt.Fprintln(w, createResult(element, "no", "Fire, Poison and Steel"))
	fmt.Fprintln(w, createResult("Dragon", "", element))
	fmt.Fprintln(w, createResult("Fighting, Bug and Dark", "no", element))
	fmt.Fprintln(w, createResult("Poison and Steel", "yes", element))
}

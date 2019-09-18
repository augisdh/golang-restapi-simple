package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Game Struct (Model | Class)
type Game struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Category  string     `json:"category"`
	Developer *Developer `json:"developer"`
}

// Developer Struct (Model | Class)
type Developer struct {
	Company string `json:"company"`
	Year    string `json:"year"`
}

// Games as a slice Game struct
var games []Game

// Get All Games
func getGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

// Get Game
func getGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Loop through games and find id
	for _, item := range games {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Game{})
}

// Get Create Game
func createGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var game Game
	_ = json.NewDecoder(r.Body).Decode(&game)

	// Mock ID
	game.ID = strconv.Itoa(rand.Intn(10000000))
	games = append(games, game)
	json.NewEncoder(w).Encode(game)
}

// Get Update Game
func updateGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range games {
		if item.ID == params["id"] {
			games = append(games[:index], games[index+1:]...)
			var game Game
			_ = json.NewDecoder(r.Body).Decode(&game)
			game.ID = params["id"]
			games = append(games, game)
			json.NewEncoder(w).Encode(game)
			return
		}
	}
	json.NewEncoder(w).Encode(games)
}

// Get Delete Game
func deleteGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range games {
		if item.ID == params["id"] {
			games = append(games[:index], games[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(games)
}

func main() {
	// Router init
	route := mux.NewRouter()

	// Mock Data
	games = append(games, Game{ID: "1", Title: "Need for speed", Category: "racing", Developer: &Developer{Company: "Ghost Games", Year: "2013-present"}})
	games = append(games, Game{ID: "2", Title: "Game number 2", Category: "adventure", Developer: &Developer{Company: "Shit Games", Year: "2017-present"}})
	games = append(games, Game{ID: "3", Title: "Game number 3", Category: "racing", Developer: &Developer{Company: "Rice Games", Year: "2018"}})
	games = append(games, Game{ID: "4", Title: "Game number 4", Category: "racing", Developer: &Developer{Company: "Poop Games", Year: "2019-present"}})
	games = append(games, Game{ID: "5", Title: "Game number 5", Category: "something", Developer: &Developer{Company: "Eat Games", Year: "2009"}})
	games = append(games, Game{ID: "6", Title: "Game number 6", Category: "else", Developer: &Developer{Company: "Cream Games", Year: "2016"}})
	games = append(games, Game{ID: "7", Title: "Game number 7", Category: "racing", Developer: &Developer{Company: "Wall Games", Year: "2017"}})
	games = append(games, Game{ID: "8", Title: "Game number 8", Category: "other", Developer: &Developer{Company: "Random Games", Year: "2011-present"}})

	// Route endpoints
	route.HandleFunc("/api/games", getGames).Methods("GET")
	route.HandleFunc("/api/games/{id}", getGame).Methods("GET")
	route.HandleFunc("/api/games", createGame).Methods("POST")
	route.HandleFunc("/api/games/{id}", updateGame).Methods("PUT")
	route.HandleFunc("/api/games/{id}", deleteGame).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", route))
}

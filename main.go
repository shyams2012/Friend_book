package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shyams2012/friend_book/seed"
	"github.com/shyams2012/friend_book/user"
)

func init() {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error getting env file. Error :", err)
	}

	// Migrate to DB
	err = seed.Migrate()
	if err == nil {
		// Seed users
		err := seed.SeedUsers()
		if err != nil {
			fmt.Println("Could not seed users")
		}
	}
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/createProfile", user.CreateProfile).Methods("POST")
	r.HandleFunc("/searchFriend/{id}", user.SearchFriends).Methods("GET")
	r.HandleFunc("/addFriend", user.AddFriends).Methods("POST")
	r.HandleFunc("/postMessage", user.PostMessage).Methods("POST")
	r.HandleFunc("/displayPost/{userId}", user.DisplayPosts).Methods("GET")
	r.HandleFunc("/likePost", user.LikePosts).Methods("POST")
	r.HandleFunc("/sharePosts", user.SharePosts).Methods("POST")

	fmt.Printf("Starting server at port 8000\n")
	http.ListenAndServe(":8000", r)

}

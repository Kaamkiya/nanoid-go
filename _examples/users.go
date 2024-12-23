package main

import (
	"net/http"
	"log"
	"encoding/json"

	"github.com/Kaamkiya/nanoid-go"
)

type User struct {
	Name string
	ID string
}

var users []User

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var u User
		switch r.Method {
		case "POST":
			err := json.NewDecoder(r.Body).Decode(&u)
			if err != nil {
				log.Printf("Failed to unmarshal request: %v\n", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			u.ID = nanoid.Nanoid(32, nanoid.DefaultCharset)
			err = json.NewEncoder(w).Encode(u)
			users = append(users, u)
			
			log.Printf("Added user %v\n", u)
		case "GET":
			if err := json.NewEncoder(w).Encode(users); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			
			log.Printf("Got all users (len: %d)\n", len(users))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	log.Println("Starting server on :4000")
	http.ListenAndServe(":4000", nil)
}

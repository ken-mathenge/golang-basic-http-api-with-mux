package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// User is a struct that represents a single user
type User struct {
	FullName string `json:"fullname"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

// Post is a struct that represents a single post
// TODO Brush up on exported fields they start with capital leters
type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author User   `json:"author"`
}

var posts []Post = []Post{}

// AddPostsHandlerFunc adds items to a slice of strings
func AddPostsHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// get the item value from the json body
	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost)

	posts = append(posts, newPost)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(posts)
}

// GetPostsHandlerFunc returns all the posts
func GetPostsHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(posts)
}

// GetPostHandlerFunc returns a single the posts
func GetPostHandlerFunc(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	// Convert the id from string to interger
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("unable to convert id to integer"))
		return
	}

	if id >= len(posts) {
		w.WriteHeader(400)
		w.Write([]byte("No post found!"))
		return
	}

	post := posts[id] // Indexed the slice..Makes sense

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// UpdatePostHandlerFunc updates a single post
func UpdatePostHandlerFunc(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	// Convert the id from string to interger
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("unable to convert id to integer"))
		return
	}

	if id >= len(posts) {
		w.WriteHeader(400)
		w.Write([]byte("No post found!"))
		return
	}

	updatedPost := Post{}
	json.NewDecoder(r.Body).Decode(&updatedPost)

	posts[id] = updatedPost

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedPost)
}

// PatchPostHandlerFunc patches a single post
func PatchPostHandlerFunc(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	// Convert the id from string to interger
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("unable to convert id to integer"))
		return
	}

	if id >= len(posts) {
		w.WriteHeader(400)
		w.Write([]byte("No post found!"))
		return
	}

	post := &posts[id]
	json.NewDecoder(r.Body).Decode(&post)

	// posts[id] = post // Workaround here

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// DeletePostHandlerFunc patches a single post
func DeletePostHandlerFunc(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	// Convert the id from string to interger
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("unable to convert id to integer"))
		return
	}

	if id >= len(posts) {
		w.WriteHeader(400)
		w.Write([]byte("No post found!"))
		return
	}

	// Go workaround to delete from slices
	posts = append(posts[:id], posts[id+1:]...)

	w.WriteHeader(200)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/posts", AddPostsHandlerFunc).Methods("POST")
	r.HandleFunc("/posts", GetPostsHandlerFunc).Methods("GET")
	r.HandleFunc("/posts/{id}", GetPostHandlerFunc).Methods("GET")
	r.HandleFunc("/posts/{id}", UpdatePostHandlerFunc).Methods("PUT")
	r.HandleFunc("/posts/{id}", PatchPostHandlerFunc).Methods("PATCH")
	r.HandleFunc("/posts/{id}", DeletePostHandlerFunc).Methods("DELETE")

	http.ListenAndServe(":5000", r)
}

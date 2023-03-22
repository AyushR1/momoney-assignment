package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Post represents the data structure of a post
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// Todo represents the data structure of a todo
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// CacheEntry represents a cached response
type CacheEntry struct {
	Response interface{}
	Expire   time.Time
}

// Cache represents the cache for responses
type Cache map[string]CacheEntry

// APIHandler represents the HTTP handler for APIs
type APIHandler struct {
	Cache      Cache
	Expiration time.Duration
}

// HandleRequest handles HTTP requests for APIs
func (h *APIHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	// Get the API endpoint and ID from the URL
	endpoint := r.URL.Path[1:]
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Check if the response is cached
	key := fmt.Sprintf("%s/%d", endpoint, id)
	cacheEntry, ok := h.Cache[key]
	if ok && cacheEntry.Expire.After(time.Now()) {
		// Serve the response from cache
		log.Println("Data fetched from cache")
		responseJSON, err := json.Marshal(cacheEntry.Response)
		if err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJSON)
		return
	}

	// Fetch the data from the API
	var response interface{}
	switch endpoint {
	case "posts":
		response, err = getPost(id)
	case "todos":
		response, err = getTodo(id)
	default:
		http.Error(w, "Invalid API endpoint", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Cache the response
	h.Cache[key] = CacheEntry{
		Response: response,
		Expire:   time.Now().Add(h.Expiration),
	}

	// Write the response to the client
	log.Println("Data fetched from API endpoint")
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

// getPost fetches a post by ID from the API
func getPost(id int) (*Post, error) {
	resp, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var post Post
	err = json.NewDecoder(resp.Body).Decode(&post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

// getTodo fetches a todo by ID from the API
func getTodo(id int) (*Todo, error) {
	resp, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var todo Todo
	err = json.NewDecoder(resp.Body).Decode(&todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

// NewAPIHandler returns a new instance of APIHandler
func NewAPIHandler() *APIHandler {
	return &APIHandler{
		Cache:      make(Cache),
		Expiration: 5 * time.Minute,
	}
}

func main() {
	apiHandler := NewAPIHandler()

	// Create a new HTTP router
	router := http.NewServeMux()

	// Register the API endpoints
	router.HandleFunc("/posts", apiHandler.HandleRequest)
	router.HandleFunc("/todos", apiHandler.HandleRequest)

	// Start the server
	// Start the server
	addr := ":8080"
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server is running at port %s", addr)
	})
	log.Printf("Server running on http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}

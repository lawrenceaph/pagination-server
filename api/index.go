package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type PaginationResult struct {
	Objects  []Object `json:"objects"`
	NextPage int      `json:"nextPage"`
}

type Object struct {
	Name    string `json:"name"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

const loremIpsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."

func generateLoremIpsum(words int) string {
	wordList := strings.Fields(loremIpsum)
	result := make([]string, words)
	for i := range result {
		result[i] = wordList[rand.Intn(len(wordList))]
	}
	return strings.Join(result, " ")
}

func simulatePagination(page int, objectsPerPage int, longContent bool) PaginationResult {
	const totalObjects = 10000
	const maxObjectsPerRequest = 1000

	if objectsPerPage > maxObjectsPerRequest {
		objectsPerPage = maxObjectsPerRequest
	}

	start := (page - 1) * objectsPerPage
	end := start + objectsPerPage

	if start >= totalObjects {
		return PaginationResult{Objects: nil, NextPage: 0}
	}

	if end > totalObjects {
		end = totalObjects
	}

	var objects []Object
	rand.Seed(time.Now().UnixNano())

	for i := start; i < end; i++ {
		content := fmt.Sprintf("random content %d", i+1)
		if longContent {
			content = generateLoremIpsum(50) // Generate 50 words of Lorem Ipsum text
		}

		objects = append(objects, Object{
			Name:    fmt.Sprintf("random name %d", i+1),
			Content: content,
			Image:   "https://placehold.co/600x400",
		})
	}

	nextPage := 0
	if end < totalObjects {
		nextPage = page + 1
	}

	return PaginationResult{Objects: objects, NextPage: nextPage}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Default parameters
	currentPage := 1
	objectsPerPage := 10
	longContent := false

	// Parse query parameters
	query := r.URL.Query()
	if page, err := strconv.Atoi(query.Get("page")); err == nil && page > 0 {
		currentPage = page
	}
	if perPage, err := strconv.Atoi(query.Get("perPage")); err == nil && perPage > 0 {
		objectsPerPage = perPage
	}
	if query.Get("longContent") == "true" {
		longContent = true
	}

	// Fetch paginated data
	data := simulatePagination(currentPage, objectsPerPage, longContent)

	// Convert data to JSON
	responseData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}

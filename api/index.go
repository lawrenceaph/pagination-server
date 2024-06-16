package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type PaginationResult struct {
	Objects  []Object `json:"objects"`
	NextPage int      `json:"nextPage"`
}

type Object struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func simulatePagination(page int, objectsPerPage int) PaginationResult {
	const totalObjects = 500
	totalPages := (totalObjects + objectsPerPage - 1) / objectsPerPage

	if page > totalPages || page < 1 {
		return PaginationResult{Objects: nil, NextPage: 0}
	}

	start := (page - 1) * objectsPerPage
	end := start + objectsPerPage

	var objects []Object

	for i := start; i < end && i < totalObjects; i++ {
		objects = append(objects, Object{
			Name:    fmt.Sprintf("random name %d", i+1),
			Content: fmt.Sprintf("random content %d", i+1),
		})
	}

	nextPage := 0
	if page < totalPages {
		nextPage = page + 1
	}

	return PaginationResult{Objects: objects, CurrentPage: page,NextPage: nextPage}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Default parameters
	currentPage := 1
	objectsPerPage := 10

	// Parse query parameters
	query := r.URL.Query()
	if page, err := strconv.Atoi(query.Get("page")); err == nil && page > 0 {
		currentPage = page
	}
	if perPage, err := strconv.Atoi(query.Get("perPage")); err == nil && perPage > 0 {
		objectsPerPage = perPage
	}

	// Fetch paginated data
	data := simulatePagination(currentPage, objectsPerPage)

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

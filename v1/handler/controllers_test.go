package handler

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sandiptorane/gin_poc/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Logger(userId string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set example variable
		c.Set("user_id", userId)

		// before request

		c.Next()

	}
}
func TestCreateBook(t *testing.T) {
	bookStore := BookStore{books: []model.Book{
		{
			Id:          0,
			Name:        "C programming",
			Author:      "Robert C. Martin",
			Description: "",
		},
	}}
	router := gin.Default()
	router.Use(Logger("user1"))
	router.POST("/v1/create", bookStore.CreateBook)
	tests := []struct {
		name   string
		body   []byte
		status int
	}{
		{
			name: "success",
			body: []byte(`{
    		"name":"C programming",
    		"author" : "Robert C. Martin",
    		"description": "This book contain tutorial of c programming"
       }`),
			status: 200,
		},
		{
			name: "validation",
			body: []byte(`{
    		"author" : "Robert C. Martin",
    		"description": "This book contain tutorial of c programming"
       }`),
			status: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/v1/create", bytes.NewReader(tt.body))
		router.ServeHTTP(w, req)

		assert.Equal(t, tt.status, w.Code)
	}
}

func TestUpdateBook(t *testing.T) {
	bookStore := BookStore{books: []model.Book{
		{
			Id:          0,
			Name:        "C programming",
			Author:      "Robert C. Martin",
			Description: "",
		},
	}}
	router := gin.Default()
	router.Use(Logger("user1"))
	router.PUT("/v1/update/:id", bookStore.UpdateBook)
	tests := []struct {
		id     int
		name   string
		body   []byte
		status int
	}{
		{
			id:   0,
			name: "success",
			body: []byte(`{
    		"name":"C programming updated",
    		"author" : "Robert C. Martin",
    		"description": "This book contain tutorial of c programming"
       }`),
			status: 200,
		},
		{
			id:   2,
			name: "not found",
			body: []byte(`{
            "name" : "C programming",
    		"author" : "Robert C. Martin",
    		"description": "This book contain tutorial of c programming"
       }`),
			status: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", fmt.Sprint("/v1/update/", tt.id), bytes.NewReader(tt.body))
		router.ServeHTTP(w, req)

		assert.Equal(t, tt.status, w.Code)
		fmt.Println(w.Body.String())
	}
}

func TestGetBook(t *testing.T) {
	bookStore := BookStore{books: []model.Book{
		{
			Id:          0,
			Name:        "C programming",
			Author:      "Robert C. Martin",
			Description: "",
		},
	}}
	router := gin.Default()
	router.Use(Logger("user1"))
	router.GET("/v1/getbook/:id", bookStore.GetBook)
	tests := []struct {
		id     int
		name   string
		query  string
		status int
	}{
		{
			id:     0,
			name:   "success with id and name",
			query:  "C programming",
			status: 200,
		},
		{
			id:     0,
			name:   "success with id and name",
			status: 200,
		},
		{
			id:     2,
			name:   "not found",
			status: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprint("/v1/getbook/", tt.id, "?name=", tt.query), nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, tt.status, w.Code)
		fmt.Println(w.Body.String())
	}
}

func TestGetBooks(t *testing.T) {
	bookStore := BookStore{books: []model.Book{
		{
			Id:          0,
			Name:        "C programming",
			Author:      "Robert C. Martin",
			Description: "",
		},
	}}
	router := gin.Default()
	router.Use(Logger("user1"))
	router.GET("/v1/getbooks", bookStore.GetBooks)
	tests := []struct {
		id     int
		name   string
		status int
	}{
		{
			id:     0,
			status: 200,
		},
	}
	for _, tt := range tests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprint("/v1/getbooks"), nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, tt.status, w.Code)
		fmt.Println(w.Body.String())
	}
}

func TestDeleteBook(t *testing.T) {
	bookStore := BookStore{books: []model.Book{
		{
			Id:          0,
			Name:        "C programming",
			Author:      "Robert C. Martin",
			Description: "",
		},
	}}
	router := gin.Default()
	router.Use(Logger("user1"))
	router.GET("/v1/delete/:id", bookStore.DeleteBook)
	tests := []struct {
		id     int
		name   string
		status int
	}{
		{
			id:     0,
			name:   "success",
			status: 200,
		},
		{
			id:     2,
			name:   "not found",
			status: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprint("/v1/delete/", tt.id), nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, tt.status, w.Code)
		fmt.Println(w.Body.String())
	}
}

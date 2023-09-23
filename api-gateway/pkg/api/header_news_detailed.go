package api

import (
	"api-gateway/pkg/models"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

func getNewsDetailed(w http.ResponseWriter, r *http.Request) {

	newsIdStr := chi.URLParam(r, "news_id")
	newsId, err := uuid.Parse(newsIdStr)

	if err != nil {
		res := &BadResponse{
			Success:   false,
			Error:     err.Error(),
			RequestID: middleware.GetReqID(r.Context()),
		}
		//return a bad request and exist the function
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(res)
		//if there's an error with encoding handle it
		if err != nil {
			log.Printf("error decoding request %v\n", err)
		}
		return
	}

	newsCh := make(chan NewsByIDResponse)

	// Запрос к сервису новостой
	go func() {
		req := &NewsByIDRequest{
			ID: newsId,
		}

		b, err := json.Marshal(req)
		if err != nil {
			log.Printf("error marshalling json: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		jsonNews := strings.NewReader(string(b))
		reqNews, err := http.NewRequest(http.MethodGet, newsIdURL, jsonNews)
		if err != nil {
			panic(err)
		}
		reqNews.Header.Set(requestIDHeader, middleware.GetReqID(r.Context()))
		resNews, err := http.DefaultClient.Do(reqNews)
		if err != nil {
			log.Printf("error sending request news service: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		newsResponse := NewsByIDResponse{}
		err = json.NewDecoder(resNews.Body).Decode(&newsResponse)
		if err != nil {
			log.Printf("error decoding response: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		newsCh <- newsResponse
	}()

	commentsCh := make(chan CommentsByIDResponse)

	// Запрос к сервису комментариев
	go func() {
		req := &CommentsByIDRequest{
			ID: newsId,
		}

		b, err := json.Marshal(req)
		if err != nil {
			log.Printf("error marshalling json: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		jsonComments := strings.NewReader(string(b))
		reqComments, err := http.NewRequest(http.MethodGet, commentURL, jsonComments)
		if err != nil {
			panic(err)
		}
		reqComments.Header.Set(requestIDHeader, middleware.GetReqID(r.Context()))
		resComments, err := http.DefaultClient.Do(reqComments)
		if err != nil {
			log.Printf("error sending request comments service: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		commentsResponse := CommentsByIDResponse{}
		err = json.NewDecoder(resComments.Body).Decode(&commentsResponse)
		if err != nil {
			log.Printf("error decoding response: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		commentsCh <- commentsResponse
	}()

	commentsResponse := <-commentsCh

	newsResponse := <-newsCh

	// подготовка ответа на запрос
	res := &NewsDetailedResponse{
		Success: true,
		Post: models.Post{
			ID:        newsResponse.Post.ID,
			Title:     newsResponse.Post.Title,
			Content:   newsResponse.Post.Content,
			CreatedAt: newsResponse.Post.CreatedAt,
			Link:      newsResponse.Post.Link,
		},
		Comments:  commentsResponse.Comments,
		RequestID: middleware.GetReqID(r.Context()),
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding response: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

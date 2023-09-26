package api

import (
	"context"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

const commentURL = "http://localhost:8081/comments"
const newsIdURL = "http://localhost:8082/news/id"
const newsRegURL = "http://localhost:8082/news/reg"
const censorURL = "http://localhost:8083/censor"
const commentTreeURL = "http://localhost:8081/comments/tree"

// Заголовок с которым передается сквозной ID запроса
const requestIDHeader = "X-Request-Id"

type key string

func StartAPI() *chi.Mux {
	//get the router
	r := chi.NewRouter()
	//add middleware
	//in this case we will store our DB to use it later
	r.Use(ctxID)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(sendRequestID)

	r.Route("/", func(r chi.Router) {
		r.Get("/news/{news_id}", getNewsDetailed)
		r.Get("/news", getNews)
		r.Get("/news/tree/{news_id}", getNewsTreeDetailed)
		r.Post("/comments", addComment)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("up and running"))
	})

	return r
}

func ctxID(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestIDStr := r.URL.Query().Get("request_id")
		requestID, err := uuid.Parse(requestIDStr)
		if err != nil {
			requestID = uuid.New()
		}

		params := url.Values{}
		params.Add("request_id", requestID.String())
		ctx := context.WithValue(r.Context(), key("request_id"), requestID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func sendRequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if w.Header().Get(requestIDHeader) == "" {
			w.Header().Add(
				requestIDHeader,
				middleware.GetReqID(r.Context()),
			)
			w.Header().Set(requestIDHeader, middleware.GetReqID(r.Context()))
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

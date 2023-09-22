package api

import (
	"encoding/json"
	"log"
	"net/http"
	"news/pkg/models"
	database "news/pkg/postgres"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func StartAPI(db *database.DB) *chi.Mux {
	//get the router
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	//add middleware
	//in this case we will store our DB to use it later
	r.Use(middleware.Logger, middleware.WithValue("DB", db))

	r.Route("/news", func(r chi.Router) {
		r.Get("/id", getPostByID)
		r.Get("/reg", getPostsByRegExp)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("up and running"))
	})

	return r
}

func getPostByID(w http.ResponseWriter, r *http.Request) {
	req := &GetPostByIDRequest{}
	err := json.NewDecoder(r.Body).Decode(req)

	if err != nil {
		res := &BadResponse{
			Success: false,
			Error:   err.Error(),
		}
		err = json.NewEncoder(w).Encode(res)
		//if there's an error with encoding handle it
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		//return a bad request and exist the function
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db, ok := r.Context().Value("DB").(*database.DB)

	if !ok {
		res := &BadResponse{
			Success: false,
			Error:   "could not get the DB from context",
		}
		err = json.NewEncoder(w).Encode(res)
		//if there's an error with encoding handle it
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		//return a bad request and exist the function
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	post, err := db.GetNewsByID(req.ID)

	if err != nil {

		res := &BadResponse{
			Success: false,
			Error:   err.Error(),
		}

		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(res)
		//if there's an error with encoding handle it
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		return
	}

	res := &GetPostByIDResponse{
		Success: true,
		Post:    post,
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding comments: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func getPostsByRegExp(w http.ResponseWriter, r *http.Request) {
	req := &GetPostsByRegExpRequest{}
	err := json.NewDecoder(r.Body).Decode(req)

	//fmt.Println(req.RequestID)
	if err != nil {
		res := &BadResponse{
			Success:   false,
			Error:     err.Error(),
			RequestID: middleware.GetReqID(r.Context()),
		}
		err = json.NewEncoder(w).Encode(res)
		//if there's an error with encoding handle it
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		//return a bad request and exist the function
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db, ok := r.Context().Value("DB").(*database.DB)

	if !ok {
		res := &BadResponse{
			Success:   false,
			Error:     "could not get the DB from context",
			RequestID: middleware.GetReqID(r.Context()),
		}
		err = json.NewEncoder(w).Encode(res)
		//if there's an error with encoding handle it
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		//return a bad request and exist the function
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Показывает 1 страницу по умолчанию
	if req.CurrentPage == 0 {
		req.CurrentPage = 1
	}

	posts, pagesNum, currentPage, err := db.GetNewsByRegExp(req.RegExp, req.CurrentPage)
	if err != nil {

		res := &BadResponse{
			Success:   false,
			Error:     err.Error(),
			RequestID: middleware.GetReqID(r.Context()),
		}

		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(res)
		//if there's an error with encoding handle it
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		return
	}

	res := &GetPostsByRegExpResponse{
		Success: true,
		Posts:   posts,
		Pagination: models.Pagination{
			CurrentPage: currentPage,
			PagesNumber: pagesNum,
			ItemsOnPage: database.ItemsOnPage,
		},
		RequestID: middleware.GetReqID(r.Context()),
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding comments: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

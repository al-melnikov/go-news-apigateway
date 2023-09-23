package api

import (
	"censor/pkg/censor"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func StartAPI() *chi.Mux {
	//get the router
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	//add middleware
	//in this case we will store our DB to use it later
	r.Use(middleware.Logger)

	r.Route("/censor", func(r chi.Router) {
		r.Put("/", censorComment)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("up and running"))
	})

	return r
}

func censorComment(w http.ResponseWriter, r *http.Request) {
	req := &CensorReqest{}
	err := json.NewDecoder(r.Body).Decode(req)

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
	fmt.Println(req.Content)

	isCensored := censor.IsCensored(req.Content)

	res := &CensorResponse{
		Success:    true,
		IsCensored: isCensored,
		RequestID:  middleware.GetReqID(r.Context()),
	}

	if isCensored {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error encoding response: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		return
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding response: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

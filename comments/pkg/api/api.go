package api

import (
	"comments/pkg/models"
	db "comments/pkg/postgres"
	"comments/pkg/tree"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

// Good tutorial
// https://bognov.tech/modern-api-design-with-golang-postgresql-and-docker

func StartAPI(db *db.DB) *chi.Mux {
	//get the router
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	//add middleware
	//in this case we will store our DB to use it later
	r.Use(middleware.Logger, middleware.WithValue("DB", db))

	r.Route("/comments", func(r chi.Router) {
		r.Post("/", createComment)
		r.Get("/", getComments)
		r.Get("/tree", getCommentsTree)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("up and running"))
	})

	return r
}

func createComment(w http.ResponseWriter, r *http.Request) {

	req := &CreateCommentRequest{}
	err := json.NewDecoder(r.Body).Decode(req)

	if err != nil {
		res := &PostCommentResponse{
			Success:   false,
			Error:     err.Error(),
			CommentID: uuid.Nil,
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

	db, ok := r.Context().Value("DB").(*db.DB)

	if !ok {
		res := &PostCommentResponse{
			Success:   false,
			Error:     "could not get the DB from context",
			CommentID: uuid.Nil,
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

	commentID, err := db.AddComment(models.Comment{
		NewsID:    req.NewsID,
		ParentID:  req.ParentID,
		CreatedAt: time.Now(),
		Content:   req.Content,
	})

	if err != nil {
		res := &PostCommentResponse{
			Success:   false,
			Error:     err.Error(),
			CommentID: uuid.Nil,
			RequestID: middleware.GetReqID(r.Context()),
		}
		fmt.Println(res)
		err = json.NewEncoder(w).Encode(res)
		//if there's an error with encoding handle it
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		//return a bad request and exist the function
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := &PostCommentResponse{
		Success:   true,
		Error:     "",
		CommentID: *commentID,
		RequestID: middleware.GetReqID(r.Context()),
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding after creating comment %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getComments(w http.ResponseWriter, r *http.Request) {
	req := &GetCommentsRequest{}
	err := json.NewDecoder(r.Body).Decode(req)

	if err != nil {
		res := &GetCommentsResponse{
			Success:   false,
			Error:     err.Error(),
			Comments:  nil,
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

	db, ok := r.Context().Value("DB").(*db.DB)

	if !ok {
		res := &GetCommentsResponse{
			Success:   false,
			Error:     "could not get the DB from context",
			Comments:  nil,
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

	comments, err := db.GetComments(req.NewsID)

	if err != nil {

		res := &GetCommentsResponse{
			Success:   false,
			Error:     err.Error(),
			Comments:  nil,
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

	res := &GetCommentsResponse{
		Success:   true,
		Error:     "",
		Comments:  comments,
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

func getCommentsTree(w http.ResponseWriter, r *http.Request) {
	req := &GetCommentsTreeRequest{}
	err := json.NewDecoder(r.Body).Decode(req)

	if err != nil {
		res := &GetCommentsTreeResponse{
			Success:   false,
			Error:     err.Error(),
			Comments:  nil,
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

	db, ok := r.Context().Value("DB").(*db.DB)

	if !ok {
		res := &GetCommentsTreeResponse{
			Success:   false,
			Error:     "could not get the DB from context",
			Comments:  nil,
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

	comments, err := db.GetComments(req.NewsID)

	if err != nil {

		res := &GetCommentsTreeResponse{
			Success:   false,
			Error:     err.Error(),
			Comments:  nil,
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

	commentsTree := tree.ArrayToTree(comments)

	res := &GetCommentsTreeResponse{
		Success:   true,
		Error:     "",
		Comments:  commentsTree,
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

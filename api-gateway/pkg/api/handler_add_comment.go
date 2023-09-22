package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5/middleware"
)

func addComment(w http.ResponseWriter, r *http.Request) {
	req := &AddCommentRequest{}
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

	reqStrCensor := &AddCommentCensorRequest{
		Content: req.Content,
	}

	b, err := json.Marshal(reqStrCensor)
	if err != nil {
		log.Printf("error marshalling json: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jsonCensor := strings.NewReader(string(b))
	reqCensor, err := http.NewRequest(http.MethodPut, censorURL, jsonCensor)
	if err != nil {
		panic(err)
	}
	reqCensor.Header.Set(requestIDHeader, middleware.GetReqID(r.Context()))

	resCensor, err := http.DefaultClient.Do(reqCensor)
	if err != nil {
		log.Printf("error sending request news service: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	censResponse := CensorResponse{}

	err = json.NewDecoder(resCensor.Body).Decode(&censResponse)

	if err != nil {
		log.Printf("error decoding censor response: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if censResponse.Success && censResponse.IsCensored {
		res := &BadResponse{
			Success:   false,
			Error:     "bad content",
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

	//err = json.NewEncoder(w).Encode(req)
	b, err = json.Marshal(req)
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

	jsonComment := strings.NewReader(string(b))

	reqComment, err := http.NewRequest(http.MethodPost, commentURL, jsonComment)
	if err != nil {
		panic(err)
	}
	reqComment.Header.Set(requestIDHeader, middleware.GetReqID(r.Context()))

	resComment, err := http.DefaultClient.Do(reqComment)
	if err != nil {
		log.Printf("error sending request news service: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	addCommentResponse := AddCommentResponse{}

	err = json.NewDecoder(resComment.Body).Decode(&addCommentResponse)

	if err != nil {
		log.Printf("error decoding response: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res := addCommentResponse

	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		log.Printf("error encoding response: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}

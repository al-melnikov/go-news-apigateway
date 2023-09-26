package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5/middleware"
)

func getNews(w http.ResponseWriter, r *http.Request) {

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}

	req := NewsRegExpRequest{
		Page:   page,
		RegExp: r.URL.Query().Get("reg_exp"),
	}

	b, err := json.Marshal(req)
	if err != nil {
		log.Printf("error marshalling json: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jsonNews := strings.NewReader(string(b))
	reqNews, err := http.NewRequest(http.MethodGet, newsRegURL, jsonNews)
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

	newsResponse := NewsRegExpResponse{}
	err = json.NewDecoder(resNews.Body).Decode(&newsResponse)
	if err != nil {
		log.Printf("error decoding response: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res := newsResponse

	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		log.Printf("error encoding response: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}

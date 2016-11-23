package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Result interface{}

type Response struct {
	Result `json:"result"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func ProfileInfo(w http.ResponseWriter, r *http.Request) {
	profileId := GetProfileId(r)

	var responseWriter http.ResponseWriter
	responseWriter = SetHeaders(w, http.StatusOK)
	result := Response{GetProfile(profileId)}
	if err := json.NewEncoder(responseWriter).Encode(result); err != nil {
		panic(err)
	}
}

func ProfileBoobs(w http.ResponseWriter, r *http.Request) {
	profileId := GetProfileId(r)
	profile := GetProfile(profileId)

	var result Result
	var hasBoobs bool
	hasBoobs = profile.Age >= 18
	result = Response{Boobs{&hasBoobs}}

	var responseWriter http.ResponseWriter
	responseWriter = SetHeaders(w, http.StatusOK)
	if err := json.NewEncoder(responseWriter).Encode(result); err != nil {
		panic(err)
	}
}

func ProfileAvatar(w http.ResponseWriter, r *http.Request) {
	profileId := GetProfileId(r)
	fmt.Fprintln(w, "Get profile avatar:", profileId)
}

func GetProfileId(r *http.Request) string {
	vars := mux.Vars(r)
	profileId := vars["profileId"]
	return profileId
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	profile := ReadProfile(w, r)
	p := UpdateRepoProfile(profile)

	SetHeaders(w, http.StatusOK)
	if err := json.NewEncoder(w).Encode(Response{p}); err != nil {
		panic(err)
	}

}

func CreateProfile(w http.ResponseWriter, r *http.Request) {
	profile := ReadProfile(w, r)
	p := SaveProfile(profile)

	SetHeaders(w, http.StatusCreated)
	if err := json.NewEncoder(w).Encode(Response{p}); err != nil {
		panic(err)
	}
}

func ReadProfile(w http.ResponseWriter, r *http.Request) Profile {
	var profile Profile
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 104856))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &profile); err != nil {
		SetHeaders(w, 422)
		if err := json.NewEncoder(w).Encode(Response{err}); err != nil {
			panic(err)
		}
	}
	return profile
}

func SetHeaders(w http.ResponseWriter, status int) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(status)
	return w
}

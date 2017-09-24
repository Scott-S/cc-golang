package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type GeneralInfo struct {
    //remove omit empty
    Email          string   `json:"email"`
    Website        string   `json:"website"`
    GithubRepoLink string   `json:"github_repo_link"`
    Name           *Name    `json:"name"`
}
type Name struct {
    First string `json:"first"`
    Last  string `json:"last"`
}

var Data = GeneralInfo{
    Email: "scott.scialabba@gmail.com",
    Name: &Name{
        First: "Scott",
        Last: "Scialabba" },
    Website: "http://scott.scialabba.com",
    GithubRepoLink: "https://github.com/Scott-S/cc-golang"}

func GetData(w http.ResponseWriter, r *http.Request) {
    //set content type
    w.Header().Set("Content-Type", "application/json; charset=utf-8")

    //set CORS header
    w.Header().Set("Access-Control-Allow-Origin", "*")

    //respond with status 201
    w.WriteHeader(http.StatusCreated)

    //json encode Data
    json.NewEncoder(w).Encode(Data)
}

// our main function
func main() {
    router := mux.NewRouter()

    //match /code/challenge or /code/challenge/
    router.StrictSlash(true)

    router.HandleFunc("/code/challenge", GetData).Methods("GET")
    log.Fatal(http.ListenAndServe(":3000", router))
}
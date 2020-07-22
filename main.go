package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// struct that holds a bullshit sentence
type bullshitSentence struct {
	adverb    string
	verb      string
	adjective string
	noun      string
}

// formats the bullshitSentence as a string for a nice printing
func (s bullshitSentence) prettyprint() string {
	return fmt.Sprintf("%s %s %s %s", s.adverb, s.verb, s.adjective, s.noun)
}

// generate and returns a bullshitSentence struct
func generateBullshit() bullshitSentence {
	s2 := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s2)
	seed := r.Intn(100)

	sentence := bullshitSentence{}

	sentence.noun = words["nouns"][seed%len(words["nouns"])]
	sentence.adverb = words["adverbs"][seed%len(words["adverbs"])]
	sentence.adjective = words["adjectives"][seed%len(words["adjectives"])]
	sentence.verb = words["verbs"][seed%len(words["verbs"])]
	return sentence
}

// handler when attacking the /bullshit endpoint
func getBullshitSentenceFromAPI(w http.ResponseWriter, r *http.Request) {
	sentenceForUser := generateBullshit().prettyprint()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"bullshit": "` + sentenceForUser + `}`))
}

func main() {
	log.Println("Starting ")
	r := mux.NewRouter()
	r.HandleFunc("/bullshit", getBullshitSentenceFromAPI).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", r))
}

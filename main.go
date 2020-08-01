package main

import (
	"log"
	"fmt"
	"strings"
	"math/rand"
	"net/http"
	"time"
	"strconv"
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

// struct for a bullshit pararaph
type bullshitParagraph struct {
	start     string
	sentences []bullshitSentence
	separator []string
	end       string
}

// returns a str formatted bullshitParagraph
func (p bullshitParagraph) prettyprint() string {
	var paragraphstr strings.Builder

	// Add the start
	paragraphstr.WriteString(p.start)

	// add sentence+sep for len(sentence)-1
	for i := 0; i < len(p.sentences)-1; i++ {
		paragraphstr.WriteString(p.sentences[i].prettyprint())
		paragraphstr.WriteString(p.separator[i])
	}
	// add last sentence
	paragraphstr.WriteString(p.sentences[len(p.sentences)-1].prettyprint())

	// add and
	paragraphstr.WriteString(p.end)

	return paragraphstr.String()
}

// generate and returns a bullshitSentence struct
func generateSentence() bullshitSentence {
	rand.Seed(time.Now().UnixNano())
	seed := rand.Intn(100)

	sentence := bullshitSentence{}

	sentence.noun = words["nouns"][seed%len(words["nouns"])]
	sentence.adverb = words["adverbs"][seed%len(words["adverbs"])]
	sentence.adjective = words["adjectives"][seed%len(words["adjectives"])]
	sentence.verb = words["verbs"][seed%len(words["verbs"])]
	return sentence
}

// given a number of sentences, it can 
func generateParagraph(num int) bullshitParagraph {
	paragraph := bullshitParagraph{}
	rand.Seed(time.Now().UnixNano())
	seed := rand.Intn(100)
	paragraph.start = words["starts"][seed%len(words["starts"])]

	rand.Seed(time.Now().UnixNano())
	seed = rand.Intn(100)
	paragraph.end = words["ends"][seed%len(words["ends"])]

	for i := 0; i < num; i++ {
		paragraph.sentences = append(paragraph.sentences, generateSentence())
		rand.Seed(time.Now().UnixNano())
		seed = rand.Intn(100)
		paragraph.separator = append(paragraph.separator,words["separators"][seed%len(words["separators"])])
	}

	return paragraph
}

// goroutine when attacking the /sentence endpoint
func getBullshitSentenceFromAPI(w http.ResponseWriter, r *http.Request) {
	log.Printf("User %s requested a sentence\n",r.RemoteAddr)
	genStart := time.Now()
	sentenceForUser := generateSentence().prettyprint()
	genDur := time.Now().Sub(genStart)
	log.Printf("Sentence for %s generated in %s\n",r.RemoteAddr,genDur.String())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"sentence": "` + sentenceForUser + `"}`))
}

// goroutine when attacking /paragraph
func getBullshitParagraphFromAPI(w http.ResponseWriter, r *http.Request) {
	numOfsentences := mux.Vars(r)["num"]
	log.Printf("User %s requested a paragraph of len %s\n",r.RemoteAddr,numOfsentences)
	intNumOfsentences, err := strconv.Atoi(numOfsentences)
	if err != nil{
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return 
	}
	genStart := time.Now()
	paragraphForUser := generateParagraph(intNumOfsentences)
	genDur := time.Now().Sub(genStart)
	log.Printf("Paragraph for %s generated in %s\n",r.RemoteAddr,genDur.String())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"paragraph": "` + paragraphForUser.prettyprint() + `"}`))
}

func main() {
	log.Println("Starting go-bs")

	port := "8080"
	r := mux.NewRouter()
	log.Printf("go-bs is listening on port %s\n",port)
	r.HandleFunc("/sentence", getBullshitSentenceFromAPI).Methods(http.MethodGet)
	r.HandleFunc("/paragraph/{num}", getBullshitParagraphFromAPI).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":"+port, r))
}

package main

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "indexを表示したしん")
}

func Show(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "movie showを表示したしん")
}

func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "movie createを表示したしん")
}

func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "movie updateを表示したしん")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "movie deleteを表示したしん")
}
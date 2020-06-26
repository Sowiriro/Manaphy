package main

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "indexを表示したしん")
	return
}

func Show(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "movie showを表示したしん")
	return
}

func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "movie createを表示したしん")
	return
}

func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "movie updateを表示したしん")
	return
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "movie deleteを表示したしん")
	return
}

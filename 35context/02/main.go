package main

import (
	"context"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ctx", ctxHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Fprintln(w, ctx)
}

func ctxHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, "userID", 777)
	ctx = context.WithValue(ctx, "fname", "Bond")

	results := dbAccess(ctx)
	fmt.Fprintln(w, results)
}

func dbAccess(ctx context.Context) int {

	uid := ctx.Value("userID").(int)
	return uid
}

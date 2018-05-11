package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
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

	results, err := dbAccess(ctx)
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}
	fmt.Fprintln(w, results)
}

func dbAccess(ctx context.Context) (int, error) {

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	ch := make(chan int)

	go func() {
		uid := ctx.Value("userID").(int)
		time.Sleep(10 * time.Second)

		if ctx.Err() != nil {
			return
		}

		ch <- uid
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case i := <-ch:
		return i, nil
	}

}

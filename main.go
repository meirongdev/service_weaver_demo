package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ServiceWeaver/weaver"
)

func main() {
	if err := weaver.Run(context.Background(), run); err != nil {
		panic(err)
	}
}

// app is the main component of our application.
type app struct {
	weaver.Implements[weaver.Main]
	r   weaver.Ref[Reverser]
	lis weaver.Listener
}

// run implements the application main.
func run(ctx context.Context, a *app) error {
	http.HandleFunc("/r", func(w http.ResponseWriter, r *http.Request) {
		str := r.URL.Query().Get("s")
		if str == "" {
			http.Error(w, "missing 's' query parameter", http.StatusBadRequest)
			return
		}
		res, err := a.r.Get().Reverse(ctx, str)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, res)
		// w.Write([]byte(fmt.Sprintf("'%s' reversed is '%s'\n", str, res)))
	})

	fmt.Println("listening on", a.lis)
	return http.Serve(a.lis, nil)
}

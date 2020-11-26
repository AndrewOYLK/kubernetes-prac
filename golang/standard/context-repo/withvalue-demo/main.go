package main

import (
	"context"
	"fmt"
)

func main() {
	type tmpKey string

	f := func(ctx context.Context, k tmpKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("nofound!")
	}

	k := tmpKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, tmpKey("color"))
}

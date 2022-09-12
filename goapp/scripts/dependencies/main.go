package main

import (
	"context"
	"fmt"

	"github.com/Khan/genqlient/graphql"
	"github.com/dagger/cloak/engine"
)

func main() {
	engine.Start(context.Background(), &engine.Config{}, func(ctx engine.Context) error {
		fmt.Println("Hello World")
		client := ctx.Client
		req := &graphql.Request{
			Query: `
			query {
				core {
				  image(ref:"alpine"){
					exec(input: {args:["ls","."]}) {
							  stdout
					}
				  }
				}
			  }
	`,
		}
		resp := ""
		err := client.MakeRequest(ctx, req, &graphql.Response{Data: &resp})
		if err != nil {
			fmt.Println(" Errooooor ", err)
		}
		return err
	})
}

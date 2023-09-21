package main

import (
	gen "github.com/rogeecn/genx"
	"github.com/rogeecn/genx/examples/dal/model"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../../dal/query",
		Mode:    gen.WithDefaultQuery,
	})

	// generate from struct in project
	g.ApplyBasic(model.Customer{})

	g.Execute()
}

package main

import (
	"context"
	"fmt"

	"github.com/rogeecn/genx/examples/biz"
	"github.com/rogeecn/genx/examples/conf"
	"github.com/rogeecn/genx/examples/dal"
	"github.com/rogeecn/genx/examples/dal/query"
)

func init() {
	dal.DB = dal.ConnectDB(conf.MySQLDSN).Debug()
}

func main() {
	// start your project here
	fmt.Println("hello world")
	defer fmt.Println("bye~")

	query.SetDefault(dal.DB)
	biz.Query(context.Background())
}

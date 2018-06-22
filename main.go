package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/repokecil/funcs"
	"github.com/melodiez14/repokecil/util/conn"
	routerv1 "github.com/melodiez14/repokecil/webserver/v1"
	routerv2 "github.com/melodiez14/repokecil/webserver/v2"
)

func main() {
	// this config load is better
	// example: https://github.com/melodiez14/meiko/blob/master/app.go
	db := conn.InitDB(conn.DBConfig{
		Database: "tokopedia-review",
		Driver:   "postgres",
		UserName: "reputation_user",
		Password: "d9tcEk9gEv4WA42j",
		Port:     "5432",
		Host:     "devel-postgre.tkpd",
	})

	// this should be
	// m := core.Init(db) {or struct of all db}
	// f := funcs.Init(m)
	// ...
	// routerv1.InitRouter(f, r)
	// routerv2.InitRouter(f, r)
	// routerv3.InitRouter(f, r)
	// ...

	f := funcs.Init(db)

	r := httprouter.New()
	routerv1.InitRouter(f, r)
	routerv2.InitRouter(f, r)
	http.ListenAndServe(":5000", r)
}

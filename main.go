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
	db := conn.InitDB(conn.DBConfig{
		Database: "tokopedia-review",
		Driver:   "postgres",
		UserName: "reputation_user",
		Password: "d9tcEk9gEv4WA42j",
		Port:     "5432",
		Host:     "devel-postgre.tkpd",
	})

	f := funcs.Init(db)

	r := httprouter.New()
	routerv1.InitRouter(f, r)
	routerv2.InitRouter(f, r)
	http.ListenAndServe(":5000", r)
}

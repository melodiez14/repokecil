package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/buaazp/fasthttprouter"
	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/repokecil/funcs"
	"github.com/melodiez14/repokecil/util/conn"
	routerv1 "github.com/melodiez14/repokecil/webserver/v1"
	routerv2 "github.com/melodiez14/repokecil/webserver/v2"
	"github.com/valyala/fasthttp"
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

	_ = conn.InitRedis(conn.RedisConfig{
		ConnectionString: "devel-redis.tkpd:6379",
		MaxActive:        10,
		MaxIdle:          3,
		IdleTimeout:      10,
		Wait:             true,
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
	r.Handler(http.MethodGet, "/debug/pprof/:item", http.DefaultServeMux)
	r.GET("/", routers)
	routerv1.InitRouter(f, r)
	routerv2.InitRouter(f, r)
	go http.ListenAndServe(":8080", r)

	fr := fasthttprouter.New()
	fr.GET("/", router)
	fasthttp.ListenAndServe(":8081", fr.Handler)
}

func routers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Write([]byte("OK\n"))
}

func router(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("X-Content-Type-Options", "nosniff")
	ctx.WriteString("ok\n")
}

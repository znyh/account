package http

import (
	"github.com/znyh/account/internal/model"
	pb "github.com/znyh/proto/account"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
)

var svc pb.AccountServer

// New new a bm server.
func New(s pb.AccountServer) (engine *bm.Engine, err error) {
	var (
		cfg bm.ServerConfig
		ct  paladin.TOML
	)
	if err = paladin.Get("http.txt").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	svc = s
	engine = bm.DefaultServer(&cfg)
	pb.RegisterAccountBMServer(engine, s)
	initRouter(engine)
	err = engine.Start()
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/account")
	{
		g.GET("/start", howToStart)
	}
}

func ping(ctx *bm.Context) {
	//if _, err := svc.Ping(ctx, nil); err != nil {
	//	log.Error("ping error(%v)", err)
	//	ctx.AbortWithStatus(http.StatusServiceUnavailable)
	//}
}

// example for http request handler.
func howToStart(c *bm.Context) {
	k := &model.Kratos{
		Hello: "Golang 大法好 !!!",
	}
	c.JSON(k, nil)
}

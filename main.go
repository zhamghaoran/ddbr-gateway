package main

import (
	"github.com/cloudwego/kitex/server"
	"log"
	gateway "zhamghaoran/ddbr-gateway/kitex_gen/ddbr/rpc/gateway/gateway"
	"zhamghaoran/ddbr-gateway/service/middware"
)

func main() {
	svr := gateway.NewServer(new(GatewayImpl), server.WithMiddleware(middware.AuthorityMiddleware))
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

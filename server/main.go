package main

import (
	"log"
	"net"

	"github.com/spadesk1991/micro/server/controller/hello_controller"
	"github.com/spadesk1991/micro/server/proto/hello"

	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	Address = "0.0.0.0:9090"
)

func main() {
	// certFile := "./server/cert/primary.pem"
	// keyFile := "./server/cert/public.pem"
	// tlsServer := gtls.Server{
	// 	CertFile: certFile,
	// 	KeyFile:  keyFile,
	// }

	// c, err := tlsServer.GetTLSCredentials()
	// if err != nil {
	// 	log.Fatalf("tlsServer.GetTLSCredentials err: %v", err)
	// }
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	// 服务注册
	hello.RegisterHelloServer(s, &hello_controller.HelloController{})

	log.Println("Listen on " + Address)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "hello world!"})
	})
	http.ListenAndServe("0.0.0.0:9000", router)
	log.Println("Listen on 0.0.0.0:9000 ")
}

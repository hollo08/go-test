package main
//https://www.cnblogs.com/baoshu/p/13507260.html#head4
//https://www.cnblogs.com/rickiyang/p/14989375.html
import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"log"
	"net"
	"test/server/product"
	"test/server/token"
)

func main() {
	cert, _ := tls.LoadX509KeyPair("server/cert/server.pem", "server/cert/server.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("server/cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

	// 1. new一个grpc的server
	rpcServer := grpc.NewServer(grpc.Creds(creds))

	// 2. 将刚刚我们新建的ProdService注册进去
	product.RegisterProdServiceServer(rpcServer, new(product.ProdService))
	token.RegisterPingServer(rpcServer, new(token.Server))

	reflection.Register(rpcServer)

	// 3. 新建一个listener，以tcp方式监听8082端口
	listener, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatal("服务监听端口失败", err)
	}

	// 4. 运行rpcServer，传入listener
	_ = rpcServer.Serve(listener)
}

package main
//https://www.cnblogs.com/baoshu/p/13507260.html#head4
//https://blog.csdn.net/cuichenghd/article/details/109230584
//https://blog.csdn.net/cuichenghd/article/details/109230584
import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net"
	"test/service/product"
)

func main() {
	//tls, err := credentials.NewServerTLSFromFile("service/cert/22server.crt", "service/cert/22server_no_password.key")
	//if err != nil {
	//	log.Fatal("服务端获取证书失败: ", err)
	//}
	cert, _ := tls.LoadX509KeyPair("service/cert/server.pem", "service/cert/server.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("service/cert/server.pem")
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

	// 3. 新建一个listener，以tcp方式监听8082端口
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal("服务监听端口失败", err)
	}

	// 4. 运行rpcServer，传入listener
	_ = rpcServer.Serve(listener)
}

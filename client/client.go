package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"test/client/service/product"
)

func main() {
	//tls, err := credentials.NewClientTLSFromFile("client/cert/ssserver.crt", "stone")

	//if err != nil {
	//	log.Fatal("客户端获取证书失败: ", err)
	//}
	cert, _ := tls.LoadX509KeyPair("client/cert/ca.crt", "client/cert/ca.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("client/cert/ca.crt")
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
		RootCAs:      certPool,
	})

	// 1. 新建连接，端口是服务端开放的8082端口
	// 并且添加grpc.WithInsecure()，不然没有证书会报错
	conn, err := grpc.Dial(":8082", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}

	// 退出时关闭链接
	defer conn.Close()

	// 2. 调用Product.pb.go中的NewProdServiceClient方法
	productServiceClient := product.NewProdServiceClient(conn)

	// 3. 直接像调用本地方法一样调用GetProductStock方法
	resp, err := productServiceClient.GetProductStock(context.Background(), &product.ProductRequest{ProdId: 2147483647})
	if err != nil {
		log.Fatal("调用gRPC方法错误: ", err)
	}

	fmt.Println("调用gRPC方法成功，ProdStock = ", resp.ProdStock)
}

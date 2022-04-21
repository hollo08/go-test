package product

import (
	"context"
	"fmt"
	token "github.com/hollo08/go-test/server/token"
)

type ProdService struct {
}

func (ps *ProdService) GetProductStock(ctx context.Context, request *ProductRequest) (*ProductResponse, error) {
	user := token.CheckAuth(ctx)
	fmt.Printf("rquest name is ：%s and request param is : %d \n",user, request.GetProdId())
	return &ProductResponse{ProdStock: request.ProdId}, nil
}

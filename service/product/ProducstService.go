package product

import (
	"context"
	"fmt"
)

type ProdService struct {
}

func (ps *ProdService) GetProductStock(ctx context.Context, request *ProductRequest) (*ProductResponse, error) {
	fmt.Printf("request param is : %d \n", request.GetProdId())
	return &ProductResponse{ProdStock: request.ProdId}, nil
}

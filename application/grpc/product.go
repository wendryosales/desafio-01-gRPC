package grpc

import (
	"context"

	"github.com/wendryosales/desafio-01-gRPC/application/grpc/pb"
	"github.com/wendryosales/desafio-01-gRPC/application/usecase"
)

type ProductGrpcService struct {
	ProductUseCase usecase.ProductUseCase
	pb.UnimplementedProductServiceServer
}

func (p *ProductGrpcService) CreateProduct(
	ctx context.Context,
	in *pb.CreateProductRequest,
) (*pb.CreateProductResponse, error) {

	product, err := p.ProductUseCase.CreateProduct(
		in.GetName(),
		in.GetDescription(),
		in.GetPrice(),
	)

	if err != nil {
		return &pb.CreateProductResponse{}, err
	}

	return &pb.CreateProductResponse{
		Product: &pb.Product{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		},
	}, nil
}

func (p *ProductGrpcService) FindProducts(
	ctx context.Context,
	in *pb.FindProductsRequest,
) (*pb.FindProductsResponse, error) {

	var products []*pb.Product

	product, err := p.ProductUseCase.FindProducts()

	if err != nil {
		return &pb.FindProductsResponse{}, err
	}

	for _, product := range product {
		products = append(products, &pb.Product{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		})
	}

	return &pb.FindProductsResponse{
		Products: products,
	}, err
}

func NewProductGrpcService(productUseCase usecase.ProductUseCase) *ProductGrpcService {
	return &ProductGrpcService{ProductUseCase: productUseCase}
}

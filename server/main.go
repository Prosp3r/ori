package main

import (
	"context"
	"fmt"
	"log"
	"net"

	ori "gitlab.com/Prosp3r/ori/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//server is an empty struct for instanciating rpc server...
type server struct {
}

func main() {

	port := ":8080"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %v with error: %v \n", port, err)
	}
	fmt.Printf("OriCalc Server started on port %v \n", port)

	s := grpc.NewServer()
	ori.RegisterORIServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed serving ORI with error : %v \n", err)
	}
}

//Divide implements division feature of ori
//returns division of two integers
func (s *server) Divide(ctx context.Context, r *ori.ORIDivideRequest) (*ori.ORIDivideResponse, error) {

	a, b := r.A, r.B
	ans := a / b
	fmt.Printf("Result of Divide operation: %v \n", ans)
	return &ori.ORIDivideResponse{Result: ans}, nil
}

//Multiply implements multiplicaiton feature of ori
//Returns multiplication of two integers
func (s *server) Multiply(ctx context.Context, r *ori.ORIMultiplyRequest) (*ori.ORIMultiplyResponse, error) {

	a, b := r.A, r.B
	ans := a * b
	fmt.Printf("Result of Multiply operation: %v \n", ans)
	return &ori.ORIMultiplyResponse{Result: ans}, nil
}

//Sum implements Sum feaure of ori
//Returns the sum of two integers
func (s *server) Sum(ctx context.Context, r *ori.ORISumRequest) (*ori.ORISumResponse, error) {

	a, b := r.A, r.B
	ans := a + b
	fmt.Printf("Result of Sum operation: %v \n", ans)
	return &ori.ORISumResponse{Result: ans}, nil
}

//Sutract implements subtraction feature of ori
//returns the subtractions of two integers
func (s *server) Sutract(ctx context.Context, r *ori.ORISutractRequest) (*ori.ORISutractResponse, error) {

	a, b := r.A, r.B
	ans := a - b
	fmt.Printf("Result of Subtract operation: %v \n", ans)
	return &ori.ORISutractResponse{Result: ans}, nil
}

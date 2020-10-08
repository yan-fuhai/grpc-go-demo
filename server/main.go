package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-go-demo/area_calculator"
	"log"
	"net"
)

const (
	port = ":40183"	// equal to localhost:40183
)

func triangleArea(_ context.Context, params *pb.TriangleParams) (*pb.AreaReply, error) {
	log.Printf("[Triangle Area] Received: height=%v, base=%v", params.GetHeight(), params.GetBase())
	return &pb.AreaReply{Area: params.GetHeight() * params.GetBase() / 2}, nil
}

func rectangleArea(_ context.Context, params *pb.RectangleParams) (*pb.AreaReply, error) {
	log.Printf("[Rectangle Area] Received: height=%v, width=%v", params.GetHeight(), params.GetWidth())
	return &pb.AreaReply{Area: params.GetHeight() * params.GetWidth()}, nil
}

func rhombusArea(_ context.Context, params *pb.RhombusParams) (*pb.AreaReply, error) {
	log.Printf("[RhombusArea Area] Received: height=%v, base=%v", params.GetHeight(), params.GetBase())
	return &pb.AreaReply{Area: params.GetHeight() * params.GetBase()}, nil
}

func squareArea(_ context.Context, params *pb.SquareParams) (*pb.AreaReply, error) {
	log.Printf("[Square Area] Received: length=%v", params.GetLength())
	return &pb.AreaReply{Area: params.GetLength() * params.GetLength()}, nil
}


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAreaCalculatorService(s, &pb.AreaCalculatorService{
		TriangleArea:  triangleArea,
		RectangleArea: rectangleArea,
		RhombusArea:   rhombusArea,
		SquareArea:    squareArea,
	})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

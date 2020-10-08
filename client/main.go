// gRPC client implementation
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-go-demo/area_calculator"
	"log"
	"strconv"
)

const (
	address = "localhost:40183"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAreaCalculatorClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var option string
	var base float64
	var height float64
	var width float64
	var length float64

	for {
		fmt.Println("Select to calculate:")
		fmt.Println("1. Triangle")
		fmt.Println("2. Rectangle")
		fmt.Println("3. Rhombus")
		fmt.Println("4. Square")
		fmt.Println("5. exit")
		fmt.Print("Input number: ")
		_, _ = fmt.Scanf("%s", &option)

		op, err := strconv.Atoi(option)
		if err != nil || op < 1 || op > 5 {
			fmt.Println("Invalid selection: ", option)
			continue
		}

		// exit
		if op == 5 {
			break
		}

		switch op {
		case 1:
			// Triangle Area
			fmt.Println("Triangle Area")
			fmt.Print("Base: ")
			_, _ = fmt.Scanf("%f", &base)
			fmt.Print("Height: ")
			_, _ = fmt.Scanf("%f", &height)

			triParams := &pb.TriangleParams{
				Base:   base,
				Height: height,
			}
			r, err := c.TriangleArea(ctx, triParams)
			if err != nil {
				log.Fatalf("could not get result from server: %v", err)
			}
			log.Printf("Triangle area: %v", r.Area)
		case 2:
			// Rectangle Area
			fmt.Println("Rectangle Area")
			fmt.Print("Width: ")
			_, _ = fmt.Scanf("%f", &width)
			fmt.Print("Height: ")
			_, _ = fmt.Scanf("%f", &height)

			r, err := c.RectangleArea(ctx, &pb.RectangleParams{
				Width:  width,
				Height: height,
			})
			if err != nil {
				log.Fatalf("could not get result from server: %v", err)
			}
			log.Printf("Rectangle area: %v", r.Area)
		case 3:
			// Rhombus Area
			fmt.Println("Rhombus Area")
			fmt.Print("Base: ")
			_, _ = fmt.Scanf("%f", &base)
			fmt.Print("Height: ")
			_, _ = fmt.Scanf("%f", &height)

			r, err := c.RhombusArea(ctx, &pb.RhombusParams{
				Base:  base,
				Height: height,
			})
			if err != nil {
				log.Fatalf("could not get result from server: %v", err)
			}
			log.Printf("Rhombus area: %v", r.Area)
		case 4:
			// Square Area
			fmt.Println("Square Area")
			fmt.Print("Length: ")
			_, _ = fmt.Scanf("%f", &length)

			r, err := c.SquareArea(ctx, &pb.SquareParams{
				Length:  length,
			})
			if err != nil {
				log.Fatalf("could not get result from server: %v", err)
			}
			log.Printf("Square area: %v", r.Area)
		}
	}
}

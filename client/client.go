package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "github.com/zanven42/corevision-HelloWorld"
	"google.golang.org/grpc"
)

func main() {
	addr := flag.String("addr", "localhost:8000", "set the address to dial for a GRPC connection, e.g `localhost:8080`")
	h := flag.Bool("h", false, "Help")
	flag.Parse()
	if *h {
		flag.Usage()
		return
	}

	grpcConn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial error: %s\n", err)
	}

	client := pb.NewHelloWorldClient(grpcConn)
	// no cancelation required so just getting background
	ctx := context.Background()
	helloInput1 := &pb.SayHelloInput{FirstName: "Core", LastName: "Vision"}
	helloResponse1, err1 := client.SayHello(ctx, helloInput1)
	helloInput2 := &pb.SayHelloInput{ForceError: true, DesiredErrorMessage: "Showing an error"}
	helloResponse2, err2 := client.SayHello(ctx, helloInput2)

	fmt.Printf("SayHello Input 1: %#v\n", *helloInput1)
	if err1 != nil {
		fmt.Printf("SayHello Output 1: Error   - %s\n", err1)
	} else {
		fmt.Printf("SayHello Output 1: Success - %#v\n", *helloResponse1)
	}

	fmt.Printf("SayHello Input 2: %#v\n", *helloInput2)
	if err2 != nil {
		fmt.Printf("SayHello Output 2: Error   - %s\n", err2)
	} else {
		fmt.Printf("SayHello Output 2: Success - %#v\n", *helloResponse2)
	}

	CustomListInput := &pb.CustomList{List: []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
	CustomListOuput, errList := client.Randomise(ctx, CustomListInput)
	fmt.Printf("Randomise Input: %#v\n", *CustomListInput)
	if err != nil {
		fmt.Printf("Randomise Output: Error - %s\n", errList)
	} else {
		fmt.Printf("Randomise Output: %#v\n", *CustomListOuput)
	}
}

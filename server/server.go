package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"net"

	"google.golang.org/grpc"
	// this alias's the package name to be pb (personal preference)
	pb "github.com/zanven42/corevision-HelloWorld"
)

type server struct{}

func (s server) SayHello(ctx context.Context, in *pb.SayHelloInput) (out *pb.SayHelloOutput, err error) {
	out = &pb.SayHelloOutput{}
	if in.ForceError {
		err = fmt.Errorf(in.DesiredErrorMessage)
		return
	}
	out.Msg = fmt.Sprintf("Hello %s %s", in.FirstName, in.LastName)
	return
}

func (s server) Randomise(ctx context.Context, in *pb.CustomList) (out *pb.CustomList, err error) {
	out = &pb.CustomList{}

	// put the incoming list of items into a random order in the output.
	for len(in.List) > 0 {
		if len(in.List) > 1 {
			i := rand.Intn(len(in.List))
			out.List = append(out.List, in.List[i])
			in.List = append(in.List[:i], in.List[i+1:]...)
		} else {
			out.List = append(out.List, in.List[0])
			in.List = in.List[:0]
		}
	}
	return
}

func main() {
	addr := flag.String("addr", "localhost:8000", "set the address to dial for a GRPC connection, e.g `localhost:8080`")
	h := flag.Bool("h", false, "Help")
	flag.Parse()
	if *h {
		flag.Usage()
		return
	}
	grpcServer := grpc.NewServer()
	// register our struct/class which needs to match the interface required
	pb.RegisterHelloWorldServer(grpcServer, &server{})
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("GRPC Listening on %s\n", *addr)
	fmt.Println(grpcServer.Serve(lis))
}

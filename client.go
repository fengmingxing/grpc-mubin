package main

import (
    "log"
    "os"

    pb "grpc-mubin/helloworld"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
)

const (
    address = "foo.bar.com:443"
)

func main() {
    address2:= os.Args[1]
    conn, err := grpc.Dial(address2, grpc.WithInsecure())

    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }

    defer conn.Close()

    c := pb.NewGreeterClient(conn)

    name := "test"
    if len(os.Args) > 1 {
	name=os.Args[1]
    }

    r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})

    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Println(r.Message)
    rtime, err := c.PrintTime(context.Background(), &pb.TimeRequest{Name: name})

    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Println(rtime.Message)
}


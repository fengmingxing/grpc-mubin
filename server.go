package main

import (
    "log"
    "net"
    "time"
    pb "grpc-mubin/helloworld"
    "google.golang.org/grpc"
    "golang.org/x/net/context"
)

const (
    PORT = ":50001"
)

type server struct {}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    log.Println("request: ", in.Name)
    return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) PrintTime(ctx context.Context, in *pb.TimeRequest) (*pb.TimeReply, error) {
    log.Println("request: ", in.Name)
    return &pb.TimeReply{Message: "Hello " + in.Name + " --> Server Time:"+ time.Now().Format("2006-01-02 15:04:05")}, nil
}

func main() {
    lis, err := net.Listen("tcp", PORT)

    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterGreeterServer(s, &server{})
    log.Println("rpc服务已经开启")
    s.Serve(lis)
}

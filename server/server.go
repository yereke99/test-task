package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	pb "test-task/proto"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct{}

func search(data string) string {
	if err := godotenv.Load("config.env"); err != nil {
		log.Fatalf("Error: %s", err)
	}
	serp_api_url := os.Getenv("SERPAPI")

	url := fmt.Sprintf(serp_api_url, data)
	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

func (s *Server) Search(ctx context.Context, r *pb.Request) (*pb.Reply, error) {
	log.Println("From proxy server: ", r.GetMsg())
	return &pb.Reply{Msg: search(r.GetMsg())}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	pb.RegisterServiceServer(srv, &Server{})
	reflection.Register(srv)

	if e := srv.Serve(listen); e != nil {
		panic(e)
	}
}

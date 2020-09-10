package main

import (
	"net"
	"log"
	//"context"
	pb"submit/submit"

	"io"
	"google.golang.org/grpc"//go mod vendor,then red become green
)


const (
	port = ":50051"
	tcp = "tcp"
)



//protoc --go_out=plugins=grpc:. ./helloworld.proto
//import{
//"context"
//pb "gRPC/helloworld"
// }
type StreamServer struct {

}
func (s *StreamServer)SubmitTransaction(trans pb.Submit_SubmitTransactionServer) error {
	n:=1
	for {
		req, err := trans.Recv()
		if err == io.EOF{
			return nil
		}
		if err != nil {
			return err
		}
		err = trans.Send(&pb.StreamResponse{
			Answer:true,
		})
		if err !=nil{
			return err
		}
		n++
		payload := req.Payload
		log.Printf("Received transaction from stream client, channel:%s  payload:%s",req.ChannelID, string(payload))
	}
}



func main() {
	lis, err := net.Listen(tcp, port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println(port + "net.Listening...")
	s := grpc.NewServer()
	pb.RegisterSubmitServer(s, &StreamServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


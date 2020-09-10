package main

import (
	"google.golang.org/grpc"
	pb "submit/submit"
	"log"
	"context"
	"io"
)

const (
	address     = "localhost:50051"
)

//var streamClient pb.SubmitClient

func send(streamClient pb.SubmitClient){
	stream, err := streamClient.SubmitTransaction(context.Background())
	if err != nil{
		log.Fatal("get SubmitTransaction stream err: %v", err)
	}
	for n:=0; n < 5; n++ {
		err := stream.Send(&pb.StreamRequest{ChannelID: "mychannel", Payload: []byte{'a'}})
		if err != nil {
			log.Fatal("stream request err: %v", err)
		}
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("stream response err: %v", err)
		}
		// 打印返回值
		log.Println("answer is: ", res.Answer)
	}

	err = stream.CloseSend()
	if err != nil {
		log.Fatal("close stream err: %v", err)
	}

}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	streamClient := pb.NewSubmitClient(conn)
	send(streamClient)
}

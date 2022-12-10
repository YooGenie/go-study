package main

import (
	"context"
	userpb "github.com/beautiful-store/grpc/v0/sunny"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

const gRPCServerPortNumber = "9000"
const portNumber = "7000"

func main() {
	// 게이트 만들어 go
	//get -t github.com/grpc-ecosystem/grpc-gateway/v2/runtime
	ctx := context.Background()
	mux := runtime.NewServeMux()

	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials())}

	// 엔드포인트로 함수하나가 자동으로 만들어 aPI(7000)가 오면 주소로 게이트웨이가 grpc(9000)를 연결해준다.
	if err := userpb.RegisterUserHandlerFromEndpoint(ctx, mux,
		"localhost:"+gRPCServerPortNumber,
		options,
	); err != nil {
		log.Fatalf("failed to register gRPC gateway: %v", err)
	}

	//http로 받기때문에 http로 열었다. 우리는 에코로 스타트 하는 거처럼
	// http로 받을거라서 알려준다.
	log.Printf("start HTTP server on %s port", portNumber)
	if err := http.ListenAndServe(":"+portNumber, mux); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

// 2022-12-06 데이터를 스트림으로 받는걸 한다. url string 받아서 주소로 가서 파일 받아서 바이트로 변경한다.
// 프론트에서는 바이트로 받으면 조금씩 받으니까

//http.get를 하면 자동으로 파일을 받는다.
//make 10으로 받으면 10바이트씩 나눠서 준다.


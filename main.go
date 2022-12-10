//package main
//
//import "github.com/YooGenie/go-study/basic"
//
//func main() {
//	//compare.CompareIfWithSwitch(5)
//	//fmt.Println("")
//	//function.Trim("          유지니             ")
//	//fmt.Println("")
//	//basic.SetStruct()
//	//fmt.Println("")
//	//basic.EmbeddedField()
//	//fmt.Println("")
//	//basic.StudyIf("red", 29)
//	//basic.StudySwitch("sunday",28)
//	//basic.StudyFor()
//	//basic.StudyForRange()
//	//basic.StudyArray()
//	//function.CheckType()
//	//basic.SliceStudy()
//	//function.FillZero(12365)
//	//function.Random()
//	//function.SliceStringAndConvertInt()
//	//basic.Map()
//	//function.ListAndMap()
//	//function.Join()
//	//basic.Channel()
//	//basic.SolutionOne()
//	//basic.SolutionTwo()
//	//basic.StudyWithTimeout()
//	basic.StringToStruct()
//}

package main

import (
	"context"
	"fmt"
	"github.com/YooGenie/go-study/data"
	userpb "github.com/beautiful-store/grpc/v0/sunny"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"net/http"
)

const portNumber = "9000"

type userServer struct {
	userpb.UserServer
}

// GetUser returns user message by user_id
func (s *userServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	userID := req.UserId

	var userMessage *userpb.UserMessage
	for _, u := range data.UserData {
		if u.UserId != userID {
			continue
		}
		userMessage = u
		break
	}

	fmt.Println(userMessage)
	return &userpb.GetUserResponse{
		UserMessage: userMessage,
	}, nil

}

type streamServer struct {
	userpb.StreamServiceServer
}

func (*streamServer) GetURL(in *userpb.GetURLRequest, srv userpb.StreamService_GetURLServer) error {
	url := in.Url
	fmt.Println("a : ", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("dd : ", resp.Body)

	defer resp.Body.Close()

	p := make([]byte, 1000)
	for i := 0; ; i++ {
		n, err := resp.Body.Read(p)
		//fmt.Printf("%d : %d 바이트 : data: %s", i, n, p[:n])
		fmt.Printf("%d", i)
		if err == io.EOF {
			break
		} else if err != nil {
			break
		}
		r := userpb.GetURLResponse{Content: p[:n]}
		if err := srv.Send(&r); err != nil {
			log.Println(err)
		}
	}

	fmt.Println("ok")
	return nil

}

// ListUsers returns all user messages
func (s *userServer) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	userMessages := make([]*userpb.UserMessage, len(data.UserData))
	for i, u := range data.UserData {
		userMessages[i] = u
	}

	return &userpb.ListUsersResponse{
		UserMessages: userMessages,
	}, nil
}

func main() {

	//gRPC server
	lis, err := net.Listen("tcp", ":"+portNumber) //포트를 열었다 //go에서 net를 만들었다. //네트워크 만들어놓은것
	// 뭐로 들어올지 모르니까 net 포트를 열어 놓는다.
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()                       //지알피씨 서버를 만들었다.
	userpb.RegisterUserServer(grpcServer, &userServer{}) //라운터 기능이랑 비슷하다 //서비스를 만들었다.
	userpb.RegisterStreamServiceServer(grpcServer, &streamServer{})
	// 지알피씨를 만들어놓은 controller

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil { //지알피에 연결해준다. //9000번을 연결 시켜라
		log.Fatalf("failed to serve: %s", err)
	}

}

//2022-12-06 "github.com/grpc-ecosystem/grpc-gateway/v2/runtime" go get

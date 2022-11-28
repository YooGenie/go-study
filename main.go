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
	"log"
	"net"

	"google.golang.org/grpc"
)

const portNumber = "9000"

func main() {
<<<<<<< Updated upstream
	//compare.CompareIfWithSwitch(5)
	//fmt.Println("")
	//function.Trim("          유지니             ")
	//fmt.Println("")
	//basic.SetStruct()
	//fmt.Println("")
	//basic.EmbeddedField()
	//fmt.Println("")
	//basic.StudyIf("red", 29)
	//basic.StudySwitch("sunday",28)
	//basic.StudyFor()
	//basic.StudyForRange()
	//basic.StudyArray()
	//function.CheckType()
	//basic.SliceStudy()
	//function.FillZero(12365)
	//function.Random()
	//function.SliceStringAndConvertInt()
	//basic.Map()
	//function.ListAndMap()
	//function.Join()
	//basic.Channel()
	//basic.SolutionOne()
	//basic.SolutionTwo()
	//basic.StudyWithTimeout()
	basic.StudyWithDeadline()
=======
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
>>>>>>> Stashed changes
}

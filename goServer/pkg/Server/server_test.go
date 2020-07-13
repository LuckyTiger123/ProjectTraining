package Server

import (
	"context"
	"fmt"
	pb "goServer/proto"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	t1 := time.Now()
	conn, err := grpc.Dial("39.98.125.235:8099", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
	}
	c := pb.NewScheduleServiceClient(conn)
	r, err := c.NewsSearch(context.Background(), &pb.NewsSearchRequest{
		Keyword: "任天堂",
		Size_:   100,
	})
	fmt.Println(time.Since(t1))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, v := range r.Result {
			fmt.Println(v.Source)
		}
	}
}

func TestFunc(t *testing.T) {
	Init()
	filter := make([]*pb.Filter, 0)
	//filter = append(filter,&pb.Filter{
	//	Type: "type",
	//	Value: "动作",
	//})
	result, err := GetGameSearch("rockstar", filter, 10)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, v := range result {
			fmt.Println(v.Source)
		}
	}
	//reg := regexp.MustCompile(`[0-9,的]`)
	//fmt.Println(reg.ReplaceAllString("鬼泣的第五5", ""))
}

func TestWord(t *testing.T) {
	fmt.Println(GetGameSearchWord("鬼泣5"))
}

func TestTime(t *testing.T) {
	Init()
	t1 := time.Now()
	//filter = append(filter,&pb.Filter{
	//	Type: "type",
	//	Value: "动作",
	//})
	result, err := GetResourceIndex(10, 1)
	fmt.Println(time.Since(t1))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, v := range result {
			fmt.Println(v)
		}
	}
	//reg := regexp.MustCompile(`[0-9,的]`)
	//fmt.Println(reg.ReplaceAllString("鬼泣的第五5", ""))
}

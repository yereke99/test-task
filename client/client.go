package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test-task/dto"
	pb "test-task/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var dataReq map[string]interface{}
var dataReply map[string]string

func main() {
	// gRPC client connection to gRPC server
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewServiceClient(conn)

	// Gin router
	r := gin.Default()
	r.GET("/get", func(ctx *gin.Context) {
		var data dto.DataRequest
		ctx.BindJSON(&data)
		in, _ := json.Marshal(data)
		json.Unmarshal(in, &dataReq)
		fmt.Println(dataReq)

		if resp, err := client.Search(ctx, &pb.Request{Msg: data.Data}); err == nil {
			dataReply = map[string]string{
				"status":  "HTTP status of 3rd-party service response",
				"headers": "headers array from 3rd-party service response",
				"result":  resp.GetMsg(),
			}
			ctx.JSON(200, gin.H{"message": dto.DataResponse{
				Status:  "HTTP status of 3rd-party service response",
				Headers: "headers array from 3rd-party service response",
				Result:  resp.GetMsg(),
			}})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
	r.Run(":8080")

}

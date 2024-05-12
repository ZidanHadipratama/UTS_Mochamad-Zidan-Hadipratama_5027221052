package main

import (
	"context"
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"strconv"

	pb "github.com/ZidanHadipratama/UTS_Mochamad-Zidan-Hadipratama_5027221052/pcmgmt" // replace with your protobuf package
)

type Processor struct {
	Id           int
	Name         string
	Manufacturer string
	Generation   int
	Core         int
	Thread       int
}

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPRCServicesClient(conn)

	r := gin.Default()

	r.Static("/assets", "/home/ika/Documents/Projects/IS/pcmgmt-grpc/client/templates/assets")

	r.LoadHTMLGlob("/home/ika/Documents/Projects/IS/pcmgmt-grpc/client/templates/pages/*")

	r.GET("/", func(c *gin.Context) {
		// List all processors
		response, err := client.ListPRCs(context.Background(), &emptypb.Empty{})
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to list PRCs")
			return
		}

		c.HTML(http.StatusOK, "dashboard.html", response.Prcs)
	})

	r.POST("/delete", func(c *gin.Context) {
		// Get the id from the query parameters
		id, _ := strconv.Atoi(c.Query("id"))

		// Delete the processor
		_, err := client.DeletePRC(context.Background(), &pb.PRCRequest{Id: int32(id)})
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to delete PRC")
			return
		}
	})

	r.GET("/update", func(c *gin.Context) {
		// Get the id from the query parameters
		id, _ := strconv.Atoi(c.Query("id"))

		// Get the current processor
		prc, err := client.ReadPRC(context.Background(), &pb.PRCRequest{Id: int32(id)})
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to get PRC")
			return
		}

		// Render the update form
		c.HTML(http.StatusOK, "update.html", prc)
	})

	r.POST("/update", func(c *gin.Context) {
		// Parse form data
		id, _ := strconv.Atoi(c.PostForm("id"))
		name := c.PostForm("name")
		manufacturer := c.PostForm("manufacturer")
		generation, _ := strconv.Atoi(c.PostForm("generation"))
		core, _ := strconv.Atoi(c.PostForm("core"))
		thread, _ := strconv.Atoi(c.PostForm("thread"))

		// Update the processor
		prc := &pb.PRC{Id: int32(id), Name: name, Manufacturer: manufacturer, Generation: int32(generation), Core: int32(core), Thread: int32(thread)}
		_, err := client.UpdatePRC(context.Background(), prc)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to update PRC")
			return
		}

		// Redirect to the list page
		c.Redirect(http.StatusSeeOther, "/")
	})

	r.GET("/add", func(c *gin.Context) {
		// Render the add form
		c.HTML(http.StatusOK, "add.html", nil)
	})

	r.POST("/add", func(c *gin.Context) {
		// Parse form data
		name := c.PostForm("name")
		manufacturer := c.PostForm("manufacturer")
		generation, _ := strconv.Atoi(c.PostForm("generation"))
		core, _ := strconv.Atoi(c.PostForm("core"))
		thread, _ := strconv.Atoi(c.PostForm("thread"))

		// Add the new processor
		prc := &pb.PRC{Name: name, Manufacturer: manufacturer, Generation: int32(generation), Core: int32(core), Thread: int32(thread)}
		_, err := client.CreatePRC(context.Background(), prc)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to add PRC")
			return
		}

		// Redirect to the list page
		c.Redirect(http.StatusSeeOther, "/")
	})

	log.Fatal(r.Run(":8080"))
}

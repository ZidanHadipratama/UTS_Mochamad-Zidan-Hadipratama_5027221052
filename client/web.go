package main

import (
	"context"
	"net/http"
	"log"

	// "os"
	// "fmt"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"html/template"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// List all processors
		response, err := client.ListPRCs(context.Background(), &emptypb.Empty{})
		if err != nil {
			http.Error(w, "Failed to list PRCs", http.StatusInternalServerError)
			return
		}

		// wd, err := os.Getwd()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println(wd)


		// Render the response in a new HTML page
		tmpl, err := template.ParseFiles("/home/ika/Documents/Projects/IS/pcmgmt-grpc/client/templates/list.html")
		if err != nil {
			log.Fatalf("Error parsing template: %v", err)
		}
		
		tmpl.Execute(w, response.Prcs)
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Get the id from the query parameters
			id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	
			// Delete the processor
			_, err := client.DeletePRC(context.Background(), &pb.PRCRequest{Id: int32(id)})
			if err != nil {
				http.Error(w, "Failed to delete PRC", http.StatusInternalServerError)
				return
			}
		}
	})

	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Parse form data
			r.ParseForm()
			id, _ := strconv.Atoi(r.Form.Get("id"))
			name := r.Form.Get("name")
			manufacturer := r.Form.Get("manufacturer")
			generation, _ := strconv.Atoi(r.Form.Get("generation"))
			core, _ := strconv.Atoi(r.Form.Get("core"))
			thread, _ := strconv.Atoi(r.Form.Get("thread"))
		
			// Update the processor
			prc := &pb.PRC{Id: int32(id), Name: name, Manufacturer: manufacturer, Generation: int32(generation), Core: int32(core), Thread: int32(thread)}
			_, err := client.UpdatePRC(context.Background(), prc)
			if err != nil {
				http.Error(w, "Failed to update PRC", http.StatusInternalServerError)
				return
			}
		
			// Redirect to the list page
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			// Get the id from the query parameters
			id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		
			// Get the current processor
			prc, err := client.ReadPRC(context.Background(), &pb.PRCRequest{Id: int32(id)})
			if err != nil {
				http.Error(w, "Failed to get PRC", http.StatusInternalServerError)
				return
			}
		
			// Render the update form
			tmpl, err := template.ParseFiles("/home/ika/Documents/Projects/IS/pcmgmt-grpc/client/templates/update.html")
			if err != nil {
				log.Fatalf("Error parsing template: %v", err)
			}
			tmpl.Execute(w, prc)
		}
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Parse form data
			r.ParseForm()
			name := r.Form.Get("name")
			manufacturer := r.Form.Get("manufacturer")
			generation, _ := strconv.Atoi(r.Form.Get("generation"))
			core, _ := strconv.Atoi(r.Form.Get("core"))
			thread, _ := strconv.Atoi(r.Form.Get("thread"))
	
			// Add the new processor
			prc := &pb.PRC{Name: name, Manufacturer: manufacturer, Generation: int32(generation), Core: int32(core), Thread: int32(thread)}
			_, err := client.CreatePRC(context.Background(), prc)
			if err != nil {
				http.Error(w, "Failed to add PRC", http.StatusInternalServerError)
				return
			}
	
			// Redirect to the list page
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			// Render the add form
			tmpl, err := template.ParseFiles("/home/ika/Documents/Projects/IS/pcmgmt-grpc/client/templates/add.html")
			if err != nil {
				log.Fatalf("Error parsing template: %v", err)
			}
			tmpl.Execute(w, nil)
		}
	})
	
	
	
	
	

	log.Fatal(http.ListenAndServe(":8080", nil))
}

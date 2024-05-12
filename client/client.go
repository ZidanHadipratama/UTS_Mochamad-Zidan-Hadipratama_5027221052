package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/ZidanHadipratama/UTS_Mochamad-Zidan-Hadipratama_5027221052/pcmgmt" // replace with your protobuf package
)

func main() {
	// Define flags for command-line arguments
	id := flag.Int("id", 0, "The id of the processor")
	name := flag.String("name", "", "Name of the processor")
	manufacturer := flag.String("manufacturer", "", "Manufacturer of the processor")
	generation := flag.Int("generation", 0, "Generation of the processor")
	core := flag.Int("core", 0, "Number of cores in the processor")
	thread := flag.Int("thread", 0, "Number of threads in the processor")
	list := flag.Bool("list", false, "List all processors")
	update := flag.Bool("update", false, "Update a processor")
	delete := flag.Bool("delete", false, "Delete a processor")

	// Parse the command-line arguments
	flag.Parse()

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
  
	client := pb.NewPRCServicesClient(conn)

	if *delete {
		// Check if id is provided
		if *id == 0 {
			fmt.Println("Usage: client -delete -id=<id>")
			os.Exit(1)
		}

		oldPRC, err := client.ReadPRC(context.Background(), &pb.PRCRequest{Id: int32(*id)})
		if err != nil {
			log.Fatalf("Failed to read PRC: %v", err)
		}
		log.Printf("The PRC: %v", oldPRC)

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Are you sure you want to update this PRC? (yes/no): ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "yes" {
			// Delete the processor
			response, err := client.DeletePRC(context.Background(), &pb.PRCRequest{Id: int32(*id)})
			if err != nil {
				log.Fatalf("Failed to delete PRC: %v", err)
			}
		
			log.Printf("Response from server: %v", response.Message)
		} else {
			fmt.Println("Delete cancelled.")
		}
	}else if *update {
		// Check if flags are provided
		if *id == 0 {
			fmt.Println("Usage: client -update -id=<id> -name=<name> -manufacturer=<manufacturer> -generation=<generation> -core=<core> -thread=<thread>")
			os.Exit(1)
		}

		prc := &pb.PRC{Id: int32(*id), Name: *name, Manufacturer: *manufacturer, Generation: int32(*generation), Core: int32(*core), Thread: int32(*thread)}
		log.Printf("PRC ID: %d, Name: %s, Manufacturer: %s, Generation: %d, Core: %d, Thread: %d",
				prc.Id, prc.Name, prc.Manufacturer, prc.Generation, prc.Core, prc.Thread)

		// Get the old data
		oldPRC, err := client.ReadPRC(context.Background(), &pb.PRCRequest{Id: int32(*id)})
		if err != nil {
			log.Fatalf("Failed to read PRC: %v", err)
		}
		log.Printf("Old PRC: %v", oldPRC)

		// Ask for confirmation
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Are you sure you want to update this PRC? (yes/no): ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "yes" {
			// Update the processor
			log.Printf("PRC ID: %d, Name: %s, Manufacturer: %s, Generation: %d, Core: %d, Thread: %d",
				prc.Id, prc.Name, prc.Manufacturer, prc.Generation, prc.Core, prc.Thread)
			response, err := client.UpdatePRC(context.Background(), prc)
			if err != nil {
				log.Fatalf("Failed to update PRC: %v", err)
			}
		
			log.Printf("Response from server: %v", response.Message)
			log.Printf("Updated PRC: %v", prc)
		} else {
			fmt.Println("Update cancelled.")
		}
	}else if *list {
		// List all processors
		response, err := client.ListPRCs(context.Background(), &emptypb.Empty{})
		if err != nil {
			log.Fatalf("Failed to list PRCs: %v", err)
		}

		for _, prc := range response.Prcs {
			log.Printf("PRC ID: %d, Name: %s, Manufacturer: %s, Generation: %d, Core: %d, Thread: %d",
				prc.Id, prc.Name, prc.Manufacturer, prc.Generation, prc.Core, prc.Thread)
		}
	} else {
		// Check if flags are provided
		if *name == "" || *manufacturer == "" || *generation == 0 || *core == 0 || *thread == 0 || *id == 0 {
			fmt.Println("Usage: client -id=<id> -name=<name> -manufacturer=<manufacturer> -generation=<generation> -core=<core> -thread=<thread>")
			os.Exit(1)
		}

		// Add a new processor
		prc := &pb.PRC{Id: int32(*id), Name: *name, Manufacturer: *manufacturer, Generation: int32(*generation), Core: int32(*core), Thread: int32(*thread)}
		response, err := client.CreatePRC(context.Background(), prc)
		if err != nil {
			log.Fatalf("Failed to create PRC: %v", err)
		}
	  
		log.Printf("Response from server: %v", response.Message)
		log.Printf("Created PRC: %v", prc)
	}
}

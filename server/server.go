package main

import (
	"context"
	"log"
	"os"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/ZidanHadipratama/UTS_Mochamad-Zidan-Hadipratama_5027221052/pcmgmt" // replace with your protobuf package
)

type server struct {
	pb.UnimplementedPRCServicesServer
	collection *mongo.Collection
	logger     *log.Logger
}

func (s *server) CreatePRC(ctx context.Context, req *pb.PRC) (*pb.PRCResponse, error) {
	s.logger.Printf("Creating PRC: %v", req)
	_, err := s.collection.InsertOne(ctx, req)
	if err != nil {
		s.logger.Printf("Failed to create PRC: %v", err)
		return nil, err
	}
	s.logger.Printf("PRC created successfully: %v", req)
	return &pb.PRCResponse{Message: "PRC created successfully"}, nil
}

func (s *server) ListPRCs(ctx context.Context, req *emptypb.Empty) (*pb.ListPRCsResponse, error) {
	s.logger.Println("Listing all PRCs")

	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		s.logger.Printf("Failed to list PRCs: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var prcs []*pb.PRC
	for cursor.Next(ctx) {
		var prc pb.PRC
		if err := cursor.Decode(&prc); err != nil {
			return nil, err
		}
		prcs = append(prcs, &prc)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.ListPRCsResponse{Prcs: prcs}, nil
}

func (s *server) ReadPRC(ctx context.Context, req *pb.PRCRequest) (*pb.PRC, error) {
	s.logger.Printf("Reading PRC: %v", req)

	filter := bson.M{"id": req.GetId()} // Change "id" to "Id"
	var prc pb.PRC
	err := s.collection.FindOne(ctx, filter).Decode(&prc)
	if err != nil {
		s.logger.Printf("Failed to read PRC: %v", err)
		return nil, err
	}

	s.logger.Printf("PRC read successfully: %v", prc)
	return &prc, nil
}

func (s *server) UpdatePRC(ctx context.Context, req *pb.PRC) (*pb.PRCResponse, error) {
	s.logger.Printf("Updating PRC: %v", req)

	filter := bson.M{"id": req.GetId()} // Change "id" to "Id"
	update := bson.M{"$set": req}

	_, err := s.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		s.logger.Printf("Failed to update PRC: %v", err)
		return nil, err
	}

	s.logger.Printf("PRC updated successfully: %v", req)
	return &pb.PRCResponse{Message: "PRC updated successfully"}, nil
}

func (s *server) DeletePRC(ctx context.Context, req *pb.PRCRequest) (*pb.PRCResponse, error) {
	s.logger.Printf("Deleting PRC: %v", req)

	filter := bson.M{"id": req.GetId()} // Change "id" to "Id"
	deleteResult, err := s.collection.DeleteOne(ctx, filter)
	if err != nil {
		s.logger.Printf("Failed to delete PRC: %v", err)
		return nil, err
	}

	if deleteResult.DeletedCount == 0 {
		return &pb.PRCResponse{Message: "No PRC found with the given id"}, nil
	}

	s.logger.Printf("PRC deleted successfully: %v", req)
	return &pb.PRCResponse{Message: "PRC deleted successfully"}, nil
}


func main() {
	// Set up MongoDB connection.
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	collection := client.Database("pcmgmt").Collection("processor")

	// Set up logging to a file.
	logFile, err := os.OpenFile("prc.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	logger := log.New(logFile, "PRC: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Set up gRPC server.
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPRCServicesServer(s, &server{collection: collection, logger: logger})
	if err := s.Serve(lis); err != nil {
		logger.Fatalf("Failed to serve: %v", err)
	}
}

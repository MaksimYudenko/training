package main

import (
	"context"
	"flag"
	"github.com/MaksimYudenko/training/finalTask/pkg/api/v1"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

func main() {
	// get configuration
	address := flag.String("server", "",
		"gRPC server in format host:port")
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := v1.NewTraineeServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Create
	req1 := v1.CreateRequest{
		Api: apiVersion,
		Trainee: &v1.Trainee{
			Name:      "goName",
			Surname:   "goSurname",
			Email:     "goEmail",
			HasAttend: true,
			Score:     100,
		},
	}
	res1, err := c.Create(ctx, &req1)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res1)
	id := res1.Id
	// Read
	req2 := v1.ReadRequest{
		Api: apiVersion,
		Id:  id,
	}
	res2, err := c.Read(ctx, &req2)
	if err != nil {
		log.Fatalf("Read failed: %v", err)
	}
	log.Printf("Read result: <%+v>\n\n", res2)
	// Update
	req3 := v1.UpdateRequest{
		Api: apiVersion,
		Trainee: &v1.Trainee{
			Id:        res2.Trainee.Id,
			Name:      res2.Trainee.Name + " + updated",
			Surname:   res2.Trainee.Surname + " + updated",
			Email:     res2.Trainee.Email + " + updated",
			HasAttend: res2.Trainee.HasAttend,
			Score:     res2.Trainee.Score,
		},
	}
	res3, err := c.Update(ctx, &req3)
	if err != nil {
		log.Fatalf("Update failed: %v", err)
	}
	log.Printf("Update result: <%+v>\n\n", res3)
	// ReadAll
	req4 := v1.ReadAllRequest{
		Api: apiVersion,
	}
	res4, err := c.ReadAll(ctx, &req4)
	if err != nil {
		log.Fatalf("ReadAll failed: %v", err)
	}
	log.Printf("ReadAll result: <%+v>\n\n", res4)
	// Delete
	req5 := v1.DeleteRequest{
		Api: apiVersion,
		Id:  id,
	}
	res5, err := c.Delete(ctx, &req5)
	if err != nil {
		log.Fatalf("Delete failed: %v", err)
	}
	log.Printf("Delete result: <%+v>\n\n", res5)
}

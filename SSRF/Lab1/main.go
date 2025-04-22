package main

import (
	"cetinboran/secretpb"
	"context"
	"fmt"
	"log"
	"net"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

// Developed by Çetin Boran Mesüm

type server struct {
	secretpb.UnimplementedSecretServiceServer
}

func (s *server) GetSecret(ctx context.Context, req *secretpb.SecretRequest) (*secretpb.SecretResponse, error) {
	return &secretpb.SecretResponse{Data: "This is the secret!"}, nil
}

func setupGRPCServer() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	secretpb.RegisterSecretServiceServer(grpcServer, &server{})

	fmt.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	app := fiber.New()

	app.Post("/fetch", func(c *fiber.Ctx) error {
		type Body struct {
			URL string `json:"url"`
		}

		var body Body
		if err := c.BodyParser(&body); err != nil {
			fmt.Println("JSON Parse Error:", err)
			return c.Status(400).SendString("Bad request: " + err.Error())
		}

		grpcURL, err := url.Parse(body.URL)
		if err != nil {
			return c.Status(500).SendString("Invalid Link " + err.Error())
		}
		conn, err := grpc.NewClient(grpcURL.Host, grpc.WithInsecure())
		if err != nil {
			return c.Status(500).SendString("Error connecting to gRPC server: " + err.Error())
		}
		defer conn.Close()

		client := secretpb.NewSecretServiceClient(conn)

		secretResp, err := client.GetSecret(context.Background(), &secretpb.SecretRequest{})
		if err != nil {
			return c.Status(500).SendString("Error calling gRPC service: " + err.Error())
		}

		return c.SendString(fmt.Sprintf("gRPC Secret: %s", secretResp.GetData()))
	})

	go func() {
		if err := app.Listen(":3001"); err != nil {
			log.Fatalf("Error starting HTTP server: %v", err)
		}
	}()

	go setupGRPCServer()

	select {}
}

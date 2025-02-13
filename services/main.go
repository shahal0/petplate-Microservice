package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"petplate-auth/models"
	"petplate-auth/proto/userpb"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type userServiceServer struct {
	userpb.UnimplementedUserServiceServer
}

var db *gorm.DB

func init() {
	var err error
	dsn := "host=localhost user=postgres password=1 dbname=petplate_auth port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}

func (s *userServiceServer) GetUserById(ctx context.Context, req *userpb.UserRequest) (*userpb.UserResponse, error) {
	var user models.User

	if err := db.WithContext(ctx).First(&user, req.UserId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to fetch user: %v", err)
	}

	return &userpb.UserResponse{
		Id:    uint64(user.ID),
		Name:  user.FirstName + " " + user.LastName,
		Email: user.Email,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, &userServiceServer{})

	log.Println("User Service running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

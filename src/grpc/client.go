package grpc

import (
	"fmt"
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/keithyw/pbuf-services/protobufs"
	"github.com/keithyw/kyw-go-docker-test/conf"
	"github.com/keithyw/kyw-go-docker-test/models"
)

type Client struct {
	config *conf.Config
	Conn *grpc.ClientConn
	client protobufs.UserClient
}

func NewGrpcClient(config *conf.Config) *Client {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(config.GrpcHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf(err.Error())
		return nil
	}
	log.Println("GRPC client connected!")
	client := protobufs.NewUserClient(conn)
	return &Client{config, conn, client}
}

func (c *Client) CreateUser(user *models.User) {
	res, err := c.client.SaveUser(context.Background(), &protobufs.UserMessage{Username: &user.Username, UserId: user.ID})
	if err != nil {
		log.Println(fmt.Sprintf("Failed saving user through grpc service %s", err))
		return
	}
	log.Println(fmt.Sprintf("GRPC SaveUser Response: %s", res.GetMsg()))
}

func (c *Client) DeleteUser(id int) {
	res, err := c.client.DeleteUser(context.Background(), &protobufs.UserMessage{UserId: int64(id)})
	if err != nil {
		log.Println(fmt.Sprintf("Failed deleting user through grpc service %s", err))
		return
	}
	log.Println(fmt.Sprintf("GRPC DeleteUser Response: %s", res.GetMsg()))
}

func (c *Client) UpdateUser(id int, user *models.User) {
	res, err := c.client.UpdateUser(context.Background(), &protobufs.UserMessage{UserId: int64(id), Username: &user.Username})
	if err != nil {
		log.Println(fmt.Sprintf("Failed updating user through grpc service %s", err))
		return
	}
	log.Println(fmt.Sprintf("GRPC UpdateUser Response: %s", res.GetMsg()))
}
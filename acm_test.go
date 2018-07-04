package goacm

import (
	"testing"
	"log"
	"os"
	"fmt"
)

func getClient() *Client {
	client, err := NewClient(func(c *Client) {
		c.AccessKey = os.Getenv("AccessKey")
		c.SecretKey = os.Getenv("SecretKey")
		c.EndPoint = "acm.aliyun.com"
		c.NameSpace = os.Getenv("NameSpace")
	})

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func RunWithSchema(t *testing.T, test func(client *Client, t *testing.T)) {
	client := getClient()
	defer func() {
		client.Delete("test", "test")
	}()

	_, err := client.Publish("test", "test", "test")

	if err != nil {
		t.Fatalf("pulish error:%s", err)
	}

	test(client, t)
}

func TestNewClient(t *testing.T) {
	client := getClient()
	servers := client.GetServers()

	if len(servers) == 0 {
		t.Error("get server error")
	}
}

func TestClient_GetConfig(t *testing.T) {
	RunWithSchema(t, func(client *Client, t *testing.T) {
		ret, err := client.GetConfig("test", "test")
		if err != nil {
			t.Error(err)
		}
		fmt.Println(ret)
	})
}

func TestClient_Subscribe(t *testing.T) {
	RunWithSchema(t, func(client *Client, t *testing.T) {
		_, err := client.Subscribe("test", "test","")
		if err != nil {
			t.Error(err)
		}
	})
}

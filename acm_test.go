package goacm

import (
	"testing"
	"fmt"
	"log"
	"os"
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

func TestNewClient(t *testing.T) {
	client := getClient()
	servers := client.GetServers()

	if len(servers) == 0 {
		t.Error("get server error")
	}

	fmt.Println(servers)
}

func TestClient_GetConfig(t *testing.T) {
	client := getClient()
	ret, err := client.GetConfig("php-test", "payment")

	if err != nil {
		t.Error(err)
	}

	fmt.Println(ret)
}

func TestClient_Subscribe(t *testing.T) {
	client := getClient()
	ret, err := client.Subscribe("php-test", "payment", "")

	if err != nil {
		t.Error(err)
	}

	fmt.Println(ret)
}

# Aliyun ACM SDK for Golang
<a href="https://travis-ci.org/goapt/goacm"><img src="https://travis-ci.org/goapt/goacm.svg" alt="Build Status"></a>
<a href="https://godoc.org/github.com/goapt/goacm"><img src="https://godoc.org/github.com/goapt/goacm?status.svg" alt="GoDoc"></a>

Aliyun acm sdk for golang, support multiple IP

## Usage

```go
import "github.com/goapt/goacm"

client, err := goacm.NewClient(func(c *Client) {
    c.AccessKey = "******"
    c.SecretKey = "******"
    c.EndPoint  = "acm.aliyun.com"
    c.NameSpace = "******"
})

if err != nil {
    log.Fatal(err)
}

//get config
ret, err := client.GetConfig("test", "DEFAULT_GROUP")

//Subscribe
ret, err := client.Subscribe("test", "DEFAULT_GROUP","")

//publish config
ret, err := client.Publish("test", "DEFAULT_GROUP","test")

//remove config
ret, err := client.Delete("test", "DEFAULT_GROUP")
```


## License
The SDK is open-sourced software licensed under the MIT license.

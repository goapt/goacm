# Aliyun ACM SDK for Golang
<a href="https://travis-ci.org/verystar/goacm"><img src="https://travis-ci.org/verystar/goacm.svg" alt="Build Status"></a>

aliyun acm sdk for golang

## Usage

```go
import "github.com/verystar/goacm"

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
package main

import (
	"fmt"
	"github.com/mediocregopher/radix.v2/cluster"
	//"github.com/mediocregopher/radix.v2/redis"
)

func main() {
	c, err := cluster.New("192.168.1.71:7000")

	key := "liqing"
	client, err := c.GetForKey(key)
	fmt.Println(err)
	//client, err := redis.Dial("tcp", "192.168.1.71:7000")
	//fmt.Println(c.GetEvery())

	// Checking Err field directly

	err = client.Cmd("SET", key, "1987", "EX", 3600).Err
	fmt.Println(err)
	foo, err := client.Cmd("GET", key).Str()
	if err != nil {
		// handle err
	}
	fmt.Println("foo is:", foo)
}

package main

import (
	"fmt"
	"log"

	// Import the redigo/redis package.
	"github.com/gomodule/redigo/redis"
)

// Define a custom struct to hold Album data. Notice the struct tags?
// These indicate to redigo how to assign the data from the reply into
// the struct.
type Album struct {
	Title  string  `redis:"title"`
	Artist string  `redis:"artist"`
	Price  float64 `redis:"price"`
	Likes  int     `redis:"likes"`
}

func main() {
	// Establish a connection to the Redis server listening on port
	// 6379 of the local machine. 6379 is the default port, so unless
	// you've already changed the Redis configuration file this should
	// work.
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}

	// Importantly, use defer to ensure the connection is always
	// properly closed before exiting the main() function.
	defer conn.Close()

	// Fetch all album fields with the HGETALL command. Wrapping this
	// in the redis.Values() function transforms the response into type
	// []interface{}, which is the format we need to pass to
	// redis.ScanStruct() in the next step.
	values, err := redis.Values(conn.Do("HGETALL", "album:2"))
	if err != nil {
		log.Fatal(err)
	}

	// Create an instance of an Album struct and use redis.ScanStruct()
	// to automatically unpack the data to the struct fields. This uses
	// the struct tags to determine which data is mapped to which
	// struct fields.
	var album Album
	err = redis.ScanStruct(values, &album)
	if err != nil {
		log.Fatal()
	}

	fmt.Printf("%+v\n", album)
}
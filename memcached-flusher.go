package main

import (
	"log"
	"os"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("no arguments")
	}
	mc := memcache.New(os.Args[1:]...)
	for {
		err := mc.FlushAll()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Millisecond * 20)
	}
}

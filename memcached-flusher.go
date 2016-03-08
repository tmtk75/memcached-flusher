package main

import "github.com/bradfitz/gomemcache/memcache"
import "time"
import "log"
import "os"

func main() {

	mc := memcache.New(os.Args[1:]...)
	for {
		err := mc.FlushAll()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Duration(20) * time.Millisecond)
	}
}

package main

import (
	"log"
	"time"

	"github.com/shifengbin/doggie"
)

func main() {
	provider := NewHttpProvider("http://aabb.com/aa")
	dog := doggie.New()
	dog.SetRemoteProvider(provider, false)
	// dog.SetConfigFile("../testdata/a.yml")
	dog.ReadConfig()
	err := dog.WatchConfig(func() {
		e := dog.Get("d.e")
		log.Println("更新", e.Idx(e.Len()-1))
	})

	if err != nil {
		panic(err)
	}

	log.Println(dog.Get("a").String())
	go func() {
		for {
			e := dog.Get("d.e")
			log.Println("获取", e.Idx(e.Len()-1))
		}
	}()
	time.Sleep(time.Minute)
}

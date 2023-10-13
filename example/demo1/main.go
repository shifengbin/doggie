package main

import (
	"log"

	"github.com/shifengbin/doggie"
)

type Worker struct {
	Name           string
	MaxWorkers     int
	MaxIdleWorkers int
	MaxQueueSize   int
	WorkerWeight   []int
}

type NamedWorker struct {
	A string    `doggie:"name"`
	B int       `doggie:"maxWorkers"`
	C int       `doggie:"maxIdleWorkers"`
	D int       `doggie:"maxQueueSize"`
	E []float32 `doggie:"workerWeight"`
}

func main() {
	dog := doggie.New()
	dog.SetConfigFile("./app.yaml")
	dog.ReadConfig()

	//直接获取配置
	log.Println("httpPort", dog.Get("httpPort").Int())
	log.Println("grpcPort", dog.Get("httpPort").Int())

	//通过默认规则获取
	worker := Worker{}
	dog.Get("worker").Unmarshal(&worker)
	log.Println(worker)

	//通过tag获取
	named := NamedWorker{}
	dog.Get("worker").Unmarshal(&named)
	log.Println(named)

	//直接获取多级
	weight := dog.Get("worker.workerWeight")
	for i := 0; i < weight.Len(); i++ {
		log.Println("wight", i, weight.Idx(i).Int())
	}

	//直接遍历对象类型
	w := dog.Get("worker")
	keys := w.Keys()

	for _, key := range keys {
		log.Println(key, w.Get(key).String())
	}

	//解码一个切片
	var uints []uint
	if err := dog.Get("worker.workerWeight").Unmarshal(&uints); err != nil {
		log.Println(err)
	}
	log.Println(uints)
}

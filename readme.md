# 🐶 Doggie

Doggy是一个配置管理工具，支持多种配置源，支持多种配置格式，支持多种配置读取方式，支持多种配置监控方式

## 快速开始

```go
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

}
```

## 🎉 特性

配置读取：
    1.读取文件配置
    2.读取远程配置
    3.读取环境变量

编码：
    1.自定义解码方法
    2.内置解码

监控：
    1.文件监控
    2.自定义监控方法(远程)

查找顺序：
    1. 命令行参数
    2. 环境变量
    3. 配置文件
    4. 远程配置
    5. 默认值

其他：
    1.设置默认值


注意：

自定义解码器转换的数据结构应为map[string]interface{}, 数组需要满足[]interface{}类型

如何实现远程配置，请参考example/remote.go

**不提供**：

    写配置文件，应用程序不应修改配置文件
    
    别名，别名会影响程序的维护，在程序中到处改key会导致无法与配置文件对应，难以理解和维护

## 思路

通过Get方法获取转换器Caster, Caster负责转换数据和Unmarshal数据

## 文档

### 包级别方法

1. RegisterDecoder 注册自定义解码器

### Doggie

1. SetConfigFile 设置配置文件
2. SetFileType 设置配置文件类型，用于解码，如果不设置则通过SetConfigFile设置文件的扩展名获取
3. SetRemoteType 设置远程配置文件类型，用于解码
4. SetRemoteProvider 设置远程配置提供者
5. SetDefault 设置默认值
6. ReadConfig 读取配置文件，配置好上述设置后，需要调用该函数进行配置文件读取
7. WatchConfig 监控文件和远程配置变换，没有配置则不会监控对应类型配置
8. GetEnv() Caster 从环境变量中获取配置
9. GetConfig() Caster 从配置文件中获取配置
10. GetRemoteConfig() Caster 从远程配置中获取配置
11. GetFlag() Caster 获取命令行参数配置
12. Get() Caster 获取配置项，通过优先级顺序获取flag > env > config > remote > default

### Caster

1. Int() 把值转换为int
2. Int8() 把值转换为int8
3. Int16() 把值转换为int16
4. Int32() 把值转换为int32
5. Int64() 把值转换为int64
6. Uint() 把值转换为uint
7. Uint8() 把值转换为uint8
8. Uint16() 把值转换为uint16
9. Uint32() 把值转换为uint32
10. Uint64() 把值转换为uint64
11. String() 把值转换为string
12. Bool() 把值转换为bool
13. Float32() 把值转换为float32
14. Float64() 把值转换为float64
15. Keys() 获取map key
16. Get(key string) Caster 获取map key对应的值
17. Unmarshal(obj interface{}) error 把值转换为struct
18. Len() 获取map或者slice的length
19. Idx(idx int) Caster 获取slice对应值


## 例子：

### flag

支持格式：

1. -a
2. -a 10
3. -a=10
4. --a
5. --a 10
6. --a=10


### 环境变量

读取系统环境变量

### 文件

```go
    //创建配置实例
    dog := doggie.New()
    //设置文件路径
	dog.SetConfigFile("./testdata/a.yml")
    //读取配置
	dog.ReadConfig()

    //获取配置项
	if v := dog.Get("a").Int(); v != 10 {
		t.Log("Get a Want 10, but:", v)
		t.Fail()
	}

    //d.e是一个数组，可以通过.Idx获取下标对应的值
    if vv := v.Get("d.e").Idx(1).Int(); vv != 2 {
		t.Log("Get d.e[1] Want 2, but:", vv)
		t.Fail()
	}
```

### 远程配置

创建远程配置提供者
```go
type HttpProvider struct {
	URL string
}

func (h HttpProvider) SetUp() ([]byte, error) {
	f, err := os.Open("../testdata/real.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return io.ReadAll(f)
}

type Config struct {
	A int    `json:"a"`
	B string `json:"b"`
	D struct {
		E []int `json:"e"`
	} `json:"d"`
}

func (h HttpProvider) Watch() ([]byte, error) {
	// time.Sleep(2 * time.Second)
	data, err := h.SetUp()
	if err != nil {
		return nil, err
	}

	c := Config{}
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}

	c.D.E = append(c.D.E, int(time.Now().Unix()))

	return json.Marshal(c)
}

func NewHttpProvider(url string) *HttpProvider {
	return &HttpProvider{
		URL: url,
	}
}
```

注册到配置中
```go

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
```

### 通过结构体解析
默认解析名称为字段名首字母小写，如果自定义请使用tag, `doggie:"xxx"`
```go
//下面的会解析key为name,age的配置
type DefaultNameConfig struct {
	Name string
	Age  int
}

//下面可以自定义, Name, Age
type CustomNameConfig struct {
	Name string `doggie:"Name"`
	Age  int `doggie:"Age"`
}
```
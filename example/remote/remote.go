package main

import (
	"encoding/json"
	"io"
	"os"
	"time"
)

/*
type RemoteProvider interface {
	//初始化执行，读取远程配置
	SetUp() ([]byte, error)

	//检测远程配置
	Watch() ([]byte, error)
}
*/

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

package doggie

import (
	"fmt"
	"io"
	"log"
	"os"
	"testing"
	"time"

	"gopkg.in/yaml.v3"
)

func TestDoggie_SetDefault(t *testing.T) {
	dog := New()
	dog.SetDefault("a", "b")
	dog.SetDefault("c", 1)
	dog.SetDefault("d.e", 1.1)
	dog.SetDefault("d.f", []interface{}{1, 2, 3})

	if a := dog.Get("a").String(); a != "b" {
		t.Log("Get a Want b, but:", a)
		t.Fail()
	}

	if c := dog.Get("c").Int(); c != 1 {
		t.Log("Get c Want 1, but:", c)
		t.Fail()
	}

	if d := dog.Get("d").Obj("e").Float64(); d != 1.1 {
		t.Log("Get d Want 1.1, but:", d)
		t.Fail()
	}

	if d := dog.Get("d.e").Float64(); d != 1.1 {
		t.Log("Get d.e Want 1.1, but:", d)
		t.Fail()
	}

	if d := dog.Get("d.e").Int(); d != 1 {
		t.Log("Get d.e Want 1, but:", d)
		t.Fail()
	}

	df := dog.Get("d.f")
	for i := 0; i < df.Len(); i++ {
		if v := df.Idx(i).Int(); v != i+1 {
			t.Logf("Get d.f[%d] Want %d, but: %d", i, i+1, v)
			t.Fail()
		}
	}
}

func TestConfig(t *testing.T) {
	dog := New()
	dog.SetConfigFile("./testdata/a.yml")
	dog.ReadConfig()
	if v := dog.Get("a").Int(); v != 10 {
		t.Log("Get a Want 10, but:", v)
		t.Fail()
	}

	if v := dog.Get("b").String(); v != "aaa" {
		t.Log("Get b Want aaa, but:", v)
		t.Fail()
	}

	if v := dog.Get("c").String(); v != "bbb" {
		t.Log("Get c Want bbb, but:", v)
		t.Fail()
	}

	v := dog.Get("d")
	if vv := v.String(); vv != "" {
		t.Log("Get d Want empty, but:", vv)
		t.Fail()
	}
	if vv := v.Obj("e").Idx(1).Int(); vv != 2 {
		t.Log("Get d.e[1] Want 2, but:", vv)
		t.Fail()
	}

}

func TestConfigType(t *testing.T) {
	dog := New()
	dog.SetConfigFile("./testdata/a.json")
	dog.SetFileType("yml")
	dog.ReadConfig()
	if v := dog.Get("a").Int(); v != 10 {
		t.Log("Get a Want 10, but:", v)
		t.Fail()
	}

	if v := dog.Get("b").String(); v != "aaa" {
		t.Log("Get b Want aaa, but:", v)
		t.Fail()
	}

	if v := dog.Get("c").String(); v != "bbb" {
		t.Log("Get c Want bbb, but:", v)
		t.Fail()
	}

	if dog.Get("c").String() != dog.GetConfig("c").String() {
		t.Errorf("Get c Want %s, but: %s", dog.GetConfig("c").String(), dog.Get("c").String())
	}

	v := dog.Get("d")
	if vv := v.String(); vv != "" {
		t.Log("Get d Want empty, but:", vv)
		t.Fail()
	}
	if vv := v.Obj("e").Idx(1).Int(); vv != 2 {
		t.Log("Get d.e[1] Want 2, but:", vv)
		t.Fail()
	}

	t.Log(v)

	if v := dog.Get("d.e").Idx(1).Int(); v != 2 {
		t.Log("Get d.e[1] Want 2, but:", v)
		t.Fail()
	}

}

func TestConfigJSON(t *testing.T) {
	dog := New()
	dog.SetConfigFile("./testdata/real.json")
	dog.ReadConfig()
	if v := dog.Get("a").Int(); v != 10 {
		t.Log("Get a Want 10, but:", v)
		t.Fail()
	}

	if v := dog.Get("b").String(); v != "aaa" {
		t.Log("Get b Want aaa, but:", v)
		t.Fail()
	}

	if v := dog.Get("c").String(); v != "bbb" {
		t.Log("Get c Want bbb, but:", v)
		t.Fail()
	}

	v := dog.Get("d")
	if vv := v.String(); vv != "" {
		t.Log("Get d Want empty, but:", vv)
		t.Fail()
	}
	if vv := v.Obj("e").Idx(1).Int(); vv != 2 {
		t.Log("Get d.e[1] Want 2, but:", vv)
		t.Fail()
	}

	t.Log(v)

	if v := dog.Get("d.e").Idx(1).Int(); v != 2 {
		t.Log("Get d.e[1] Want 2, but:", v)
		t.Fail()
	}

}

func TestConfigTOML(t *testing.T) {
	dog := New()
	dog.SetConfigFile("./testdata/a.toml")
	dog.ReadConfig()
	if v := dog.Get("a.a").Int(); v != 10 {
		t.Log("Get a Want 10, but:", v)
		t.Fail()
	}

	if v := dog.Get("a.b").String(); v != "bbb" {
		t.Log("Get b Want aaa, but:", v)
		t.Fail()
	}

	if v := dog.Get("a.c").Idx(1).Int(); v != 2 {
		t.Log("Get d.e[1] Want 2, but:", v)
		t.Fail()
	}

}

type HttpProvider struct {
	URL string
}

func (h HttpProvider) SetUp() ([]byte, error) {
	f, err := os.Open("./testdata/a.yml")
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
	if err := yaml.Unmarshal(data, &c); err != nil {
		return nil, err
	}

	c.D.E = append(c.D.E, 10)
	// time.Sleep(500 * time.Millisecond)
	return yaml.Marshal(c)
}

func NewHttpProvider(url string) *HttpProvider {
	return &HttpProvider{
		URL: url,
	}
}

func TestConfigRemote(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	dog := New()
	dog.SetConfigFile("./testdata/a.yml")
	dog.SetRemoteType("yml")
	dog.SetRemoteProvider(NewHttpProvider(""), false)
	dog.ReadConfig()
	dog.WatchConfig(nil)

	time.Sleep(1 * time.Second)
	e := dog.GetRemoteConfig("d.e").(*Cast)
	if v := e.Idx(e.Len() - 1).Int(); v != 10 {
		t.Log("Get d.e Want 10, but:", v, e.Value())
		t.Fail()
	}
	time.Sleep(1 * time.Second)

}

type configUnmarshal struct {
	A int    `doggie:"a"`
	B string `doggie:"b"`
	C string `doggie:"c"`
	D struct {
		E []int `doggie:"e"`
	} `doggie:"d"`
}

func TestFlats(t *testing.T) {
	os.Args = append(os.Args, "-a=10")
	dog := New()
	dog.ReadConfig()
	if dog.GetFlag("a").Int() != 10 {
		t.Error("GetFlag a not equal:", dog.GetFlag("a").Int())
	}
}

type remoteProvider struct {
}

func (remoteProvider) SetUp() ([]byte, error) {
	return nil, fmt.Errorf("format, a")
}

func (remoteProvider) Watch() ([]byte, error) {
	return nil, fmt.Errorf("123123")
}

func TestSetConfigFileErr(t *testing.T) {
	dog := New()

	if err := dog.SetConfigFile("./testdata/"); err == nil {
		t.Errorf("SetConfigFile err: %v", err)
	}

	if err := dog.SetConfigFile("./testdata/aabb.c"); err == nil {
		t.Errorf("SetConfigFile err: %v", err)
	}

	dog.SetRemoteProvider(remoteProvider{}, false)
	if err := dog.ReadConfig(); err == nil {
		t.Errorf("ReadConfig err: %v", err)
	}

	dog.SetRemoteType("a")
	if err := dog.ReadConfig(); err == nil {
		t.Errorf("ReadConfig err: %v", err)
	}

	dog.configFile = "./testdata/aabb.c"
	if err := dog.ReadConfig(); err == nil {
		t.Errorf("ReadConfig err: %v", err)
	}

	os.Args = append(os.Args, "a=10")
	if err := dog.ReadConfig(); err == nil {
		t.Errorf("ReadConfig err: %v", err)
	}

}

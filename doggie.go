package doggie

import (
	"doggie/internal/decode"
	"doggie/internal/flags"
	"io"
	"log"
	"path/filepath"

	"fmt"
	"os"
	"strings"
	"sync/atomic"

	"github.com/fsnotify/fsnotify"
)

type RemoteProvider interface {
	//初始化执行，读取远程配置
	SetUp() ([]byte, error)

	//检测远程配置
	Watch() ([]byte, error)
}

type Doggie struct {
	//保存命令行参数
	flags map[string]string

	//保存默认值
	defaults map[string]interface{}

	//保存配置信息 内部也是一个map[string]interface{}, 因为需要监控变化所以需要，防止同时读写导致panic
	config atomic.Value

	//远程配置信息 内部也是一个map[string]interface{}
	remoteConfig atomic.Value

	//配置文件地址
	configFile string

	//配置文件类型，decode使用，不设置默认使用配置文件后缀
	fileType string

	//远程配置文件格式，用来decode使用, 默认json
	remoteType string

	//名称分隔符， 默认为.
	sep string

	//远程更新全部，默认为false 为增量更新
	refreshAll bool

	remoteProvider RemoteProvider
}

func New() *Doggie {
	d := &Doggie{
		flags:        make(map[string]string),
		defaults:     make(map[string]interface{}),
		config:       atomic.Value{},
		remoteConfig: atomic.Value{},
		remoteType:   "json",
		sep:          ".",
	}
	d.config.Store(make(map[string]interface{}))
	d.remoteConfig.Store(make(map[string]interface{}))
	return d
}

// SetConfigFile 设置配置文件地址
func (d *Doggie) SetConfigFile(fpath string) error {
	fp, err := os.Stat(fpath)
	if err != nil {
		return err
	}

	if fp.IsDir() {
		return fmt.Errorf("not support dir: %s", fpath)
	}

	d.configFile = fpath

	if d.fileType == "" {
		d.fileType = strings.TrimPrefix(filepath.Ext(fpath), ".")
	}

	return nil
}

// SetFileType 设置配置文件类型
func (d *Doggie) SetFileType(fileType string) {
	d.fileType = fileType
}

// SetRemoteType 设置远程配置类型
func (d *Doggie) SetRemoteType(remoteType string) {
	d.remoteType = remoteType
}

// SetRemoteProvider 设置远程配置更新者
// refreshAll 是否为全量更新， true为全量， false为增量
func (d *Doggie) SetRemoteProvider(provider RemoteProvider, refreshAll bool) {
	d.remoteProvider = provider
	d.refreshAll = refreshAll
}

// 设置默认值
func (d *Doggie) SetDefault(key string, value interface{}) {
	path := strings.Split(key, d.sep)
	if len(path) == 1 {
		d.defaults[key] = value
		return
	}
	d.defaults[key] = value
	m := deepSearch(d.defaults, path[:len(path)-1])
	m[path[len(path)-1]] = value
}

// ReadConfig 读取配置
func (d *Doggie) ReadConfig() error {
	if err := d.readFlags(); err != nil {
		return err
	}

	if err := d.readConfigFile(); err != nil {
		return err
	}

	if err := d.readRemoteConfig(); err != nil {
		return err
	}
	return nil

}

func (d *Doggie) WatchConfig(f func()) error {
	if err := d.watchFile(f); err != nil {
		return err
	}

	if err := d.watchRemote(f); err != nil {
		return err
	}

	return nil
}

func (d *Doggie) watchFile(f func()) error {
	if d.configFile == "" {
		return nil
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println("error: watch fail", err)
		return err
	}

	if err := watcher.Add(d.configFile); err != nil {
		log.Println("error: add file fail", err)
		return err
	}

	go func() {
		for evt := range watcher.Events {
			if evt.Has(fsnotify.Remove) {
				break
			}

			if !evt.Has(fsnotify.Write) {
				continue
			}

			if err := d.readConfigFile(); err != nil {
				log.Println("error: watch file error:", err)
				continue
			}

			if f == nil {
				continue
			}
			f()
		}
	}()
	return nil
}

func (d *Doggie) watchRemote(f func()) error {
	if d.remoteProvider == nil {
		return nil
	}
	go func() {
		for {
			data, err := d.remoteProvider.Watch()
			if err != nil {
				log.Println("error: watch remote config error:", err)
				continue
			}
			dec := decode.Get(d.remoteType)
			if dec == nil {
				log.Println("error: unsupport remote type:", d.remoteType)
				break
			}
			m := make(map[string]interface{})
			if err := dec.Decode(data, &m); err != nil {
				log.Println("error: decode error:", err)
				continue
			}

			//全量更新
			if d.refreshAll {
				d.remoteConfig.Store(m)
				if f != nil {
					f()
				}
				continue
			}

			org := d.remoteConfig.Load().(map[string]interface{})
			org = deepCopy(org)
			merge(m, org)
			d.remoteConfig.Store(org)

			if f == nil {
				continue
			}
			f()
		}

	}()
	return nil
}

func (d *Doggie) readFlags() error {
	var err error
	d.flags, err = flags.Parse(os.Args[1:])
	return err
}

func (d *Doggie) readConfigFile() error {
	if d.configFile == "" {
		return nil
	}

	f, err := os.Open(d.configFile)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	dec := decode.Get(d.fileType)
	if dec == nil {
		return fmt.Errorf("unsupported remote config type: %s", d.remoteType)
	}

	m := make(map[string]interface{})
	if err := dec.Decode(data, &m); err != nil {
		return err
	}

	d.config.Store(m)

	return nil
}

func (d *Doggie) readRemoteConfig() error {
	if d.remoteProvider == nil {
		return nil
	}

	data, err := d.remoteProvider.SetUp()
	if err != nil {
		return err
	}

	// 获取解码器，解析配置
	dec := decode.Get(d.remoteType)
	if dec == nil {
		return fmt.Errorf("unsupported remote config type: %s", d.remoteType)
	}
	m := make(map[string]interface{})
	if err := dec.Decode(data, &m); err != nil {
		return err
	}

	d.remoteConfig.Store(m)
	return nil
}

func (d *Doggie) GetEnv(key string) Caster {
	return NewCast(os.Getenv(key))
}

func (d *Doggie) GetConfig(key string) Caster {
	path := strings.Split(key, d.sep)
	conf := d.config.Load().(map[string]interface{})
	val := searchMap(conf, path)
	return NewCast(val)
}

func (d *Doggie) GetRemoteConfig(key string) Caster {
	path := strings.Split(key, d.sep)
	conf := d.remoteConfig.Load().(map[string]interface{})
	val := searchMap(conf, path)
	return NewCast(val)
}

func (d *Doggie) GetFlag(key string) Caster {
	return NewCast(d.flags[key])
}

// Get 获取配置
// 优先级：
//
//	flag > env > config > remote > default
func (d *Doggie) Get(key string) Caster {
	path := strings.Split(key, d.sep)
	//获取flag
	if v, ok := d.flags[key]; ok {
		return NewCast(v)
	}

	//获取环境变量
	if v, ok := os.LookupEnv(key); ok {
		return NewCast(v)
	}

	//获取配置文件
	conf := d.config.Load().(map[string]interface{})
	if v := searchMap(conf, path); v != nil {
		return NewCast(v)
	}

	//获取远程配置
	remote := d.remoteConfig.Load().(map[string]interface{})
	if v := searchMap(remote, path); v != nil {
		return NewCast(v)
	}

	//获取默认值
	return NewCast(searchMap(d.defaults, path))
}

// RegisterDecoder 注册解码器
func RegisterDecoder(name string, dec decode.Decoder) {
	decode.Register(name, dec)
}

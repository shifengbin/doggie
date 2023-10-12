package doggie

import (
	"fmt"
	"reflect"

	"github.com/spf13/cast"
)

type Caster interface {
	// Int returns the value as an int.
	Int() int
	// Int8 returns the value as an int8.
	Int8() int8
	// Int16 returns the value as an int16.
	Int16() int16
	// Int32 returns the value as an int32.
	Int32() int32
	// Int64 return the value as an int64
	Int64() int64
	// Uint return the value as an uint
	Uint() uint
	// Int8 returns the value as an uint8.
	Uint8() uint8
	// Int16 returns the value as an uint16.
	Uint16() uint16
	// Int32 returns the value as an uint32.
	Uint32() uint32
	// Int64 return the value as an uint64
	Uint64() uint64

	String() string
	Bool() bool
	Float32() float32
	Float64() float64

	//返回map key
	Keys() []string
	//获取map key对应的值
	Get(key string) Caster
	//map转struct
	Unmarshal(obj interface{}) error

	//获取map或者slice的length
	Len() int

	//获取slice对应值
	Idx(idx int) Caster
}

var _ Caster = &Cast{}

type Cast struct {
	value any
}

func NewCast(value any) *Cast {
	return &Cast{value: value}
}

func (c *Cast) Int() int {
	return cast.ToInt(c.value)
}

func (c *Cast) Int8() int8 {
	return cast.ToInt8(c.value)
}

func (c *Cast) Int16() int16 {
	return cast.ToInt16(c.value)
}

func (c *Cast) Int32() int32 {
	return cast.ToInt32(c.value)
}

func (c *Cast) Int64() int64 {
	return cast.ToInt64(c.value)
}

func (c *Cast) Uint() uint {
	return cast.ToUint(c.value)
}

func (c *Cast) Uint8() uint8 {
	return cast.ToUint8(c.value)
}

func (c *Cast) Uint16() uint16 {
	return cast.ToUint16(c.value)
}

func (c *Cast) Uint32() uint32 {
	return cast.ToUint32(c.value)
}

func (c *Cast) Uint64() uint64 {
	return cast.ToUint64(c.value)
}

func (c *Cast) String() string {
	return cast.ToString(c.value)
}

func (c *Cast) Float32() float32 {
	return cast.ToFloat32(c.value)
}

func (c *Cast) Float64() float64 {
	return cast.ToFloat64(c.value)
}

func (c *Cast) Bool() bool {
	return cast.ToBool(c.value)
}

// Keys 如果对象是map,返回所有key
func (c *Cast) Keys() []string {
	m := cast.ToStringMap(c.value)
	keys := []string{}
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// Obj 获取map key对应的值
func (c *Cast) Get(key string) Caster {
	m := cast.ToStringMap(c.value)
	return &Cast{m[key]}
}

func (c *Cast) Len() int {
	switch t := c.value.(type) {
	case map[string]interface{}:
		return len(t)
	case []interface{}:
		return len(t)
	default:
		return 0
	}
}

// Idx Slice下标
func (c *Cast) Idx(idx int) Caster {
	s := cast.ToSlice(c.value)
	if idx < len(s) {
		return &Cast{s[idx]}
	}
	return &Cast{}
}

func (c *Cast) Unmarshal(v interface{}) error {
	var err error
	switch t := c.value.(type) {
	case map[string]interface{}:
		err = unmarshal(t, v)
	default:
		err = fmt.Errorf("unsupport unmarsharl %s", reflect.TypeOf(v).Name())
	}
	return err
}

func (c *Cast) Value() any {
	return c.value
}

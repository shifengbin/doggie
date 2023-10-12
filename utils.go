package doggie

import (
	"fmt"
	"reflect"

	"github.com/spf13/cast"
)

// deepSearch 查不到会自动创建
func deepSearch(m map[string]any, path []string) map[string]any {
	for _, k := range path {
		m2, ok := m[k]
		if !ok {
			// intermediate key does not exist
			// => create it and continue from there
			m3 := make(map[string]any)
			m[k] = m3
			m = m3
			continue
		}
		m3, ok := m2.(map[string]any)
		if !ok {
			// intermediate key is a value
			// => replace with a new map
			m3 = make(map[string]any)
			m[k] = m3
		}
		// continue search from here
		m = m3
	}
	return m
}

// deepCopy 深度copy,
func deepCopy(m map[string]interface{}) map[string]interface{} {
	m2 := make(map[string]interface{})
	for k, v := range m {
		v1, ok := v.(map[string]interface{})
		if ok {
			m2[k] = deepCopy(v1)
			continue
		}
		s1, ok := v.([]interface{})
		if ok {
			sv := make([]interface{}, len(s1))
			copy(sv, s1)
			v = sv
		}
		m2[k] = v
	}
	return m2
}

func merge(src, des map[string]interface{}) {
	for k, v := range src {
		sv, ok := v.(map[string]interface{})
		if !ok { //不是对象，是基础类型，直接复制
			des[k] = v
			continue
		}

		dv, ok := des[k]
		if !ok { //不存在直接复制
			des[k] = deepCopy(sv)
			continue
		}

		dm, ok := dv.(map[string]interface{})
		if !ok { //不是对象，不用合并，直接复制
			des[k] = deepCopy(sv)
			continue
		}

		//源和目标都是对象，进行合并
		merge(sv, dm)
	}
}

// searchMap 查找map, 不创建新数据
func searchMap(m any, path []string) any {
	if len(path) == 0 {
		return nil
	}
	mm, ok := m.(map[string]any)
	if !ok {
		return nil
	}

	if len(path) == 1 {
		return mm[path[0]]
	}

	return searchMap(mm[path[0]], path[1:])
}

func unmarshal(data map[string]interface{}, v interface{}) error {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Ptr || vv.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("v must be a pointer to a struct")
	}
	vv = vv.Elem()

	tv := reflect.TypeOf(v).Elem()
	for i := 0; i < vv.NumField(); i++ {
		vf := vv.Field(i)
		if !vf.CanSet() {
			continue
		}
		tf := tv.Field(i)
		key := tf.Name
		tag := tf.Tag.Get("doggie")
		if tag != "" {
			key = tag
		}

		val, ok := data[key]
		if !ok {
			continue
		}

		//如果是对象，看下是否是结构体
		obj, ok := val.(map[string]interface{})
		if ok && tf.Type.Kind() == reflect.Struct {
			//如果是对象，递归解析
			if err := unmarshal(obj, vf.Addr().Interface()); err != nil {
				return err
			}
			continue
		}

		var setVal interface{}
		var err error
		switch vf.Interface().(type) {
		case int:
			setVal, err = cast.ToIntE(val)
		case int8:
			setVal, err = cast.ToInt8E(val)
		case int16:
			setVal, err = cast.ToInt16E(val)
		case int32:
			setVal, err = cast.ToInt32E(val)
		case int64:
			setVal, err = cast.ToInt64E(val)
		case uint:
			setVal, err = cast.ToUintE(val)
		case uint8:
			setVal, err = cast.ToUint8E(val)
		case uint16:
			setVal, err = cast.ToUint16E(val)
		case uint32:
			setVal, err = cast.ToUint32E(val)
		case uint64:
			setVal, err = cast.ToUint64E(val)
		case float32:
			setVal, err = cast.ToFloat32E(val)
		case float64:
			setVal, err = cast.ToFloat64E(val)
		case bool:
			setVal, err = cast.ToBoolE(val)
		case string:
			setVal, err = cast.ToStringE(val)
		default: // slice
			if vf.Kind() != reflect.Slice {
				err = fmt.Errorf("unmarshalSlice: v is not slice: %s", key)
				break
			}
			v, ok := val.([]interface{})
			if !ok {
				continue
			}
			setVal, err = unmarshalSlice(key, v, vf.Interface())

		}
		if err != nil {
			return err
		}
		vf.Set(reflect.ValueOf(setVal))
	}
	return nil
}

func unmarshalSlice(name string, data []interface{}, v any) (interface{}, error) {
	vf := reflect.TypeOf(v)

	if vf.Kind() != reflect.Slice {
		return nil, fmt.Errorf("unmarshalSlice: v is not a slice: %s (%s)", name, vf.Kind())
	}

	item := vf.Elem()

	var setVal interface{}
	var err error

	switch item.Kind() {
	case reflect.Int:
		setVal, err = ToXIntXSliceE[int](data)
	case reflect.Int8:
		setVal, err = ToXIntXSliceE[int8](data)
	case reflect.Int16:
		setVal, err = ToXIntXSliceE[int16](data)
	case reflect.Int32:
		setVal, err = ToXIntXSliceE[int32](data)
	case reflect.Int64:
		setVal, err = ToXIntXSliceE[int64](data)
	case reflect.Uint:
		setVal, err = ToXIntXSliceE[uint](data)
	case reflect.Uint8:
		setVal, err = ToXIntXSliceE[uint8](data)
	case reflect.Uint16:
		setVal, err = ToXIntXSliceE[uint16](data)
	case reflect.Uint32:
		setVal, err = ToXIntXSliceE[uint32](data)
	case reflect.Uint64:
		setVal, err = ToXIntXSliceE[uint64](data)
	case reflect.Float32:
		setVal, err = ToFloatXSliceE[float32](data)
	case reflect.Float64:
		setVal, err = ToFloatXSliceE[float64](data)
	case reflect.Bool:
		setVal, err = ToBoolSlice(data)
	case reflect.String:
		setVal, err = ToStringSlice(data)
	default:
		return nil, fmt.Errorf("not support: %s", item.Name())
	}

	return setVal, err
}

type XIntX interface {
	int |
		int8 |
		int16 |
		int32 |
		int64 |
		uint |
		uint8 |
		uint16 |
		uint32 |
		uint64
}

type FloatX interface {
	float32 |
		float64
}

func ToXIntXSliceE[T XIntX](data []interface{}) ([]T, error) {
	val, err := cast.ToIntSliceE(data)
	if err != nil {
		return nil, err
	}
	ret := make([]T, 0, len(val))

	var tmp T
	for _, i := range val {
		err = ToXIntXE(i, &tmp)
		if err != nil {
			return nil, err
		}
		ret = append(ret, tmp)
	}
	return ret, nil
}

func ToXIntXE(data interface{}, v interface{}) error {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Ptr || vv.IsNil() {
		return fmt.Errorf("v is nil")
	}
	val, err := cast.ToInt64E(data)
	if err != nil {
		return err
	}

	switch vv.Elem().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		vv.Elem().SetInt(val)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		vv.Elem().SetUint(uint64(val))
	}
	return nil
}

func ToFloatXSliceE[T FloatX](data interface{}) ([]T, error) {
	val := cast.ToSlice(data)

	ret := make([]T, 0, len(val))

	var tmp T
	for _, i := range val {
		if err := ToFloatX(i, &tmp); err != nil {
			return nil, err
		}
		ret = append(ret, tmp)
	}
	return ret, nil
}

func ToFloatX(data interface{}, v interface{}) error {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Ptr || vv.IsNil() {
		return fmt.Errorf("v is nil")
	}
	val, err := cast.ToFloat64E(data)
	if err != nil {
		return err
	}
	switch vv.Elem().Kind() {
	case reflect.Float32, reflect.Float64:
		vv.Elem().SetFloat(val)
	}
	return nil
}

func ToBoolSlice(data interface{}) ([]bool, error) {
	val := cast.ToSlice(data)

	ret := make([]bool, 0, len(val))

	for _, i := range val {
		b, err := cast.ToBoolE(i)
		if err != nil {
			return nil, err
		}
		ret = append(ret, b)
	}
	return ret, nil
}

func ToStringSlice(data interface{}) ([]string, error) {
	val := cast.ToSlice(data)

	ret := make([]string, 0, len(val))

	for _, i := range val {
		b, err := cast.ToStringE(i)
		if err != nil {
			return nil, err
		}
		ret = append(ret, b)
	}
	return ret, nil
}

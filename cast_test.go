package doggie

import (
	"reflect"
	"testing"
)

func TestCast_Int(t *testing.T) {
	type fields struct {
		value any
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "int",
			fields: fields{value: 1},
			want:   1,
		},
		{
			name:   "string",
			fields: fields{value: "1"},
			want:   1,
		},
		{
			name:   "float",
			fields: fields{value: 1.0},
			want:   1,
		},
		{
			name:   "bool",
			fields: fields{value: true},
			want:   1,
		},
		{
			name:   "nil",
			fields: fields{value: nil},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cast{
				value: tt.fields.value,
			}
			if got := c.Int(); got != tt.want {
				t.Errorf("Cast.Int() = %v, want %v", got, tt.want)
			}
			if got := c.Int8(); got != int8(tt.want) {
				t.Errorf("Cast.Int8() = %v, want %v", got, tt.want)
			}
			if got := c.Int16(); got != int16(tt.want) {
				t.Errorf("Cast.Int16() = %v, want %v", got, tt.want)
			}
			if got := c.Int32(); got != int32(tt.want) {
				t.Errorf("Cast.Int32() = %v, want %v", got, tt.want)
			}
			if got := c.Int64(); got != int64(tt.want) {
				t.Errorf("Cast.Int64() = %v, want %v", got, tt.want)
			}

			if got := c.Uint(); got != uint(tt.want) {
				t.Errorf("Cast.Uint() = %v, want %v", got, tt.want)
			}
			if got := c.Uint8(); got != uint8(tt.want) {
				t.Errorf("Cast.Uint8() = %v, want %v", got, tt.want)
			}
			if got := c.Uint16(); got != uint16(tt.want) {
				t.Errorf("Cast.Uint16() = %v, want %v", got, tt.want)
			}
			if got := c.Uint32(); got != uint32(tt.want) {
				t.Errorf("Cast.Uint32() = %v, want %v", got, tt.want)
			}
			if got := c.Uint64(); got != uint64(tt.want) {
				t.Errorf("Cast.Uint64() = %v, want %v", got, tt.want)
			}
			if got := c.Float32(); got != float32(tt.want) {
				t.Errorf("Cast.Float32() = %v, want %v", got, tt.want)
			}
			if got := c.Float64(); got != float64(tt.want) {
				t.Errorf("Cast.Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCast_Bool(t *testing.T) {
	type fields struct {
		value any
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "string",
			fields: fields{value: "true"},
			want:   true,
		},
		{
			name:   "int",
			fields: fields{value: 1},
			want:   true,
		},
		{
			name:   "float",
			fields: fields{value: 1.0},
			want:   true,
		},
		{
			name:   "bool",
			fields: fields{value: true},
			want:   true,
		},
		{
			name:   "nil",
			fields: fields{value: nil},
			want:   false,
		},
		{
			name:   "empty string",
			fields: fields{value: ""},
			want:   false,
		},
		{
			name:   "empty slice",
			fields: fields{value: []string{}},
			want:   false,
		},
		{
			name:   "empty map",
			fields: fields{value: map[string]string{}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cast{
				value: tt.fields.value,
			}
			if got := c.Bool(); got != tt.want {
				t.Errorf("Cast.Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCast_Keys(t *testing.T) {
	type fields struct {
		value any
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name:   "string",
			fields: fields{value: "true"},
			want:   []string{},
		},
		{
			name:   "int",
			fields: fields{value: 1},
			want:   []string{},
		},
		{
			name:   "float",
			fields: fields{value: 1.0},
			want:   []string{},
		},
		{
			name:   "bool",
			fields: fields{value: true},
			want:   []string{},
		},
		{
			name:   "nil",
			fields: fields{value: nil},
			want:   []string{},
		},
		{
			name:   "empty string",
			fields: fields{value: ""},
			want:   []string{},
		},
		{
			name:   "empty slice",
			fields: fields{value: []string{}},
			want:   []string{},
		},
		{
			name:   "empty map",
			fields: fields{value: map[string]interface{}{}},
			want:   []string{},
		},
		{
			name:   "map",
			fields: fields{value: map[string]interface{}{"a": "b"}},
			want:   []string{"a"},
		},
		{
			name:   "slice",
			fields: fields{value: []string{"a", "b"}},
			want:   []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cast{
				value: tt.fields.value,
			}
			if got := c.Keys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cast.Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCast_Obj(t *testing.T) {
	type fields struct {
		value any
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Cast
	}{
		{
			name:   "empty string",
			fields: fields{value: ""},
			args:   args{key: "a"},
			want:   &Cast{},
		},
		{
			name:   "map",
			fields: fields{value: map[string]interface{}{"a": 1}},
			args:   args{key: "a"},
			want:   &Cast{value: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cast{
				value: tt.fields.value,
			}
			if got := c.Obj(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cast.Obj() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCast_Len(t *testing.T) {
	type fields struct {
		value any
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "empty string",
			fields: fields{value: ""},
			want:   0,
		},
		{
			name:   "map",
			fields: fields{value: map[string]interface{}{"a": 1}},
			want:   1,
		},
		{
			name:   "slice",
			fields: fields{value: []interface{}{1, 2, 3}},
			want:   3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cast{
				value: tt.fields.value,
			}
			if got := c.Len(); got != tt.want {
				t.Errorf("Cast.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCast_Idx(t *testing.T) {
	type fields struct {
		value any
	}
	type args struct {
		idx int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Cast
	}{
		{
			name:   "empty string",
			fields: fields{value: ""},
			args:   args{idx: 0},
			want:   &Cast{},
		},
		{
			name:   "map",
			fields: fields{value: map[string]interface{}{"a": 1}},
			args:   args{idx: 0},
			want:   &Cast{},
		},
		{
			name:   "slice",
			fields: fields{value: []interface{}{1, 2, 3}},
			args:   args{idx: 0},
			want:   &Cast{value: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cast{
				value: tt.fields.value,
			}
			if got := c.Idx(tt.args.idx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cast.Idx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCast_String(t *testing.T) {
	type fields struct {
		value any
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "empty string",
			fields: fields{value: ""},
			want:   "",
		},
		{
			name:   "map",
			fields: fields{value: map[string]interface{}{"a": 1}},
			want:   "",
		},
		{
			name:   "slice",
			fields: fields{value: []interface{}{1, 2, 3}},
			want:   "",
		},
		{
			name: "int",
			fields: fields{
				value: 1,
			},
			want: "1",
		},
		{
			name: "float",
			fields: fields{
				value: 1.1,
			},
			want: "1.1",
		},
		{
			name: "bool",
			fields: fields{
				value: true,
			},
			want: "true",
		},
		{
			name: "string",
			fields: fields{
				value: "a",
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cast{
				value: tt.fields.value,
			}
			c.Value()
			if got := c.String(); got != tt.want {
				t.Errorf("Cast.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCast_Unmarshal(t *testing.T) {
	m := map[string]any{
		"a": 1,
		"b": int8(2),
		"c": int16(3),
		"d": int32(4),
		"e": int64(5),
		"f": uint(6),
		"g": uint8(7),
		"h": uint16(8),
		"i": uint32(9),
		"j": uint64(10),
		"k": "string",
		"l": true,
		"m": []interface{}{"1", 1.2, uint(3)},
		"n": map[string]interface{}{
			"a": 1,
			"b": int8(2),
		},
		"o": "1.2",
		"p": 3.4,
		"q": []interface{}{true, false, 1},
		"r": []interface{}{1.2, 3.4, "5.6"},
		"s": []interface{}{1.2, 3.4, "5.6"},
	}

	v := ConfigTest{}
	target := ConfigTest{
		A: 1,
		B: 2,
		C: 3,
		D: 4,
		E: 5,
		F: 6,
		G: 7,
		H: 8,
		I: 9,
		J: 10,
		K: "string",
		L: true,
		M: []int{1, 1, 3},
		N: struct {
			A int  "doggie:\"a\""
			B int8 "doggie:\"b\""
		}{
			A: 1,
			B: 2,
		},
		O: 1.2,
		P: 3.4,
		Q: []bool{true, false, true},
		R: []float32{1.2, 3.4, 5.6},
		S: []float64{1.2, 3.4, 5.6},
	}

	c := &Cast{
		value: m,
	}
	if err := c.Unmarshal(&v); err != nil {
		t.Errorf("Unmarshal() error = %v", err)
	}

	if !reflect.DeepEqual(v, target) {
		t.Errorf("unmarshal() = %v, want %v", v, target)
	}

	c = &Cast{}
	if err := c.Unmarshal(&v); err == nil {
		t.Error("Unmarshal() error = nil, want error")
	}
}

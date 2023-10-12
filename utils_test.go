package doggie

import (
	"reflect"
	"testing"
)

func Test_deepSearch(t *testing.T) {
	type args struct {
		m    map[string]any
		path []string
	}
	tests := []struct {
		name string
		args args
		want map[string]any
	}{
		{
			name: "test1",
			args: args{
				m: map[string]any{
					"a": map[string]any{
						"b": map[string]any{
							"c": "d",
						},
					},
				},
				path: []string{"a", "b"},
			},
			want: map[string]any{"c": "d"},
		},
		{
			name: "test2",
			args: args{
				m:    map[string]any{},
				path: []string{"a", "b"},
			},
			want: map[string]any{},
		},
		{
			name: "test3",
			args: args{
				m:    map[string]any{},
				path: []string{"a", "b"},
			},
			want: map[string]any{},
		}, {
			name: "test4",
			args: args{
				m: map[string]any{
					"a": map[string]any{
						"b": map[string]any{
							"c": "d",
						},
					},
				},
				path: []string{"a", "b", "c"},
			},
			want: map[string]any{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deepSearch(tt.args.m, tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deepSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deepCopy(t *testing.T) {
	type args struct {
		m map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "test1",
			args: args{
				m: map[string]interface{}{
					"a": map[string]interface{}{
						"b": map[string]interface{}{
							"c": "d",
						},
					},
				},
			},
			want: map[string]interface{}{
				"a": map[string]interface{}{
					"b": map[string]interface{}{
						"c": "d",
					},
				},
			},
		},
		{
			name: "test2",
			args: args{
				m: map[string]interface{}{},
			},
			want: map[string]interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deepCopy(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_merge(t *testing.T) {
	type args struct {
		src map[string]interface{}
		des map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "test1",
			args: args{
				src: map[string]interface{}{
					"a": map[string]interface{}{
						"b": map[string]interface{}{
							"c": "d",
						},
					},
				},
				des: map[string]interface{}{},
			},
			want: map[string]interface{}{
				"a": map[string]interface{}{
					"b": map[string]interface{}{
						"c": "d",
					},
				},
			},
		}, {
			name: "test2",
			args: args{
				src: map[string]interface{}{},
				des: map[string]interface{}{
					"a": map[string]interface{}{
						"b": map[string]interface{}{
							"c": "d",
						},
					},
				},
			},
			want: map[string]interface{}{
				"a": map[string]interface{}{
					"b": map[string]interface{}{
						"c": "d",
					},
				},
			},
		}, {
			name: "test3",
			args: args{
				src: map[string]interface{}{
					"a": map[string]interface{}{
						"b": map[string]interface{}{
							"c": map[string]interface{}{
								"c": "d",
							},
						},
					},
				},
				des: map[string]interface{}{
					"a": map[string]interface{}{
						"b": map[string]interface{}{
							"c": "d",
						},
					},
				},
			},
			want: map[string]interface{}{
				"a": map[string]interface{}{
					"b": map[string]interface{}{
						"c": map[string]interface{}{
							"c": "d",
						},
					},
				},
			},
		}, {
			name: "test4",
			args: args{
				src: map[string]interface{}{
					"a": map[string]interface{}{
						"b": map[string]interface{}{
							"c": "d",
						},
					},
				},
				des: map[string]interface{}{
					"a": map[string]interface{}{
						"b": map[string]interface{}{
							"c": map[string]interface{}{
								"c": "d",
							},
						},
					},
				},
			},
			want: map[string]interface{}{
				"a": map[string]interface{}{
					"b": map[string]interface{}{
						"c": "d",
					},
				},
			},
		}, {
			name: "test5",
			args: args{
				src: map[string]interface{}{
					"a": map[string]interface{}{
						"b": map[string]interface{}{
							"c": "d",
						},
					},
				},
				des: map[string]interface{}{},
			},
			want: map[string]interface{}{
				"a": map[string]interface{}{
					"b": map[string]interface{}{
						"c": "d",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.src, tt.args.des)
			if !reflect.DeepEqual(tt.args.des, tt.want) {
				t.Errorf("merge() = %v, want %v", tt.args.des, tt.want)
			}
		})
	}
}

func Test_searchMap(t *testing.T) {
	type args struct {
		m    map[string]any
		path []string
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "test1",
			args: args{
				m: map[string]any{
					"a": map[string]any{
						"b": map[string]any{
							"c": "d",
						},
					},
				},
				path: []string{"a", "b", "c"},
			},
			want: "d",
		},
		{
			name: "test2",
			args: args{
				m: map[string]any{
					"a": map[string]any{
						"b": map[string]any{
							"c": "d",
						},
					},
				},
				path: []string{"a", "b"},
			},
			want: map[string]any{
				"c": "d",
			},
		},
		{
			name: "test3",
			args: args{
				m: map[string]any{
					"a": map[string]any{
						"b": map[string]any{
							"c": "d",
						},
					},
				},
				path: []string{"a"},
			},
			want: map[string]any{
				"b": map[string]any{
					"c": "d",
				},
			},
		},
		{
			name: "test4",
			args: args{
				m: map[string]any{
					"a": map[string]any{
						"b": map[string]any{
							"c": "d",
						},
					},
				},
				path: []string{"a", "d"},
			},
			want: nil,
		},
		{
			name: "test5",
			args: args{
				m: map[string]any{
					"a": map[string]any{
						"b": map[string]any{
							"c": "d",
						},
					},
				},
				path: []string{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMap(tt.args.m, tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToStringSlice(t *testing.T) {
	type args struct {
		data []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				data: []interface{}{true, false},
			},
			want:    []string{"true", "false"},
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				data: []interface{}{1, 2},
			},
			want:    []string{"1", "2"},
			wantErr: false,
		},
		{
			name: "test3",
			args: args{
				data: []interface{}{"1", "2"},
			},
			want:    []string{"1", "2"},
			wantErr: false,
		},
		{
			name: "test4",
			args: args{
				data: []interface{}{1i, 2i},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToStringSlice(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToStringSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBoolSlice(t *testing.T) {
	type args struct {
		data []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []bool
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				data: []interface{}{true, false},
			},
			want:    []bool{true, false},
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				data: []interface{}{"true", "false"},
			},
			want:    []bool{true, false},
			wantErr: false,
		},
		{
			name: "test3",
			args: args{
				data: []interface{}{1, 0},
			},
			want:    []bool{true, false},
			wantErr: false,
		},
		{
			name: "test4",
			args: args{
				data: []interface{}{map[string]string{}, 0},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToBoolSlice(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToBoolSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBoolSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloatX(t *testing.T) {
	type args struct {
		data interface{}
		v    interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "float32",
			args: args{
				data: float32(1),
				v:    float32(1),
			},
			wantErr: false,
		},
		{
			name: "int",
			args: args{
				data: int(1),
				v:    float32(1),
			},
			wantErr: false,
		},
		{
			name: "string",
			args: args{
				data: "float32",
				v:    float32(1),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ToFloatX(tt.args.data, &tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("ToFloatX() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestToFloat32SliceE(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []float32
		wantErr bool
	}{
		{
			name: "int",
			args: args{
				data: []interface{}{1, 2, 3},
			},
			want:    []float32{1, 2, 3},
			wantErr: false,
		},
		{
			name: "float32",
			args: args{
				data: []interface{}{float32(1), float32(2), float32(3)},
			},
			want:    []float32{1, 2, 3},
			wantErr: false,
		},
		{
			name: "float64",
			args: args{
				data: []interface{}{float64(1), float64(2), float64(3)},
			},
			want:    []float32{1, 2, 3},
			wantErr: false,
		},
		{
			name: "string",
			args: args{
				data: []interface{}{"1", "2", "3"},
			},
			want:    []float32{1, 2, 3},
			wantErr: false,
		},
		{
			name: "err",
			args: args{
				data: []interface{}{"a"},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToFloatXSliceE[float32](tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToFloatXSliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToFloatXSliceE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat64SliceE(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		wantErr bool
	}{
		{
			name: "int",
			args: args{
				data: []interface{}{1, 2, 3},
			},
			want:    []float64{1, 2, 3},
			wantErr: false,
		},
		{
			name: "float32",
			args: args{
				data: []interface{}{float32(1), float32(2), float32(3)},
			},
			want:    []float64{1, 2, 3},
			wantErr: false,
		},
		{
			name: "float64",
			args: args{
				data: []interface{}{float64(1), float64(2), float64(3)},
			},
			want:    []float64{1, 2, 3},
			wantErr: false,
		},
		{
			name: "string",
			args: args{
				data: []interface{}{"1", "2", "3"},
			},
			want:    []float64{1, 2, 3},
			wantErr: false,
		},
		{
			name: "err",
			args: args{
				data: []interface{}{"a"},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToFloatXSliceE[float64](tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToFloatXSliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToFloatXSliceE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToXIntXE(t *testing.T) {
	type args struct {
		data interface{}
		v    interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "int64",
			args: args{
				data: 1,
				v:    int64(1),
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "uint64",
			args: args{
				data: 1,
				v:    uint64(1),
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ToXIntXE(tt.args.data, &tt.args.v)
			if (err != nil) != tt.wantErr && !reflect.DeepEqual(tt.args.v, tt.want) {
				t.Errorf("ToXIntXE() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestToXIntXSliceE(t *testing.T) {
	args := []interface{}{
		1,
		"2",
		true}
	ints := []int{1, 2, 1}
	i, err := ToXIntXSliceE[int](args)
	if err != nil {
		t.Errorf("ToXIntXSliceE() error = %v", err)
		return
	}
	if !reflect.DeepEqual(i, ints) {
		t.Errorf("ToXIntXSliceE() = %v, want %v", i, ints)
	}

	int8s := []int8{1, 2, 1}
	i8, err := ToXIntXSliceE[int8](args)
	if err != nil {
		t.Errorf("ToXIntXSliceE() error = %v", err)
		return
	}
	if !reflect.DeepEqual(i8, int8s) {
		t.Errorf("ToXIntXSliceE() = %v, want %v", i8, int8s)
	}

	int16s := []int16{1, 2, 1}
	i16, err := ToXIntXSliceE[int16](args)
	if err != nil {
		t.Errorf("ToXIntXSliceE() error = %v", err)
		return
	}
	if !reflect.DeepEqual(i16, int16s) {
		t.Errorf("ToXIntXSliceE() = %v, want %v", i16, int16s)
	}

	uint16s := []int16{1, 2, 1}
	ui16, err := ToXIntXSliceE[int16](args)
	if err != nil {
		t.Errorf("ToXIntXSliceE() error = %v", err)
		return
	}
	if !reflect.DeepEqual(ui16, uint16s) {
		t.Errorf("ToXIntXSliceE() = %v, want %v", ui16, uint16s)
	}
}

func Test_unmarshalSlice(t *testing.T) {
	type args struct {
		name string
		data []interface{}
		v    any
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{name: "test1", args: args{name: "test1", data: []interface{}{uint(1), uint8(2), uint64(3), 4, int8(5), int16(6), int32(7), int64(8)}, v: []int{}}, want: []int{1, 2, 3, 4, 5, 6, 7, 8}, wantErr: false},
		{name: "test2", args: args{name: "test2", data: []interface{}{1, int32(2), int64(3)}, v: []int8{}}, want: []int8{1, 2, 3}, wantErr: false},
		{name: "test3", args: args{name: "test3", data: []interface{}{"1", 2, true}, v: []uint8{}}, want: []uint8{1, 2, 1}, wantErr: false},
		{name: "test4", args: args{name: "test4", data: []interface{}{"1", uint16(2), true}, v: []string{}}, want: []string{"1", "2", "true"}, wantErr: false},
		{name: "test5", args: args{name: "test5", data: []interface{}{"1", 2, 3}, v: []int16{}}, want: []int16{1, 2, 3}, wantErr: false},
		{name: "test6", args: args{name: "test6", data: []interface{}{"1", 2, 3}, v: []int32{}}, want: []int32{1, 2, 3}, wantErr: false},
		{name: "test7", args: args{name: "test7", data: []interface{}{"1", 2, 3}, v: []int64{}}, want: []int64{1, 2, 3}, wantErr: false},
		{name: "test8", args: args{name: "test5", data: []interface{}{"1", 2, 3}, v: []uint16{}}, want: []uint16{1, 2, 3}, wantErr: false},
		{name: "test9", args: args{name: "test6", data: []interface{}{"1", 2, 3}, v: []uint32{}}, want: []uint32{1, 2, 3}, wantErr: false},
		{name: "test10", args: args{name: "test7", data: []interface{}{"1", 2, 3}, v: []uint64{}}, want: []uint64{1, 2, 3}, wantErr: false},
		{name: "test11", args: args{name: "test6", data: []interface{}{"1", 2, 3}, v: []uint{}}, want: []uint{1, 2, 3}, wantErr: false},
		{name: "test12", args: args{name: "test7", data: []interface{}{"1", 2, 3}, v: []int{}}, want: []int{1, 2, 3}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unmarshalSlice(tt.args.name, tt.args.data, tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("unmarshalSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unmarshalSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

type ConfigTest struct {
	A int    `doggie:"a"`
	B int8   `doggie:"b"`
	C int16  `doggie:"c"`
	D int32  `doggie:"d"`
	E int64  `doggie:"e"`
	F uint   `doggie:"f"`
	G uint8  `doggie:"g"`
	H uint16 `doggie:"h"`
	I uint32 `doggie:"i"`
	J uint64 `doggie:"j"`
	K string `doggie:"k"`
	L bool   `doggie:"l"`
	M []int  `doggie:"m"`
	N struct {
		A int  `doggie:"a"`
		B int8 `doggie:"b"`
	} `doggie:"n"`
	O float32   `doggie:"o"`
	P float64   `doggie:"p"`
	Q []bool    `doggie:"q"`
	R []float32 `doggie:"r"`
	S []float64 `doggie:"s"`
}

func Test_unmarshal(t *testing.T) {
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

	if err := unmarshal(m, &v); err != nil {
		t.Errorf("unmarshal() error = %v", err)
	}
	if !reflect.DeepEqual(v, target) {
		t.Errorf("unmarshal() = %v, want %v", v, target)
	}
}

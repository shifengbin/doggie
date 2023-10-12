package flags

import (
	"reflect"
	"testing"
)

func Test_countPrefixByte(t *testing.T) {
	type args struct {
		flag   string
		prefix byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty", args{"empty", 'a'}, 0},
		{"a1", args{"a1", 'a'}, 1},
		{"aa2", args{"aa2", 'a'}, 2},
		{"", args{"", 'a'}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrefixByte(tt.args.flag, tt.args.prefix); got != tt.want {
				t.Errorf("countPrefixByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseFlag1(t *testing.T) {
	type args struct {
		flag string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		{"invalid", args{"empty"}, "", "", true},
		{"-a", args{"-a"}, "a", "", false},
		{"--a", args{"--a"}, "a", "", false},
		{"-a=10", args{"-a=10"}, "a", "10", false},
		{"-a=", args{"-a="}, "a", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := parseFlag1(tt.args.flag)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseFlag1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseFlag1() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseFlag1() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_parseOneFlag(t *testing.T) {
	type args struct {
		cmd []string
	}
	tests := []struct {
		name       string
		args       args
		wantKey    string
		wantVal    string
		wantRetCmd []string
		wantErr    bool
	}{
		{"invalid", args{[]string{"val"}}, "", "", nil, true},
		{"-a -b", args{[]string{"-a", "-b"}}, "a", "", []string{"-b"}, false},
		{"-a", args{[]string{"-a"}}, "a", "", []string{}, false},
		{"", args{[]string{}}, "", "", nil, true},
		{"---a", args{[]string{"---a", "-b"}}, "", "", nil, true},
		{"--a", args{[]string{"--a", "--b"}}, "a", "", []string{"--b"}, false},
		{"-a=10 --b", args{[]string{"-a=10", "--b"}}, "a", "10", []string{"--b"}, false},
		{"-a 10 --b", args{[]string{"-a", "10", "--b"}}, "a", "10", []string{"--b"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, gotVal, gotRetCmd, err := parseOneFlag(tt.args.cmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseOneFlag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotKey != tt.wantKey {
				t.Errorf("parseOneFlag() gotKey = %v, want %v", gotKey, tt.wantKey)
			}
			if gotVal != tt.wantVal {
				t.Errorf("parseOneFlag() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
			if !reflect.DeepEqual(gotRetCmd, tt.wantRetCmd) {
				t.Errorf("parseOneFlag() gotRetCmd = %v, want %v", gotRetCmd, tt.wantRetCmd)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		cmd []string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{"", args{[]string{}}, map[string]string{}, false},
		{"invalid", args{[]string{"val"}}, map[string]string{}, true},
		{"-a", args{[]string{"-a", "-b"}}, map[string]string{"a": "", "b": ""}, false},
		{"--a", args{[]string{"--a", "--b"}}, map[string]string{"a": "", "b": ""}, false},
		{"-a=10 --b", args{[]string{"-a=10", "--b"}}, map[string]string{"a": "10", "b": ""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.cmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

package flags

import (
	"fmt"
	"strings"
)

//参数格式为(EBNF)：
//FALG1 ::= "-"{1,2} IDENT ("=" LETTER+)?
//FLAG2 ::= "-"{1,2} IDENT  SP LETTERS
//IDENT ::= [A-Z] [0-9A-Z.]*

//-a
//--a
//-a=10
//-a 10

func Parse(cmd []string) (map[string]string, error) {
	fdata := make(map[string]string)
	if len(cmd) == 0 {
		return fdata, nil
	}
	var err error
	for len(cmd) > 0 {
		key, val, rCmd, er := parseOneFlag(cmd)
		if er != nil {
			err = er
			break
		}
		fdata[key] = val
		cmd = rCmd
	}
	return fdata, err
}

// parseOneFlag 解析一个flag
// cmd 要解析的参数
// 返回：
//
//	key flag名称
//	val flag 值
//	cmd 剩余cmd
//	err 格式错误返回
func parseOneFlag(cmd []string) (key string, val string, retCmd []string, err error) {
	if len(cmd) == 0 {
		return "", "", nil, fmt.Errorf("command not found")
	}

	key, val, err = parseFlag1(cmd[0])
	if err != nil {
		return
	}
	cmd = cmd[1:]
	retCmd = cmd
	//-a=10
	if val != "" {
		return
	}

	//-a
	if len(cmd) == 0 {
		return
	}

	//-a -b
	if countPrefixByte(cmd[0], '-') > 0 {
		return
	}

	//-a 10
	retCmd = cmd[1:]
	val = cmd[0]
	return
}

func parseFlag1(flag string) (string, string, error) {
	// 检查格式
	prefixCount := countPrefixByte(flag, '-')
	if prefixCount == 0 || prefixCount > 2 {
		return "", "", fmt.Errorf("command error:%s", flag)
	}

	flag = flag[prefixCount:]

	i := strings.Index(flag, "=")
	if i == -1 {
		return flag, "", nil
	}

	return flag[:i], flag[i+1:], nil

}

func countPrefixByte(flag string, prefix byte) int {
	if flag == "" {
		return 0
	}

	var c int
	for i := 0; i < len(flag); i++ {
		if flag[i] != prefix {
			break
		}
		c++
	}
	return c
}

package json

import (
	"fmt"
	"gitee.com/cristiane/go-push-sdk/push/common/convert"
	jsoniter "github.com/json-iterator/go"
)

func Marshal(v interface{}) ([]byte, error) {
	if v == nil {
		return nil,fmt.Errorf("\u001B[34m%v\u001B[0m: \x1b[31m%v\x1b[0m\n", "[go-push-sdk]", "invalid memory address or nil pointer dereference")
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(v)
}

func MarshalToString(v interface{}) (string, error) {
	_byte, err := Marshal(v)
	return convert.Byte2Str(_byte), err
}

func MarshalToStringNoError(v interface{}) string {
	str, _ := MarshalToString(v)
	return str
}

func Unmarshal(data string, v interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(convert.Str2ByteByReflect(data), v)
}

func UnmarshalByte(data []byte, v interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(data, v)
}


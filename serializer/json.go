/*
@Time : 2022/1/15 18:33
@Author : Hwdhy
@File : json
@Software: GoLand
*/
package serializer

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

// protobuf 转json
func ProtobufToJSON(message proto.Message) (string, error) {
	mershaler := jsonpb.Marshaler{
		EnumsAsInts:  false, //枚举类型输出格式  false：字符串  true：int数字
		EmitDefaults: true,
		Indent:       "  ",
		OrigName:     true, //字段名： false：驼峰   true: 下划线
	}

	return mershaler.MarshalToString(message)
}

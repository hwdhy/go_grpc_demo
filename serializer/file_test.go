/*
@Time : 2022/1/15 18:11
@Author : Hwdhy
@File : file_test
@Software: GoLand
*/
package serializer_test

import (
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"grpc_project/pb"
	"grpc_project/sample"
	"grpc_project/serializer"
	"testing"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}

	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.WriteProtobufToJSON(laptop2, jsonFile)
	require.NoError(t, err)
}

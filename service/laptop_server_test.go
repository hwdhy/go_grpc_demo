/*
@Time : 2022/1/15 22:41
@Author : Hwdhy
@File : laptop_server_test
@Software: GoLand
*/
package service_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc_project/pb"
	"grpc_project/sample"
	"grpc_project/service"
	"testing"
)

func TestServerCreateLaptop(t *testing.T) {
	t.Parallel()

	LaptopNoID := sample.NewLaptop()
	LaptopNoID.Id = ""

	laptopInvalidID := sample.NewLaptop()
	laptopInvalidID.Id = "invalid_uuid"

	laptopDuplicateID := sample.NewLaptop()
	storeDuplicateID := service.NewInmemoryLaptopStore()
	err := storeDuplicateID.Save(laptopDuplicateID)
	require.NoError(t, err)

	testCases := []struct {
		name   string
		laptop *pb.Laptop
		store  service.LaptopStore
		code   codes.Code
	}{
		{
			name:   "success_with_id",
			laptop: sample.NewLaptop(),
			store:  service.NewInmemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "success_no_id",
			laptop: LaptopNoID,
			store:  service.NewInmemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "failure_invlid_id",
			laptop: laptopInvalidID,
			store:  service.NewInmemoryLaptopStore(),
			code:   codes.InvalidArgument,
		},
		{
			name:   "failure_duplicate_id",
			laptop: laptopDuplicateID,
			store:  storeDuplicateID,
			code:   codes.AlreadyExists,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := &pb.CreateLaptopRequest{
				Laptop: tc.laptop,
			}

			server := service.NewLaptopServer(tc.store, nil, nil)
			res, err := server.CreateLaptop(context.Background(), req)
			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)
				if len(res.Id) > 0 {
					require.Equal(t, tc.laptop.Id, res.Id)
				}
			} else {
				require.Error(t, err)
				require.Nil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, st.Code(), tc.code)
			}
		})
	}
}

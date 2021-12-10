package service

import (
	"fmt"
	"myEntry/pkg/entity"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	//type args struct {
	//	req entity.GetUserInfoReq
	//}
	//
	//
	//tests := []struct {
	//	name string
	//	args args
	//	want entity.GetUserInfoResp
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := GetUserInfo(tt.args.req); !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("GetUserInfo() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}

	req := entity.GetUserInfoReq{}
	req.Username = "wei"
	resp := GetUserInfo(req)
	fmt.Print(resp)
}

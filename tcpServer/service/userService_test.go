package service

import (
	"myEntry/pkg/entity"
	"reflect"
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
}

func TestGetUserInfo1(t *testing.T) {
	type args struct {
		req entity.GetUserInfoReq
	}
	tests := []struct {
		name     string
		args     args
		wantResp entity.GetUserInfoResp
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := GetUserInfo(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("GetUserInfo() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

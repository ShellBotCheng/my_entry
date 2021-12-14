package service

import (
	"github.com/spf13/pflag"
	"myEntry/pkg/entity"
	"myEntry/pkg/log"
	"myEntry/pkg/mysql"
	"myEntry/pkg/redis"
	"myEntry/tcpServer/conf"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	InitServer()
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
		{
			name: "getUserInfo",
			args: args{
				req: entity.GetUserInfoReq{
					Username: "admin",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := GetUserInfo(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResp.Status != 200 {
				t.Errorf("GetUserInfo() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
			t.Logf("GetUserInfo() gotResp = %v", gotResp)
		})
	}
}

var (
	tcpCfgFile = pflag.StringP("config", "c", "../../tcpServer/conf/dev.yml", "config file path.")
)

func InitServer() error {
	err := log.Init()
	if err != nil {
		return err
	}
	// init config
	cfg := conf.Init(*tcpCfgFile)

	// init redis
	redis.Init(&cfg.Redis)

	// init mysql
	mysql.Init(&cfg.Mysql)
	return nil
}

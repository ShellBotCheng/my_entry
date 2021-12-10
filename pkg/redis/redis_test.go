package redis

import (
	"context"
	"testing"
)

func TestInitTestRedis(t *testing.T) {

	err := Client.Ping(context.Background()).Err()
	if err != nil {
		t.Error("ping redis server err: ", err)
		return
	}
	t.Log("ping redis server pass")
}

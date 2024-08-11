package gfunc

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func Test_GetSerName(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		name := GetSerName("grpc","v1/jwt/*")
		t.Logf("name:%s", name)
	})
}




package btExtend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func TestGetWebList(t *testing.T) {
	jsonData, err := CreateBtRequest(gctx.New(), "CuPXPASiwWQuyDuNuU2nqxPvW7iwhyiv", "http://127.0.0.1:9944").Resp("/data?action=getData&table=sites", g.Map{
		"limit": 15,
	})
	gtest.AssertNil(err)
	g.DumpWithType(jsonData.Map())
}

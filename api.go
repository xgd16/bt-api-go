package btExtend

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gtime"
	"time"
)

type BtRequest struct {
	ctx       context.Context
	token     string
	domainUrl string
}

// CreateBtRequest create bt request client
func CreateBtRequest(ctx context.Context, token, domainUrl string) *BtRequest {
	return &BtRequest{ctx: ctx, token: token, domainUrl: domainUrl}
}

// Resp get resp data
func (t *BtRequest) Resp(url string, data g.Map) (jsonData *gjson.Json, err error) {
	// init resp data
	var (
		post *gclient.Response
	)
	// create base http client
	client := g.Client().SetTimeout(60 * time.Second)
	// create request token
	nowTimeUnix := gtime.Now().Unix()
	if data == nil {
		data = g.Map{}
	}
	data["request_token"] = gmd5.MustEncrypt(fmt.Sprintf(
		"%d%s",
		nowTimeUnix,
		gmd5.MustEncrypt(t.token),
	))
	data["request_time"] = nowTimeUnix
	// send request
	if data == nil {
		post, err = client.Post(t.ctx, fmt.Sprintf("%s/%s", t.domainUrl, url))
	} else {
		post, err = client.Post(t.ctx, fmt.Sprintf("%s/%s", t.domainUrl, url), data)
	}
	if err != nil {
		return
	}
	// decode to json obj
	jsonData, err = gjson.DecodeToJson(post.ReadAllString())
	if err != nil {
		return
	}
	return
}

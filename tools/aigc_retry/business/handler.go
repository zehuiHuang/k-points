package business

import (
	"context"
	"fmt"
	"net/http"
)

type ApiBase struct {
	Ctx      context.Context
	TaskInfo *Info
}

type Api interface {
	Models() []string //该接口支持的模型参数
}

type ApiFactory func(base *ApiBase) Api

// ApiDispatcher 调度器, 注册第三方接口处理器
// 根据唯一标识确定使用哪个处理器进行处理, 同步模式和异步模式均可调用
type ApiDispatcher struct {
	workers map[string]ApiFactory
	client  *http.Client
}

func (d *ApiDispatcher) Dispatch(ctx context.Context, taskInfo *Info) error {
	taskInfo.ExecuteCount++
	fmt.Printf("Dispatch handler productName:%s,executeCount:%s\n", taskInfo.ProductName, taskInfo.ExecuteCount)
	return nil
}

type RespPayload struct {
	ChanExtend    string `json:"chan_extend"`
	UniversalResp string `json:"universal_resp"`
}

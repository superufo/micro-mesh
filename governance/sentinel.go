package service

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/flow"
)

// todo  此组件的调用宜放在网关端和 grpc客户端调用
var Sentinel = sentinelService{}

type sentinelService struct{
	ctx      context.Context
	log      *glog.Logger
	Rule     []*flow.Rule
}

func NewSentinelService(ctx context.Context,log *glog.Logger,Rule []*flow.Rule) *sentinelService{
	return &sentinelService{
		ctx:  ctx,
		log:  log,
		Rule: Rule,
	}
}

// 根据参数判断 限流
func (s sentinelService)Limit()  bool  {
	err := sentinel.InitDefault()
	if err != nil {
		s.log.Fatalf(s.ctx,"Unexpected error: %+v", err)
	}

	_, err = flow.LoadRules(s.Rule)
	if err != nil {
		s.log.Fatalf(s.ctx,"Unexpected error: %+v", err)
	}

	if  len(s.Rule)<1 {
		s.log.Info(s.ctx,"没有配置限流规则")
		return false
	}

	if s.Rule[0].TokenCalculateStrategy == flow.Direct {
		e, b := sentinel.Entry(s.Rule[0].Resource, sentinel.WithTrafficType(base.Inbound))
		if b != nil {
			s.log.Notice(s.ctx,"限流了")
			return true
		} else {
			s.log.Info(s.ctx,"限流通过")
			// 业务函数
			e.Exit()
			return  false
		}
	}

	return  false
}

// 根据参数判断 熔断




// 统计




package data

import (
	"github.com/gogf/gf/v2/os/gtime"
	"golang.org/x/net/context"
	"myshop/internal/dao"
	"myshop/internal/model"
	"myshop/internal/service"
	"myshop/utility"
)

type sData struct {
}

func init() { //注册服务
	service.RegisterData(New())
}

func New() *sData {
	return &sData{}
}
func (s *sData) DataHead(ctx context.Context) (out *model.DataHeadOutput, err error) {
	return &model.DataHeadOutput{
		TodayOrderCount: todayOrderCount(ctx),
		DAU:             utility.RandInt(1000),
		ConversionRate:  utility.RandInt(50),
	}, nil
}

// 查询今天订单总数
func todayOrderCount(ctx context.Context) (count int) {
	count, err := dao.OrderInfo.Ctx(ctx).
		WhereBetween(dao.OrderInfo.Columns().CreatedAt, gtime.Now().StartOfDay(), gtime.Now().EndOfDay()).
		Count(dao.OrderInfo.Columns().Id)
	if err != nil {
		return -1
	}
	return count
}

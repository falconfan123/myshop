package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DataHeadReq struct { //请求
	g.Meta `path:"/backend/data/head" method:"get" tags:"数据data" desc:"数据大屏头部信息"`
}

type DataHeadRes struct { //响应
	TodayOrderCount int `json:"today_order_count" desc:"今日订单总数"`
	DAU             int `json:"dau" desc:"今日日活"`
	ConversionRate  int `json:"conversion-rate" desc:"转化率"`
}

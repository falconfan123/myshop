// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EventeECardExtendDao is the data access object for table evente_e_card_extend.
type EventeECardExtendDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns EventeECardExtendColumns // columns contains all the column names of Table for convenient usage.
}

// EventeECardExtendColumns defines and stores column names for table evente_e_card_extend.
type EventeECardExtendColumns struct {
	Id           string //
	OrgId        string // 主办id
	CardId       string // E通卡id
	EventeId     string // 活动id
	ScreeningsId string // 场次id
	PriceId      string // 票品ID集合 如 :9988,2235,6667
	CreateDate   string // 创建时间
	UpdateDate   string // 修改时间
}

// eventeECardExtendColumns holds the columns for table evente_e_card_extend.
var eventeECardExtendColumns = EventeECardExtendColumns{
	Id:           "id",
	OrgId:        "org_id",
	CardId:       "card_id",
	EventeId:     "evente_id",
	ScreeningsId: "screenings_id",
	PriceId:      "price_id",
	CreateDate:   "create_date",
	UpdateDate:   "update_date",
}

// NewEventeECardExtendDao creates and returns a new DAO object for table data access.
func NewEventeECardExtendDao() *EventeECardExtendDao {
	return &EventeECardExtendDao{
		group:   "default",
		table:   "evente_e_card_extend",
		columns: eventeECardExtendColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EventeECardExtendDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EventeECardExtendDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EventeECardExtendDao) Columns() EventeECardExtendColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EventeECardExtendDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EventeECardExtendDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EventeECardExtendDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

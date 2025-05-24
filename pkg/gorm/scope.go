package gorm

import (
	"fmt"

	"gorm.io/gorm"
)

const (
	ORDER_ASC  = "ASC"
	ORDER_DESC = "DESC"
)

// 分页
func PaginateScope(page, pageSize int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page < 1 {
			page = 1
		}
		offset := (page - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}

// 自定义排序
func OrderByScope(field string, order string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(fmt.Sprintf("%s %s", field, order))
	}
}

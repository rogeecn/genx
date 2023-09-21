// Code generated by github.com/rogeecn/genx. DO NOT EDIT.
// Code generated by github.com/rogeecn/genx. DO NOT EDIT.
// Code generated by github.com/rogeecn/genx. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNamePerson = "people"

// Person mapped from table <people>
type Person struct {
	ID             int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"-"`
	Name           *string        `gorm:"column:name" json:"-"`
	Alias_         *string        `gorm:"column:alias" json:"-"`
	Age            *int32         `gorm:"column:age" json:"-"`
	Flag           *bool          `gorm:"column:flag" json:"-"`
	AnotherFlag    *int32         `gorm:"column:another_flag" json:"-"`
	Commit         *string        `gorm:"column:commit" json:"-"`
	First          *bool          `gorm:"column:First" json:"-"`
	Bit            *[]uint8       `gorm:"column:bit" json:"-"`
	Small          *int32         `gorm:"column:small" json:"-"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at" json:"-"`
	Score          *float64       `gorm:"column:score" json:"-"`
	Number         *int32         `gorm:"column:number" json:"-"`
	Birth          *time.Time     `gorm:"column:birth;default:CURRENT_TIMESTAMP" json:"-"`
	XMLHTTPRequest *string        `gorm:"column:xmlHTTPRequest;default:' '" json:"-"`
	JStr           *string        `gorm:"column:jStr" json:"-"`
	Geo            *string        `gorm:"column:geo" json:"-"`
	Mint           *int32         `gorm:"column:mint" json:"-"`
	Blank          *string        `gorm:"column:blank;default:' '" json:"-"`
	Remark         *string        `gorm:"column:remark" json:"-"`
	LongRemark     *string        `gorm:"column:long_remark" json:"-"`
}

// TableName Person's table name
func (*Person) TableName() string {
	return TableNamePerson
}

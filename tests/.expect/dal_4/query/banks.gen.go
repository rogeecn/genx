// Code generated by github.com/rogeecn/genx. DO NOT EDIT.
// Code generated by github.com/rogeecn/genx. DO NOT EDIT.
// Code generated by github.com/rogeecn/genx. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"github.com/rogeecn/genx"
	"github.com/rogeecn/genx/field"

	"gorm.io/plugin/dbresolver"

	"github.com/rogeecn/genx/tests/.gen/dal_4/model"
)

func newBank(db *gorm.DB, opts ...gen.DOOption) bank {
	_bank := bank{}

	_bank.bankDo.UseDB(db, opts...)
	_bank.bankDo.UseModel(&model.Bank{})

	tableName := _bank.bankDo.TableName()
	_bank.ALL = field.NewAsterisk(tableName)
	_bank.ID = field.NewInt64(tableName, "id")
	_bank.Name = field.NewString(tableName, "name")
	_bank.Address = field.NewString(tableName, "address")
	_bank.Scale = field.NewInt64(tableName, "scale")

	_bank.fillFieldMap()

	return _bank
}

type bank struct {
	bankDo bankDo

	ALL     field.Asterisk
	ID      field.Int64
	Name    field.String
	Address field.String
	Scale   field.Int64

	fieldMap map[string]field.Expr
}

func (b bank) Table(newTableName string) *bank {
	b.bankDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b bank) As(alias string) *bank {
	b.bankDo.DO = *(b.bankDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *bank) updateTableName(table string) *bank {
	b.ALL = field.NewAsterisk(table)
	b.ID = field.NewInt64(table, "id")
	b.Name = field.NewString(table, "name")
	b.Address = field.NewString(table, "address")
	b.Scale = field.NewInt64(table, "scale")

	b.fillFieldMap()

	return b
}

func (b *bank) WithContext(ctx context.Context) IBankDo { return b.bankDo.WithContext(ctx) }

func (b bank) TableName() string { return b.bankDo.TableName() }

func (b bank) Alias() string { return b.bankDo.Alias() }

func (b bank) Columns(cols ...field.Expr) gen.Columns { return b.bankDo.Columns(cols...) }

func (b *bank) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *bank) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 4)
	b.fieldMap["id"] = b.ID
	b.fieldMap["name"] = b.Name
	b.fieldMap["address"] = b.Address
	b.fieldMap["scale"] = b.Scale
}

func (b bank) clone(db *gorm.DB) bank {
	b.bankDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b bank) replaceDB(db *gorm.DB) bank {
	b.bankDo.ReplaceDB(db)
	return b
}

type bankDo struct{ gen.DO }

type IBankDo interface {
	gen.SubQuery
	Debug() IBankDo
	WithContext(ctx context.Context) IBankDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBankDo
	WriteDB() IBankDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBankDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBankDo
	Not(conds ...gen.Condition) IBankDo
	Or(conds ...gen.Condition) IBankDo
	Select(conds ...field.Expr) IBankDo
	Where(conds ...gen.Condition) IBankDo
	Order(conds ...field.Expr) IBankDo
	Distinct(cols ...field.Expr) IBankDo
	Omit(cols ...field.Expr) IBankDo
	Join(table schema.Tabler, on ...field.Expr) IBankDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBankDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBankDo
	Group(cols ...field.Expr) IBankDo
	Having(conds ...gen.Condition) IBankDo
	Limit(limit int) IBankDo
	Offset(offset int) IBankDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBankDo
	Unscoped() IBankDo
	Create(values ...*model.Bank) error
	CreateInBatches(values []*model.Bank, batchSize int) error
	Save(values ...*model.Bank) error
	First() (*model.Bank, error)
	Take() (*model.Bank, error)
	Last() (*model.Bank, error)
	Find() ([]*model.Bank, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Bank, err error)
	FindInBatches(result *[]*model.Bank, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Bank) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBankDo
	Assign(attrs ...field.AssignExpr) IBankDo
	Joins(fields ...field.RelationField) IBankDo
	Preload(fields ...field.RelationField) IBankDo
	FirstOrInit() (*model.Bank, error)
	FirstOrCreate() (*model.Bank, error)
	FindByPage(offset int, limit int) (result []*model.Bank, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBankDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (b bankDo) Debug() IBankDo {
	return b.withDO(b.DO.Debug())
}

func (b bankDo) WithContext(ctx context.Context) IBankDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b bankDo) ReadDB() IBankDo {
	return b.Clauses(dbresolver.Read)
}

func (b bankDo) WriteDB() IBankDo {
	return b.Clauses(dbresolver.Write)
}

func (b bankDo) Session(config *gorm.Session) IBankDo {
	return b.withDO(b.DO.Session(config))
}

func (b bankDo) Clauses(conds ...clause.Expression) IBankDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b bankDo) Returning(value interface{}, columns ...string) IBankDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b bankDo) Not(conds ...gen.Condition) IBankDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b bankDo) Or(conds ...gen.Condition) IBankDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b bankDo) Select(conds ...field.Expr) IBankDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b bankDo) Where(conds ...gen.Condition) IBankDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b bankDo) Order(conds ...field.Expr) IBankDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b bankDo) Distinct(cols ...field.Expr) IBankDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b bankDo) Omit(cols ...field.Expr) IBankDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b bankDo) Join(table schema.Tabler, on ...field.Expr) IBankDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b bankDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBankDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b bankDo) RightJoin(table schema.Tabler, on ...field.Expr) IBankDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b bankDo) Group(cols ...field.Expr) IBankDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b bankDo) Having(conds ...gen.Condition) IBankDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b bankDo) Limit(limit int) IBankDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b bankDo) Offset(offset int) IBankDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b bankDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBankDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b bankDo) Unscoped() IBankDo {
	return b.withDO(b.DO.Unscoped())
}

func (b bankDo) Create(values ...*model.Bank) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b bankDo) CreateInBatches(values []*model.Bank, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b bankDo) Save(values ...*model.Bank) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b bankDo) First() (*model.Bank, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bank), nil
	}
}

func (b bankDo) Take() (*model.Bank, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bank), nil
	}
}

func (b bankDo) Last() (*model.Bank, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bank), nil
	}
}

func (b bankDo) Find() ([]*model.Bank, error) {
	result, err := b.DO.Find()
	return result.([]*model.Bank), err
}

func (b bankDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Bank, err error) {
	buf := make([]*model.Bank, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b bankDo) FindInBatches(result *[]*model.Bank, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b bankDo) Attrs(attrs ...field.AssignExpr) IBankDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b bankDo) Assign(attrs ...field.AssignExpr) IBankDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b bankDo) Joins(fields ...field.RelationField) IBankDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b bankDo) Preload(fields ...field.RelationField) IBankDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b bankDo) FirstOrInit() (*model.Bank, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bank), nil
	}
}

func (b bankDo) FirstOrCreate() (*model.Bank, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bank), nil
	}
}

func (b bankDo) FindByPage(offset int, limit int) (result []*model.Bank, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b bankDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b bankDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b bankDo) Delete(models ...*model.Bank) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *bankDo) withDO(do gen.Dao) *bankDo {
	b.DO = *do.(*gen.DO)
	return b
}

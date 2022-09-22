package repositories

import (
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . DBer

// ユーザ
type DBer interface {
	// finisher(https://github.com/go-gorm/gorm/blob/master/finisher_api.go)
	Create(value interface{}) (tx *gorm.DB)
	CreateInBatches(value interface{}, batchSize int) (tx *gorm.DB)
	Save(value interface{}) (tx *gorm.DB)
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Take(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Last(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB
	FirstOrInit(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	FirstOrCreate(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Update(column string, value interface{}) (tx *gorm.DB)
	Updates(values interface{}) (tx *gorm.DB)
	UpdateColumn(column string, value interface{}) (tx *gorm.DB)
	UpdateColumns(values interface{}) (tx *gorm.DB)
	Delete(value interface{}, conds ...interface{}) (tx *gorm.DB)
	Count(count *int64) (tx *gorm.DB)
	Row() *sql.Row
	Rows() (*sql.Rows, error)
	Scan(dest interface{}) (tx *gorm.DB)
	Pluck(column string, dest interface{}) (tx *gorm.DB)
	ScanRows(rows *sql.Rows, dest interface{}) error
	Connection(fc func(tx *gorm.DB) error) (err error)
	Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error)
	Begin(opts ...*sql.TxOptions) *gorm.DB
	Commit() *gorm.DB
	Rollback() *gorm.DB
	SavePoint(name string) *gorm.DB
	RollbackTo(name string) *gorm.DB
	Exec(sql string, values ...interface{}) (tx *gorm.DB)

	// chainable(https://github.com/go-gorm/gorm/blob/master/chainable_api.go)
	Model(value interface{}) (tx *gorm.DB)
	Clauses(conds ...clause.Expression) (tx *gorm.DB)
	Table(name string, args ...interface{}) (tx *gorm.DB)
	Distinct(args ...interface{}) (tx *gorm.DB)
	Select(query interface{}, args ...interface{}) (tx *gorm.DB)
	Omit(columns ...string) (tx *gorm.DB)
	Where(query interface{}, args ...interface{}) (tx *gorm.DB)
	Not(query interface{}, args ...interface{}) (tx *gorm.DB)
	Or(query interface{}, args ...interface{}) (tx *gorm.DB)
	Joins(query string, args ...interface{}) (tx *gorm.DB)
	Group(name string) (tx *gorm.DB)
	Having(query interface{}, args ...interface{}) (tx *gorm.DB)
	Order(value interface{}) (tx *gorm.DB)
	Limit(limit int) (tx *gorm.DB)
	Offset(offset int) (tx *gorm.DB)
	Scopes(funcs ...func(*gorm.DB) *gorm.DB) (tx *gorm.DB)
	Preload(query string, args ...interface{}) (tx *gorm.DB)
	Attrs(attrs ...interface{}) (tx *gorm.DB)
	Assign(attrs ...interface{}) (tx *gorm.DB)
	Unscoped() (tx *gorm.DB)
	Raw(sql string, values ...interface{}) (tx *gorm.DB)
}

var (
	// インターフェースが期待通りに宣言されているか確認
	_ DBer = (*gorm.DB)(nil)
)

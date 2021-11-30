package dao

import (
	"errors"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq" //pg driver
)

var (
	ErrDataNotExist  = errors.New("data not exist")
	ErrDataExist     = errors.New("data exist")
	ErrStatusInvalid = errors.New("status invalid")
)

type DataEngine struct {
	engine *xorm.Engine
}

func NewDataEngine() *DataEngine {
	return &DataEngine{}
}

// Start database engine, dbUrl sample: postgres://postgres:root@localhost:5432/testdb?sslmode=disable
func (d *DataEngine) Start(dbUrl string) error {
	engine, err := xorm.NewEngine("postgres", dbUrl)
	if err != nil {
		println("xxx1")
		return err
	}
	d.engine = engine
	d.engine.SetMaxIdleConns(3)
	//d.engine.ShowSQL(true)

	return nil
}

func (d *DataEngine) Stop() error {
	return d.engine.Close()
}

// Transaction 数据库事务
type Transaction struct {
	sess *xorm.Session
}

// NewTransaction open a transaction with engine
func (d *DataEngine) NewTransaction() (*Transaction, error) {
	t := &Transaction{
		sess: d.engine.NewSession(),
	}
	err := t.sess.Begin()
	return t, err
}

func (t *Transaction) Commit() error {
	return t.sess.Commit()
}

func (t *Transaction) Rollback() error {
	return t.sess.Rollback()
}

func (t *Transaction) Close() {
	t.sess.Close()
}

package schema

/**
数据库表结构定义
*/

import (
	"time"
)

const (
	TNameUsers    = "users"
	TNameProjects = "projects"
)

type Users struct {
	Id        uint64    `xorm:"bigint pk autoincr 'id'"` //自增主键
	Name      string    `xorm:"varchar(64) 'name'"`      //用户名称
	Account   string    `xorm:"varchar(64) 'account'"`   //账号
	password  string    `xorm:"varchar(32) 'password'"`  //密码
	CreatedAt time.Time `xorm:"created 'created_at'"`    //创建时间
	UpdatedAt time.Time `xorm:"updated 'updated_at'"`    //更新时间
}

func (u *Users) TableName() string {
	return TNameUsers
}

type Projects struct {
	Id          uint64    `xorm:"bigint pk autoincr 'id'"`   //自增主键
	UserId      uint64    `xorm:"bigint 'user_id'"`          //用户ID
	Name        string    `xorm:"varchar(64) 'name'"`        //项目名称
	Description string    `xorm:"varchar(64) 'description'"` //项目描述
	Status      uint8     `xorm:"int 'status'"`              //项目状态
	CreatedAt   time.Time `xorm:"created 'created_at'"`      //创建时间
	UpdatedAt   time.Time `xorm:"updated 'updated_at'"`      //更新时间
}

func (p *Projects) TableName() string {
	return TNameProjects
}

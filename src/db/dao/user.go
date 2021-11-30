package dao

import "apusic/go-webserver/src/db/schema"

type User struct{}

func (u *User) GetUserList() {

}

func (u *User) GetUserInfoById() *schema.Users {
	return &schema.Users{}
}

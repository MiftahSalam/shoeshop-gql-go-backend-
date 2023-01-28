package user

import (
	"fmt"

	"shoeshop-backend/src/domain/user"
	"shoeshop-backend/src/infrastructure/database"
)

type repo struct {
	master database.ORM
	slave  database.ORM
}

func NewRepository(master database.ORM, slave database.ORM) user.Repository {
	if master == nil {
		panic("please provide sql DB")
	}
	if slave == nil {
		panic("please provide sql DB Slave")
	}
	return &repo{master: master, slave: slave}
}

func (r *repo) AutoMigrate() {
	err := r.master.Migrate(&user.User{})
	if err != nil {
		fmt.Println("error auto migrate user domain with error:", err)
	}
}

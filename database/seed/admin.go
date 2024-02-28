package seed

import (
	"TugasRPL/database/repository"
	"TugasRPL/util"
	"context"
	"fmt"
	"log"
)

type MigrateAdmin struct {
	name     string
	username string
	role_id  int64
	email    string
	pass     string
}

var admin = []MigrateAdmin{
	{
		name:     "admin",
		username: "admin",
		role_id:  repository.ADMIN,
		email:    "admin@gmail.com",
		pass:     "admin123",
	},
	{
		name:     "admin2",
		username: "admin2",
		role_id:  repository.ADMIN,
		email:    "admin2@gmail.com",
		pass:     "admin123",
	},
}

func SeedAdmin() {
	ctx := context.Background()
	for _, v := range admin {
		hash, err := util.CreateHashPass(v.pass)
		if err != nil {
			log.Fatal(err.Error())
		}
		id, err := repository.Query.MigrateAdmin(ctx, repository.MigrateAdminParams{
			Name:     v.name,
			UserName: v.username,
			RoleID:   v.role_id,
			Email:    v.email,
			Password: hash,
		})
		err = repository.Query.CreateUserDetail(ctx, id)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			fmt.Println("Succes migrate admin")
		}
	}
}

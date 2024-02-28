package service

import (
	r "TugasRPL/database/repository"
	"TugasRPL/src/ui"
	"TugasRPL/util"
	"context"
	"database/sql"
	"errors"
	"time"
)

func Login(ctx context.Context) error {
	util.ClearScreen()
	ui.Magenta.Println("Login to your account")
	email := ui.Promt(ui.EMAIL)
	password := ui.Promt(ui.PASSWORD)
	user, err := r.Query.SelectUserByEmail(ctx, email)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return errors.New("email or password wrong")
		default:
			return err
		}
	}
	if err := util.CompareHashPass(user.Password, password); err != nil {
		return errors.New("email or password wrong")
	}
	ctx = context.WithValue(ctx, "id", user.ID)
	util.ClearScreen()
	ServiceMenu(ctx)
	return nil
}

func Register(ctx context.Context) error {
	util.ClearScreen()
	ui.Magenta.Println("Register New Account")
	name := ui.Promt("name")
	user_name := ui.Promt("username")
	email := ui.Promt(ui.EMAIL)
	password := ui.Promt(ui.PASSWORD)
	ui.PromtVerify(password)
	hash, err := util.CreateHashPass(password)
	if err != nil {
		return err
	}
	id, err := r.Query.CreateUser(ctx, r.CreateUserParams{
		Name:     name,
		RoleID:   r.USER,
		UserName: user_name,
		Email:    email,
		Password: hash,
	})
	if err != nil {
		return err
	}
	user, err := r.Query.GetUserById(ctx, id)
	if err != nil {
		return err
	}
	err = r.Query.CreateUserDetail(ctx, user.ID)
	if err != nil {
		return err
	}
	ui.GreenBg.Println("Success Register ", user.Name)
	time.Sleep(2 * time.Second)
	return nil
}

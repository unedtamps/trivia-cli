package service

import (
	"TugasRPL/database/repository"
	"TugasRPL/src/ui"
	"TugasRPL/util"
	"context"
	"os"
)

var (
	GAME    = "Start Trivia Quiz"
	MAKE    = "Make a Tvia Question"
	RANKING = "See Triva Rank"
	ADMIN   = "Admin Menu"
	LOGOUT  = "Logout"
	EXIT    = "Exit"
	USER    = "View My Trivia"
)

func ServiceMenu(ctx context.Context) {
	id := ctx.Value("id").(int64)
	menu := []string{GAME, MAKE, RANKING}
	role_id, err := repository.Query.CheckIfAdmin(ctx, id)
	if err != nil {
		ui.RedBg.Println(err.Error())
		return
	}
	if role_id == repository.ADMIN {
		menu = append(menu, ADMIN)
	} else {
		menu = append(menu, USER)
	}
	menu = append(menu, EXIT)
	menu = append(menu, LOGOUT)
	for {
		choice := ui.Select("Select Menu !", menu)
		util.ClearScreen()
		switch choice {
		case GAME:
			err := StartGame(ctx)
			if err != nil {
				ui.RedBg.Println(err.Error())
			}
		case RANKING:
			err := SeeRanking(ctx)
			if err != nil {
				ui.RedBg.Println(err.Error())
			}
		case USER:
			err := SeeMyTrivia(ctx)
			if err != nil {
				ui.RedBg.Println(err.Error())
			}
		case MAKE:
			err := crateTrivia(ctx)
			if err != nil {
				ui.RedBg.Println(err.Error())
			}
		case ADMIN:
			err := adminMenu(ctx)
			if err != nil {
				ui.RedBg.Println(err.Error())
			}
		case EXIT:
			ans := ui.Promt(ui.YESORNO)
			if ans == "n" {
				util.ClearScreen()
				continue
			}
			util.ClearScreen()
			os.Exit(1)
		case LOGOUT:
			ans := ui.Promt(ui.YESORNO)
			if ans == "n" {
				util.ClearScreen()
				continue
			}
			ctx = context.WithValue(ctx, "id", nil)
			util.ClearScreen()
			return
		}
	}
}

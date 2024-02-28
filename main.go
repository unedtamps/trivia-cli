package main

import (
	"TugasRPL/config"
	"TugasRPL/database/seed"
	"TugasRPL/src/service"
	"TugasRPL/src/ui"
	"TugasRPL/util"
	"context"
	"os"

	"github.com/gookit/slog"
)

var (
	LOGIN    = "Login"
	EXIT     = "Exit"
	REGISTER = "Register"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "seed_admin" && config.ENV == "development" {
			seed.SeedAdmin()
		} else if os.Args[1] == "seed_trivia" && config.ENV == "development" {
			ctx := context.Background()
			err := seed.SeedTrivia(ctx)
			if err != nil {
				slog.Fatal(err)
			}
		} else {
			ui.RedBg.Println("Command not found")
		}
		return
	}

	util.ClearScreen()
	ctx := context.Background()
	for {
		choice := ui.Select("Select Menu !", []string{LOGIN, REGISTER, EXIT})
		util.ClearScreen()
		switch choice {
		case LOGIN:
			err := service.Login(ctx)
			util.ClearScreen()
			if err != nil {
				ui.RedBg.Println(err)
			}
		case REGISTER:
			err := service.Register(ctx)
			util.ClearScreen()
			if err != nil {
				ui.RedBg.Println(err)
			}
		case EXIT:
			ans := ui.Promt(ui.YESORNO)
			if ans == "n" {
				util.ClearScreen()
				continue
			}
			util.ClearScreen()
			return
		}
	}
}

package service

import (
	"TugasRPL/database/repository"
	"TugasRPL/src/ui"
	"TugasRPL/util"
	"context"
	"fmt"
	"sort"
)

func crateTrivia(ctx context.Context) error {
	title := ui.Promt("title")
	question := ui.Promt("question")
	answere := ui.Promt("answer")
	dificulity := ui.SelectDifficulity()
	limit := limitQuestion(dificulity, 1)
	options := []string{}
	options = append(options, answere)
	ui.BlueBg.Println("Create option different from answer")
	for i := 0; i < (limit - 1); i++ {
		lable := fmt.Sprintf("option %d", i+1)
		lablePromt := ui.Promt(lable)
		for {
			if lablePromt == answere {
				ui.RedBg.Println("Option can't be the same as the answer")
				lablePromt = ui.Promt(lable)
				continue
			}
			break
		}
		options = append(options, lablePromt)
	}
	sort.Strings(options)
	err := repository.Query.TrviaCreateWithTx(ctx, repository.CreateTriviaParams{
		Title:        title,
		Question:     question,
		Answer:       answere,
		DificulityID: dificulity,
	}, options)
	util.ClearScreen()
	return err
}

func SeeMyTrivia(ctx context.Context) error {
	id := ctx.Value("id").(int64)
	trivias, err := repository.Query.FindMyTrivia(ctx, id)
	if err != nil {
		return err
	}
	l := len(trivias)
	for i := 0; i < l; i++ {
		ui.Magenta.Println(
			i+1,
			" ",
			trivias[i].Title,
			" ",
			repository.MPSTATUS[trivias[i].TriviaStatusID],
		)
	}

	for {
		menu := ui.Select("Select Menu !", []string{"Delete Rejected", "Back"})
		switch menu {
		case "Delete Rejected":
			confirm := ui.Promt(ui.YESORNO)
			if confirm == "y" {
				err := repository.Query.DeleteTrviaWithTx(ctx, id)
				if err != nil {
					return err
				}
			}
			util.ClearScreen()
			return nil
		case "Back":
			util.ClearScreen()
			return nil

		}
	}
}

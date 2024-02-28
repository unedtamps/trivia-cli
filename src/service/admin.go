package service

import (
	"TugasRPL/database/repository"
	"TugasRPL/src/ui"
	"TugasRPL/util"
	"context"
	"fmt"
	"strconv"
	"strings"
)

var (
	SUB  = "Triva Submission"
	RANK = "See Triva Rank"
)

type trivas struct {
	title      string
	question   string
	aswere     string
	status     int64
	dificulity int64
}

type choice struct {
	choice string
	id     int64
}

func adminMenu(ctx context.Context) error {
	opts := []string{SUB, "Back"}
	for {
		selectMenu := ui.Select("Select Menu !", opts)
		switch selectMenu {
		case SUB:
			subMenu := []string{"See All Subs", "Review Trivia", "Back"}
			selectSub := ui.Select("Select Menu !", subMenu)
			switch selectSub {
			case "See All Subs":
				if err := seeAllSubs(ctx); err != nil {
					return err
				}
			case "Back":
				util.ClearScreen()
				break
			case "Review Trivia":
				err := review(ctx)
				if err != nil {
					return err
				}
			}
		case "Back":
			util.ClearScreen()
			return nil
		}
	}
}

func review(ctx context.Context) error {
	options := []string{"Next Review", "Back"}
	for {
		selectMenu := ui.Select("Select Menu !", options)
		switch selectMenu {
		case "Back":
			return nil
		case "Next Review":
			triviaSub, err := repository.Query.FindReviewSubmission(ctx)
			if err != nil {
				return nil
			}
			// map choice
			choices := make(map[int64][]choice)
			trivia := make(map[int64]trivas)
			for _, v := range triviaSub {
				choices[v.ID] = append(choices[v.ID], choice{v.Choice, v.ChoiceID})
				trivia[v.ID] = trivas{
					title:      v.Title,
					question:   v.Question,
					dificulity: v.DificulityID,
					aswere:     v.Answer,
					status:     v.TriviaStatusID,
				}
			}
			var optstoReview []string
			for i, v := range trivia {
				optstoReview = append(optstoReview, fmt.Sprintf("%d %s", i, v.title))
			}
			optstoReview = append(optstoReview, "Back")
			selectToreview := ui.Select("Select Trivia to Review", optstoReview)
			if selectToreview == "Back" {
				break
			}
			idReveiw, err := strconv.Atoi(strings.Split(selectToreview, " ")[0])
			if err != nil {
				return err
			}
			err = reviewTrivia(
				ctx,
				int64(idReveiw),
				choices[int64(idReveiw)],
				trivia[int64(idReveiw)],
			)
			if err != nil {
				return err
			}
		}
	}
}

func reviewTrivia(ctx context.Context, id int64, cho []choice, triva trivas) error {
	options := []string{"Approve", "Reject", "Edit", "Back"}
	for {
		util.ClearScreen()
		ui.Cyan.Println("Details")
		ui.Yellow.Println("Title: ", triva.title)
		ui.Blue.Println("Question: ", triva.question)
		ui.Green.Println("Answer: ", triva.aswere)
		ui.Blue.Println("Status: ", repository.MPSTATUS[triva.status])
		ui.Magenta.Println("Dificulity: ", repository.MPDIFFREV[triva.dificulity])
		for i, v := range cho {
			ui.Blue.Printf("Choice %d: %s\n", i+1, v.choice)
		}
		selectMenu := ui.Select("Select Menu !", options)
		switch selectMenu {
		case "Back":
			util.ClearScreen()
			return nil
		case "Approve":
			triva.status = repository.OK
			err := repository.Query.UpdateTrivaStatus(ctx, repository.UpdateTrivaStatusParams{
				TriviaStatusID: repository.OK,
				TriviaID:       id,
			})
			if err != nil {
				util.ClearScreen()
				return err
			}
			util.ClearScreen()
		case "Reject":
			triva.status = repository.REJECTED
			err := repository.Query.UpdateTrivaStatus(ctx, repository.UpdateTrivaStatusParams{
				TriviaStatusID: repository.REJECTED,
				TriviaID:       id,
			})
			if err != nil {
				util.ClearScreen()
				return err
			}
			util.ClearScreen()
		case "Edit":
			ui.BlueBg.Println("Edit Trivia enter to skip")
			title := ui.PromtEmpty("title")
			if title != "" {
				triva.title = title
			}
			question := ui.PromtEmpty("question")
			if question != "" {
				triva.question = question
			}
			aswere := ui.PromtEmpty("answer")
			if aswere != "" {
				triva.aswere = aswere
			}
			diffID := ui.SelectDifficulity()
			limitDif := limitQuestion(diffID, 1)
			limitPrev := limitQuestion(triva.dificulity, 1)
			for {
				diffID = ui.SelectDifficulity()
				limitDif = limitQuestion(diffID, 1)
				limitPrev = limitQuestion(triva.dificulity, 1)
				if limitDif < limitPrev {
					ui.RedBg.Println("You can't change to lower difficulity")
					continue
				}
				triva.dificulity = diffID
				break
			}

			ui.BlueBg.Println("Edit Choice enter to skip")
			var new_choice []repository.Choice
			for i, v := range cho {
				c := ui.PromtEmpty(fmt.Sprintf("Choice %d", i+1))
				if c != "" {
					new_choice = append(new_choice, repository.Choice{
						Id:     v.id,
						Choice: c,
					})
				} else {
					new_choice = append(new_choice, repository.Choice{
						Id:     v.id,
						Choice: v.choice,
					})
				}
			}
			if limitDif > limitPrev {
				ui.BlueBg.Println("Add new choice because you change difficulity")
				x := limitDif - limitPrev
				a := len(new_choice)
				for i := 0; i < x; i++ {
					newc := ui.Promt(fmt.Sprintf("Choice %d", a+1))
					err := repository.Query.CreateTriviaChoice(
						ctx,
						repository.CreateTriviaChoiceParams{
							TriviaID: id,
							Choice:   newc,
						},
					)
					if err != nil {
						return err
					}
					a++
				}
			}
			err := repository.Query.UpdateTrivaWithTx(ctx, id, repository.Triva{
				Title:      triva.title,
				Question:   triva.question,
				Aswere:     triva.aswere,
				Dificulity: triva.dificulity,
			}, new_choice)
			util.ClearScreen()
			return err
		}

	}
}

func seeAllSubs(ctx context.Context) error {
	trivia, err := repository.Query.FindAllSubmission(ctx)
	if err != nil {
		return err
	}
	ui.Magenta.Println("No\tTitle\tDificulity\tAuthor")
	for i, v := range trivia {
		ui.Red.Printf("%d\t", i+1)
		ui.Yellow.Printf("%s\t", v.Title)
		ui.BlueBg.Printf("%s\t", v.Diffname)
		ui.GreenBg.Println(v.Name)
	}
	fmt.Print("\n")
	return nil
}

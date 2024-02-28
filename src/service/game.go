package service

import (
	"TugasRPL/database/repository"
	"TugasRPL/src/ui"
	"TugasRPL/util"
	"context"
	"strconv"
	"time"
)

func StartGame(ctx context.Context) error {

loop:
	for {
		opts := []string{"Start Game", "Back"}
		selectMenu := ui.Select("Select Menu !", opts)
		switch selectMenu {
		case "Start Game":
			choices := make(map[int64][]string)
			questions := make(map[int64]*repository.FindTriviaWithSpecsRow)
			id := ctx.Value("id").(int64)
			diffId := ui.SelectDifficulity()
			limitStr := ui.Promt(ui.NUMBER)
			limit, _ := strconv.Atoi(limitStr)
			ui.BlueBg.Println("Gap for 3 seconds in each question")
			trivias, err := repository.Query.FindTriviaWithSpecs(
				ctx,
				repository.FindTriviaWithSpecsParams{
					DificulityID: diffId,
					Limit:        int32(limit),
					UserID:       id,
				},
			)
			if err != nil {
				return err
			}
			for _, v := range trivias {
				c, err := repository.Query.FindChoicesOfTrivia(ctx, v.ID)
				if err != nil {
					return err
				}
				choices[v.ID] = c
				questions[v.ID] = v
			}
			util.ClearScreen()
			var j, k int = 0, 0
			for i, v := range choices {
				j++
				answer := ui.Select(questions[i].Question, v)
				if answer == questions[i].Answer {
					k++
					ui.GreenBg.Println("Correct")
				} else {
					ui.RedBg.Println("Incorrect")
				}
				time.Sleep(3 * time.Second)
				util.ClearScreen()
			}
			ui.BlueBg.Println("Game Over")
			var score float32 = float32(k) / float32(j) * 100
			if k == 0 && j == 0 {
				score = 0
			}
			ui.Yellow.Printf("Skor: %2f\n", score)
			user_detail, err := repository.Query.FindUserDetailById(ctx, id)
			if err != nil {
				return err
			}
			ui.Yellow.Println("Please Wait ...")
			err = repository.Query.UpdateUserDetail(ctx, repository.UpdateUserDetailParams{
				UserID:      id,
				TriviaRound: user_detail.TriviaRound + 1,
				RightAnswer: int64(int(user_detail.RightAnswer) + k),
				WrongAnswer: int64(int(user_detail.WrongAnswer) + (j - k)),
			})
			if err != nil {
				return err
			}
			time.Sleep(3 * time.Second)
			util.ClearScreen()
		case "Back":
			break loop
		}
	}
	return nil
}

func limitQuestion(diffId int64, limit int) int {
	if diffId == repository.EASY {
		return limit * 3
	} else if diffId == repository.MEDIUM {
		return limit * 4
	} else {
		return limit * 5
	}
}

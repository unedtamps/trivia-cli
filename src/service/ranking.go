package service

import (
	"TugasRPL/database/repository"
	"TugasRPL/src/ui"
	"context"
	"fmt"
)

var (
	DESC    = "Descending"
	ASC     = "Ascending"
	AVGUP   = "Score Above Average"
	AVGDOWN = "Score Below Average"
)

func SeeRanking(ctx context.Context) error {
	options := []string{DESC, ASC}
	id := ctx.Value("id").(int64)
	role_id, err := repository.Query.CheckIfAdmin(ctx, id)
	if err != nil {
		return err
	}
	if role_id == repository.ADMIN {
		options = append(options, AVGUP, AVGDOWN)
	}
	options = append(options, "Back")
	for {
		switch ui.Select("Select Menu !", options) {
		case DESC:
			ranks, err := repository.Query.SelectUserRankDesc(ctx)
			if err != nil {
				return err
			}
			var ranksList = []string{}
			for i, v := range ranks {
				ranksList = append(ranksList, fmt.Sprintf("%d %s (%d)", i+1, v.Name, v.RightAnswer))
			}
			ranksList = append(ranksList, "Back")
			for {
				res := ui.SelectWithSearch("Ranking By Highest Score", ranksList)
				if res == "Back" {
					break
				}
			}
		case ASC:
			ranks, err := repository.Query.SelectUserRankAsc(ctx)
			if err != nil {
				return err
			}
			var ranksList = []string{}
			for i, v := range ranks {
				ranksList = append(ranksList, fmt.Sprintf("%d %s (%d)", i+1, v.Name, v.RightAnswer))
			}
			ranksList = append(ranksList, "Back")
			for {
				res := ui.SelectWithSearch("Ranking By Lowest Score", ranksList)
				if res == "Back" {
					break
				}
			}
		case AVGUP:
			ranks, err := repository.Query.SelectUserAboveAverage(ctx)
			if err != nil {
				return err
			}
			var ranksList = []string{}
			for _, v := range ranks {
				ranksList = append(
					ranksList,
					fmt.Sprintf(
						"%s (%2f)",
						v.Name,
						float32(v.RightAnswer)/float32(v.RightAnswer+v.WrongAnswer)*100,
					),
				)
			}
			ranksList = append(ranksList, "Back")
			for {
				res := ui.SelectWithSearch("User with Above Average Score", ranksList)
				if res == "Back" {
					break
				}
			}
		case AVGDOWN:
			ranks, err := repository.Query.SelectUserBelowAverage(ctx)
			if err != nil {
				return err
			}
			var ranksList = []string{}
			for _, v := range ranks {
				ranksList = append(
					ranksList,
					fmt.Sprintf(
						"%s (%2f)",
						v.Name,
						float32(v.RightAnswer)/float32(v.RightAnswer+v.WrongAnswer)*100,
					),
				)
			}
			ranksList = append(ranksList, "Back")
			for {
				res := ui.SelectWithSearch("User with Below Average Score", ranksList)
				if res == "Back" {
					break
				}
			}
		case "Back":
			return nil
		}
	}
}

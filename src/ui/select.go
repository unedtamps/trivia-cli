package ui

import (
	"TugasRPL/database/repository"
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

func Select(lable string, options []string) string {

	lable = Yellow.Sprintf("%s", lable)
	templates := promptui.SelectTemplates{
		Label:    "   {{ . | bold}}",
		Active:   "⏩ {{ . | cyan | bold}}",
		Inactive: "  {{ . | green }}",
		Selected: "",
	}
	selector := promptui.Select{
		HideSelected: true,
		Label:        lable,
		Items:        options,
		CursorPos:    0,
		HideHelp:     true,
		Templates:    &templates,
		IsVimMode:    true,
		Size:         10,
	}
	_, res, err := selector.Run()
	if err != nil {
		fmt.Println(err)
	}
	return res
}

func SelectDifficulity() int64 {
	diff := Select("Select Difficulty !", []string{"Easy", "Medium", "Hard"})
	diffId := repository.MPDIFF[diff]
	return diffId
}

func SelectWithSearch(lable string, options []string) string {

	searcher := func(input string, index int) bool {
		pepper := options[index]
		name := strings.Replace(strings.ToLower(pepper), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)
		return strings.Contains(name, input)
	}

	lable = Yellow.Sprintf("%s", lable)
	templates := promptui.SelectTemplates{
		Label:    "   {{ . | bold}}",
		Active:   "⏩ {{ . | cyan | bold}}",
		Inactive: "  {{ . | green }}",
		Selected: "",
	}
	selector := promptui.Select{
		HideSelected:      true,
		Label:             lable,
		Items:             options,
		CursorPos:         0,
		HideHelp:          true,
		Templates:         &templates,
		StartInSearchMode: true,
		IsVimMode:         false,
		Searcher:          searcher,
		Size:              20,
	}
	_, res, err := selector.Run()
	if err != nil {
		fmt.Println(err)
	}
	return res
}

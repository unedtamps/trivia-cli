package ui

import (
	"errors"
	"fmt"
	"net/mail"
	"strconv"

	"github.com/manifoldco/promptui"
)

var (
	PASSWORD = "password"
	EMAIL    = "email"
	NUMBER   = "number of items"
	YESORNO  = "Are you sure [y/n]"
)

func Promt(label string) string {
	var validate func(string) error
	switch label {
	case EMAIL:
		validate = func(input string) error {
			_, err := mail.ParseAddress(input)
			return err
		}
	case PASSWORD:
		validate = func(input string) error {
			if len(input) < 8 {
				return errors.New("Password must be at least 8 characters")
			}
			return nil
		}
	case NUMBER:
		validate = func(input string) error {
			_, err := strconv.Atoi(input)
			return err
		}
	case YESORNO:
		validate = func(input string) error {
			if input != "y" && input != "n" {
				return errors.New("Please enter y/n")
			}
			return nil
		}
	default:
		validate = func(input string) error {
			if input == "" {
				return errors.New("Please enter a value")
			}
			return nil
		}
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . | bold }} ",
		Valid:   "{{ . | blue | bold }} ",
		Invalid: "{{ . | red | bold | italic }} ",
		Success: "{{ . | green | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     fmt.Sprintf("%s:", label),
		Templates: templates,
		Validate:  validate,
	}

	if label == PASSWORD {
		prompt.Mask = '*'
	}

	res, err := prompt.Run()
	if err != nil {
		RedBg.Println(err)
	}
	return res
}

func PromtVerify(password string) {
	validate := func(input string) error {
		if len(input) < 8 {
			return errors.New("Password must be at least 8 characters")
		} else if input != password {
			return errors.New("Password not match")
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . | bold }} ",
		Valid:   "{{ . | blue | bold }} ",
		Invalid: "{{ . | red | bold | italic }} ",
		Success: "{{ . | green | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     "verify password:",
		Templates: templates,
		Validate:  validate,
		Mask:      '*',
	}
	_, err := prompt.Run()
	if err != nil {
		RedBg.Println(err)
	}
}

func PromtEmpty(label string) string {
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . | bold }} ",
		Valid:   "{{ . | blue | bold }} ",
		Invalid: "{{ . | red | bold | italic }} ",
		Success: "{{ . | green | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     fmt.Sprintf("%s:", label),
		Templates: templates,
	}

	res, err := prompt.Run()
	if err != nil {
		RedBg.Println(err)
	}
	return res

}

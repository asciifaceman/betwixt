package csl

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/fatih/color"
)

type KV struct {
	K string
	V interface{}
}

func (kv *KV) Len() int {
	t := fmt.Sprintf("%s%v", kv.K, kv.V)
	return len(t)
}

func (kv *KV) Print(space int) {
	buff := strings.Repeat(" ", space-kv.Len())
	msg := fmt.Sprintf("%s:%s%v", kv.K, buff, kv.V)
	Success(msg)
}

func Info(message string, a ...interface{}) {
	if a != nil {
		color.White(message, a)
	} else {
		color.White(message)
	}

}

func Error(message string, a ...interface{}) {
	if a != nil {
		color.Red(message, a)
	} else {
		color.Red(message)
	}
}

func Success(message string, a ...interface{}) {
	if len(a) > 0 {
		color.Green(message, a)
	} else {
		color.Green(message)
	}
}

func KeyValue(key string, value string) {
	color.Green("*--- %s: %s", key, value)
}

func PrintStruct(s interface{}) {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	values := make([]*KV, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() {

			values[i] = &KV{
				K: t.Field(i).Name,
				V: v.Field(i).Interface(),
			}
		}
	}

	longestPair := 0

	for _, val := range values {
		thisLen := val.Len()
		if thisLen > longestPair {
			longestPair = thisLen
		}
	}

	Success(strings.Repeat("*", longestPair+2))
	for _, val := range values {
		val.Print(longestPair + 1)
	}
	Success(strings.Repeat("*", longestPair+2))
}

// YesNoPrompt gives the user a y/n option loop and returns a bool
func YesNoPrompt() bool {
	var confirmation string
	for {
		Info("(y/n) or ctl+c to exit: ")
		fmt.Scanln(&confirmation)

		confirm := strings.ToLower(confirmation)

		if confirm == "" || confirm != "y" && confirm != "n" {
			Info("Input must be y/Y or n/N")
			continue
		}

		if confirm == "y" {
			return true
		}

		return false
	}
}

// OptionsPrompt allows a user to select from several valid
// options available. options[0] is provided as the default option.
func OptionsPrompt(prompt string, options []string) string {
	var response string

	for {
		Info(fmt.Sprintf("%s. Valid options include (%s)", prompt, strings.Join(options, ",")))
		fmt.Printf(fmt.Sprintf("(%s): ", options[0]))
		fmt.Scanln(&response)

		if response == "" {
			return options[0]
		}

		for _, val := range options {
			if response == val {
				return response
			}
		}

		continue

	}
}

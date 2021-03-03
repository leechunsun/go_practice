package service

import (
	"errors"
	"strings"
)

type StringService interface {
	UpperCase(string) (string, error)
	Count(string) int
}


var EmptyError = errors.New("string is empty !!!")


type StringServiceImpl struct {

}

func (s *StringServiceImpl) UpperCase(str string) (string, error) {
	if str == ""{
		return "", EmptyError
	}
	return strings.ToUpper(str), nil
}


func (s *StringServiceImpl) Count(str string) int {
	return len([]rune(str))
}

package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "atom"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("atom")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("atom")=%q,%v,want match for %#q, nil`, msg, err, want)
	}
}

//无输入参数调用greetings.Hello 检查错误
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("")=%q,%v,want "",error`, msg, err)
	}
}

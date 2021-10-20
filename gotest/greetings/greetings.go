package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	//如果没有名字输入，报错
	if name == "" {
		return "", errors.New("错误：无指定姓名")
	}
	//如果有指定姓名输入
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

//返回“人-问候语”map
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

//初始化
func init() {
	rand.Seed(time.Now().UnixNano())
}

//随机返回greeting值
func randomFormat() string {
	formats := []string{
		"Hi %v, welcome",
		"How are you, %v",
		"Nice to have you %v",
	}

	return formats[rand.Intn(len(formats))]
}

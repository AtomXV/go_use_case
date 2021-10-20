package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	//设置预定义 Logger 的属性，包括日志条目前缀和用于禁用打印时间、源文件和行号的标志。
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"atom", "ashely", "tim"}

	//request a greeting message
	messages, err := greetings.Hellos(names) //设空报错

	//如果报错，在控制台打印并退出程序
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}

package main

import (
	"fmt"
)

//学生结构体
type Student struct {
	name   string
	gender string
	age    uint
	id     uint
	score  float64
}

func (student *Student) say() string { //传指针速度快
	stuInfo := fmt.Sprintf("name is %v, gender is %v, age is %v, id is %v, score is %v",
		student.name, student.age, student.gender, student.id, student.score)

	return stuInfo
}

//输入长宽高，计算立方体体积
type Cube struct {
	length uint
	width  uint
	height uint
}

func (cube *Cube) culc() uint {
	V_of_cube := cube.length * cube.width * cube.height
	return V_of_cube
}

func main() {
	var stu = Student{
		name:   "yyp",
		gender: "male",
		age:    40,
		id:     3,
		score:  100.01,
	}
	fmt.Println(stu.say())

	var cub = Cube{
		length: 10,
		width:  20,
		height: 100,
	}
	fmt.Printf("立方体体积为：%v", cub.culc())
}

package main

import "fmt"

//定义一个接口
type testInterface interface {
	testString() string
	testInt() int
}

//定义结构体
type testStruct struct {
	name string
	age  int
}

//定义结构体方法
func (t *testStruct) testString() string {
	return t.name
}

func (t *testStruct) testInt() int {
	return t.age
}

func main() {
	t1 := &testStruct{"abc", 111} //指针类型
	t2 := testStruct{"efg", 222}  //值类型

	var test1 testInterface //接口只能是值类型
	test1 = t1              //testStruct类的指针类型实例传给接口
	fmt.Println(test1.testString())
	fmt.Println(test1.testInt())

	test2 := t2 //testInt类的指针类型实例传给接口
	fmt.Println(test2.testString())
	fmt.Println(test2.testInt())
}

package main

import "fmt"

type Number struct {
	num int
}

func (n Number) T1() {
	n.num++
}

func (n *Number) T2() {
	n.num += 2
}

func main() {
	t1 := Number{1}  //值接收
	t2 := &Number{1} //指针接收

	t1.T2()
	fmt.Printf("值接收返回为：%d\n", t1.num)

	t1.T1() //T1 ++ 操作后 不影响t1中num的原值
	fmt.Printf("值接收返回为：%d\n", t1.num)

	t2.T2() //T2 +=2 操作后 改变了t2中num的原值
	fmt.Printf("指针接受返回为：%d", t2.num)
}

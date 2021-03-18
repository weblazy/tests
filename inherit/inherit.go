package main

import (
	"github.com/weblazy/tests/parent"
)

//子类
type Student05 struct {
	parent.Person05
	id int
}

// func (p Student05) printInfo() {
// 	fmt.Printf("名称1：%s, 性别：%s, 年龄：%d\n", p.Name, p.sex, p.age)
// }

func main() {
	var s Student05 = Student05{Person05: parent.Person05{Name: "haha"}}

	//使用父类方法
	s.PrintInfo()
}

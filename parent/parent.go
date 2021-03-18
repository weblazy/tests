package parent

import "fmt"

//父类
type Person05 struct {
	Name string
	sex  string
	age  int
}

func (p Person05) PrintInfo() {
	fmt.Printf("名称：%s, 性别：%s, 年龄：%d\n", p.Name, p.sex, p.age)
}

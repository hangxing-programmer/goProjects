package basic

import "fmt"

type Person struct {
	name   string
	age    int
	scores []float64
	ptr    *int              //指针
	slice  []int             //切片
	map1   map[string]string //map
}
type Cat struct {
	name  string
	color string
}

//方法的声明
//func (绑定对象,对象类型)action(参数列表)(返回值列表){
//	方法体
//	return 返回值
//}

// 绑定方法
func (cat Cat) action() {
	fmt.Println("Cat的名字", cat.name)
}

// 方法体
func (cat Cat) add() {
	start := 0
	for i := 1; i < 100; i++ {
		start += i
	}
	fmt.Println(cat.name, "计算结果==>", start)
}

// 传值
func (cat Cat) getAndAdd(n int) {
	start := 0
	for i := 0; i < n; i++ {
		start += i
	}
	fmt.Println(cat.name, "计算结果==>", start)
}

// 返回值
func (cat Cat) getAndAdd2(n, n1 int) int {
	return n + n1
}

// 提高效率，绑定对象指针
func (cat *Cat) getAndAdd3(n1, n2 int) int {
	return n1 + n2
}

// int类型也可以绑定方法
type intenger int

func (i *intenger) get() {
	*i = *i + 1
}

type student struct {
	Name  string
	Age   int
	score float64
}

// 工厂模式
func NewStudent(name string, age int, score float64) *student {
	return &student{
		Name:  name,
		Age:   age,
		score: score,
	}
}

// 如果需要把字段首字母小写,且可以被其他包访问
func (s *student) GetStu() float64 {
	return s.score
}

// json序列化
//type student struct {
//	name string `json:"name"`
//	age  int    `json:"age"`
//}

func (stu *student) getName() string {
	return stu.Name
}

func main14() {

	//面向对象,四种创建对象方式
	var person Person
	//切片和map必须要make
	person.slice = make([]int, 5)
	person.slice[0] = 100
	person.map1 = make(map[string]string)
	person.map1["friend"] = "zs"
	fmt.Println(person)

	var cat2 *Cat = new(Cat)
	(*cat2).name = "黑猫"
	(*cat2).color = "黑色"
	fmt.Println(*cat2)
	cat2.name = "黄猫"  //底层有简化处理
	cat2.color = "黄色" //底层有简化处理
	fmt.Println(*cat2)

	var cat3 *Cat = &Cat{}
	(*cat3).name = "橙猫"
	(*cat3).color = "橙色"
	fmt.Println(*cat3)
	cat3.name = "蓝猫"
	cat3.color = "蓝色"
	fmt.Println(*cat3)

	var cat Cat = Cat{
		"白猫",
		"白色",
	}
	fmt.Println(cat)
	cat.action()
	cat.add()
	cat.getAndAdd(10)
	add2 := cat.getAndAdd2(1, 10)
	fmt.Println("计算结果是", add2)
	add3 := (&cat).getAndAdd3(11, 22) //等价于add3 := cat.getAndAdd3(11, 22) , 编译器会自动添加指针
	fmt.Println("指针传递==>", add3)

	var i intenger = 1
	i.get()
	fmt.Println("i==>", i)

	stu := student{"zs", 11, 99.9}
	name := (&stu).getName()
	fmt.Println("该学生的名字==>", name)

	//Go的结构体(结构体首字母小写无法实现共享)中没有构造函数，可以使用工厂模式
}

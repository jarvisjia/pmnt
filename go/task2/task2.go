package task2

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 1.指针
// 编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
// 考察点 ：指针的使用、值传递与引用传递的区别
func modifyPointer(num *int) {
	*num += 10
}

// 2.切片及指针
// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。
func modifySlice(nums *[]int) {
	for i := range *nums {
		(*nums)[i] *= 2
	}
}

// 3.Go线程goroutine
// 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func goRoutineFun() {
	//管理并行任务的工具，能够等待一组goroutine完成后再继续执行主线程。
	//它通过计数器机制实现对goroutine的同步控制
	var wg sync.WaitGroup

	//启动两个goroutine,所有waitgroup加2
	wg.Add(2)

	go func() {
		defer wg.Done() //当goroutine完成，技术减1
		for i := 1; i <= 10; i += 2 {
			fmt.Println("奇数=", i)
			time.Sleep(100 * time.Microsecond)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			fmt.Println("偶数=", i)
			time.Sleep(100 * time.Microsecond)
		}
	}()

	wg.Wait() //阻塞知道计数器为零，即等待所有goroutine都执行完成
	fmt.Println("operation finished")
}

// 4.设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。
type Task struct {
	TaskName     string        //任务名称
	Func         func()        //任务函数
	StartTime    time.Time     //开始时间
	EndTime      time.Time     //结束事件
	ExcutionTime time.Duration //执行时间
}
type TaskScheduler struct {
	tasks []*Task        //任务列表, 指针形式
	wg    sync.WaitGroup //等待组
}

func (ts *TaskScheduler) AddTask(taskName string, taskFun func()) {
	task := &Task{ //&task地址
		TaskName: taskName,
		Func:     taskFun,
	}
	ts.tasks = append(ts.tasks, task)
}

func NewTaskScheduler(tasks ...func()) *TaskScheduler {
	ts := &TaskScheduler{}
	for i, task := range tasks {
		ts.AddTask(fmt.Sprintf("Task%d", i+1), task)
	}
	return ts
}

func (ts *TaskScheduler) Run() {
	for _, task := range ts.tasks {
		ts.wg.Add(1)
		go func(task *Task) {
			defer ts.wg.Done()
			task.StartTime = time.Now()
			task.Func()
			task.EndTime = time.Now()
			task.ExcutionTime = task.EndTime.Sub(task.StartTime)
		}(task)
	}
	ts.wg.Wait()
	fmt.Println("All tasks finished")

	for _, task := range ts.tasks {
		fmt.Printf("TaskName=%s, ExcutionTime=%v", task.TaskName, task.ExcutionTime)
	}
}

// 5.面向对象1
// 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。
type Shap interface {
	Area()
	Perimeter()
}
type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	R float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}
func (r Rectangle) Perimeter() float64 {
	return (r.Height + r.Width) * 2
}
func (c Circle) Area() float64 {
	return c.R * c.R * 3.14
}
func (c Circle) Perimeter() float64 {
	return 2 * c.R * 3.14
}

//6.面向对象2
//使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
//为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
//考察点 ：组合的使用、方法接收者。

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person
	EmployeeID int
}

func (e Employee) PrintInfo() {
	fmt.Printf("Employee info, EmployeeID:%d, Name:%s, Age:%d", e.EmployeeID, e.Name, e.Age)
}

// 7.Channel1
// 编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。
func channelFun1() {
	c := make(chan int)
	go func(c chan int) {
		for i := range 10 {
			fmt.Println("c<-", i)
			c <- i
			time.Sleep(1000 * time.Microsecond)
		}
		close(c) //要关闭通道，否则，下面的for循环就不会结束，从而在接收第11个数据的时候就阻塞了
	}(c)
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("channelFun finished")
}

// 8.Channel2
// 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。
func channelFun2() {
	c := make(chan int, 10)
	go func(c chan int) {
		for i := range 100 {
			fmt.Println("c<-", i)
			c <- i
		}
		close(c) //要关闭通道，否则，下面的for循环就不会结束，从而在接收第11个数据的时候就阻塞了
	}(c)
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("channelFun finished")
}

// 9.锁机制1
// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。
func lockMutex() {
	var wg sync.WaitGroup
	var lock sync.Mutex
	var num int
	wg.Add(10)
	for i := range 10 {
		go func(index int) {
			defer wg.Done()
			j := 0
			for ; j < 1000; j++ {
				lock.Lock()
				num++
				lock.Unlock()
			}
			fmt.Println("Thread", i, "add", j)
		}(i)
	}
	wg.Wait()
	fmt.Println("Total add", num)
}

// 10.锁机制2
// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。
func atomicOp() {
	var wg sync.WaitGroup
	var num int64
	wg.Add(10)
	for i := range 10 {
		go func(index int) {
			defer wg.Done()
			j := 0
			for ; j < 1000; j++ {
				atomic.AndInt64(&num, 1)
			}
			fmt.Println("Thread", i, "add", j)
		}(i)
	}
	wg.Wait()
	fmt.Println("Total add", num)
}

func Task2() {
	//指针
	var num = 10
	modifyPointer(&num)
	fmt.Println("1.指针", num)

	//切片及指针
	nums := []int{2, 4, 6, 8}
	modifySlice(&nums)
	fmt.Println("2.切片及指针", nums)

	//goroutine
	fmt.Println("3.goroutine")
	goRoutineFun()

	//任务调度器
	task1 := func() {
		time.Sleep(500 * time.Microsecond)
		fmt.Println("Task 1 finished.")
	}
	task2 := func() {
		time.Sleep(1600 * time.Microsecond)
		fmt.Println("Task 2 finished.")
	}
	task3 := func() {
		time.Sleep(4000 * time.Microsecond)
		fmt.Println("Task 3 finished.")
	}
	task4 := func() {
		time.Sleep(1000 * time.Microsecond)
		fmt.Println("Task4 finished.")
	}
	fmt.Println("4.任务调度器")
	ts := NewTaskScheduler(task1, task2, task3, task4)
	ts.Run()

	//面向对象1
	r := Rectangle{Width: 10, Height: 5}
	c := Circle{R: 6}
	fmt.Println("5.面向对象1")
	fmt.Println(r.Area(), r.Perimeter())
	fmt.Println(c.Area(), c.Perimeter())

	//面向对象2
	fmt.Println("6.面向对象2")
	e := Employee{
		Person: Person{
			Name: "张三",
			Age:  18,
		},
		EmployeeID: 123,
	}
	e.PrintInfo()

	fmt.Println("7.Channel1")
	channelFun1()

	fmt.Println("8.Channel2")
	channelFun2()

	fmt.Println("9.锁机制1")
	lockMutex()

	fmt.Println("10.锁机制2")
	atomicOp()
}

package task3

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// 1.基本CRUD操作
// 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、grade （学生年级，字符串类型）。
// 要求 ：
// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
type Student struct {
	gorm.Model
	Name  string `gorm:"type:varchar(64)"`
	Age   uint8
	Grade string `gorm:"type:varchar(256)"`
}

func crudOp() {
	db := DbConnet()

	//创建表
	db.AutoMigrate(&Student{})
	fmt.Println("Create new table")

	//插入数据
	// std1 := Student{Name: "张三", Age: 20, Grade: "三年级"}
	// db.Create(&std1)
	// std2 := Student{Name: "李四", Age: 13, Grade: "一年级"}
	// db.Create(&std2)
	// stdsInse := []*Student{{Name: "王五", Age: 20, Grade: "三年级"}, {Name: "马六", Age: 18, Grade: "二年年级"}}
	// db.Create(stdsInse)
	// fmt.Println("Insert new data")

	//查询 students 表中所有年龄大于 18 岁的学生信息
	var stds []Student
	// queryRes := db.Where("age > ?", 18).Find(&stds)
	queryRes := db.Debug().Find(&stds, "age > ?", 18)
	fmt.Println("Query", stds, queryRes.RowsAffected, queryRes.Statement.Model)
	fmt.Println(queryRes.Statement.Dest)

	//将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"
	result := db.Model(&Student{}).Where("name = ?", "张三").Update("Grade", "五年级")
	fmt.Println("Update", result.RowsAffected, result.Error == nil)
	uStd := Student{Name: "张三"}
	result1 := db.Model(&uStd).Where("age = ?", "20").Update("Grade", "六年级")
	fmt.Println("Update", result1.RowsAffected, result1.Error == nil)

	//删除 students 表中年龄小于 15 岁的学生记录
	delResult := db.Delete(&Student{}, "age < ?", 15)
	fmt.Println("Delete data", delResult.RowsAffected, delResult.Error == nil)

	//延时关闭数据库连接
	// defer db.Close()
}

// 2.事务语句
// 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID，
// to_account_id 转入账户ID， amount 转账金额）。
// 要求 ：
// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，
// 向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
type Account struct {
	ID      uint `gorm:primarykey`
	Balance float64
}

type Transaction struct {
	gorm.Model
	FromAccountId uint
	ToAccountId   uint
	Amount        float64
}

func transfer(db *gorm.DB, fAccId uint, tAccId uint, amount float64) {
	tx := db.Begin()
	var f, t Account
	fRes := tx.First(&f, fAccId) //主键查询
	if fRes.Error != nil {
		fmt.Println("FromAccountId no exists.")
		tx.Rollback()
		return
	}
	tRes := tx.First(&t, tAccId) //主键查询
	if tRes.Error != nil {
		fmt.Println("toAccountId no exists.")
		tx.Rollback()
		return
	}
	if f.Balance < amount {
		fmt.Println("transfer failed")
		tx.Rollback()
		return
	}
	f.Balance -= amount
	t.Balance += amount
	updateFRes := tx.Model(&Account{}).Where("id=?", fAccId).Update("balance", f.Balance)
	if updateFRes.Error != nil {
		tx.Rollback()
		return
	}
	updateTRes := tx.Model(&Account{}).Where("id=?", tAccId).Update("balance", t.Balance)
	if updateTRes.Error != nil {
		tx.Rollback()
		return
	}
	trans := Transaction{FromAccountId: fAccId, ToAccountId: tAccId, Amount: amount}
	createRes := tx.Create(&trans)
	if createRes.Error != nil {
		tx.Rollback()
		return
	}
	commitRes := tx.Commit()
	if commitRes.Error != nil {
		tx.Rollback()
		return
	}
	fmt.Println("transfer success")
}
func txOp() {
	db := DbConnet()
	fmt.Println("create tables")
	// db.AutoMigrate(&Account{})
	db.AutoMigrate(&Transaction{})

	// a := Account{Balance: 200}
	// b := Account{Balance: 100}
	// fmt.Println("insert two line datas")
	// db.CreateInBatches([]Account{a, b}, 2)

	fmt.Println("transfer amount")
	transfer(db, 1, 2, 50)
}

// 使用SQL扩展库进行查询
// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
// 要求 ：
// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
type Employee struct {
	ID         uint   `gorm:primaryKey`
	Name       string `gorm:"type:varchar(64)"`
	Department string `gorm:"type:varchar(256)"`
	Salary     float64
}

func queryByDepartment() []Employee {
	db := DbConnet()
	var emplees []Employee
	db.Debug().Where("Department=?", "技术部").Find(&emplees)
	return emplees
}

func queryMaxSalary() Employee {
	db := DbConnet()
	emplee := Employee{}
	db.Debug().Order("Salary desc").First(&emplee)
	return emplee
}

func sqlExtQuery() {
	// db := DbConnet()
	// db.AutoMigrate(&Employee{})
	// employees := []Employee{
	// 	{Name: "张三", Salary: 1000, Department: "运维部"},
	// 	{Name: "李四", Salary: 2000, Department: "技术部"},
	// 	{Name: "马五", Salary: 3000, Department: "技术部"},
	// }
	// db.Create(employees)

	empls := queryByDepartment()
	fmt.Println("查询技术部同事", empls)

	empl := queryMaxSalary()
	fmt.Println("查询工资最高的同事", empl)
}

// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
// 要求 ：
// 定义一个 Book 结构体，包含与 books 表对应的字段。
// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
type Book struct {
	ID     uint   `gorm:primaryKey`
	Title  string `gorm:"type:varchar(256)"`
	Author string `gorm:"type:varchar(64)"`
	Price  float64
}

func queryByPrice(price uint) []Book {
	db := DbConnet()
	var books []Book
	db.Debug().Where("price>?", price).Find(&books)
	return books
}

func sqlExtQuery2() {
	// db := DbConnet()
	// db.AutoMigrate(&Book{})
	// books := []Book{
	// 	{Title: "哈哈世界", Price: 10, Author: "张三"},
	// 	{Title: "是的你好", Price: 20, Author: "李四"},
	// 	{Title: "丫丫", Price: 80, Author: "马五"},
	// }
	// db.Create(books)

	empls := queryByPrice(50)
	fmt.Println("查询价格大于50", empls)
}

// 模型定义
// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 要求 ：
// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
// 编写Go代码，使用Gorm创建这些模型对应的数据库表。
type User struct {
	ID        uint   `gorm:primarykey`
	Name      string `gorm:"type:varchar(64)"`
	PostCount uint
	Posts     []Post `gorm:foreignKey:UserID`
}

type Post struct {
	ID        uint   `gorm:primarykey`
	Title     string `gorm:"type:varchar(512)"`
	Content   string
	CreatedAt time.Time
	UserID    uint
	Comments  []Comment `gorm:foreignKey:PostID`
}

type Comment struct {
	ID        uint `gorm:primarykey`
	Content   string
	CreatedAt time.Time
	PostID    uint
}

func creaTables() {
	db := DbConnet()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})
}

// 关联查询
// 基于上述博客系统的模型定义。
// 要求 ：
// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
func queryBlogData() {
	db := DbConnet()
	users := []User{
		{Name: "张三"},
		{Name: "李四"},
	}
	db.Create(users)

	posts := []Post{
		{Title: "哈哈哈雅虎", Content: "水电费哈多发点", UserID: 2},
		{Title: "大法师代发", Content: "贵航股份涣发大号好", UserID: 1},
	}
	db.Create(posts)

	comments := []Comment{
		{Content: "很好，非常好", PostID: 2},
		{Content: "沙发发", PostID: 1},
		{Content: "发vFDA发斯蒂芬", PostID: 2},
	}
	db.Create(comments)

	u := []User{}
	db.Preload("Posts.Comments").Find(&u, 1)
	fmt.Println("查询某个用户发布的所有文章及其对应的评论信息:")
	fmt.Println(u)

	postOfMaxCom := Post{}
	db.Model(&Post{}).Where("id=(?)", db.Table("comments").Limit(1).Select("post_id").Group("post_id").Order(" count(*) desc")).First(&postOfMaxCom)
	fmt.Println("查询评论数量最多的文章信息：")
	fmt.Println(postOfMaxCom)
}

// 钩子函数
// 继续使用博客系统的模型。
// 要求 ：
// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (p *Post) BeforeCreate(db *gorm.DB) (err error) {
	fmt.Println("创建文章之前执行。。。")
	return nil
}

func (p *Post) AfterCreate(db *gorm.DB) (err error) {
	fmt.Println("创建完文章后执行", p)
	userId := p.UserID
	var postCount int64
	db.Model(&Post{}).Where("user_id=?", userId).Count(&postCount)
	db.Model(&User{}).Where("id=?", userId).UpdateColumn("post_count", postCount)
	fmt.Println("更新用户文章数", postCount)
	return nil
}

func (c *Comment) AfterDelete(db *gorm.DB) (err error) {
	fmt.Println("删除评论后执行", c)
	postId := c.PostID
	var commentCount int64
	db.Model(&Comment{}).Where("post_id=?", postId).Count(&commentCount)
	if commentCount == 0 {
		noComment := Comment{
			Content: "无评论",
			PostID:  postId,
		}
		db.Create(&noComment)
		fmt.Println("插入无评论")
	}
	return nil
}

func delComment() {
	db := DbConnet()
	c := Comment{}
	db.First(&c, 2)
	db.Delete(&c)
}

func Task3() {
	// crudOp()
	// txOp()
	sqlExtQuery()
	sqlExtQuery2()
	// creaTables()
	// queryBlogData()
	delComment()
}

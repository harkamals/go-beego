package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	_ "portal/routers"
	"time"
)

type User struct {
	Id      string
	Name    string
	Address *Address
}

func (u *User) TableName() string {
	return "custom_user"
}

type Address struct {
	Id      int
	Country string
	User    *User
}

type Product struct {
	Id           int       `orm:"pk;auto"`
	Name         string    `orm:"index"`
	Description  string    `orm:"column(product_desc)"`
	SerialNumber int       `orm:"size(15)"`
	Value        float32   `orm:"digits(10);decimals(2)"`
	Inventory    int       `orm:"-"`
	LastOrdered  time.Time `orm:"auto_now_add;type(date)"`
}

func (p *Product) TableName() string {
	return "Product"
}

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default",
		"postgres",
		"user=postgres password=postgres host=127.0.0.1 dbname=beego port=5432 sslmode=disable")

	orm.RunSyncdb("default", false, true)

	//orm.RegisterModel(new(User), new(Product))
	fmt.Println("INIT completed")

}

func main() {

	if beego.BConfig.RunMode == "dev" {
		fmt.Print(beego.BConfig.RunMode)
	}

	beego.SetStaticPath("/static", "static")
	beego.BConfig.WebConfig.DirectoryIndex = true

	beego.Run()
}

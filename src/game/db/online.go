package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"time"
)

/*
type User struct {
	uid      int
	name     string
	password string
}
*/

var _online = map[int]*Users{}

var engine *xorm.Engine
var max int

type Users struct {
	Uid      int    `xorm:"notnull pk UUID"`
	UserName string `xorm:"notnull unique VARCHAR(30)"`
	Password string `xorm:"notnull VARCHAR(44)"`

	CreatedAt time.Time `xorm:"notnull TIMESTAMPZ created"`
	UpdatedAt time.Time `xorm:"notnull TIMESTAMPZ updated"`
}

func init() {
	engine1, _ := xorm.NewEngine("mysql", "dev:dev@tcp(127.0.0.1:3306)/game1")
	engine = engine1
	engine.ShowSQL(true)

	one := GetUser(1)
	log.Print(one.UserName)

	/*
		engine.Sync2(new(Users))
		engine.Insert(&Users{Uid: 1, UserName: "player1", Password: "1"})
		engine.Insert(&Users{Uid: 2, UserName: "player2", Password: "1"})
		engine.Insert(&Users{Uid: 3, UserName: "player3", Password: "1"})
	*/
}

func insert(name string, password string) {
	max += 1
	engine.Insert(&Users{Uid: max, UserName: name, Password: password})
}

func GetUser(uid int) *Users {
	user := &Users{}
	_, err := engine.Where("uid = ?", uid).Get(user)
	if err != nil {
		log.Print(err)
	}
	return user
}

func TestUser() {
	player := GetUser(1)
	log.Println(player)
}

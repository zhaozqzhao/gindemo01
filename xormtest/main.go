package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

type PersonTable struct {
	Id         int64     `xorm:"pk autoincr"`
	PersonName string    `xorm:"varchar(24)"`
	PersonAge  int       `xorm:"int default 0"`
	PersonSex  int       `xorm:"not null"`
	City       CityTable `xorm:"-"`
}

type CityTable struct {
	CityName      string
	CityLongitude float32
	CityLatitude  float32
}

func main() {
	engine, err := xorm.NewEngine("mysql", "root:root@/gotest?charset=utf8")
	if err != nil {
		panic(err.Error())
	}

	engine.SetMapper(core.SnakeMapper{})
	engine.Sync2(new(PersonTable))

	personExist, err := engine.IsTableExist(new(PersonTable))
	if err != nil {
		panic(err.Error())
	}
	if personExist {
		fmt.Println("人员表存在")
	} else {
		fmt.Println("人员表不存在")
	}

	personEmpty, err := engine.IsTableEmpty(new(PersonTable))
	if err != nil {
		panic(err.Error())
	}
	if personEmpty {
		fmt.Println("人员表数据为空")
	} else {
		fmt.Println("人员表数据bu为空")
	}

	sql := "truncate table person_table"
	engine.Exec(sql)

	per := PersonTable{
		PersonName: "qq",
		PersonAge:  18,
		PersonSex:  1,
	}
	engine.Insert(&per)

	per1 := []PersonTable{
		{PersonName: "xiaoming", PersonAge: 20, PersonSex: 1},
		{PersonName: "xiaoli", PersonAge: 28, PersonSex: 0},
		{PersonName: "wuhua", PersonAge: 27, PersonSex: 1},
		{PersonName: "xizi", PersonAge: 20, PersonSex: 1},
		{PersonName: "cs", PersonAge: 27, PersonSex: 1},
	}
	engine.Insert(&per1)

	var person PersonTable
	engine.ID(1).Get(&person)
	fmt.Println(person.PersonName)
	fmt.Println()

	var person1 PersonTable
	engine.Where("person_age=? and person_sex=?", 27, 1).Get(&person1)
	fmt.Println(person1)
	fmt.Println()

	var persons []PersonTable
	err = engine.Where("person_age=?", 27).And("person_sex=?", 1).Find(&persons)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(persons)
	fmt.Println()

	var persons1 []PersonTable
	err = engine.SQL("select * from person_table where person_name like 't%'").Find(&persons1)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(person1)

	//6、排序条件查询
	var personsOrderBy []PersonTable
	engine.OrderBy(" person_age desc ").Find(&personsOrderBy)
	fmt.Println(personsOrderBy)
	fmt.Println()

	//7、查询特定字段
	var personsCols []PersonTable
	engine.Cols("person_name", "person_age").Find(&personsCols)
	for _, col := range personsCols {
		fmt.Println(col)
	}

	personsArray := []PersonTable{
		PersonTable{
			PersonName: "jack",
			PersonAge:  28,
			PersonSex:  1,
		},
		PersonTable{
			PersonName: "mali",
			PersonAge:  28,
			PersonSex:  1,
		},
		PersonTable{
			PersonName: "ruby",
			PersonAge:  28,
			PersonSex:  1,
		},
	}

	session := engine.NewSession()
	session.Begin()
	for i := 0; i < len(personsArray); i++ {
		_, err := session.Insert(personsArray[i])
		if err != nil {
			session.Rollback()
			session.Close()
		}
	}
	err = session.Commit()
	session.Close()
	if err != nil {
		panic(err.Error())
	}

	personInsert := PersonTable{
		PersonName: "hello",
		PersonAge:  19,
		PersonSex:  1,
	}
	rowNum, err := engine.ID(7).Update(&personInsert)
	fmt.Println(rowNum)
	if err != nil {
		panic(err.Error())
	}

}

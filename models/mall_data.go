package models

import (
	"fmt"
)

type Userinfo struct {
	Uid        int
	Username   string
	Departname string
	Created    string
}

type MallData struct {
	Id         int
	SendNum    string
	DealerNum  string
	DealerName string
	Code       string
	Batch      string
	ItemCode   string
}

func TestFun() []Userinfo {
	db = GetSqlDb()
	defer db.Close()
	rows, err := db.Query("select uid, username from userinfo")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()
	var alluser []Userinfo
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		var tempInfo Userinfo
		tempInfo.Uid = id
		tempInfo.Username = name
		alluser = append(alluser, tempInfo)
	}
	return alluser
}

func ShowGetProduct() []Userinfo {

	db = GetSqlDb()
	defer db.Close()
	rows, err := db.Query("usp_GetDemo")
	fmt.Println(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()
	var alluser []Userinfo
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		var tempInfo Userinfo
		tempInfo.Uid = id
		tempInfo.Username = name
		alluser = append(alluser, tempInfo)
	}
	return alluser

}

func GetPageList() []MallData {
	db = GetSqlDb()
	defer db.Close()

	ins := make(map[string]string)
	ins["@CurrentPage"] = "1"
	ins["@PageSize"] = "10"

	sqlStr := GetPageProcSql("usp_GetMallDataPageList", ins)
	fmt.Printf(sqlStr)
	//rows, err := db.Query("usp_GetMallDataPageList", sql.Named("CurrentPage", 1), sql.Named("PageSize", 10), sql.Named("TotalCount", sql.Out{Dest: &totalCount}))
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//defer rows.Close()
	var mallData []MallData
	for rows.Next() {
		var Id int
		var SendNum string
		rows.Scan(&Id, &SendNum)
		var tempInfo MallData
		tempInfo.Id = Id
		tempInfo.SendNum = SendNum
		mallData = append(mallData, tempInfo)
	}
	if rows.NextResultSet() {
		for rows.Next() {
			var returnValue int
			rows.Scan(&returnValue)
			fmt.Printf("测试结果 %v", returnValue)
		}
	}
	return mallData
}

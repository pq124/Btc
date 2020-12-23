package models

import (
	"BtWeb/db_mysql"
)

type User struct {
	Id int `form:"id"`
	Telephone string `form:"telephone"`
	Password string `form:"password"`
}

func (u User)SavaData()(int64,error)  {
	result,err:=db_mysql.Db.Exec("insert into user (telephone,password)"+"values(?,?)",u.Telephone,u.Password)
	if err!=nil {
		return -1 ,err
	}
	row,err:=result.RowsAffected()
	if err!=nil {
		return -1,err
	}
	return row,nil
}

func (u User)QueryData()(*User,error)  {
	row:=db_mysql.Db.QueryRow("select  telephone from user where telephone=? and password=? ",u.Telephone,u.Password)
	err:= row.Scan(&u.Telephone)
	if err!=nil {
		return nil,err
	}
	return &u,nil
}
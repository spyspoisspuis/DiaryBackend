package db

import (
	"web-server/internal/util"

	_ "github.com/go-sql-driver/mysql"
)

func AddDiary(d util.DiaryStruct) error {
	db := GetDatabase()
	_, err := db.Exec(`INSERT INTO diary (writer,week,header,context,footer,status) VALUES(?,?,?,?,?,?);`, 
	d.Writer, d.Week,d.Header,d.Context,d.Footer,d.Status)
	return err
}

func GetDiary(writer string, week string) (util.DiaryStruct, error) {
	d := util.DiaryStruct{}
	d.Writer = writer
	d.Week = week
	db := GetDatabase()

	err := db.QueryRow(`SELECT header,context,footer,status FROM diary WHERE 
	(writer = ? AND week = ?)`, writer, week).Scan(&d.Header,&d.Context,&d.Footer,&d.Status)
	
	return d,err
}

func DeleteDiary(d util.DiaryStruct) error {
	db := GetDatabase()
	_,err := db.Exec(`DELETE FROM diary WHERE (writer = ? AND week = ?)`,d.Writer,d.Week)
	return err
}
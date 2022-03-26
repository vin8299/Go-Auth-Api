package database

import ("database/sql"
	_ "github.com/lib/pq"
	)
var err error
var pgObj *sql.DB

func GetPG()(*sql.DB,error){
	connString := "host=localhost  user=postgres dbname=postgres sslmode=disable password=Vinit@123"

	pgObj,err = sql.Open("postgres", connString)

	if err==nil{
		err = pgObj.Ping();
	}

	return pgObj , err
}




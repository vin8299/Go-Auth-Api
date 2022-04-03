package database

import(
	"fmt"
	"testing"
)

func TestPass(t *testing.T){
	conn,err:=GetPG()
	if err!=nil{
		fmt.Println(err.Error())
	}else{
		fmt.Println(conn)
	}

}
package models

import (
	"main/database"
	"strconv"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(data map[string]interface{}) map[string]interface{}{

	var response map[string]interface{} = map[string]interface{}{
		"Code":"",
		"Status":"",
		"Message":"",
		"User Id":"",
	}


	var usr_id int
	pgConn,err:=database.GetPG()
    if err!=nil{
		response["Message"]="Unable to connect to DB "+err.Error()
        response["Status"]="failure"
		response["Code"]="503"
		
    }else{
		usr_name:=data["name"]
		usr_email:=data["email"]
		usr_age:=data["age"]
		
		usr_password,_:=bcrypt.GenerateFromPassword([]byte(data["password"].(string)),14)

		sqlStatement := `INSERT INTO usr_table (usr_name,usr_email,usr_age,usr_password) VALUES ($1,$2,$3,$4) RETURNING usr_id`
		err1 := pgConn.QueryRow(sqlStatement,usr_name,usr_email,usr_age,usr_password).Scan(&usr_id)
		if err1!=nil{
		response["Message"]=err1.Error()
        response["Status"]="failure"
		response["Code"]="502"
		}else{
			response["Message"]="User Successfully Created"
			response["Status"]="Success"
			response["Code"]="200"
			response["User Id"]=strconv.Itoa(usr_id)
		}
	}

	return response
}

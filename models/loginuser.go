package models

import (
	"main/database"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(data map[string]interface{}) map[string]interface{}{

	var response map[string]interface{} = map[string]interface{}{
		"Code":"",
		"Status":"",
		"Message":"",
	}
	var usr_id int
	var usr_password string
	pgConn,err:=database.GetPG()
    if err!=nil{
		response["Message"]=err.Error()
        response["Status"]="failure"
		response["Code"]="503"
		
    }else{
		
		usr_email:=data["email"]

		sqlStatement := `select coalesce((select usr_id from usr_table where usr_email=$1), 0),coalesce((select usr_password from usr_table where usr_email=$1), '')`
		err1 := pgConn.QueryRow(sqlStatement,usr_email).Scan(&usr_id,&usr_password)
		if err1!=nil{
		response["Message"]=err1.Error()
        response["Status"]="failure"
		response["Code"]="502"
		}else{
			if usr_id>0{

				if err:=bcrypt.CompareHashAndPassword([]byte(usr_password),[]byte(data["password"].(string))); err!=nil{
					response["Message"]="Incorrect Password"
					response["Status"]="Success"
					response["Code"]="401"
				}else{
					response["Message"]="Successfully loggedin"
					response["Status"]="Success"
					response["Code"]="200"
				}
			}else{
				response["Message"]="user not found"
				response["Status"]="Success"
				response["Code"]="404"
			}
		}
	}

	return response
}

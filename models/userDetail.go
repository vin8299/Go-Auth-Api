package models


import (
	"main/database"
	"strconv"
)

func UserDetails(data map[string]interface{}) (map[string]interface{}){
	var response map[string]interface{} = map[string]interface{}{
		"Code":"",
		"Status":"",
		"Message":"",
	}

	
	pgConn,err:=database.GetPG()
    if err!=nil{
		response["Message"]=err.Error()
        response["Status"]="failure"
		response["Code"]="503"
		
    }else{
		var usr_id = data["id"].(string)
		usr_id1,_:=strconv.Atoi(usr_id)
		var usr_age int
		var usr_name,usr_email string
		sqlStatement := `select usr_name,usr_email,usr_age from usr_table where usr_id=$1`
		err1 := pgConn.QueryRow(sqlStatement,usr_id1).Scan(&usr_name,&usr_email,&usr_age)
		if err1!=nil{
			response["Message"]=err1.Error()
			response["Status"]="failure"
			response["Code"]="502"
		}else{
			response["Message"]="Fetched data"
			response["Status"]="success"
			response["Code"]="200"
			response["Data"] = "Id="+usr_id+",Name="+usr_name+",Age="+strconv.Itoa(usr_age)+",Email="+usr_email
		}
	}
	return response
}
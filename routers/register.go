package routers

import (
	"context"
	"encoding/json"
	"fmt"

	"SocialNetworkBackend/db"
	"SocialNetworkBackend/models"
)

func Register(ctx context.Context) (models.RespAPI) {
	var t models.User
	var res models.RespAPI
	res.Status = 400

	fmt.Println("Hello from the handler")

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		res.Message=err.Error()
		fmt.Println(res.Message)
		return res
	}

	if len(t.Email) == 0 {
		res.Message="Email is required"
		fmt.Println(res.Message)
		return res
	}

	if len(t.Password) < 8 {
		res.Message="Password must be at least 8 characters"
		fmt.Println(res.Message)
		return res
	}

	_, found, _ := db.UserExists(t.Email)
	if found {
		res.Message="User already exists"
		fmt.Println(res.Message)
		return res
	}

	_, status, err := db.InsertUser(t)
	if err != nil {
		res.Message="Error inserting user: " + err.Error()
		fmt.Println(res.Message)
		return res
	}

	if !status {
		res.Message="Error inserting user"
		fmt.Println(res.Message)
		return res
		
	}

	res.Status = 200
	res.Message = "User inserted"
	fmt.Println(res.Message)

	return res

}
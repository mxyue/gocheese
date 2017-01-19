package main

import (
	"encoding/json"
	"fmt"
	"gocheese/apis"
	"gocheese/db"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestCreateUser(t *testing.T) {
	params := url.Values{"email": {"test@126.com"}, "password": {"123456"}}
	res, err := http.PostForm(fmt.Sprintf("%s/users", server.URL), params)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var data apis.ResponseBody
	err = json.Unmarshal(body, &data)
	fmt.Println("body:", data)
	var users []db.User
	err = db.UserColl().Find(nil).All(&users)
	if res.StatusCode == 200 && err == nil && len(users) == 1 {
		t.Log("通过")
	} else {
		t.Log(res.StatusCode)
		t.Error(err)
	}
}

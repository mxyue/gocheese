package main

import (
	"encoding/json"
	"fmt"
	"gocheese/apis"
	"gocheese/util"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestCreateSession(t *testing.T) {
	params := url.Values{"email": {firstUser.Email}, "password": {"123456"}}
	res, err := http.PostForm(fmt.Sprintf("%s/sessions", server.URL), params)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var data apis.ResponseBody
	err = json.Unmarshal(body, &data)
	fmt.Println("body:", data)
	if res.StatusCode == 200 && err == nil {
		tokenString := data.Content["token"]
		claims, err := util.Decrypt(fmt.Sprintf("%s", tokenString))
		if err != nil {
			t.Error("decrypt失败:", err)
		} else if claims["id"] == firstUser.Id.Hex() {
			t.Log("pass")
		} else {
			t.Error("user id不正确")
		}
	} else {
		t.Log(res.StatusCode)
		t.Error(err)
	}
}

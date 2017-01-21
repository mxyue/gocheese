package main

import (
    "encoding/json"
    "fmt"
    "gocheese/apis"
    "io/ioutil"
    "net/http"
    "net/url"
    "testing"
)

func TestCreateSession(t *testing.T) {
    params := url.Values{"email": {firstUser.Email}, "password": {string(firstUser.Password)}}
    res, err := http.PostForm(fmt.Sprintf("%s/sessions", server.URL), params)
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    var data apis.ResponseBody
    err = json.Unmarshal(body, &data)
    fmt.Println("body:", data)
    if res.StatusCode == 200 && err == nil {
        t.Log("通过")
    } else {
        t.Log(res.StatusCode)
        t.Error(err)
    }
}

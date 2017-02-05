package middleware

import (
    "net/http"
)

func Validate(protectedPage http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

        // http.NotFound(res, req)
        protectedPage(res, req)
    })
}

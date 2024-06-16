package token

import "github.com/IanZC0der/go-myblog/exception"

var AuthFailed = exception.NewAuthFailed("Incorrect Username/Password")

var TokenExpired = exception.NewTokenExpired("Token Expired")

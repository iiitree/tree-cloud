package define

import "github.com/golang-jwt/jwt/v4"

type UserClain struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "tree-cloud-key"
var EmailPassword = "xxx"

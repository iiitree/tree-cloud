package helper

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"net/smtp"
	"tree-cloud/core/define"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GenerateToken(id int, identity string, name string) (string, error) {
	uc := define.UserClain{
		Id:       id,
		Identity: identity,
		Name:     name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func MailSendCode(mail, code string) error {
	e := email.NewEmail()
	e.From = "no-reply <no-reply@xxx.com>"
	e.To = []string{mail}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("您的验证码为：<h1>" + code + "</h1>")
	err := e.SendWithTLS("smtphz.qiye.163.com:465", smtp.PlainAuth("", "no-reply@xxx.com", define.EmailPassword, "smtphz.qiye.163.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtphz.qiye.163.com"})
	if err != nil {
		return err
	}
	return nil
}

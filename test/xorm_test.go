package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
	"tree-cloud/core/models"

	_ "github.com/go-sql-driver/mysql"

	"xorm.io/xorm"
)

func TestXormTest(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/db_cloud?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		t.Fatal(err)
	}
	data := make([]*models.UserBasic, 0)
	engine.Find(&data)
	if err != nil {
		t.Fatal(err)
	}
	b, err := json.Marshal(data)
	if err != nil {
		return
	}
	dst := new(bytes.Buffer)
	err = json.Indent(dst, b, "", " ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dst.String())
}

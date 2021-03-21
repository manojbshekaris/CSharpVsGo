package home

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/valyala/fasthttp"
	_ "github.com/ziutek/mymysql/godrv"
)

func Text(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json; charset=utf-8")
	fmt.Fprint(ctx, `"hello world"`)
}

func Email(ctx *fasthttp.RequestCtx) {
	iD, _ := strconv.Atoi(string(ctx.QueryArgs().Peek("id")))
	ctx.SetContentType("application/json; charset=utf-8")
	var email string

	dataBaseObj, err := sql.Open("mysql", "user:password@tcp(192.168.29.113:3306)/databasename")

	if err != nil {
		fmt.Println(err)
	}
	err = dataBaseObj.QueryRow(`SELECT Email from User where Userid =?`, iD).Scan(&email)
	if err != nil {
		fmt.Fprint(ctx, err.Error())
		ctx.SetStatusCode(500)
		fmt.Println(err)
	}
	defer dataBaseObj.Close()
	fmt.Fprint(ctx, `"`, email, `"`)
}

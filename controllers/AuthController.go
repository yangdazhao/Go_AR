package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

//SESSION_USER_KEY
const SessionUserKey string = "bgsessionID"

type AuthController struct {
	beego.Controller
	isLogin bool
}

func (c *AuthController) Prepare() {
	userLogin := c.GetSession(SessionUserKey)
	if userLogin == nil{
		c.isLogin = false
		if "/home/login" != c.Ctx.Request.URL.Path {
			c.Ctx.Redirect(302, "/home/login")
		}
	} else  {
		c.isLogin = true
	}
	c.Data["isLogin"] = c.isLogin
}


func (c *AuthController) Login() {
	name := c.Ctx.GetCookie("name")
	//password := c.Ctx.GetCookie("password")
	if name == "12364" {

		c.Ctx.Redirect(302, "/statisticalEx")
		//c.Ctx.WriteString("Username:" + name + "password:" + password)
	} else {
		//
		c.Data["Email"] = "yangdazhao@live.com"
		//c.Data["Param"] = ope
		c.TplName = "login.tpl"
	}
}

type User struct{
	Username string
	Password string
}

func (c *AuthController) PostData() {
	userLogin := c.GetSession("bgsessionID")
	fmt.Print(userLogin)

	u := User{}
	if err := c.ParseForm(&u); err != nil {

	}
	//c.Ctx.SetCookie("name", u.Username, 100, "/")  // 设置cookie
	//c.Ctx.SetCookie("name", u.Username, 100, "/")  // 设置cookie
	//c.Ctx.SetCookie("password", u.Password, 100, "/")  // 设置cookie
	c.SetSession(SessionUserKey, u)
	c.Ctx.Redirect(302, "/index")
}

//退出
type LogoutController struct {
	beego.Controller
}

//登录退出功能
func (c *LogoutController) Logout() {
	v := c.GetSession("bgsessionID")
	isLogin:=false
	if v != nil {
		//删除指定的session
		c.DelSession("bgsessionID")
		//销毁全部的session
		c.DestroySession()
		isLogin=true

		fmt.Println("当前的session:")
		fmt.Println(c.CruSession)
	}
	c.Data["json"]=map[string]interface{}{"islogin":isLogin}
	c.ServeJSON()
}

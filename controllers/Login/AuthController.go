package Login

import (
	"fmt"
	. "github.com/astaxie/beego"
	"log"
)

// SESSION_USER_KEY
const SessionUserKey string = "bgsessionID"

type AuthController struct {
	Controller
	isLogin bool
}

type User struct {
	Username string
	Password string
}

//退出
type LogoutController struct {
	Controller
}

func (c *AuthController) Prepare() {
	userLogin := c.GetSession(SessionUserKey)
	if userLogin == nil {
		c.isLogin = false
		if "/home/login" != c.Ctx.Request.URL.Path {
			c.Ctx.Redirect(302, "/home/login")
		}
	} else {
		c.isLogin = true
	}
	c.Data["isLogin"] = c.isLogin
}

func (c *AuthController) Login() {
	name := c.Ctx.GetCookie("name")
	if name == "12364" {
		c.Ctx.Redirect(302, "/statisticalEx")
	} else {
		c.Data["Email"] = "yangdazhao@live.com"
		c.TplName = "login.tpl"
	}
}

func (c *AuthController) PostData() {
	userLogin := c.GetSession("bgsessionID")
	fmt.Print(userLogin)

	u := User{}
	if err := c.ParseForm(&u); err != nil {

	}
	c.SetSession(SessionUserKey, u)
	c.Ctx.Redirect(302, "/index")
}

//登录退出功能
func (c *LogoutController) Logout() {
	v := c.GetSession("bgsessionID")
	isLogin := false
	if v != nil {
		//删除指定的session
		c.DelSession("bgsessionID")
		//销毁全部的session
		c.DestroySession()
		isLogin = true
		log.Fatalf("当前的session: %v", c.CruSession)
	}
	c.Data["json"] = map[string]interface{}{"islogin": isLogin}
	c.ServeJSON()
}

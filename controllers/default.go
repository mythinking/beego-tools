/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-20 09:58:09
 * @LastEditTime: 2019-08-21 16:02:58
 * @LastEditors: Please set LastEditors
 */
package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = beego.AppConfig.String("Website")
	c.Data["Email"] = beego.AppConfig.String("Email")
	c.TplName = "index.tpl"
}

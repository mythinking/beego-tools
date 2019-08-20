package crawlers

import (
	"beego-tools/controllers"
	"fmt"
	"github.com/astaxie/beego/logs"
)

type TmallController struct {
	controllers.CrawlerController
}

func (c *TmallController) Index() {
	c.Data["Website"] = "zhaozhe.ren"
	c.Data["Email"] = "546604336@qq.com"
	c.TplName = "crawlers/eshop.tpl"
}

func (c *TmallController) CreateJob()  {
	storename := c.GetString("storename")
	cookie := c.GetString("cookie")
	title := c.GetString("title")

	logs.Notice("请求参数如下: ")
	logs.Notice("storename: ", storename)
	logs.Notice("cookie: ", cookie)
	logs.Notice("title: ", title)
	logs.Notice("开始执行任务: ")

	res := c.Sucess();
	if storename == "" {
		res = c.Error("storename不能为空哦！")
	}
	if cookie == "" {
		res = c.Error("cookie不能为空哦！")
	}
	if title == "" {
		res = c.Error("title不能为空哦！")
	}
	result, err := c.CmdShell("python", "bats/t.py", storename, cookie, title)
	if err != nil {
		res = c.Error(fmt.Sprintf("%v", err))
	}
	logs.Notice(result)

	c.Data["json"] = res;
	c.ServeJSON();
}
/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-20 21:10:06
 * @LastEditTime: 2019-08-21 17:19:59
 * @LastEditors: Please set LastEditors
 */
package crawlers

import (
	"beego-tools/controllers"
	"fmt"
	"os"
	"io"
	"strings"
	"crypto/md5"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
)

type TmallController struct {
	controllers.CrawlerController
}

func (c *TmallController) Index() {
	c.Data["Website"] = beego.AppConfig.String("Website")
	c.Data["Email"] = beego.AppConfig.String("Email")
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

	res := c.Success();
	if storename == "" {
		res = c.Error("storename不能为空哦！")
	}
	if cookie == "" {
		res = c.Error("cookie不能为空哦！")
	}
	if title == "" {
		res = c.Error("title不能为空哦！")
	}
	result, err := c.CmdShell("python", "bats/tmall.py", storename, cookie, title)
	if err != nil {
		res = c.Error(fmt.Sprintf("%v", err))
	}
	// logs.Notice(res)
	if strings.Index(result, "success") != -1 {
		w := md5.New()
		io.WriteString(w, title)   //将str写入到w中
		md5str := fmt.Sprintf("%x", w.Sum(nil))  //w.Sum(nil)将w的hash转成[]byte格式
		url := beego.AppConfig.String("url")
		res = c.SuccessData(fmt.Sprintf("%s%s_%s.csv", url, storename, md5str));
	} else {
		res = c.Error(result[strings.LastIndex(result, "]")+1:])
	}

	c.Data["json"] = res;
	c.ServeJSON();
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
 func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
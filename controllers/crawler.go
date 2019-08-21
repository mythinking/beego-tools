/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-20 21:44:56
 * @LastEditTime: 2019-08-21 15:04:37
 * @LastEditors: Please set LastEditors
 */
package controllers

import (
	"beego-tools/controllers/common"
	"bytes"
	"errors"
	"fmt"
	"time"
	"github.com/astaxie/beego"
	"os/exec"
)

type CrawlerController struct {
	beego.Controller
	common.Common
}

//func (c *CrawlerController) Get() {
//	c.Data["Website"] = "beego.me"
//	c.Data["Email"] = "astaxie@gmail.com"
//	c.TplName = "index.tpl"
//}

func (c *CrawlerController) Log()  {
	res := c.Success();

	result, err := c.CmdShell("bash", "-c", "tail -n 100 bats/tmall.log." + time.Now().Format("2006-01-02"))
	if err != nil {
		res = c.Error(fmt.Sprintf("%v", err))
	}

	res = c.SuccessData(result)

	c.Data["json"] = res;
	c.ServeJSON();
}


//执行脚本
func (c *CrawlerController) CmdShell(cmdname string, args ...string) (result string, err error) {
	//out, err := exec.Command("python", args...).Output()
	//logs.Error(out)
	//if err != nil {
	//    return
	//}
	cmd := exec.Command(cmdname, args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	errs := cmd.Run()
	if errs != nil {
		err = errors.New(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	result = out.String()
	//result := string(out)
	//if strings.Index(result, "success") != 0 {
	//    err = errors.New(fmt.Sprintf("天猫脚本 error：%s", result))
	//}
	return
}
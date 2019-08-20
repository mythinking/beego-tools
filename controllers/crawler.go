package controllers

import (
	"beego-tools/controllers/common"
	"bytes"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"os/exec"
	"strconv"
	"strings"
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
	line, _ := c.GetInt("line", 0)

	itotal := 0
	result := ""
	res := c.Sucess();

	total, err := c.CmdShell("bash", "-c", "cat logs/project.log |grep [N] |wc -l")
	if err != nil {
		res = c.Error(fmt.Sprintf("%v", err))
	}
	itotal, err = strconv.Atoi(strings.Replace(total, "\n", "", -1))
	if line > 0 {
		total = string(itotal-line)
	}
	fmt.Println(total)
	result, _ =  c.CmdShell("bash", "-c", "cat logs/project.log |grep [N] |tail -n " + total)

	res = c.SucessData(map[string]interface{}{"line":itotal,"content":result})

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
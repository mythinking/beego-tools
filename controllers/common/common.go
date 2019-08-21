/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-20 17:54:45
 * @LastEditTime: 2019-08-21 11:21:45
 * @LastEditors: Please set LastEditors
 */
package common

type Common struct {
}

type JSONStruct struct {
	Code int
	Msg string
	Data interface{}
}

func (c *Common) Error(msg string)  (j JSONStruct) {
	return JSONStruct{1, msg, struct{}{}}
}

func (c *Common) SuccessData(data interface{})  (j JSONStruct) {
	return JSONStruct{0, "操作成功", data}
}

func (c *Common) Success()  (j JSONStruct) {
	return c.SuccessData(struct{}{})
}
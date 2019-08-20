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

func (c *Common) SucessData(data interface{})  (j JSONStruct) {
	return JSONStruct{0, "操作成功", data}
}

func (c *Common) Sucess()  (j JSONStruct) {
	return c.SucessData(struct{}{})
}
package controllers

import (
	"encoding/json"

	"erpweb/logs"
	"erpweb/models"
	"erpweb/util"
)

type CustomerController struct {
	BaseController
}

func (c *CustomerController) QueryCustomer() {
	querystr := c.GetString("querystr")
	rets := models.QueryCustomer(querystr)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = rets
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *CustomerController) GetCustomers() {
	var (
		param = make(map[string]int64)
	)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.FileLogs.Error("param is err", string(c.Ctx.Input.RequestBody))
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	}
	pageNum := param["pageNum"]
	pageSize := param["pageSize"]
	if pageNum > 0 {
		pageNum = pageNum - 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	rets := models.GetCustomerBypage(pageNum, pageSize)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = rets
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *CustomerController) GetCustomerById() {
	var (
		param = make(map[string]int64)
	)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.FileLogs.Error("param is err", string(c.Ctx.Input.RequestBody))
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	}

	id := param["id"]

	logs.FileLogs.Info("%v ---id = ", id)
	ret, _ := models.GetCustomerById(id)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = ret
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *CustomerController) EditCustomerById() {
	var (
		param models.Customer
	)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.FileLogs.Error("%s", err)
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return

	}
	code := models.EditCustomerById(param)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

//{"cardid":"123","name":"小李","sex":0,"compID":10,"deptID":1,"dutyID":1,"health":"very good","height":"170","nativeplace":"安徽省","nation":"汉","maritalstatus":"未婚","education":"本科","university":"大连东软","major":"计算  机","qualification":"666","trialsalary":2000,"salary":6000,"idnumber":"341182","address1":"1111","postcode1":"www111","address2":"2222","postcode2":"www222","contactnumber":"17615002988","phonenumber":"110","email":"www666","emergencycontact":"1323654222","c  ontactnumber1":"1323654222","address3":"33333","trialexpired":"2019-05-28 15:03:03","entrydate":"2019-05-28   15:03:03","birthday":"2019-05-28 15:03:03","contractbegindate":"2019-05-28 15:03:03","contractenddate":"2019-05-28 15:03:03"}
func (c *CustomerController) AddCustomer() {
	var (
		param models.Customer
	)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.FileLogs.Error("%s", err)
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	} else {
		logs.FileLogs.Info("%v", param)
	}
	code := models.AddCustomer(param)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *CustomerController) DeleteCustomer() {
	var (
		param = make(map[string]int64)
	)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.FileLogs.Error("param is err", string(c.Ctx.Input.RequestBody))
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	}

	id := param["id"]

	logs.FileLogs.Info("%v ---", id)
	code := models.DeleteCustomer(id)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

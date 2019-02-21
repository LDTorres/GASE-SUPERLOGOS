package controllers

import (
	"GASE/controllers/services/mails"
	"GASE/controllers/services/payments"
	"GASE/models"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/validation"
)

// LeadsController operations for Leads
type LeadsController struct {
	BaseController
}

// URLMapping ...
func (c *LeadsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("GetAllFromTrash", c.GetAllFromTrash)
	c.Mapping("RestoreFromTrash", c.RestoreFromTrash)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)

	c.Mapping("Newsletter", c.Newsletter)
}

// Test ...
// @Title Post
// @router /test [post]
func (c *LeadsController) Test() {

	token, err := payments.SafetyPayCreateExpressToken("USD", 100.00, 23432, "http://liderlogos.com/success", "http://liderlogos.com/error", "online")

	fmt.Println(err)
	fmt.Println(token)

	c.ServeJSON()
}

// Post ...
// @Title Post
// @Description create Leads
// @Param	body		body 	models.Leads	true		"body for Leads content"
// @Success 201 {int} models.Leads
// @Failure 400 body is empty
// @router / [post]
func (c *LeadsController) Post() {
	var v models.Leads

	//Validate Iso Country
	countryIso := c.Ctx.Input.Header("Country-Iso")

	country, err := models.GetCountriesByIso(countryIso)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v.Country = country

	valid := validation.Validation{}

	b, err := valid.Valid(&v)

	if !b {
		c.BadRequestErrors(valid.Errors, "Leads")
		return
	}

	_, err = models.AddLeads(&v)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	go func() {

		HTMLParams := &mails.HTMLParams{
			Lead:    &v,
			Country: country,
		}

		mailNotification := &mails.Email{
			To:         []string{mails.DefaultEmail},
			Subject:    "Nuevo Lead " + v.Email,
			HTMLParams: HTMLParams,
		}

		err = mails.SendMail(mailNotification, "002")

		if err != nil {
			beego.Debug(err)
		}
	}()

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = v

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Leads by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Leads
// @Failure 403 :id is empty
// @router /:id [get]
func (c *LeadsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetLeadsByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Leads
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Leads
// @Failure 403
// @router / [get]
func (c *LeadsController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllLeads(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = l
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Leads
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Leads	true		"body for Leads content"
// @Success 200 {object} models.Leads
// @Failure 403 :id is not int
// @router /:id [put]
func (c *LeadsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Leads{ID: id}
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	err = models.UpdateLeadsByID(&v)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = MessageResponse{
		Message:       "Updated element",
		PrettyMessage: "Elemento Actualizado",
	}

	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Leads
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *LeadsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	trash := false

	if c.Ctx.Input.Query("trash") != "" {
		trash = true
	}

	err = models.DeleteLeads(id, trash)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = MessageResponse{
		Message:       "Deleted element",
		PrettyMessage: "Elemento Eliminado",
	}

	c.ServeJSON()
}

// GetAllFromTrash ...
// @Title Get All From Trash
// @Description Get All From Trash
// @router /trashed [get]
func (c *LeadsController) GetAllFromTrash() {

	v, err := models.GetLeadsFromTrash()

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()

}

// RestoreFromTrash ...
// @Title Restore From Trash
// @Description Restore From Trash
// @router /:id/restore [put]
func (c *LeadsController) RestoreFromTrash() {

	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := &models.Leads{ID: id}

	err = models.RestoreFromTrash(v.TableName(), v.ID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()

}

// Newsletter ...
// @Title Newsletter
// @Description create Newsletter
// @router /newsletter [post]
func (c *LeadsController) Newsletter() {
	var v models.Leads

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	if v.Email == "" {
		err = errors.New("email is missing")
		c.BadRequest(err)
		return
	}

	//Validate Iso Country
	countryIso := c.Ctx.Input.Header("Country-Iso")

	country, err := models.GetCountriesByIso(countryIso)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	go func() {
		HTMLParams := &mails.HTMLParams{
			Lead:    &v,
			Country: country,
		}

		mailNotification := &mails.Email{
			To:         []string{mails.DefaultEmail, v.Email},
			Subject:    "Subscripción a newsletter - " + v.Email,
			HTMLParams: HTMLParams,
		}

		err = mails.SendMail(mailNotification, "003")

		if err != nil {
			beego.Debug(err)
		}

	}()

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = MessageResponse{
		Message:       "Mail Sent",
		PrettyMessage: "Correo Enviado",
	}

	c.ServeJSON()
}

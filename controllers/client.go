package controllers

import (
	"github.com/astaxie/beego"
)

// ClientController operations for Client
type ClientController struct {
	beego.Controller
}

// URLMapping ...
func (c *ClientController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Client
// @Param	body		body 	models.Client	true		"body for Client content"
// @Success 201 {object} models.Client
// @Failure 403 body is empty
// @router / [post]
func (c *ClientController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Client by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Client
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ClientController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Client
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Client
// @Failure 403
// @router / [get]
func (c *ClientController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Client
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Client	true		"body for Client content"
// @Success 200 {object} models.Client
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ClientController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Client
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ClientController) Delete() {

}

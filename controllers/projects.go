package controllers

import (
	"GASE-SUPERLOGOS/controllers/services/token"
	"GASE-SUPERLOGOS/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego/validation"
)

// ProjectsController operations for Prices
type ProjectsController struct {
	BaseController
}

// URLMapping ...
func (c *ProjectsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GenerateUploadToken", c.GenerateUploadToken)
	c.Mapping("VerifyUploadToken", c.VerifyUploadToken)
}

// Post ...
// @Title Post
// @Description create Projects
// @Param	body		body 	models.Projects	true		"body for Projects content"
// @Success 201 {object} models.Projects
// @Failure 400 body is empty
// @router / [post]
func (c *ProjectsController) Post() {
	var v models.Projects

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	valid := validation.Validation{}

	b, err := valid.Valid(&v)

	if !b {
		c.BadRequestErrors(valid.Errors, v.TableName())
		return
	}

	/* foreignsModels := map[string]int{
		"Currencies": v.Currency.ID,
		"Services":   v.Service.ID,
	}

	resume := c.doForeignModelsValidation(foreignsModels)

	if !resume {
		return
	} */

	_, err = models.AddProjects(&v)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = v

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Projects by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Projects
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ProjectsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetProjectsByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GenerateUploadToken ...
// @Title Generate Upload Token
// @Success 200 {object} models.Projects
// @Failure 403 :id is empty
// @router /:id/token [post]
func (c *ProjectsController) GenerateUploadToken() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetProjectsByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	token, err := token.GenerateTimeToken(v.ID, "project", 7)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v.Token = token

	c.Data["json"] = v
	c.ServeJSON()
}

// VerifyUploadToken ...
// @Title Verify Upload Token
// @Success 200 {object} models.Projects
// @Failure 403 :id is empty
// @router /token/:token [get]
func (c *ProjectsController) VerifyUploadToken() {

	tokenString := c.Ctx.Input.Param(":token")

	_, err := token.ValidateTimeToken(tokenString, "project")

	if err != nil {
		c.BadRequest(err)
		return
	}

	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Projects
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Prices
// @Failure 403
// @router / [get]
func (c *ProjectsController) GetAll() {
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

	l, err := models.GetAllProjects(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = l
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Projects
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Projects	true		"body for Projects content"
// @Success 200 {object} models.Projects
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ProjectsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Projects{ID: id}
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	valid := validation.Validation{}

	b, err := valid.Valid(&v)

	if !b {
		c.BadRequestErrors(valid.Errors, v.TableName())
		return
	}

	/* foreignsModels := map[string]int{
		"Currencies": v.Currency.ID,
		"Services":   v.Service.ID,
	}

	resume := c.doForeignModelsValidation(foreignsModels)

	if !resume {
		return
	} */

	err = models.UpdateProjectsByID(&v)

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
// @Description delete the Projects
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ProjectsController) Delete() {
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

	err = models.DeleteProjects(id, trash)

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

package controllers

import (
	"GASE/controllers/services/files"
	"GASE/controllers/services/mails"
	"GASE/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

// CommentsController operations for Prices
type CommentsController struct {
	BaseController
}

// URLMapping ...
func (c *CommentsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetAttachmentsByUUID", c.GetAttachmentsByUUID)
}

// Post ...
// @Title Post
// @Description create Comments
// @Param	body		body 	models.Comments	true		"body for Comments content"
// @Success 201 {object} models.Comments
// @Failure 400 body is empty
// @router / [post]
func (c *CommentsController) Post() {

	err := c.Ctx.Input.ParseFormOrMulitForm(128 << 20)

	if err != nil {
		c.Ctx.Output.SetStatus(413)
		c.ServeJSON()
		return
	}

	var r = c.Ctx.Request

	var v models.Comments

	err = json.Unmarshal([]byte(r.FormValue("comments")), &v)

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

	foreignsModels := map[string]int{
		"Sketchs": v.Sketch.ID,
	}

	resume := c.doForeignModelsValidation(foreignsModels)

	if !resume {
		return
	}

	_, fileFh, err := c.GetFile("files")
	if err != nil {
		c.BadRequest(err)
		return
	}

	fileUUID, fileMimetype, err := files.CreateFile(fileFh, "project_comments")

	if err != nil {
		c.BadRequest(err)
		return
	}

	v.AttachmentMime = fileMimetype
	v.AttachmentUUID = fileUUID

	_, err = models.AddComments(&v)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	if v.Type == "client" {
		go func() {

			sketchID := v.Sketch.ID

			sketch, err := models.GetSketchsByID(sketchID)
			if err != nil {
				return
			}

			project, err := models.GetProjectsByID(sketch.Project.ID)
			if err != nil {
				return
			}

			HTMLParams := &mails.HTMLParams{
				Project: project,
				Service: sketch.Service,
				Sketch:  sketch,
				Client:  project.Client,
				Comment: &v,
			}

			mailNotification := &mails.Email{
				To:         []string{sketch.Project.NotificationsEmail},
				Subject:    "Comentario en Boceto #" + strconv.Itoa(sketch.Version) + " del Proyecto " + project.Name,
				HTMLParams: HTMLParams,
			}

			err = mails.SendMail(mailNotification, "004")

			if err != nil {
				beego.Debug(err)
			}
		}()
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = v

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Comments by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Comments
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CommentsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetCommentsByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Comments
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Prices
// @Failure 403
// @router / [get]
func (c *CommentsController) GetAll() {
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

	l, err := models.GetAllComments(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = l
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Comments
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Comments	true		"body for Comments content"
// @Success 200 {object} models.Comments
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CommentsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Comments{ID: id}
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

	err = models.UpdateCommentsByID(&v)

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
// @Description delete the Comments
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CommentsController) Delete() {
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

	err = models.DeleteComments(id, trash)

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

// GetAttachmentsByUUID ...
// @Title Get  By UUID
// @Description Get file By UUID
// @router /attachment/:uuid [get]
func (c *CommentsController) GetAttachmentsByUUID() {

	uuid := c.Ctx.Input.Param(":uuid")
	if uuid == "" {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte{})
		return
	}

	imageBytes, mimeType, err := files.GetFile(uuid, "project_comments")
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Ctx.Output.Body([]byte{})
		return
	}

	c.Ctx.Output.Header("Content-Type", mimeType)
	c.Ctx.Output.SetStatus(200)
	c.Ctx.Output.Body(imageBytes)

}

package controllers

import (
	"GASE/controllers/services/files"
	"GASE/controllers/services/token"
	"GASE/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego/validation"
)

// SketchsController operations for Prices
type SketchsController struct {
	BaseController
}

// URLMapping ...
func (c *SketchsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("NewPublicSketch", c.NewPublicSketch)
}

// Post ...
// @Title Post
// @Description create Sketchs
// @Param	body		body 	models.Sketchs	true		"body for Sketchs content"
// @Success 201 {object} models.Sketchs
// @Failure 400 body is empty
// @router / [post]
func (c *SketchsController) Post() {

	err := c.Ctx.Input.ParseFormOrMulitForm(128 << 20)

	if err != nil {
		c.Ctx.Output.SetStatus(413)
		c.ServeJSON()
		return
	}

	var r = c.Ctx.Request

	var v models.Sketchs

	err = json.Unmarshal([]byte(r.FormValue("sketchs")), &v)

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
		"Projects": v.Project.ID,
		"Services": v.Service.ID,
	}

	resume := c.doForeignModelsValidation(foreignsModels)

	if !resume {
		return
	}

	filesFh, err := c.GetFiles("files")

	if err != nil {
		c.BadRequest(err)
		return
	}

	filesData := [][]string{}

	for _, fileFh := range filesFh {

		fileUUID, fileMime, err := files.CreateFile(fileFh, "sketch_files")

		if err != nil {
			c.BadRequest(err)
			return
		}

		fileData := []string{fileUUID, fileMime}

		filesData = append(filesData, fileData)
	}

	_, err = models.AddSketchs(&v)

	if err != nil {

		for _, fileData := range filesData {
			files.DeleteFile(fileData[0], "sketchs_files")
		}

		c.ServeErrorJSON(err)
		return
	}

	for _, fileData := range filesData {

		sketchFile := &models.SketchsFiles{UUID: fileData[0], Mimetype: fileData[1], Sketch: &v}

		_, err := models.AddSketchsFiles(sketchFile)

		if err != nil {
			c.ServeErrorJSON(err)
			return
		}
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = v

	c.ServeJSON()
}

// NewPublicSketch ...
// @Title New Public Sketchs
// @router /token/:token [post]
func (c *SketchsController) NewPublicSketch() {

	tokenString := c.Ctx.Input.Param(":token")

	decodedToken, err := token.ValidateTimeToken(tokenString, "project")

	if err != nil {
		c.BadRequest(err)
		return
	}

	err = c.Ctx.Input.ParseFormOrMulitForm(128 << 20)

	if err != nil {
		c.Ctx.Output.SetStatus(413)
		c.ServeJSON()
		return
	}

	var r = c.Ctx.Request

	var v models.Sketchs

	err = json.Unmarshal([]byte(r.FormValue("sketchs")), &v)

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

	if decodedToken.ID != v.Project.ID {
		err = errors.New("Project IDs dont matchs")
		c.BadRequest(err)
		return
	}

	foreignsModels := map[string]int{
		"Projects": v.Project.ID,
		"Services": v.Service.ID,
	}

	resume := c.doForeignModelsValidation(foreignsModels)

	if !resume {
		return
	}

	filesFh, err := c.GetFiles("sketchs_files")

	if err != nil {
		c.BadRequest(err)
		return
	}

	filesData := [][]string{}

	for _, fileFh := range filesFh {

		fileUUID, fileMime, err := files.CreateFile(fileFh, "sketchs_files")

		if err != nil {
			c.BadRequest(err)
			return
		}

		fileData := []string{fileUUID, fileMime}

		filesData = append(filesData, fileData)
	}

	_, err = models.AddSketchs(&v)

	if err != nil {

		for _, fileData := range filesData {
			files.DeleteFile(fileData[0], "sketchs_files")
		}

		c.ServeErrorJSON(err)
		return
	}

	for _, fileData := range filesData {

		sketchFile := &models.SketchsFiles{UUID: fileData[0], Mimetype: fileData[1], Sketch: &v}

		_, err := models.AddSketchsFiles(sketchFile)

		if err != nil {
			c.ServeErrorJSON(err)
			return
		}
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = v

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Sketchs by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Sketchs
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SketchsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetSketchsByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Success 200 {array} models.Sketchs
// @Failure 403
// @router / [get]
func (c *SketchsController) GetAll() {
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

	l, err := models.GetAllSketchs(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = l
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Success 200 {object} models.Sketchs
// @Failure 400 :id is not int
// @router /:id [put]
func (c *SketchsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Sketchs{ID: id}
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

	err = models.UpdateSketchsByID(&v)

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
// @Description delete the Sketchs
// @Success 200 {string} delete success!
// @Failure 400 id is empty
// @router /:id [delete]
func (c *SketchsController) Delete() {
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

	err = models.DeleteSketchs(id, trash)

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

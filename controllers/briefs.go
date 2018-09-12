package controllers

import (
	"GASE/models"
	"encoding/json"
	"errors"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/validation"
	"github.com/globalsign/mgo/bson"
)

// BriefsController definiton.
type BriefsController struct {
	BaseController
}

// URLMapping ...
func (c *BriefsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Get", c.GetOne)
	c.Mapping("GetOneByCookie", c.GetOneByCookie)
	c.Mapping("GetAll", c.GetAll)
}

// Post ...
// @Title Post
// @Description create Briefs
// @Param	body		body 	models.Briefs	true		"body for Briefs content"
// @Success 201 {int} models.Briefs
// @Failure 400 body is empty
// @router / [post]
func (c *BriefsController) Post() {
	var v models.Briefs

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	// Validate context body

	valid := validation.Validation{}

	b, err := valid.Valid(&v)

	if !b {
		c.BadRequestErrors(valid.Errors, "Briefs")
		return
	}

	err = v.Insert()

	if err != nil {
		c.BadRequest(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Briefs by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Briefs
// @Failure 403 :id is empty
// @router /:id [get]
func (c *BriefsController) GetOne() {
	v := models.Briefs{}

	idStr := c.Ctx.Input.Param(":id")

	if idStr == "" {
		c.BadRequest(errors.New("El campo id no púede ser vacio"))
		return
	}

	validID := bson.IsObjectIdHex(idStr)

	if !validID {
		c.BadRequest(errors.New("El Id no es valido"))
		return
	}

	err := v.GetBriefsByID(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get all Briefs
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Briefs
// @Failure 403 :id is empty
// @router /:id [get]
func (c *BriefsController) GetAll() {
	var v models.Briefs

	Briefs, err := v.GetAllBriefs()

	if err != nil {
		c.BadRequest(err)
		return
	}

	if len(Briefs) == 0 {
		c.ServeErrorJSON(errors.New("No hubo resultados"))
		return
	}

	c.Data["json"] = Briefs
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description Put Briefs
// @Param	body		body 	models.Briefs	true		"body for Briefs content"
// @Success 201 {ObjectId} models.Briefs
// @Failure 400 body is empty
// @router /:id [put]
func (c *BriefsController) Put() {
	var v models.Briefs

	idStr := c.Ctx.Input.Param(":id")

	if idStr == "" {
		c.BadRequest(errors.New("El campo id no púede ser vacio"))
		return
	}

	validID := bson.IsObjectIdHex(idStr)

	if !validID {
		c.BadRequest(errors.New("El Id no es valido"))
		return
	}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	err = v.Update()

	if err != nil {
		c.BadRequest(err)
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
// @Description Delete Briefs
// @Param	body		body 	models.Briefs	true		"body for Briefs content"
// @Success 201 {ObjectId} models.Briefs
// @Failure 400 body is empty
// @router /:id [delete]
func (c *BriefsController) Delete() {
	var v models.Briefs

	idStr := c.Ctx.Input.Param(":id")

	if idStr == "" {
		c.BadRequest(errors.New("El campo id no púede ser vacio"))
		return
	}

	validID := bson.IsObjectIdHex(idStr)

	if !validID {
		c.BadRequest(errors.New("El Id no es valido"))
		return
	}

	err := v.Delete(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	c.Data["json"] = MessageResponse{
		Message:       "Deleted element",
		PrettyMessage: "Elemento Eliminado",
	}
	c.ServeJSON()
}

// GetOneByCookie ...
// @Title Get One
// @Description get Briefs by cookie
// @Param	cookie		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Briefs
// @Failure 403 :cookie, :form, :service  is empty
// @router /:service/:cookie [get]
func (c *BriefsController) GetOneByCookie() {
	v := models.Briefs{}

	cookie := c.Ctx.Input.Param(":cookie")

	if cookie == "" {
		c.BadRequest(errors.New("El campo cookie no púede ser vacio"))
		return
	}

	service := c.Ctx.Input.Param(":service")

	if service == "" {
		c.BadRequest(errors.New("El campo service no púede ser vacio"))
		return
	}

	err := v.GetBriefsByCookie(cookie, service)

	if err != nil {

		if err.Error() == "not found" {
			// Verify if exists a form
			form := models.ServiceForms{}

			err = form.GetServiceFormsByServiceSlug(service)

			if err != nil {
				c.BadRequest(err)
				return
			}

			v.ServiceForm = &form

			v.Cookie = ""

			err = v.Insert()

			beego.Debug(err)

			if err != nil {
				c.BadRequest(err)
				return
			}
		}

	}

	c.Data["json"] = v
	c.ServeJSON()
}

package controllers

import (
	"GASE/models"
	"encoding/json"
	"errors"

	"github.com/astaxie/beego/validation"
	"github.com/globalsign/mgo/bson"
)

// UsersController definiton.
type UsersController struct {
	BaseController
}

// URLMapping ...
func (c *UsersController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Get", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Login", c.Login)
	c.Mapping("ChangePassword", c.ChangePassword)
}

// Post ...
// @Title Post
// @Description create User
// @Param	body		body 	models.User	true		"body for User content"
// @Success 201 {int} models.User
// @Failure 400 body is empty
// @router / [post]
func (c *UsersController) Post() {
	var user models.User

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)

	if err != nil {
		c.BadRequest(err)
		return
	}

	// Validate context body

	valid := validation.Validation{}

	b, err := valid.Valid(&user)

	if !b {
		c.BadRequestErrors(valid.Errors, "User")
		return
	}

	user.ID = bson.NewObjectId()

	err = user.Insert()

	if err != nil {
		c.BadRequest(err)
		return
	}

	c.Data["json"] = user
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Users by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Users
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UsersController) GetOne() {
	user := models.User{}

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

	err := user.GetUsersByID(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	c.Data["json"] = user
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get all Users
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Users
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UsersController) GetAll() {
	var user models.User

	users, err := user.GetAllUsers()

	if err != nil {
		c.BadRequest(err)
		return
	}

	c.Data["json"] = users
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description Put User
// @Param	body		body 	models.User	true		"body for User content"
// @Success 201 {ObjectId} models.User
// @Failure 400 body is empty
// @router /:id [put]
func (c *UsersController) Put() {
	var user models.User

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

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)

	if err != nil {
		c.BadRequest(err)
		return
	}

	err = user.Update(idStr)

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
// @Description Delete User
// @Param	body		body 	models.User	true		"body for User content"
// @Success 201 {ObjectId} models.User
// @Failure 400 body is empty
// @router /:id [delete]
func (c *UsersController) Delete() {
	var user models.User

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

	err := user.Delete(idStr)

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

// Login ...
// @Title Post
// @Description Login for Users
// @Param	body		body 	models.Users	true		"body for Users content"
// @Success 200 {Object} models.Users
// @Failure 400 body is empty
// @router /login [post]
func (c *UsersController) Login() {
	var user models.User

	// Validate empty body

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)

	if err != nil {
		c.BadRequest(err)
		return
	}

	// Validate context body

	valid := validation.Validation{}

	valid.Required(user.Email, "email")
	valid.Required(user.Password, "password")

	if valid.HasErrors() {
		c.BadRequestErrors(valid.Errors, "User")
		return
	}

	err = user.LoginUsers()

	if err != nil {
		c.BadRequest(errors.New("No hubo resultados"))
		return
	}

	user.Token = c.GenerateToken("Admin", user.Token)

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = user

	c.ServeJSON()

}

// ChangePassword ...
// @Title Post
// @Description ChangePassword for Users
// @Param	body		body 	models.Users	true		"body for Users content"
// @Success 200 {Object} models.Users
// @Failure 400 body is empty
// @router /login [post]
func (c *UsersController) ChangePassword() {
	var user models.User

	// Validate empty body

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)

	if err != nil {
		c.BadRequest(err)
		return
	}

	// Validate context body
	valid := validation.Validation{}

	valid.Required(user.Email, "email")
	valid.Required(user.Password, "password")

	if valid.HasErrors() {
		c.BadRequestErrors(valid.Errors, "User")
		return
	}

	err = user.ChangePassword()

	if err != nil {
		c.BadRequest(err)
		return
	}

	user.Token = c.GenerateToken("Admin", user.Token)

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = user

	c.ServeJSON()

}

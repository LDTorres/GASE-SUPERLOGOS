package controllers

import (
	"GASE/models"
	"encoding/json"
	"time"

	"github.com/astaxie/beego/validation"
)

// UsersController definiton.
type UsersController struct {
	BaseController
}

// URLMapping ...
func (c *UsersController) URLMapping() {
	c.Mapping("Post", c.Post)
	/* 	c.Mapping("GetOne", c.GetOne)
	   	c.Mapping("GetAll", c.GetAll)
	   	c.Mapping("Put", c.Put)
	   	c.Mapping("Delete", c.Delete) */
}

// Post ...
// @Title Post
// @Description create User
// @Param	body		body 	models.User	true		"body for Portfolios content"
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

	user.RegDate = time.Now()

	_, err = user.InsertOrUpdate()

	//user.ID = id

	/* user.ID = id */

	if err != nil {
		c.BadRequest(err)
		return
	}

	c.Data["json"] = user
	c.ServeJSON()
}

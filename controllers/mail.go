package controllers

import (
	"GASE/models"
)

// EmailController operations for Email
type EmailController struct {
	BaseController
}

// URLMapping ...
func (c *EmailController) URLMapping() {

	c.Mapping("Post", c.Post)

}

// Post ...
// @Title Post
// @Description create Email
// @Param	body		body 	models.Email	true		"body for Email content"
// @Success 201 {int} models.Email
// @Failure 400 body is empty
// @router / [post]
func (c *EmailController) Post() {

	Email := models.Email{
		Subject: "Manalo  becerro",
		To:      []string{"xarias13@gmail.com", "luisdtc2696@gmail.com"},
		HTML:    "<h1>Prueba</h1>",
	}

	err := models.SendMail(&Email)

	if err != nil {
		c.BadRequest(err)
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = MessageResponse{
		Message:       "Sent Email",
		PrettyMessage: "Correo enviado",
	}

	c.ServeJSON()
}

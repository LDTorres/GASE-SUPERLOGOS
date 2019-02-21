package controllers

import (
	"GASE/controllers/services/files"
	"GASE/controllers/services/mails"
	"GASE/models"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/globalsign/mgo/bson"
)

// BriefsController definiton.
type BriefsController struct {
	BaseController
}

// URLMapping ...
func (c *BriefsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Get", c.GetOne)
	c.Mapping("GetOneByCookie", c.GetOneByCookie)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Options", c.Options)
}

// Post ...
// @router / [post]
func (c *BriefsController) Post() {

	err := c.Ctx.Input.ParseFormOrMulitForm(128 << 20)

	if err != nil {
		c.Ctx.Output.SetStatus(413)
		c.ServeJSON()
		return
	}

	var r = c.Ctx.Request

	client := &models.Clients{}

	err = json.Unmarshal([]byte(r.FormValue("client")), client)

	if err != nil {
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

	client.Country = country

	clientID, err := models.CreateOrUpdateUser(client)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	client.Token, err = c.GenerateToken("Client", strconv.Itoa(clientID))

	if err != nil {
		c.BadRequest(err)
		return
	}

	dataBrief := &map[string]interface{}{}

	err = json.Unmarshal([]byte(r.FormValue("data")), dataBrief)

	if err != nil {
		c.BadRequest(err)
		return
	}

	filesUUIDs := []string{}

	_, fileFh, err := c.GetFile("files")

	if err == nil {

		fileUUID, _, err := files.CreateFile(fileFh, "briefs")

		if err != nil {
			c.BadRequest(err)
			return
		}

		filesUUIDs = append(filesUUIDs, fileUUID)

	}

	v := models.Briefs{
		ID:          bson.NewObjectId(),
		Client:      client,
		Data:        *dataBrief,
		Attachments: filesUUIDs,
		Country:     country,
	}

	//imagesURLs := []string{}

	if c.Ctx.Input.IsUpload() {
		_, err := c.GetFiles("files")
		if err != nil && err != http.ErrMissingFile {
			c.BadRequest(err)
			return
		}
	}

	err = v.Insert()

	if err != nil {
		c.BadRequest(err)
		return
	}

	go func() {

		HTMLParams := &mails.HTMLParams{
			Client:  client,
			BriefID: v.ID.Hex(),
			Country: country,
		}

		mailNotification := &mails.Email{
			To:         []string{mails.DefaultEmail},
			Subject:    "Nuevo Brief - " + v.ID.Hex(),
			HTMLParams: HTMLParams,
		}

		err = mails.SendMail(mailNotification, "005")

		if err != nil {
			beego.Debug(err)
		}

	}()

	c.Data["json"] = v
	c.ServeJSON()
}

// GetOne ...
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
// @router /:id [get]
func (c *BriefsController) GetAll() {
	//var v models.Briefs

	Briefs, err := models.GetAllBriefs()

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

// Delete ...
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
// @router /:service/:cookie [get]
func (c *BriefsController) GetOneByCookie() {
	v := models.Briefs{}

	cookie := c.Ctx.Input.Param(":cookie")

	if cookie == "" {
		c.BadRequest(errors.New("El campo cookie no púede ser vacio"))
		return
	}

	err := v.GetBriefsByCookie(cookie)

	if err != nil {

		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetImageByUUID ...
// @Title Get  By UUID
// @Description Get file By UUID
// @router /image/:uuid [get]
func (c *BriefsController) GetImageByUUID() {

	uuid := c.Ctx.Input.Param(":uuid")
	if uuid == "" {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte{})
		return
	}

	imageBytes, mimeType, err := files.GetFile(uuid, "briefs")
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Ctx.Output.Body([]byte{})
		return
	}

	c.Ctx.Output.Header("Content-Type", mimeType)
	c.Ctx.Output.SetStatus(200)
	c.Ctx.Output.Body(imageBytes)

}

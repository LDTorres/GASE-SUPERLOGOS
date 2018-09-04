package controllers

import (
	"GASE/models"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/astaxie/beego/validation"
)

// PortfoliosController operations for Portfolios
type PortfoliosController struct {
	BaseController
}

// URLMapping ...
func (c *PortfoliosController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Portfolios
// @Param	body		body 	models.Portfolios	true		"body for Portfolios content"
// @Success 201 {int} models.Portfolios
// @Failure 400 body is empty
// @router / [post]
func (c *PortfoliosController) Post() {

	err := c.Ctx.Input.ParseFormOrMulitForm(128 << 20)

	if err != nil {
		c.Ctx.Output.SetStatus(413)
		c.ServeJSON()
		return
	}

	var r = c.Ctx.Request

	stringsInts := &map[string]string{
		"location": r.FormValue("location[id]"),
		"service":  r.FormValue("service[id]"),
		"activity": r.FormValue("activity[id]"),
		"priority": r.FormValue("priority"),
	}

	intValues, err := stringIsValidInt(stringsInts)

	if err != nil {
		c.BadRequest()
		return
	}

	v := models.Portfolios{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Client:      r.FormValue("client"),
		Priority:    int8(intValues["priority"]),
		Service:     &models.Services{ID: intValues["service"]},
		Location:    &models.Locations{ID: intValues["location"]},
		Activity:    &models.Activities{ID: intValues["activity"]},
	}

	valid := validation.Validation{}

	b, err := valid.Valid(&v)

	if !b {
		c.BadRequestErrors(valid.Errors, v.TableName())
		return
	}

	foreignsModels := map[string]int{
		"Locations":  v.Location.ID,
		"Services":   v.Service.ID,
		"Activities": v.Activity.ID,
	}

	resume := c.doForeignModelsValidation(foreignsModels)

	if !resume {
		return
	}

	fmt.Println(filepath.Abs("./"))

	if c.Ctx.Input.IsUpload() {

		images, err := c.GetFiles("images[]")

		if err != nil {
			fmt.Println("Error aquí:")
			fmt.Println(err)
		}

		for _, image := range images {

			//fmt.Println(image.Filename)
			fileType := image.Header["Content-Type"][0]

			if fileType != "image/jpeg" && fileType != "image/png" {

				//TODO: c.BadRequest()
				c.Data["json"] = MessageResponse{
					Message: "Incorrect file type, expected 'image/jpeg' or 'image/png', '" + fileType + "' type was given",
				}
				c.Ctx.Output.SetStatus(409)
				return
			}

			/*
				file, err := image.Open()

				if err != nil {

					//TODO: c.BadRequest()

					c.Data["json"] = MessageResponse{
						Message: "File opening error",
					}
					c.Ctx.Output.SetStatus(409)
					return
				}

				defer file.Close()


				// write the whole body at once
				err = ioutil.WriteFile("output.txt", b, 0644)

			*/

		}

		//r := c.Ctx.Request

		//err = r.ParseMultipartForm(128 << 20)
		/*
			if err != nil {
				c.Ctx.Output.SetStatus(413)
				c.ServeJSON()
				return
			}

			fileHeaders := r.MultipartForm.File["images"]

			for _, fileHeader := range fileHeaders {
				_, err := fileHeader.Open()
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		*/
		/*reader, err := r.MultipartReader()

		if err != nil {
			fmt.Println(err)
			return
		}

		for {
			part, err := reader.NextPart()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(part)
			break
		}
		*/
	}

	_, err = models.AddPortfolios(&v)

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
// @Description get Portfolios by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Portfolios
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PortfoliosController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest()
		return
	}

	v, err := models.GetPortfoliosByID(id)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Portfolios
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Portfolios
// @Failure 403
// @router / [get]
func (c *PortfoliosController) GetAll() {
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

	l, err := models.GetAllPortfolios(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = l
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Portfolios
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Portfolios	true		"body for Portfolios content"
// @Success 200 {object} models.Portfolios
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PortfoliosController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest()
		return
	}

	v := models.Portfolios{ID: id}
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	valid := validation.Validation{}

	b, err := valid.Valid(&v)

	if !b {
		c.BadRequestErrors(valid.Errors, v.TableName())
		return
	}

	foreignsModels := map[string]int{
		"Locations":  v.Location.ID,
		"Services":   v.Service.ID,
		"Activities": v.Activity.ID,
	}

	resume := c.doForeignModelsValidation(foreignsModels)

	if !resume {
		return
	}

	err = models.UpdatePortfoliosByID(&v)

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
// @Description delete the Portfolios
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *PortfoliosController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest()
		return
	}

	err = models.DeletePortfolios(id)

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

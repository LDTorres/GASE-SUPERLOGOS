package controllers

import (
	"GASE/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

// ServicesController operations for Services
type ServicesController struct {
	BaseController
}

// URLMapping ...
func (c *ServicesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetPricesServiceByCountry", c.GetPricesServiceByCountry)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("GetAllFromTrash", c.GetAllFromTrash)
	c.Mapping("RestoreFromTrash", c.RestoreFromTrash)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Services
// @Param	body		body 	models.Services	true		"body for Services content"
// @Success 201 {int} models.Services
// @Failure 400 body is empty
// @router / [post]
func (c *ServicesController) Post() {
	var v models.Services

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	prices := v.Prices
	v.Prices = nil

	valid := validation.Validation{}

	b, _ := valid.Valid(&v)

	if !b {
		c.BadRequestErrors(valid.Errors, "Services")
		return
	}

	_, err = models.AddServices(&v)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	p := make(map[int]*models.Prices)
	var currenciesID []int

	for _, price := range prices {

		currency := price.Currency

		if currency == nil {
			err := errors.New("Currencies is missing")
			c.BadRequest(err)
			return
		}

		_, err := models.GetCurrenciesByID(currency.ID)

		if err != nil {
			c.BadRequestDontExists("Currencies")
			return
		}

		if _, ok := p[currency.ID]; ok {

			err := errors.New("Currencies duplicated")
			c.BadRequest(err)
			return
		}

		price.Service = &v
		p[currency.ID] = price

		price.Currency = nil
		b, _ := valid.Valid(price)

		if !b {
			c.BadRequestErrors(valid.Errors, "Prices")
			return
		}

		price.Currency = currency

		currenciesID = append(currenciesID, currency.ID)

	}

	missingCurrencies, err := models.GetMissingCurrencies(currenciesID...)

	if err != nil {
		beego.Debug("hola becerro")
		c.ServeErrorJSON(err)
		return
	}

	for _, missingCurrency := range missingCurrencies {

		p[missingCurrency.ID] = &models.Prices{
			Value:    0,
			Currency: missingCurrency,
			Service:  &v,
		}

	}

	/* v.Prices = []*models.Prices{} */

	for _, price := range p {
		_, err = models.AddPrices(price)

		if err != nil {
			c.ServeErrorJSON(err)
			return
		}

		/* v.Prices = append(v.Prices, price) */
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = v

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Services by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Services
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ServicesController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetServicesByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Services
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Services
// @Failure 403
// @router / [get]
func (c *ServicesController) GetAll() {
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

	l, err := models.GetAllServices(query, fields, sortby, order, offset, limit)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	CountryIso := c.Ctx.Input.Header("Country-Iso")

	if CountryIso != "" {

		country, err := models.GetCountriesByIso(CountryIso)

		if err != nil {
			c.ServeErrorJSON(err)
			return
		}

		for i, service := range l {

			s := service.(models.Services)
			prices := s.Prices

			for _, price := range prices {
				if country.Currency.ID == price.Currency.ID {

					s.Price = price
					l[i] = s
					break
				}
			}

		}
	}

	c.Data["json"] = l
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Services
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Services	true		"body for Services content"
// @Success 200 {object} models.Services
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ServicesController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Services{ID: id}
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if err != nil {
		c.BadRequest(err)
		return
	}

	err = models.UpdateServicesByID(&v)

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
// @Description delete the Services
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ServicesController) Delete() {
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

	err = models.DeleteServices(id, trash)

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

// GetPricesServiceByCountry ...
// @Title GetPricesServiceByCountry
// @Description get Prices Services by Country Iso
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Services
// @Failure 403 :id is empty
// @router /:id/prices [get]
func (c *ServicesController) GetPricesServiceByCountry() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	CountryIso := c.Ctx.Input.Header("Country-Iso")

	if CountryIso == "" {
		CountryIso = "US"
	}

	_, err = models.GetCountriesByIso(CountryIso)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	v := &models.Services{ID: id}

	err = v.GetPricesServicesByID(CountryIso)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetAllFromTrash ...
// @Title Get All From Trash
// @Description Get All From Trash
// @router /trashed [get]
func (c *ServicesController) GetAllFromTrash() {

	v, err := models.GetServicesFromTrash()

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()

}

// RestoreFromTrash ...
// @Title Restore From Trash
// @Description Restore From Trash
// @router /:id/restore [put]
func (c *ServicesController) RestoreFromTrash() {

	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := &models.Services{ID: id}

	err = models.RestoreFromTrash(v.TableName(), v.ID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()

}

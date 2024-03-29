package controllers

import (
	"GASE-SUPERLOGOS/controllers/services/mails"
	"GASE-SUPERLOGOS/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

// OrdersController operations for Orders
type OrdersController struct {
	BaseController
}

// URLMapping ...
func (c *OrdersController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("GetAllFromTrash", c.GetAllFromTrash)
	c.Mapping("RestoreFromTrash", c.RestoreFromTrash)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetSelf", c.GetSelf)
}

// Post ...
// @Title Post
// @Description create Orders
// @Param	body		body 	models.Orders	true		"body for Orders content"
// @Success 201 {int} models.Orders
// @Failure 400 body is empty
// @router / [post]
func (c *OrdersController) Post() {

	var orderPayment struct {
		Payment map[string]interface{} `json:"payment"`
		Order   *models.Orders         `json:"order"`
	}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &orderPayment)

	if err != nil {
		c.BadRequest(err)
		return
	}

	Order := orderPayment.Order

	//Validate Cookie
	cookie := c.Ctx.Input.Header("Cart-Id")

	if cookie == "" {
		err := errors.New("No se ha recibido el id del Carrito")
		c.BadRequest(err)
		return
	}

	//Validate Iso Country
	countryIso := c.Ctx.Input.Header("Country-Iso")

	if countryIso == "" {
		countryIso = "US"
	}

	country, err := models.GetCountriesByIso(countryIso)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	//Validate Client
	decodedToken, err := VerifyToken(c.Ctx.Input.Header("Authorization"), "Client")

	if err != nil {
		c.BadRequest(err)
		return
	}

	clientID, err := strconv.Atoi(decodedToken.ID)

	if err != nil {
		c.BadRequest(err)
		return
	}

	client, err := models.GetClientsByID(clientID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	Order.Client = client
	Order.Country = country

	// Validate Gateway
	if Order.Gateway == nil || Order.Gateway.ID == 0 {
		err := errors.New("Gateway data is empty")
		c.BadRequest(err)
		return
	}

	gateway, err := models.GetGatewaysByID(Order.Gateway.ID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	Order.Gateway = gateway

	//Get cart
	beego.Debug("Cookie del carrito", cookie)
	cart, err := models.GetOrCreateCartsByCookie(cookie, country.Iso)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	Order.InitialValue = cart.InitialValue
	Order.FinalValue = cart.FinalValue

	// Validate Coupons exists
	var couponsRelationsIDs []int
	for _, el := range Order.Coupons {

		exists := models.ValidateExists("Coupons", el.ID)

		if !exists {
			c.BadRequestDontExists("Coupons")
			return
		}

		couponsRelationsIDs = append(couponsRelationsIDs, el.ID)

		// Getting the discount
		Order.InitialDiscount = (Order.InitialValue * el.Percentage) / 100
		Order.FinalDiscount = (Order.FinalDiscount * el.Percentage) / 100
	}

	Order.Status = "PENDING"

	// Create the new order
	_, err = models.AddOrders(Order)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	price, err := models.GetPricesByID(cart.Services[0].Price.ID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	// Add Prices relations
	var pricesRelations []map[string]int
	for _, service := range cart.Services {
		service.Price.Symbol = price.Currency.Symbol

		relation := map[string]int{"id": service.Price.ID, "quantity": service.Quantity}

		pricesRelations = append(pricesRelations, relation)
	}

	_, err = models.AddRelationsM2MQuantity("orders", Order.ID, "prices", pricesRelations)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	// Add Coupons relations

	_, err = models.RelationsM2M("INSERT", "orders", Order.ID, "coupons", couponsRelationsIDs)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	paymentAmount, err := Order.GetInitialPaymentAmount()

	if err != nil {
		c.BadRequest(err)
		return
	}

	paid, redirect, paymentID, err := paymentsHandler(Order.ID, Order.Gateway, paymentAmount, country, orderPayment.Payment, client)

	if err != nil {
		c.BadRequest(err)
		return
	}

	beego.Debug(paid, redirect)

	if paid {
		Order.Status = "COMPLETED"
		Order.PaymentID = paymentID
		models.UpdateOrdersByID(Order)
	}

	Order.Redirect = redirect

	/* Set used to true */
	for _, el := range Order.Coupons {
		coupon, _ := models.GetCouponsByID(el.ID)

		coupon.Used = true

		err := models.UpdateCouponsByID(coupon)

		if err != nil {
			c.ServeErrorJSON(err)
			return
		}
	}

	go func() {

		HTMLParams := &mails.HTMLParams{
			Client:   client,
			Order:    Order,
			Gateway:  gateway,
			Country:  country,
			Symbol:   price.Currency.Symbol,
			Services: cart.Services,
		}

		mailNotification := &mails.Email{
			To:         []string{client.Email, mails.DefaultEmail},
			Subject:    "Gracias por su pedido - " + strconv.Itoa(Order.ID),
			HTMLParams: HTMLParams,
		}

		err = mails.SendMail(mailNotification, "001")

		if err != nil {
			beego.Debug("Parametros de email", HTMLParams)
			beego.Debug("Error", err)
		}
	}()

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = Order

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Orders by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Orders
// @Failure 403 :id is empty
// @router /:id [get]
func (c *OrdersController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v, err := models.GetOrdersByID(id)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()
}

// GetSelf ...
// @Title Get Self
// @router /self [get]
func (c *OrdersController) GetSelf() {

	token := c.Ctx.Input.Header("Authorization")

	//Validate Client
	decodedToken, err := VerifyToken(token, "Client")

	if err != nil {
		c.BadRequest(err)
		return
	}

	clientID, err := strconv.Atoi(decodedToken.ID)

	if err != nil {
		c.BadRequest(err)
		return
	}

	client, err := models.GetClientsByID(clientID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	orders, err := models.GetOrdersByClientID(client.ID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = orders
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Orders
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Orders
// @Failure 403
// @router / [get]
func (c *OrdersController) GetAll() {
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

	l, err := models.GetAllOrders(query, fields, sortby, order, offset, limit)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = l
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Orders
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Orders	true		"body for Orders content"
// @Success 200 {object} models.Orders
// @Failure 403 :id is not int
// @router /:id [put]
func (c *OrdersController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := models.Orders{ID: id}
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

	foreignsModels := map[string]int{
		"Clients": v.Client.ID,
	}

	resume := c.doForeignModelsValidation(foreignsModels)

	if !resume {
		return
	}

	//TODO: VERITICAR DATOS DE TARJETA DE CREDITO

	err = models.UpdateOrdersByID(&v)

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
// @Description delete the Orders
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *OrdersController) Delete() {
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

	err = models.DeleteOrders(id, trash)

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

// GetAllFromTrash ...
// @Title Get All From Trash
// @Description Get All From Trash
// @router /trashed [get]
func (c *OrdersController) GetAllFromTrash() {

	v, err := models.GetOrdersFromTrash()

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
func (c *OrdersController) RestoreFromTrash() {

	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.BadRequest(err)
		return
	}

	v := &models.Orders{ID: id}

	err = models.RestoreFromTrash(v.TableName(), v.ID)

	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	c.Data["json"] = v
	c.ServeJSON()

}

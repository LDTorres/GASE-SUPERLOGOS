package controllers

// MainController operations for Admin
type MainController struct {
	BaseController
}

// URLMapping ...
func (c *MainController) URLMapping() {
	c.Mapping("Home", c.Home)
}

// Main ...
// @Title Main
// @router /home [get]
func (c *MainController) Home() {
	c.Layout = "layouts/main.tpl"
	c.TplName = "home/index.tpl"
}

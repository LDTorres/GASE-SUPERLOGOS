package controllers

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"

	"GASE/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/go-sql-driver/mysql"
)

// BaseController operations for Activities
type BaseController struct {
	beego.Controller
}

//MessageResponse =
type MessageResponse struct {
	Message       string              `json:"message,omitempty"`
	Code          uint16              `json:"code,omitempty"`
	PrettyMessage string              `json:"pretty_message,omitempty"`
	Errors        []map[string]string `json:"errors,omitempty"`
	Error         string              `json:"error,omitempty"`
}

var rootDir string

func init() {

	rootDir, _ = filepath.Abs("./")

	validation.SetDefaultMessage(map[string]string{
		"Required":     "This field is required",
		"Min":          "The min length requred is %d",
		"Max":          "The max length requred is %d",
		"Range":        "The range of the values is %d до %d",
		"MinSize":      "Longitud mínima permitida %d",
		"MaxSize":      "Minimum length allowed %d",
		"Length":       "The length should be equal to %d",
		"Alpha":        "Must consist of letters",
		"Numeric":      "Must consist of numbers",
		"AlphaNumeric": "Must consist of letters or numbers",
		"Match":        "Must coincide with %s",
		"NoMatch":      "It should not coincide with %s",
		"AlphaDash":    "Must consist of letters, numbers or symbols (-_)",
		"Email":        "Must be in the correct email format",
		"IP":           "Must be a valid IP address",
		"Base64":       "Must be presented in the correct format base64",
		"Mobile":       "Must be the correct mobile number",
		"Tel":          "Must be the correct phone number",
		"Phone":        "Must be the correct phone or mobile number",
		"ZipCode":      "Must be the correct zip code",
	})
}

//BadRequest =
func (c *BaseController) BadRequest(err error) {
	c.Ctx.Output.SetStatus(400)
	c.Data["json"] = MessageResponse{
		Message:       "Bad request body",
		PrettyMessage: "Peticion mal formada",
		Error:         err.Error(),
	}
	c.ServeJSON()
}

//BadRequestDontExists =
func (c *BaseController) BadRequestDontExists(message string) {
	c.Ctx.Output.SetStatus(404)
	c.Data["json"] = MessageResponse{
		Message:       "Dont exist the element " + message,
		PrettyMessage: "No existe el elemento a relacionar " + message,
	}
	c.ServeJSON()
}

//ServeErrorJSON : Serve Json error
func (c *BaseController) ServeErrorJSON(err error) {

	if err == orm.ErrNoRows {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = MessageResponse{
			Message:       "No results",
			PrettyMessage: "No se encontraron resultados",
		}
		c.ServeJSON()

		return
	}

	if driverErr, ok := err.(*mysql.MySQLError); ok {

		switch driverErr.Number {
		case 1062:
			c.Ctx.Output.SetStatus(409)
			c.Data["json"] = MessageResponse{
				Message:       "The element already exists",
				Code:          driverErr.Number,
				PrettyMessage: "El elemento de la base de datos ya existe",
			}
		case 1054:
			c.Ctx.Output.SetStatus(409)
			c.Data["json"] = MessageResponse{
				Message:       "Unknown column in 'Field List'",
				Code:          driverErr.Number,
				PrettyMessage: "Columna desconocida",
			}
		default:
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = MessageResponse{
				Message:       "An error has ocurred",
				Code:          driverErr.Number,
				PrettyMessage: "Un error ha ocurrido",
			}
		}
	}

	c.ServeJSON()
}

//BadRequestErrors = Validar
func (c *BaseController) BadRequestErrors(errors []*validation.Error, entity string) {

	var errorsMessages []map[string]string

	for _, err := range errors {

		var errorMessage = map[string]string{
			"entity":     entity,
			"message":    err.Message,
			"field":      strings.ToLower(err.Field),
			"validation": strings.ToLower(err.Name),
		}

		errorsMessages = append(errorsMessages, errorMessage)
	}

	c.Ctx.Output.SetStatus(400)
	c.Data["json"] = MessageResponse{
		Errors: errorsMessages,
	}

	c.ServeJSON()
}

func (c *BaseController) doForeignModelsValidation(foreignModels map[string]int) (resume bool) {

	for foreignModelName, foreignModelID := range foreignModels {

		exists := models.ValidateExists(foreignModelName, foreignModelID)

		if !exists {
			c.BadRequestDontExists(foreignModelName)
			return
		}

	}

	resume = true

	return

}

func stringIsValidInt(stringIDs *map[string]string) (IDs map[string]int, err error) {

	intIDs := make(map[string]int)

	for key, id := range *stringIDs {

		intID, err := strconv.Atoi(id)

		if err != nil {
			return nil, err
		}

		intIDs[key] = intID

	}

	IDs = intIDs

	return
}

func checkOrCreateImagesFolder(rootDir string, imagesFolderPath string) (err error) {

	if _, err := os.Stat(rootDir); os.IsNotExist(err) {
		//os.MkdirAll(folderPath, os.ModePerm);

	}

	return

}

package middlewares

import (
	"GASE/controllers"
	"encoding/json"

	"github.com/astaxie/beego/context"
)

var (
	// ControllersNames ...
	ControllersNames = []string{
		"portfolios", "activities", "carts", "clients", "countries", "coupons", "currencies", "gateways", "images", "locations", "orders", "prices", "sectors", "services", "briefs", "users",
	}
	// ExcludeUrls ...
	ExcludeUrls = map[string][]string{
		"login":           {"POST"},
		"clients":         {"POST"},
		"carts":           {"POST", "PUT"},
		"orders":          {"POST"},
		"change-password": {"POST"},
		"custom-search":   {"GET"},
	}
)

// DenyAccess ...
func DenyAccess(ctx *context.Context, err error) {

	ctx.Output.SetStatus(401)
	ctx.Output.Header("Content-Type", "application/json")

	message := controllers.MessageResponse{
		Message:       "Permission Deny",
		PrettyMessage: "Permiso Denegado",
		Code:          002,
		Error:         err.Error(),
	}

	res, _ := json.Marshal(message)

	ctx.Output.Body([]byte(string(res)))
	return
}

// MwPattern ...
type MwPattern struct {
	URL       string
	Methods   []string
	UserTypes []string
}

// GetControllerPatterns ...
func GetControllerPatterns(route string) []*MwPattern {

	AccessGetAll := []*MwPattern{
		{
			URL:       "/",
			Methods:   []string{"GET"},
			UserTypes: []string{"Guest"},
		},
		{
			URL:       "/:id",
			Methods:   []string{"All"},
			UserTypes: []string{"Admin"},
		},
	}

	OnlyAdmin := []*MwPattern{
		{
			URL:       "/",
			Methods:   []string{"GET"},
			UserTypes: []string{"Admin"},
		},
		{
			URL:       "/:id",
			Methods:   []string{"All"},
			UserTypes: []string{"Admin"},
		},
	}

	/** Only Get All **/

	countries := AccessGetAll

	coupons := AccessGetAll

	currencies := AccessGetAll

	gateways := AccessGetAll

	locations := AccessGetAll

	sectors := AccessGetAll

	services := AccessGetAll

	portfolios := AccessGetAll

	activities := AccessGetAll

	/** Only Admin **/

	prices := OnlyAdmin

	users := OnlyAdmin

	/** Custom **/
	images := []*MwPattern{
		{
			URL:       "/",
			Methods:   []string{"GET"},
			UserTypes: []string{"Admin"},
		},
		{
			URL:       "/:id",
			Methods:   []string{"GET", "DELETE", "PUT"},
			UserTypes: []string{"Admin"},
		},
		{
			URL:       "/slug/:slug",
			Methods:   []string{"GET"},
			UserTypes: []string{"Guest"},
		},
	}

	clients := []*MwPattern{
		{
			URL:       "/",
			Methods:   []string{"GET"},
			UserTypes: []string{"Admin"},
		},
		{
			URL:       "/:id",
			Methods:   []string{"GET", "DELETE", "PUT"},
			UserTypes: []string{"Admin", "Client"},
		},
		{
			URL:       "/email/:email",
			Methods:   []string{"GET"},
			UserTypes: []string{"Admin", "Client"},
		},
	}

	carts := []*MwPattern{
		{
			URL:       "/",
			Methods:   []string{"GET"},
			UserTypes: []string{"Admin"},
		},
		{
			URL:       "/:id",
			Methods:   []string{"GET", "DELETE"},
			UserTypes: []string{"Admin"},
		},
	}

	orders := []*MwPattern{
		{
			URL:       "/",
			Methods:   []string{"GET"},
			UserTypes: []string{"Admin"},
		},
		{
			URL:       "/:id",
			Methods:   []string{"GET", "PUT", "DELETE"},
			UserTypes: []string{"Client", "Admin"},
		},
	}

	briefs := []*MwPattern{
		{
			URL:       "/:id",
			Methods:   []string{"PUT", "DELETE"},
			UserTypes: []string{"Admin"},
		},
	}

	validations := map[string][]*MwPattern{
		"carts":      carts,
		"clients":    clients,
		"countries":  countries,
		"coupons":    coupons,
		"currencies": currencies,
		"gateways":   gateways,
		"images":     images,
		"locations":  locations,
		"orders":     orders,
		"prices":     prices,
		"sectors":    sectors,
		"services":   services,
		"briefs":     briefs,
		"users":      users,
		"portfolios": portfolios,
		"activities": activities,
	}

	return validations[route]
}

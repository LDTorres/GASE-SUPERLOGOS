package models

import (
	"io"
	"net/http"
	"os"

	"github.com/astaxie/beego/orm"
	"github.com/globalsign/mgo/bson"
	"github.com/gosimple/slug"
)

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

//AddDefaultDataCurrencies on init app
func AddDefaultDataCurrencies() (count int64, err error) {

	o := orm.NewOrm()

	dummyData := []*Currencies{
		{
			Symbol: "€",
			Name:   "Euro",
			Iso:    "EUR",
		},
		{
			Symbol: "$",
			Name:   "Dólar estadounidense",
			Iso:    "USD",
		},
	}

	count, err = o.InsertMulti(100, dummyData)

	return
}

// AddRelationsGatewaysCurrencies ...
func AddRelationsGatewaysCurrencies() (result int, errors []error) {

	o := orm.NewOrm()

	dummyData := map[string][]string{
		"01": {
			"USD",
			"EUR",
		},
		"02": {
			"USD",
			"EUR",
		},
		"03": {
			"USD",
			"EUR",
		},
		"04": {
			"USD",
		},
		"05": {
			"USD",
			"EUR",
		},
	}

	for key, dummyGateway := range dummyData {

		gateway := Gateways{Code: key}

		err := o.Read(&gateway, "code")

		if err != nil {
			continue
		}

		var relationsIDs []int

		for _, iso := range dummyGateway {

			currency := Currencies{Iso: iso}

			err := o.Read(&currency, "iso")

			if err != nil {
				continue
			}

			relationsIDs = append(relationsIDs, currency.ID)
		}

		RelationsM2M("INSERT", "gateways", gateway.ID, "currencies", relationsIDs)
	}

	return
}

func AddDefaultDataCountries() (count int64, err error) {

	o := orm.NewOrm()

	dummyData := []map[string]interface{}{
		{
			"name":     "USA",
			"iso":      "US",
			"phone":    "0",
			"currency": "USD",
			"tax":      21.0,
			"email":    "liderlogo@gmail.com",
			"skype":    "1024234",
		},
		{
			"name":     "España",
			"iso":      "ES",
			"phone":    "0",
			"currency": "USD",
			"tax":      0,
			"email":    "info@liderlogo.com",
			"skype":    "info@liderlogo.com",
		},
		{
			"name":     "Puerto Rico",
			"iso":      "PR",
			"phone":    "0",
			"currency": "USD",
			"tax":      0,
			"email":    "info@liderlogo.com",
			"skype":    "info@liderlogo.com",
		},
		{
			"name":     "Panama",
			"iso":      "PA",
			"phone":    "0",
			"currency": "USD",
			"tax":      0,
			"email":    "info@liderlogo.com",
			"skype":    "info@liderlogo.com",
		},
		{
			"name":     "Ecuador",
			"iso":      "EC",
			"phone":    "0",
			"currency": "USD",
			"tax":      0,
			"email":    "info@liderlogo.com",
			"skype":    "info@liderlogo.com",
		},
		{
			"name":     "Costa Rica",
			"iso":      "CR",
			"phone":    "0",
			"currency": "USD",
			"tax":      0,
			"email":    "info@liderlogo.com",
			"skype":    "info@liderlogo.com",
		},
		{
			"name":     "Argentina",
			"iso":      "AR",
			"phone":    "0",
			"currency": "USD",
			"tax":      0,
			"email":    "info@liderlogo.com",
			"skype":    "info@liderlogo.com",
		},
		{
			"name":     "Bolivia",
			"iso":      "BO",
			"phone":    "0",
			"currency": "USD",
			"tax":      0,
			"email":    "info@liderlogo.com",
			"skype":    "info@liderlogo.com",
		},
		{
			"name":     "Chile",
			"iso":      "CL",
			"phone":    "0",
			"currency": "USD",
			"tax":      0,
			"email":    "info@liderlogo.com",
			"skype":    "info@liderlogo.com",
		},
		{
			"name":     "Colombia",
			"iso":      "CO",
			"phone":    "0",
			"currency": "USD",
			"tax":      0,
			"email":    "info@liderlogo.com",
			"skype":    "info@liderlogo.com",
		},
		{
			"name":     "Dominicana",
			"iso":      "DM",
			"phone":    "0",
			"currency": "USD",
			"tax":      0,
			"email":    "info@liderlogo.com",
			"skype":    "info@liderlogo.com",
		},
		{
			"name":     "Guatemala",
			"iso":      "GT",
			"phone":    "0",
			"currency": "USD",
			"tax":      0,
			"email":    "info@liderlogo.com",
			"skype":    "info@liderlogo.com",
		},
		{
			"name":     "Honduras",
			"iso":      "HN",
			"phone":    "0",
			"currency": "USD",
			"tax":      0,
			"email":    "info@liderlogo.com",
			"skype":    "info@liderlogo.com",
		},
		{
			"name":     "Mexico",
			"iso":      "MX",
			"phone":    "0",
			"currency": "USD",
			"tax":      0,
			"email":    "info@liderlogo.com",
			"skype":    "info@liderlogo.com",
		},
		{
			"name":     "Peru",
			"iso":      "PE",
			"phone":    "0",
			"currency": "USD",
			"tax":      0,
			"email":    "info@liderlogo.com",
			"skype":    "info@liderlogo.com",
		},
		{
			"name":     "Paraguay",
			"iso":      "PY",
			"phone":    "0",
			"currency": "USD",
			"tax":      0,
			"email":    "info@liderlogo.com",
			"skype":    "info@liderlogo.com",
		},
		{
			"name":     "Uruguay",
			"iso":      "UY",
			"phone":    "0",
			"currency": "USD",
			"tax":      0,
			"email":    "info@liderlogo.com",
			"skype":    "info@liderlogo.com",
		},
		{
			"name":     "Nicaragua",
			"iso":      "NI",
			"phone":    "0",
			"currency": "USD",
			"tax":      0,
			"email":    "info@liderlogo.com",
			"skype":    "info@liderlogo.com",
		},
		{
			"name":     "Salvador",
			"iso":      "SV",
			"phone":    "0",
			"currency": "USD",
			"tax":      0,
			"email":    "info@liderlogo.com",
			"skype":    "info@liderlogo.com",
		},
	}

	var dummyCountries []Countries

	for _, dummyCountry := range dummyData {

		currency := Currencies{Iso: dummyCountry["currency"].(string)}

		err := o.Read(&currency, "Iso")

		if err != nil {
			continue
		}

		country := Countries{
			Name:     dummyCountry["name"].(string),
			Iso:      dummyCountry["iso"].(string),
			Phone:    dummyCountry["phone"].(string),
			Currency: &currency,
			Slug:     slug.Make(dummyCountry["name"].(string)),
			Email:    dummyCountry["email"].(string),
			Skype:    dummyCountry["skype"].(string),
		}

		o.ReadOrCreate(&country, "Iso")

		dummyCountries = append(dummyCountries, country)

	}

	count, err = o.InsertMulti(100, dummyCountries)

	return
}

//AddDefaultDataSectors on init app
func AddDefaultDataSectors() (count int64, err error) {

	o := orm.NewOrm()

	dummyData := []*Sectors{
		{
			Name: "Figuras Geometricas y Abstractas",
			Code: "01",
		},
		{
			Name: "Agricultura y ganaderia",
			Code: "02",
		},
		{
			Name: "Animales y Mascotas",
			Code: "03",
		},
		{
			Name: "Arte y fotografia",
			Code: "04",
		},
		{
			Name: "Industria automotriz",
			Code: "05",
		},
		{
			Name: "Accesorios y glamour",
			Code: "06",
		},
		{
			Name: "Transporte y logistica",
			Code: "07",
		},
		{
			Name: "Asesoria y Consultoria",
			Code: "08",
		},
		{
			Name: "Construccion y arquitectura",
			Code: "09",
		},
		{
			Name: "Ropa y Moda",
			Code: "10",
		},
		{
			Name: "Educacion y Formacion",
			Code: "01",
		},
		{
			Name: "Hogar y Jardin",
			Code: "11",
		},
		{
			Name: "Alimentos y Bebidas",
			Code: "12",
		},
		{
			Name: "Belleza y Cuidado personal",
			Code: "13",
		},
		{
			Name: "Salud y Medicina",
			Code: "14",
		},
		{
			Name: "Deportes y ejercicios",
			Code: "15",
		},
		{
			Name: "Hobies y Entretenimiento",
			Code: "16",
		},
		{
			Name: "Organización sin fin de lucro",
			Code: "17",
		},
		{
			Name: "Tecnologia y telecomunicacion",
			Code: "18",
		},
		{
			Name: "Turismo y viajes",
			Code: "19",
		},
		{
			Name: "Reparaciones y Mantenimiento",
			Code: "20",
		},
		{
			Name: "Niños",
			Code: "21",
		},
		{
			Name: "Medio ambiente",
			Code: "22",
		},
		{
			Name: "Social Media",
			Code: "23",
		},
		{
			Name: "Seguridad",
			Code: "24",
		},
		{
			Name: "Personas",
			Code: "25",
		},
		{
			Name: "Espiritualidad",
			Code: "26",
		},
	}

	for _, v := range dummyData {
		v.Slug = slug.Make(v.Name)
	}

	count, err = o.InsertMulti(100, dummyData)

	return
}

// AddDefaultDataActivities ...
func AddDefaultDataActivities() (count int64, errors []error) {

	o := orm.NewOrm()

	dummyData := map[string][]Activities{
		/*"01": "figuras-geometricas-y-abstractas  {
			{
				Name:        "Figuras",
				Description: "",
			},
		},
		"02": {
			{
				Name:        "Agricultura",
				Description: "",
			},
		}, */
	}

	for key, dummySector := range dummyData {

		sector := Sectors{Code: key}

		err := o.Read(&sector, "code")

		if err != nil {
			continue
		}

		for key, dummyActivity := range dummySector {

			dummySector[key].Slug = slug.Make(dummyActivity.Name)
			dummySector[key].Sector = &sector

		}

		result, err := o.InsertMulti(100, dummySector)

		if err != nil {

			errors = append(errors, err)
			continue
		}

		count += result

	}

	return
}

//AddDefaultDataServices on init app
func AddDefaultDataServices() (result int64, err error) {

	o := orm.NewOrm()

	dummyData := []*Services{
		/* {
			Name:       "Logo a Medida",
			Percertage: 10.0,
			Code:       "01",
		}, */
	}

	for _, dummyService := range dummyData {
		dummyService.Slug = slug.Make(dummyService.Name)
	}

	result, err = o.InsertMulti(100, dummyData)

	return
}

//GetPricesServicesByID retrieves Services by Id. Returns error if Id doesn't exist
func (m *Services) GetPricesServicesByID(iso string) (err error) {
	o := orm.NewOrm()

	err = o.Read(m)

	if err != nil {
		return
	}

	//Get countries by Iso
	country, err := GetCountriesByIso(iso)

	if err != nil {
		return
	}

	price := &Prices{Currency: country.Currency, Service: m}

	err = o.Read(price, "Currency", "Service")

	if err != nil {
		return
	}

	err = searchFK(price.TableName(), price.ID).One(price)

	price.Service = nil
	m.Price = price

	return
}

//AddDefaultDataPrices on init app
func AddDefaultDataPrices() (count int64, errors []error) {

	o := orm.NewOrm()

	dummyData := map[string]map[string]Prices{
		"01": {
			"USD": {
				Value: 10.0,
			},
			"EUR": {
				Value: 10.0,
			},
		},
	}

	for code, dummyService := range dummyData {

		service := Services{Code: code}

		err := o.Read(&service, "code")

		if err != nil {
			continue
		}

		for iso, dummyPriceByCurrency := range dummyService {

			currency := Currencies{Iso: iso}

			err := o.Read(&currency, "iso")

			if err != nil {
				continue
			}

			dummyPriceByCurrency.Service = &service
			dummyPriceByCurrency.Currency = &currency

			_, result, err := o.ReadOrCreate(&dummyPriceByCurrency, "Currency", "Service")

			if err != nil {
				continue
			}

			count += result
		}

		errors = append(errors, err)
	}

	return
}

//AddDefaultDataUsers ...
func AddDefaultDataUsers() (id *bson.ObjectId, err error) {
	u := User{ID: bson.NewObjectId(), Name: defaultUsername, Email: defaultMail, Password: defaultPassword}

	err = u.Insert()

	if err != nil {
		return nil, err
	}

	return &u.ID, nil
}

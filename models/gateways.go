package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

//Gateways Model
type Gateways struct {
	ID           int           `orm:"column(id);auto" json:"id"`
	Name         string        `orm:"column(name);size(255)" json:"name,omitempty" valid:"Required"`
	Description  string        `orm:"column(description)" json:"description,omitempty" valid:"Required"`
	Instructions string        `orm:"column(instructions)" json:"instructions,omitempty" valid:"Required"`
	Code         string        `orm:"column(code);size(255)" json:"code,omitempty" valid:"Required; AlphaNumeric"`
	Currencies   []*Currencies `orm:"rel(m2m)" json:"currencies,omitempty"`
	Orders       []*Orders     `orm:"reverse(many)" json:"orders,omitempty"`
	CreatedAt    time.Time     `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt    time.Time     `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt    time.Time     `orm:"column(deleted_at);type(datetime);null"  json:"-"`
}

//TableName =
func (t *Gateways) TableName() string {
	return "gateways"
}

func (t *Gateways) loadRelations() {

	o := orm.NewOrm()

	relations := []string{"Orders", "Currencies"}

	for _, relation := range relations {
		o.LoadRelated(t, relation)
	}

	return

}

// AddGateways insert a new Gateways into database and returns
// last inserted Id on success.
func AddGateways(m *Gateways) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)

	return
}

// GetGatewaysByID retrieves Gateways by Id. Returns error if
// Id doesn't exist
func GetGatewaysByID(id int) (v *Gateways, err error) {
	v = &Gateways{ID: id}
	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	return
}

// GetAllGateways retrieves all Gateways matches certain condition. Returns empty list if
// no records exist
func GetAllGateways(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Gateways)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Gateways
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).Filter("deleted_at__isnull", true).RelatedSel().All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				v.loadRelations()
				v.Orders = nil
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				v.loadRelations()
				v.Orders = nil
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateGatewaysByID updates Gateways by Id and returns error if
// the record to be updated doesn't exist
func UpdateGatewaysByID(m *Gateways) (err error) {
	o := orm.NewOrm()
	v := Gateways{ID: m.ID}
	// ascertain id exists in the database

	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}

	}
	return
}

// DeleteGateways deletes Gateways by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGateways(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := Gateways{ID: id}
	// ascertain id exists in the database
	err = o.Read(&v)

	if err != nil {
		return
	}

	if trash {
		_, err = o.Delete(&v)
	} else {
		v.DeletedAt = time.Now()
		_, err = o.Update(&v)
	}

	if err != nil {
		return
	}

	return
}

//AddDefaultDataGateways on init app
func AddDefaultDataGateways() (count int64, err error) {

	o := orm.NewOrm()

	dummyData := []*Gateways{
		{
			Name:         "Paypal",
			Code:         "01",
			Description:  "Paga con PayPal; puedes pagar con tu tarjeta de crédito si no tienes una cuenta PayPal.",
			Instructions: "1: Verifique su bandeja de entrada para ver si ha recibido un mensaje de correo electrónico que incluya la solicitud de pago o el formato de pago .<br><br> 2: Haga clic en el botón Pagar ahora en el correo electrónico. (Si no lo ve, haga clic en el link del correo electrónico o copie y pegue el link en la barra de direcciones del navegador).<br><br> 3: Si ya tiene una cuenta PayPal, ingrese su contraseña y haga clic en Iniciar sesión.<br><br> Revise la solicitud de pago o el formato de pago y, a continuación, haga clic en Pagar ahora para completar la transacción.",
		},
		{
			Name:         "Stripe",
			Code:         "02",
			Description:  "Paga con tu tarjeta de crédito a través de Stripe, es rápido y seguro.",
			Instructions: "1: Ingresa los datos de tu tarjeta de crédito habilitada para pagos internacionales <br><br> 2: Has clic en pagar y ¡LISTO!",
		},
		{
			Name:         "Transferencia Bancaría Banco Santander",
			Code:         "03",
			Description:  "Pago a través de Transferencia o Ingreso en cuenta en el Banco Santander. Por favor use su nombre de proyecto como referencia en su pago. Su orden será tramitada una vez se confirme su pago.",
			Instructions: "Pago a través de Transferencia o Ingreso en cuenta en Banco Santander Central Hispano<br> Numero de cuenta bancaria: 0049 1555 14 2810178350 <br> Beneficiario: Solmax Europe SLU<br> CIF: B64493018<br>",
		},
		{
			Name:         "Deposito o transferencia",
			Code:         "04",
			Description:  "Pago a través de Transferencia o Ingreso en cuenta en el Banco Santander. Por favor use su nombre de proyecto como referencia en su pago. Su orden será tramitada una vez se confirme su pago.",
			Instructions: "Pago a través de Transferencia o Ingreso en cuenta:<br> Banco santander<br> Beneficiario: Solmax Europe SLU<br> Identificación Fiscal y Tributaria: B64493018<br> Número de Cuenta Bancaria: 65-50362108-8<br> CLABE 014180655036210882 <br>",
		},
		{
			Name:         "Paypal",
			Code:         "05",
			Description:  "SafetyPay te permite realizar pagos en línea desde tu cuenta bancaria. La dirección de facturación debe estar en Brasil, México, Costa Rica, Perú, España o Austria para pagar mediante SafetyPay.",
			Instructions: "1: Se te remitirá al sitio web de SafetyPay. <br><br> 2: En el sitio web de SafetyPay, selecciona el banco desde el que deseas efectuar el pago. <br><br> 3: Aparecerá un mensaje que indica las horas durante el cual se aceptarán los pagos. Sigue las instrucciones del sitio para hacer un pago.<br><br> 4: SafetyPay autoriza la transacción una vez que se ha completado el pago. El tiempo que tarda este proceso puede variar de un banco a otro. <br><br> 5: Recibirás una confirmación cuando el pago se realice correctamente.<br><br>",
		},
	}

	count, err = o.InsertMulti(100, dummyData)

	return
}

//GetGatewaysFromTrash return Gateways soft Deleted
func GetGatewaysFromTrash() (gateways []*Gateways, err error) {

	o := orm.NewOrm()

	var v []*Gateways

	_, err = o.QueryTable("gateways").Filter("deleted_at__isnull", false).RelatedSel().All(&v)

	if err != nil {
		return
	}

	for _, currency := range v {
		currency.loadRelations()
	}

	gateways = v

	return

}

<div
  style="
  width: 600px;
  background: white;
  margin: auto;
"
>
  <div>
    <h1
      style="
    margin-bottom: 0;
    font-family: sans-serif;
    color: #414143;
    font-size: 26px;
    font-weight: bolder;
    margin-top: 30px;
    text-align: center;
    "
    >
      GRACIAS POR SU PEDIDO
    </h1>
    <h4
      style="  
      color: #414143;
      font-size: 16px;
      font-weight: 500;
      font-family: sans-serif;
      margin-top: 8px;
      margin-bottom: 8px;
      "
    >
      En breve, uno de nuestros asesores de diseño se comunicará contigo.
    </h4>
    <hr
    style="
    margin-bottom: 0;
    height: 3px;
    background: #007ab5;
    border: none;
    "
    />
  </div>
  <br /><br />
  <div>
    <h5
      style="
      color: #414143;
      font-size: 20px;
      font-weight: bold;
      font-family: sans-serif;
      margin-top: 8px;
      margin-bottom: 0px;
      "
    >
      RESUMEN DE SU PEDIDO
    </h5>
  </div>
  <br>
  <div>
    <div>
    <table style="width: 100%;">
      <tr>
        <th
          style="    
        padding: 8px 8px 8px 0px;
        font-family: sans-serif;
        font-size: 16px;
        font-weight: bold;
        margin: 10px 20px;
        color: #414143;
        width: 50%;
        text-align: left;
        "
        >
          Servicio
        </th>
        <th
        style="    
        padding: 8px 8px 8px 0px;
        text-align: right;
        font-family: sans-serif;
        font-size: 16px;
        font-weight: bold;
        margin: 10px 20px;
        color: #414143;
        width: 50%;
        "
        >
          Precio
        </th>
      </tr>
      {{ range .Services }}
        <tr>
          <td
            style="    
          padding: 8px 8px 8px 0px;
          font-family: sans-serif;
          font-size: 14px;
          font-weight: 500;
          margin: 10px 20px;
          color: #414143;
          width: 50%;
          text-align: left;
          "
          >
            {{.Name}}
          </td>
          <td
            style="    
          padding: 8px 8px 8px 0px;
          font-family: sans-serif;
          font-size: 14px;
          font-weight: 500;
          margin: 10px 20px;
          color: #414143;
          width: 50%;
          text-align: right;
          "
          >
          {{.Price.Symbol}} {{.Price.Value}} 
          </td>
        </tr>
      {{ end }}
    </table>
    <hr
    style="
    margin-bottom: 0;
    height: 3px;
    background: #007ab5;
    border: none;
    "
    />
    <table style="width: 100%;">
      <tr>
        <th
        style="    
            padding: 8px 8px 8px 0px;
        text-align: center;
        font-family: sans-serif;
        font-size: 16px;
        font-weight: bold;
        margin: 10px 20px;
        color: #414143;
        width: 50%;
        " colspan="2"
        >
          <div style="    width: 270px;
          display: flex;
          justify-content: space-between;
          margin-left: auto;
      ">
            <span style="width: 135px;
            text-align: left;">
              Pago Inicial
            </span>
            <span style="width: 135px;
            text-align: right;">
              Pago Total
            </span>
          </div>        
        </th>
      </tr>
      <tr>
          <td
          style="    
            padding: 8px 8px 8px 0px;
        font-family: sans-serif;
        font-size: 14px;
        font-weight: 500;
        margin: 10px 20px;
        color: #414143;
        width: 50%;
        word-break: break-word;
        text-align: center;
        " colspan="2"
        >
          <div style="    width: 270px;
          display: flex;
          justify-content: space-between;
          margin-left: auto;
      ">
            <span style="width: 135px;
            text-align: left;">
              {{.Symbol}} {{.Order.GetInitialPaymentAmount}} 
            </span>
            <span style="width: 135px;
            text-align: right;">
              {{.Symbol}} {{.Order.GetFinalPaymentAmount}}
            </span>
          </div>
        </td>
      </tr>
    </table>
  </div>

  <br />

  <h4 style="    
  padding: 8px 8px 8px 0px;
  font-family: sans-serif;
  font-size: 14px;
  font-weight: 500;
  color: #414143;
  width: 50%;
  word-break: break-word;
  margin: 0;
  ">
    <b>Referencia:</b> {{.Order.ID}}
  </h4>
  
  <h4 style="    
  padding: 8px 8px 8px 0px;
  font-family: sans-serif;
  font-size: 14px;
  font-weight: 500;
  color: #414143;
  width: 50%;
  word-break: break-word;
  margin: 0;
  ">
    <b>Método de Pago:</b> {{.Order.Gateway.Name}}
  </h4>
  
  <br />
  <br />
  <br />
</div>
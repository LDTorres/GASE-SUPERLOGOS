<div
  style="
  width: 600px;
  background: white;
  margin: auto;
"
>
  <div>
    <img
      style="display: block; width: 40%; margin: auto;"
      src="http://liderlogos.com/_nuxt/img/logo.f23ffaf.png"
    />
  </div>
  <div>
    <h1
      style="
    margin-bottom: 0;
    font-family: sans-serif;
    color: #414143;
    font-size: 26px;
    font-weight: bolder;
    margin-top: 30px;
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
      margin-bottom: 8px;
      "
    >
      RESUMEN DE SU PEDIDO
    </h5>
    <hr
      style="
      margin-bottom: 0;
      height: 3px;
      background: #007ab5;
      border: none;
      "
      />
  </div>
  <br>
  <div>
    <div>
    <table style="width: 100%;">
      <tr>
        <th
          style="    
        padding: 8px;
        font-family: sans-serif;
        font-size: 16px;
        font-weight: bold;
        margin: 10px 20px;
        color: #414143;
        width: 50%;
        "
        >
          Nombre
        </th>
        <th
        style="    
        padding: 8px;
        text-align: center;
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
          padding: 8px;
          font-family: sans-serif;
          font-size: 14px;
          font-weight: 500;
          margin: 10px 20px;
          color: #414143;
          width: 50%;
          text-align: center;
          "
          >
            {{.Name}}
          </td>
          <td
            style="    
          padding: 8px;
          font-family: sans-serif;
          font-size: 14px;
          font-weight: 500;
          margin: 10px 20px;
          color: #414143;
          width: 50%;
          text-align: center;
          "
          >
            {{.Price.Value}} {{.Price.Currency.Symbol}}
          </td>
        </tr>
      {{ end }}

      <tr>
        <th
          style="    
        padding: 8px;
        font-family: sans-serif;
        font-size: 16px;
        font-weight: bold;
        margin: 10px 20px;
        color: #414143;
        width: 50%;
        
        "
        >
          Pago Inicial
        </th>
        <th
        style="    
        padding: 8px;
        text-align: center;
        font-family: sans-serif;
        font-size: 16px;
        font-weight: bold;
        margin: 10px 20px;
        color: #414143;
        width: 50%;
        "
        >
          Pago Total
        </th>
      </tr>
      <tr>
        <td
          style="    
        padding: 8px;
        font-family: sans-serif;
        font-size: 14px;
        font-weight: 500;
        margin: 10px 20px;
        color: #414143;
        width: 50%;
        word-break: break-word;
        text-align: center;
        "
        >
        {{.Services[0].Price.Currency.Symbol}}{{.Order.GetInitialPaymentAmount}}
        </td>
        <td
          style="    
        padding: 8px;
        font-family: sans-serif;
        font-size: 14px;
        font-weight: 500;
        margin: 10px 20px;
        color: #414143;
        width: 50%;
        word-break: break-word;
        text-align: center;
        "
        >
        {{.Services[0].Price.Currency.Symbol}}{{.Order.GetFinalPaymentAmount}}
        </td>
      </tr>
    </table>
  </div>

  <br />
  <br />

  <h4 style="    
  padding: 8px;
  font-family: sans-serif;
  font-size: 14px;
  font-weight: 500;
  color: #414143;
  width: 50%;
  word-break: break-word;
  ">
    <b>Referencia:</b> {{.Order.ID}}
  </h4>
  
  <h4 style="    
  padding: 8px;
  font-family: sans-serif;
  font-size: 14px;
  font-weight: 500;
  color: #414143;
  width: 50%;
  word-break: break-word;
  ">
    <b>Método de Pago:</b> {{.Order.Gateway.Name}}
  </h4>
  
  <br />
  <br />
  <!-- Not important -->
  <div
    style="text-align: center;width:600px;margin: auto;font-size: 13pt;color: #969393;font-family: sans-serif"
  >
    <b>¿Consultas?</b> Escribanos a
    <a
      style="color: #0081c1;text-decoration: none !important"
      href="mailto:info@liderlogo.com"
      target="_top"
      >info@liderlogo.com</a
    >
    o visitenos
    <a
      style="color: #0081c1;text-decoration: none !important"
      href="https://www.liderlogo.com/contacto/"
      target="_new"
      >liderlogo.com/contacto</a
    >
  </div>

  <div
    style="text-align: center;margin: auto;font-size: 13pt;color: #969393;font-family: sans-serif;margin-top: 10px"
  >
    <div style="margin-bottom: 10px"><b>Siganos</b></div>
    <div style="
        color: transparent;margin: auto;
        margin: auto;
        margin-bottom: 15px;
        width: 300px;
        display: block;">
      <a
        href="https://www.facebook.com/liderlogo"
        class="icon"
        target="_blank"
        style="      
        border: 1px solid silver;
        width: 50px;
        color: transparent;
        height: 50px;
        text-decoration: none !important;
        margin-right: 10px;
        border-radius: 50%;
        display: inline-block;"
        ><img
          style="
          width: 32px;
          height: 32px;
          margin-top: 9px;
          "
          src="http://www.liderlogo.info/wp-content/uploads/facebook-logo2.png"
      /></a>
      <a href="https://twitter.com/lider_logo" class="icon"
      style="      
        border: 1px solid silver;
        width: 50px;
        color: transparent;
        height: 50px;
        text-decoration: none !important;
        margin-right: 10px;
        border-radius: 50%;
        display: inline-block;"

        ><img style="
        width: 32px;
        height: 32px;
        margin-top: 9px;
        "
        src="http://www.liderlogo.info/wp-content/uploads/twitter1.png" />
      </a>
      <a
        href="https://plus.google.com/114230428674579789847"
        class="icon"
        target="_blank"
        style="      
        border: 1px solid silver;
        width: 50px;
        color: transparent;
        height: 50px;
        text-decoration: none !important;
        margin-right: 10px;
        border-radius: 50%;
        display: inline-block;"
        ><img
          style="
          width: 32px;
          height: 32px;
          margin-top: 9px;
          "
          src="http://www.liderlogo.info/wp-content/uploads/google-plus1.png"
        />
      </a>
      <a
        href="https://www.linkedin.com/company-beta/2369935"
        class="icon"
        target="_blank"
        style="      
        border: 1px solid silver;
        width: 50px;
        color: transparent;
        height: 50px;
        text-decoration: none !important;
        margin-right: 10px;
        border-radius: 50%;
        display: inline-block;"
      >
        <img
          style="
          width: 32px;
          height: 32px;
          margin-top: 9px;
          "
          src="http://www.liderlogo.info/wp-content/uploads/linkedin-letters1.png"
        />
      </a>
    </div>
  </div>
</div>
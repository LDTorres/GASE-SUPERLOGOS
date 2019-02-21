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
    NUEVO BRIEF
  </h1>
  <hr
    style="
  margin-bottom: 0;
  height: 3px;
  background: #007ab5;
  border: none;
  "
  />
  <h4
    style="  
      color: #414143;
      font-size: 14px;
      font-weight: 500;
      font-family: sans-serif;
      margin-top: 8px;
      margin-bottom: 8px;
      "
  >
    <b>ID:</b> <a href="http://admin.liderlogos.com/#/briefs?q={{.BriefID}}" target="_new">{{.BriefID}}</a>
  </h4>

  <h4
    style="  
      color: #414143;
      font-size: 14px;
      font-weight: 500;
      font-family: sans-serif;
      margin-top: 8px;
      margin-bottom: 8px;
      "
  >
    <b>Nombre:</b> {{.Client.Name}}
  </h4>

  <h4
    style="  
      color: #414143;
      font-size: 14px;
      font-weight: 500;
      font-family: sans-serif;
      margin-top: 8px;
      margin-bottom: 8px;
      "
  >
    <b>Email:</b> {{.Client.Email}}
  </h4>

  <h4
    style="  
      color: #414143;
      font-size: 14px;
      font-weight: 500;
      font-family: sans-serif;
      margin-top: 8px;
      margin-bottom: 8px;
      "
  >
    <b>Telefono:</b> {{.Client.Phone}}
  </h4>

  <h4
    style="  
      color: #414143;
      font-size: 14px;
      font-weight: 500;
      font-family: sans-serif;
      margin-top: 8px;
      margin-bottom: 8px;
      "
  >
    <b>Pais:</b> {{.Country.Name}}
  </h4>
</div>

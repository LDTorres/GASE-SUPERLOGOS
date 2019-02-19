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
    <b>Datos de contacto:</b>
  </div>
  <br />
  <div><b>Nombre:</b> {{.Lead.Name}}</div>
  <div><b>Email:</b> {{.Lead.Email}}</div>
  <div><b>Mensaje:</b> {{.Lead.Message}}</div>
  {{if .Lead.Phone}}
  <div><b>Telefono:</b> {{.Lead.Phone}}</div>
  <div><b>Horario:</b> {{.Lead.Schedule}}</div>
  {{ end }}
  <div><b>Pais:</b> {{.Country.Name}}</div>
</div>

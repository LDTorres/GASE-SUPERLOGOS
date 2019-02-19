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
    <p>¡Hola!.</p>
    <br />
    <p>
      El <b>Boceto Version #{{.Sketch.Version}}</b> del
      <b>Proyecto {{.Project.Name}}</b> ha sido comentado por el cliente
      {{.Client.Name}}.
    </p>
  </div>
  <br />
  <div>
    <p>"{{.Comment.Description}}"</p>
  </div>
</div>

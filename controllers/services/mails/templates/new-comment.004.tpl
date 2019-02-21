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
    !Hola¡
  </h1>
  <hr
  style="
  margin-bottom: 0;
  height: 3px;
  background: #007ab5;
  border: none;
  "
  />
  <div>
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

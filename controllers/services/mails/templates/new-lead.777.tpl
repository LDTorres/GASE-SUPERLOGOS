<div>
    <div><b>Datos de contacto:</b></div>
</div>
<br>
<div><b>Nombre:</b> {{.Lead.Name}}</div>
<div><b>Email:</b> {{.Lead.Email}}</div> 
<div><b>Mensaje:</b> {{.Lead.Message}}</div> 
{{if .Lead.Phone}}
<div><b>Telefono:</b> {{.Lead.Phone}}</div> 
<div><b>Horario:</b> {{.Lead.Schedule}}</div> 
{{end}}
<div><b>Pais:</b> {{.Country.Name}}</div> 

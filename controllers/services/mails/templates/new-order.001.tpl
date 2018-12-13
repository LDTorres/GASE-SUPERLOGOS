<div>
    <p>Gracias por su Pedido!. </p>
    <br>
    <p>En breve, uno de nuestros asesores de diseño se comunicará contigo.</p>
</div>
<br>
<div>A continuación, le presentamos el resumen de su pedido:</div>
<br>
<div>
    <table>
        <thead>
            <tr>
                <th>Referencia</th>
                <th>Precio</th>
                <th>Pago Inicial</th>
                <th>Estado</th>
                <th>Método de Pago</th>
                
            </tr>
        </thead>
        <tbody>
           <tr>
               <td>{{.ID}}</td>
               <td>{{.GetFinalPaymentAmount}}</td>
               <td>{{.GetInitialPaymentAmount}}</td>
               <td>{{.Status}}</td>
               <td>{{.Gateway.Name}}</td>
           </tr>
        </tbody>
    </table>
</div>

<br>

<div>Los servicios que contrató:</div>
<br>
<div>
    <table>
        <thead>
            <tr>
                <th>Nombre</th>                
            </tr>
        </thead>
        <tbody>
            {{range .Services}}
           <tr>
               <td>{{.Name}}</td>
           </tr>
           {{end}}
        </tbody>
    </table>
</div>
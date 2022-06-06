{{define "create"}}
{{template "header"}}

<div class="card">

{{template "headercreateorder"}}


<div class="card">

    <div class="card-header">

        INGRESAR VALORES DEL NUEVO PEDIDO

    </div>


    <div class="card-body">

        <form method="post" action="/registerorder">

        <div class="form-group">
            <label for=""> DIRECCION: </label>
            <input type="text"
            class="form-control" name="direction" id="direction" aria-describedby="helpId" placeholder="">
            <small id="helpId" class="form-text text-muted"> DIRECCION + DEPARTAMENTO </small>
        </div>

       <div class="form-group">
            <label for=""> CONTACTO:  </label>
            <input type="text"
            class="form-control" name="contact" id="contact" aria-describedby="helpId" placeholder="">
            <small id="helpId" class="form-text text-muted"> NUMERO DE TELEFONO </small>
        </div>

        

        <button type="submit" class="btn btn-primary"> REGISTRAR PEDIDO </button>

        <a name="" id="" class="btn btn-success" href="/home" role="button"> REGRESAR </a>
        
        </form>

    </div>
  


  

    <div class="card-footer text-muted">

        PRODUCTOS AÑADIDOS HASTA EL MOMENTO: 

    </div>

  


{{template "footerorder" .}}
{{end}}
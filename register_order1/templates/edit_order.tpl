{{define "editorder"}}
{{template "header"}}


<div class="card">

    <div class="card-header">

        INGRESAR VALORES DEL PEDIDO A MODIFICAR

    </div>

    <div class="card-body">

        <form method="post" action="/updateorder">

        <div class="form-group">
            <label class="sr-only" for="inputName"> ID PEDIDO </label>
            <input type="text" value={{.Id}} class="form-control" name="id" id="id" placeholder="">
        </div>


        <div class="form-group">
            <label for=""> NOMBRE: </label>
            <input type="text"
            class="form-control" name="name" value={{.Name}} id="name" aria-describedby="helpId" placeholder="">
            <small id="helpId" class="form-text text-muted"></small>
        </div>

        <div class="form-group">
            <label for=""> RELLENO / SABOR: </label>
            <input type="text"
            class="form-control" name="amount" value={{.Info}} id="amount" aria-describedby="helpId" placeholder="">
            <small id="helpId" class="form-text text-muted"></small>
        </div>

        <div class="form-group">
            <label for=""> CANTIDAD: </label>
            <input type="text"
            class="form-control" name="amount" value={{.Amount}} id="amount" aria-describedby="helpId" placeholder="">
            <small id="helpId" class="form-text text-muted"></small>
        </div>

        <div class="form-group">
            <label for=""> TOTAL: </label>
            <input type="text"
            class="form-control" name="total" value={{.Total}} id="total" aria-describedby="helpId" placeholder="">
            <small id="helpId" class="form-text text-muted"></small>
        </div>

        <button type="submit" class="btn btn-primary"  href="/updateorder?id={{.Id}}" > CONFIRMAR  </button>

        <a name="" id="" class="btn btn-success" href="/homeorder" role="button"> REGRESAR </a>

</div>






{{template "footer"}}
{{end}}
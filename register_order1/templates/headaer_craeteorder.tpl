   {{define "headercreateorder"}}
   
    <div class="card-header">

        INGRESAR VALORES EL PRODUCTO

    </div>

 <form method="post" action="/insertproduct">
 <form method="post" action="/insertproduct">
    <div class="card-body">

        
        <div class="form-group">
            <label for=""> CÓDIGO DE PRODUCTO: </label>
            <input type="text"
            class="form-control" name="id" id="id" aria-describedby="helpId" placeholder="">
            <small id="helpId" class="form-text text-muted"></small>
        </div>

        <div class="form-group">
            <label for=""> RELLENO / SABOR: </label>
            <input type="text"
            class="form-control" name="info" id="info" aria-describedby="helpId" placeholder="">
            <small id="helpId" class="form-text text-muted"></small>
        </div>


         <div class="form-group">
            <label for=""> CANTIDAD: </label>
            <input type="text"
            class="form-control" name="amount" id="amount" aria-describedby="helpId" placeholder="">
            <small id="helpId" class="form-text text-muted"></small>
        </div>

        <button type="submit" class="btn btn-primary"> AGREGAR PRODUCTO </button>

        </form>

    </div>

</form>

{{end}}
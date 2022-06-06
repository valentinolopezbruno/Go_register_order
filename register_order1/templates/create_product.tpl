{{define "createproduct"}}

{{template "header"}}


<div class="card">

    <div class="card-header">

        INGRESAR VALORES DEL NUEVO PRODUCTO

    </div>



    <div class="card-body">

        <form method="post" action="/createproduct">

        <div class="form-group">
            <label for=""> NOMBRE: </label>
            <input type="text"
            class="form-control" name="name" id="name" aria-describedby="helpId" placeholder="">
            <small id="helpId" class="form-text text-muted"></small>
        </div>

                <div class="form-group">
            <label for=""> PRECIO: </label>
            <input type="text"
            class="form-control" name="price" id="price" aria-describedby="helpId" placeholder="">
            <small id="helpId" class="form-text text-muted"></small>
        </div>

        <div class="form-group">
            <label for=""> ID:  </label>
            <input type="text"
            class="form-control" name="id" id="id" aria-describedby="helpId" placeholder="">
            <small id="helpId" class="form-text text-muted"></small>
        </div>

        <button type="submit" class="btn btn-primary"> REGISTRAR PRODUCTO </button>

        <a name="" id="" class="btn btn-success" href="/homeproduct" role="button"> REGRESAR </a>




{{template "footer"}}

{{end}}
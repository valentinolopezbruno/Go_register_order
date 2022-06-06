{{define "editproduct"}}
{{template "header"}}


<div class="card">

    <div class="card-header">

        INGRESAR VALORES DEL PRODUCTO A MODIFICAR

    </div>

    <div class="card-body">

        <form method="post" action="/updateproduct">

        <div class="form-group">
        <label for=""> ID PRODUCTO: </label>
            <label class="sr-only" for="inputName"> ID PRODUCTO </label>
            <input type="text" value={{.Id}} class="form-control" name="id" id="id" placeholder="">
        </div>


        <div class="form-group">
            <label for=""> NOMBRE: </label>
            <input type="text"
            class="form-control" name="name" value={{.Name}} id="name" aria-describedby="helpId" placeholder="">
            <small id="helpId" class="form-text text-muted"></small>
        </div>

        <div class="form-group">
            <label for=""> PRECIO: </label>
            <input type="text"
            class="form-control" name="price" value={{.Price}} id="price" aria-describedby="helpId" placeholder="">
            <small id="helpId" class="form-text text-muted"></small>
        </div>

        <button type="submit" class="btn btn-primary"  href="/updateproduct?id={{.Id}}" > CONFIRMAR  </button>

        <a name="" id="" class="btn btn-success" href="/homeproduct" role="button"> REGRESAR </a>

</div>






{{template "footer"}}
{{end}}
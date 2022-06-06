{{define "showorderinfo"}}

{{template "header"}}

            <table class="table">

                <thead>
                    <tr>
                        <th> NOMBRE </th>
                        <th> RELLENO / SABOR </th>
                        <th> CANTIDAD </th>
                        <th> TOTAL </th>
                    </tr>
                </thead>

                <tbody>

                {{range.}}
                    <tr>
                        <td>{{.Name}}</td>
                        <td>{{.Info}}</td>
                        <td>{{.AmountProduct}}</td>
                        <td>{{.Total}}</td>
                    </tr>
                {{end}}

                </tbody>

                                        
                        

             </table>
             <a name="" id="" class="btn btn-success w-100 m-5 h-5" href="/home" role="button"> REGRESAR </a>  
{{end}}
package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var temp = template.Must(template.ParseGlob("templates/*"))

type ProductInOrder struct {
	Id            int
	IdOrder       int
	IdProduct     int
	AmountProduct int
	Total         int
	Name          string
	TotalPrice    int
	Info          string
}

type Product struct {
	Id         int
	Name       string
	Price      int
	Amount     int
	TotalPrice int
}

type Order struct {
	Id        int
	Direction string
	Price     int
	Contact   string
}

func main() {
	//HANDLERS DE PEDIDOS
	http.HandleFunc("/", Home)
	http.HandleFunc("/createorder", Createorder)
	http.HandleFunc("/deleteorder", DeleteOrder)
	http.HandleFunc("/registerorder", RegisterOrder)
	http.HandleFunc("/showorderinfo", ShowOrderInfo)
	//http.HandleFunc("/editorder", EditOrder)

	// HANDLERS DE PRODUCTOS:
	http.HandleFunc("/insertproduct", InsertProductInOrder)
	http.HandleFunc("/createproduct", CreateproductInDB)
	http.HandleFunc("/deleteproductintable", DeleteProductInTable)
	http.HandleFunc("/editproduct", Editproduct)
	http.HandleFunc("/deleteproduct", Deleteproduct)
	http.HandleFunc("/updateproduct", Updateproduct)
	http.HandleFunc("/homeproduct", Homeproduct)

	fmt.Println("SERVIDOR INICIADO")
	http.ListenAndServe(":8080", nil)
}

func connectDB() (conn *sql.DB) {
	Name := "orders"
	conn, err := sql.Open("mysql", "root:root@tcp(127.0.0.1)/"+Name) //nameBD == product

	if err != nil {
		panic(err.Error())
	}
	return conn
}

func Home(w http.ResponseWriter, r *http.Request) {
	connect := connectDB()
	getOrder, err := connect.Query("SELECT * FROM orders")

	if err != nil {
		panic(err.Error())
	}

	order := Order{}
	sliceOrder := []Order{}

	for getOrder.Next() {
		var id, price int
		var direction, contact string

		err = getOrder.Scan(&id, &direction, &price, &contact)

		if err != nil {
			panic(err.Error())
		}

		order.Id = id
		order.Direction = direction
		order.Price = price
		order.Contact = contact

		sliceOrder = append(sliceOrder, order)
	}

	temp.ExecuteTemplate(w, "home", sliceOrder)
}

func Createorder(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "create", nil)
}

var product = Product{}
var SliceProducts = []Product{}
var TotalPrice = 0

func getOrderID() int {
	connect := connectDB()
	getOrderId, err := connect.Query("SELECT * FROM orders")

	if err != nil {
		panic(err.Error())
	}

	order := Order{}

	for getOrderId.Next() {
		var id, price int
		var direction, contact string
		err = getOrderId.Scan(&id, &direction, &price, &contact)

		if err != nil {
			panic(err.Error())
		}
		order.Id = id
	}

	a := order.Id
	return a
}

func RegisterOrder(w http.ResponseWriter, r *http.Request) {
	connect := connectDB()

	getInfo, err := connect.Query("SELECT * FROM productsinorder")

	if err != nil {
		panic(err.Error())
	}

	productinorder := ProductInOrder{}

	for getInfo.Next() {
		var id, idorder, idproduct, amount, total int
		var name, info string
		err = getInfo.Scan(&id, &idorder, &idproduct, &amount, &total, &name, &info)

		if err != nil {
			panic(err.Error())
		}

		productinorder.Id = id
		productinorder.IdOrder = idorder
		productinorder.IdProduct = idproduct
		productinorder.AmountProduct = amount
		productinorder.Total = total
		productinorder.Name = name
		productinorder.Info = info

		if productinorder.IdOrder == (getOrderID() + 1) {
			productinorder.TotalPrice = (productinorder.TotalPrice + productinorder.Total)
		}
	}

	direction := r.FormValue("direction")
	contact := r.FormValue("contact")

	registerOrderInDB, err := connect.Prepare("INSERT INTO orders(direction,price,contact) VALUES(?,?,?)")

	if err != nil {
		panic(err.Error())
	}

	registerOrderInDB.Exec(direction, productinorder.TotalPrice, contact)
	http.Redirect(w, r, "/home", http.StatusPermanentRedirect)
}

func InsertProductInOrder(w http.ResponseWriter, r *http.Request) {

	connect := connectDB()

	idd := r.FormValue("id")
	info := r.FormValue("info")
	cantidad := r.FormValue("amount")

	iddd, err := strconv.Atoi(idd)

	if err != nil {
		panic(err.Error())
	}

	amount, err := strconv.Atoi(cantidad)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("VALORES INSERTADOS: ", "ID: ", iddd, " AMOUNT: ", amount, "INFO: ", info)

	if r.Method == "POST" {

		if err != nil {
			panic(err.Error())
		}

		getProduct, err := connect.Query("SELECT * FROM products")

		if err != nil {
			panic(err.Error())
		}

		for getProduct.Next() {
			var id, price int
			var name string

			err = getProduct.Scan(&id, &name, &price)

			if err != nil {
				panic(err.Error())
			}

			product.Id = id
			product.Name = name
			product.Price = price
			product.Amount = amount
			product.TotalPrice = (product.Price * product.Amount)

			if product.Id == iddd {
				insertIntoProductsInOrder, err := connect.Prepare("INSERT INTO productsinorder(idorder,idproduct,amountproduct,total,name,info) VALUES(?,?,?,?,?,?)")

				if err != nil {
					panic(err.Error())
				}

				insertIntoProductsInOrder.Exec((getOrderID() + 1), product.Id, product.Amount, product.TotalPrice, product.Name, info)

				fmt.Println("REGISTRO REALIZADO EN LA TABLA productsinorder: ", (getOrderID() + 1), product.Id, product.Amount, product.TotalPrice, product.Name, info)
			}
		}

		getInfo, err := connect.Query("SELECT * FROM productsinorder")

		if err != nil {
			panic(err.Error())
		}

		productinorder := ProductInOrder{}
		Sliceproductinorder := []ProductInOrder{}

		for getInfo.Next() {
			var id, idorder, idproduct, amount, total int
			var name, info string
			err = getInfo.Scan(&id, &idorder, &idproduct, &amount, &total, &name, &info)

			if err != nil {
				panic(err.Error())
			}

			productinorder.Id = id
			productinorder.IdOrder = idorder
			productinorder.IdProduct = idproduct
			productinorder.AmountProduct = amount
			productinorder.Total = total
			productinorder.Name = name
			productinorder.Info = info

			if productinorder.IdOrder == (getOrderID() + 1) {
				productinorder.TotalPrice = (productinorder.TotalPrice + productinorder.Total)
				Sliceproductinorder = append(Sliceproductinorder, productinorder)
				fmt.Println("EL Sliceproductinorder CONTIENE LOS PRODUCTOS: ", Sliceproductinorder)
			}
		}

		temp.ExecuteTemplate(w, "create", Sliceproductinorder)
	}
}

func CreateproductInDB(w http.ResponseWriter, r *http.Request) {
	connect := connectDB()
	if r.Method == "POST" {
		name := r.FormValue("name")

		precio := r.FormValue("price")
		price, err := strconv.Atoi(precio)
		if err != nil {
			panic(err.Error())
		}

		idd := r.FormValue("id")
		id, err := strconv.Atoi(idd)
		if err != nil {
			panic(err.Error())
		}

		createProduct, err := connect.Prepare("INSERT INTO products(id,name,price) VALUES(?,?,?)")

		if err != nil {
			panic(err.Error())
		}

		createProduct.Exec(id, name, price)

		http.Redirect(w, r, "/homeproduct", http.StatusPermanentRedirect)
	}
	temp.ExecuteTemplate(w, "createproduct", nil)
}

func Editproduct(w http.ResponseWriter, r *http.Request) {
	connect := connectDB()
	getIdProduct := r.URL.Query().Get("id")
	editProduct, err := connect.Query("SELECT * FROM products WHERE id=?", getIdProduct)

	if err != nil {
		panic(err.Error())
	}

	product := Product{}

	for editProduct.Next() {
		var id, price int
		var name string
		err = editProduct.Scan(&id, &name, &price)

		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Price = price
	}

	fmt.Println(product)
	temp.ExecuteTemplate(w, "editproduct", product)
}

func Updateproduct(w http.ResponseWriter, r *http.Request) {
	connect := connectDB()
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		price := r.FormValue("price")

		updateProduct, err := connect.Prepare("UPDATE products SET name=?, price=? WHERE id=?")

		if err != nil {
			panic(err.Error())
		}

		updateProduct.Exec(name, price, id)

		http.Redirect(w, r, "homeproduct", http.StatusPermanentRedirect)
	}
}

func Deleteproduct(w http.ResponseWriter, r *http.Request) {
	getIdProduct := r.URL.Query().Get("id")
	connect := connectDB()
	deleteProduct, err := connect.Prepare("DELETE FROM products WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(getIdProduct)
	http.Redirect(w, r, "/homeproduct", http.StatusPermanentRedirect)
}

func Homeproduct(w http.ResponseWriter, r *http.Request) {
	connect := connectDB()
	getProduct, err := connect.Query("SELECT * FROM products")

	if err != nil {
		panic(err.Error())
	}

	product := Product{}
	sliceProduct := []Product{}

	for getProduct.Next() {
		var id, price int
		var name string
		err = getProduct.Scan(&id, &name, &price)

		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Price = price

		sliceProduct = append(sliceProduct, product)
	}
	temp.ExecuteTemplate(w, "homeproduct", &sliceProduct)
}
func DeleteProductInTable(w http.ResponseWriter, r *http.Request) {
	getIdProduct := r.URL.Query().Get("id")
	fmt.Println("EL ID CAPTADO ES: ", getIdProduct)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	getIdOrder := r.URL.Query().Get("id")

	fmt.Println("ID ORDER: ", getIdOrder)

	connect := connectDB()
	deleteOrder, err := connect.Prepare("DELETE FROM orders WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	deleteOrder.Exec(getIdOrder)
	fmt.Println("REGISTRO ELIMINADO EN LA TABLA ORDER CON EL ID : ", getIdOrder)
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

/* func EditOrder(w http.ResponseWriter, r *http.Request) {
	connect := connectDB()
	getIdOrder := r.URL.Query().Get("id")
	editOrder, err := connect.Query("SELECT * FROM productsinorder WHERE idorder=?", getIdOrder)

	if err != nil {
		panic(err.Error())
	}

	productinorder := ProductInOrder{}
	Sliceproductinorder := []ProductInOrder{}

	for editOrder.Next() {
		var id, idorder, idproduct, amount, total int
		var name, infoo string
		err = getInfo.Scan(&id, &idorder, &idproduct, &amount, &total, &name, &infoo)

		if err != nil {
			panic(err.Error())
		}

		productinorder.Id = id
		productinorder.IdOrder = idorder
		productinorder.IdProduct = idproduct
		productinorder.AmountProduct = amount
		productinorder.Total = total
		productinorder.Name = name
		productinorder.Info = infoo
	}

	fmt.Println(product)

	temp.ExecuteTemplate(w, "editorder", nil)
} */

func ShowOrderInfo(w http.ResponseWriter, r *http.Request) {
	connect := connectDB()
	getIdOrder := r.URL.Query().Get("id")

	idActualOrder, err := strconv.Atoi(getIdOrder)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("idActualOrder: ", idActualOrder)

	getInfo, err := connect.Query("SELECT * FROM productsinorder WHERE idorder=?", getIdOrder)

	if err != nil {
		panic(err.Error())
	}
	productinorder := ProductInOrder{}
	Sliceproductinorder := []ProductInOrder{}

	for getInfo.Next() {
		var id, idorder, idproduct, amount, total int
		var name, infoo string
		err = getInfo.Scan(&id, &idorder, &idproduct, &amount, &total, &name, &infoo)

		if err != nil {
			panic(err.Error())
		}

		productinorder.Id = id
		productinorder.IdOrder = idorder
		productinorder.IdProduct = idproduct
		productinorder.AmountProduct = amount
		productinorder.Total = total
		productinorder.Name = name
		productinorder.Info = infoo
		fmt.Println("productinorder.IdOrder: ", productinorder.IdOrder)

		if productinorder.IdOrder == idActualOrder {
			fmt.Println("ENTRO AL IF")
			Sliceproductinorder = append(Sliceproductinorder, productinorder)
			fmt.Println("EL Sliceproductinorder CONTIENE LOS PRODUCTOS: ", Sliceproductinorder)
		}
	}
	temp.ExecuteTemplate(w, "showorderinfo", Sliceproductinorder)

}

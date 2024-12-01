/*
Behavioural startegy Pattern

To decide the startergy during run time
also know as policy design pattern 

*/

package main
import "fmt"

type IDBconnection interface { // can be assigned during runtime // I stands for interface 
	Connect()
}

type DBConnection struct {
	Db IDBconnection // Compatible to accept any type in runtime  // inherited the above interface
}

func (con DBConnection) DBConnect(){
	con.Db.Connect()
}
type MySqlConnection struct{
	ConnectionString string
}

func (con MySqlConnection) Connect(){
	fmt.Println(("MySql")+con.ConnectionString)
}
type PostgressConnection struct {
	ConnectionString string
}

func (con PostgressConnection) Connect(){
	fmt.Println(("Postgress")+con.ConnectionString)
}

type MongoDbConnection struct{
	ConnectionString string
}

func (con MongoDbConnection) Connect(){
	fmt.Println(("MongoDb")+con.ConnectionString)
}

func main(){
	MySqlConnection:= MySqlConnection {ConnectionString:" :Connected to mysql server"}
	con:=DBConnection{Db:MySqlConnection}
	con.DBConnect()

	PostgressConnection:=PostgressConnection{ConnectionString:" :Connected to Postgress server "}
	con1:=DBConnection{Db:PostgressConnection} // what is being returned here?
	con1.DBConnect()

	MongoDbConnection:=MongoDbConnection{ConnectionString:" :Connected MonGod server "}
	con2:=DBConnection{Db:MongoDbConnection} // what is being returned here?
	con2.DBConnect()

}
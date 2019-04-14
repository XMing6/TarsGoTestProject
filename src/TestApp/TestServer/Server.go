package main

import (
	"github.com/TarsCloud/TarsGo/tars"

	"TestApp/TestServer/TestApp"
	"TestApp/TestServer/hello"
)

func main() { //Init servant
	imp := new(hello.HelloWorldImp)                                    //New Imp
	app := new(TestApp.HelloWorld)                                 //New init the A Tars
	cfg := tars.GetServerConfig()                               //Get Config File Object
	app.AddServant(imp, cfg.App+"."+cfg.Server+".HelloWorldObj") //Register Servant
	tars.Run()
}

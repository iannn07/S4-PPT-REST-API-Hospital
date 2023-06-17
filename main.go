package main

import (
	"HospitalFinpro/Database"
	"HospitalFinpro/Route"
)

func main() {
	Database.ConnectDB()
	Route.MakeRoute()
}
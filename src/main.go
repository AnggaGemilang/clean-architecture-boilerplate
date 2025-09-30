package main

// entry point, file yang akan dipanggil ketika menjalankan project
// file ini kemudian akan memanggil file app/app.go


import (
	"[service_name]]/src/app"
	_ "time/tzdata"
)

func main() {
	application := app.NewApp()
	application.Run()
}

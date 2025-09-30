package app

// jantung service, isinya instansiasi dan wiring antar objek

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	println("App is running...")
}

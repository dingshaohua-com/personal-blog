package main

import (
	"backend/internal/bootstrap"
	"log"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	app, err := bootstrap.NewApp()
	if err != nil {
		log.Fatalf("应用启动失败: %v", err)
	}
	app.Run()
}

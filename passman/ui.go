package passman

import (
	"log"
	"errors"
	"fmt"
	"os"

	"bitbucket.org/rj/goey"
	"bitbucket.org/rj/goey/loop"
	"bitbucket.org/rj/goey/base"

)

var mainWindow *goey.Window

func startApp() {
	win, err := goey.NewWindow("passman", render())
	if err != nil {
		panic(err)
	}
	
} 

func showDBFileDialog() {
	dialog := mainWindow.OpenFileDialog().WithFilename("passman.db")
	db, err := dialog.Show()
	if err != nil {
		fmt.Printf("%w\n", err)
	}
	dbf, err := os.Open(db)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf(fmt.Errorf("can't open %s: file doesn't exist\n", db))
		}
	}
	if dbf.
	
}
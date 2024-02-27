package gui

import (
	"github.com/rajveermalviya/gamen/display"
	"time"
	"fmt"
)

func Run() {
	d, err := display.NewDisplay()
	if err != nil {
		panic(err)
	}
	defer d.Destroy() // destroy after run

	w, err := display.NewWindow(d)
	if err != nil {
		panic(err)
	}
	defer w.Destroy()

	w.SetCloseRequestedCallback(func() { d.Destroy() })

	fmt.Println("Sleeping...")
	time.Sleep(5 * time.Second)
}

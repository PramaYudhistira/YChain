package main

import (
	"fmt"

	"github.com/go-rod/rod"

	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {

	fmt.Println("Starting...")
	browser := rod.New().ControlURL(launcher.New().Headless(true).MustLaunch()).MustConnect()
	defer browser.MustClose()

	fmt.Println("connected to browser, opening page...")
	page := browser.MustPage("https://camelcamelcamel.com/").MustWaitStable()

	searchTerm := "iphone"
	page.MustElement("#sq").MustInput(searchTerm).MustType(input.Enter)

	page.MustWaitLoad()

	html, err := page.HTML()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Page content:\n", html)
}

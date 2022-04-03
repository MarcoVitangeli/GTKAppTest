package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

const ENTER_KEY uint = 65293

func buildSearchUrl(text string) string {
	return fmt.Sprintf("https://golangbyexample.com/?s=%s", text)
}

func buildPackageUrl(text string) string {
	return fmt.Sprintf("https://pkg.go.dev/search?q=%s", text)
}

func handleClick(searchBar *gtk.SearchEntry, urlFactory func(string) string) func() {
	return func() {
		text, err := searchBar.GetText()

		if err != nil {
			log.Fatal(err)
		}

		var url = urlFactory(text)
		log.Println(url)
		err = exec.Command("xdg-open", url).Start()

		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)
	err := gtk.WindowSetDefaultIconFromFile("./go.png")

	if err != nil {
		log.Fatal(err)
	}

	searchBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)

	if err != nil {
		log.Fatal(err)
	}

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Go searcher")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	searchBar, err := gtk.SearchEntryNew()

	if err != nil {
		log.Fatal("Unable to create search entry:", err)
	}

	exampleButton, err := gtk.ButtonNew()
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}

	var handleInput = handleClick(searchBar, buildSearchUrl)

	exampleButton.SetLabel("Go search by example")
	exampleButton.Connect("clicked", handleInput)

	searchBar.Connect("key-press-event", func(s *gtk.SearchEntry, ev *gdk.Event) {
		keyEvent := &gdk.EventKey{Event: ev}

		if keyEvent.KeyVal() == ENTER_KEY {
			handleInput()
		}
	})

	packageButton, err := gtk.ButtonNew()

	if err != nil {
		log.Fatal(err)
	}

	packageButton.SetLabel("Go search by package")
	packageButton.Connect("clicked", handleClick(searchBar, buildPackageUrl))

	searchBox.Add(searchBar)
	searchBox.Add(exampleButton)
	searchBox.Add(packageButton)
	searchBox.SetSpacing(20)
	searchBox.SetBorderWidth(20)

	win.Add(searchBox)
	// Set the default window size.
	win.SetDefaultSize(600, 600)
	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}

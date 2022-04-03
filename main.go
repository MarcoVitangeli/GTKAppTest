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

func handleClick(searchBar *gtk.SearchEntry) func() {
	return func() {
		text, err := searchBar.GetText()

		if err != nil {
			log.Fatal(err)
		}

		var url = buildSearchUrl(text)
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

	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)

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

	inputButton, err := gtk.ButtonNew()
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}

	var handleInput = handleClick(searchBar)

	inputButton.SetLabel("GO search")
	inputButton.Connect("clicked", handleInput)

	searchBar.Connect("key-press-event", func(s *gtk.SearchEntry, ev *gdk.Event) {
		keyEvent := &gdk.EventKey{Event: ev}

		if keyEvent.KeyVal() == ENTER_KEY {
			handleInput()
		}
	})

	box.Add(searchBar)
	box.Add(inputButton)
	box.SetSpacing(20)
	box.SetBorderWidth(20)
	win.Add(box)
	// Set the default window size.
	win.SetDefaultSize(800, 600)
	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}

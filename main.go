package main

import ( // import the modules
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

//Create a struct that holds information to be displayed in our HTML file
type Welcome struct {
	Name   string
	Length string
}

//Go application entrypoint
func main() {
	// create empty prompt message
	welcome := Welcome{"User", "Please enter your name below."}
	// basically tell go where the html template is and wrap it inside Must() function which aborts the program if smthg goes wrong
	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))
	// tell go where to search for the css file
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	// this function is the main pipeline between go server and html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if name := r.FormValue("name"); name != "" { // check if the name is empty, if so set the welcome struct's variables back to the original ones
			welcome.Name = name
			trimmed_name := strings.ReplaceAll(name, " ", "") // separate actual letters from white characters
			mod_name := []rune(trimmed_name)                  // polish letters are not regular ascii characters, so in order to get accurate len we have to use rune type data.
			welcome.Length = "Your name has: " + strconv.Itoa(len(mod_name)) + " characters."
		} else {
			welcome = Welcome{"User", "Please enter your name below."}
		}
		// handle errors
		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	// start the web server, set the port to listen to 8080.
	fmt.Println(http.ListenAndServe(":8080", nil))
}

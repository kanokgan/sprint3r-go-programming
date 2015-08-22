package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)

type Contact struct {
	Email 	string
	Subject string
	Message string
}

var ContactList = []Contact{
	Contact {
		Email: "kanokgan@gmail.com",
		Subject: "GoLang",
		Message: "Test",
	},
	Contact {
		Email: "kanokgan@gmail.com",
		Subject: "C#",
		Message: "Test2",
	},
}

func verbose(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}


func main() {
	
	http.Handle("/static/", 
		http.StripPrefix("/static", verbose(http.FileServer(http.Dir("./www")))))
	
	http.HandleFunc("/contact/save", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 Not found")
			return
		}

		err := r.ParseForm()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 Internal Server Error")
		}

		email := r.FormValue("email")
		subject := r.FormValue("subject")
		message := r.FormValue("message")


		newContact := Contact {
			Email: email,
			Subject: subject,
			Message: message,
		}

		ContactList = append(ContactList, newContact)

		fmt.Fprint(w, "saved")
	})
	
	http.HandleFunc("/contact/list", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.New("list").Parse(contactListTpl)
		
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 Internal Server Error")
		} 

		err = t.Execute(w, struct {
			ContactList []Contact
		}{
			ContactList: ContactList,
		})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 Internal Server Error")
		}
	})


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, indexTpl)
	})
	
	http.ListenAndServe(":8000", nil)
}




var indexTpl = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Contact</title>
	<link rel="stylesheet" href="/static/css/main.css">
</head>
<body>
	<section id="contact-header" class="wrapper">
		<div class="content">
			<h1>Geeky Base</h1>
			<p>Meet us at 2549/41-43 Phaholyothin Road, Lat Yao, Chatuchak, Bangkok.</p>
		</div>
	</section>
	<section id="send-message" class="wrapper">
		<div class="content">
			<h2>Send us a message</h2>
			<form name="contactForm" method="POST" action="/contact/save" novalidate>
				<label for="email">Email:</label>
				<input type="email" id="email" name="email" spellcheck="false" required/>
				<label for="subject">Subject:</label>
				<input type="text" id="subject" name="subject" required/>
				<label for="message">Message:</label>
				<textarea name="message" id="message" rows="5" required></textarea>
				<button id="send-button" ng-disabled="contactForm.$invalid">Send</button>
			</form>
		</div>
	</section>
</body>
</html>
`

var contactListTpl = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Contact List</title>
	<link rel="stylesheet" href="/static/css/main.css">
</head>
<body>
	<table>
		<thead>
			<tr>
				<th>Email</th>
				<th>Subject</th>
				<th>Message</th>
			</tr>
		</thead>
		<tbody>
			{{range .ContactList}}
				<tr>
					<td>{{.Email}}</td>
					<td>{{.Subject}}</td>
					<td>{{.Message}}</td>
				</tr>
			{{end}}
		</tbody>
	</table>
</body>
</html>
`













































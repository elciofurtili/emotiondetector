package controllers

import (
	"emotion-detector/models"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("views/index.gohtml"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Emotion string
		Error   string
	}{}

	if r.Method == http.MethodPost {
		r.ParseForm()
		text := r.FormValue("text")
		if text == "" {
			data.Error = "Por favor, insira algum texto."
		} else {
			emotion, err := models.DetectEmotion(text)
			if err != nil {
				data.Error = err.Error()
			} else {
				data.Emotion = emotion
			}
		}
	}

	tmpl.Execute(w, data)
}

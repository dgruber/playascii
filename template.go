package main

import (
	"bytes"
	"html/template"
)

var i = `
	<html>
	<head>
  	<link rel="stylesheet" type="text/css" href="asciinema-player.css" />
	</head>
	<body>
  	<asciinema-player src="{{.CastFile}}"></asciinema-player>
  	<script src="asciinema-player.js"></script>
	</body>
	</html>
`

func CreateIndexTemplate() (string, error) {
	tmpl, err := template.New("index.html").Parse(i)
	if err != nil {
		return "", err
	}

	data := struct {
		CastFile string
	}{
		CastFile: "demo.cast",
	}

	var out bytes.Buffer
	err = tmpl.Execute(&out, data)
	return out.String(), err
}

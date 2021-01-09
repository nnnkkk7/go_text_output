package main

import (
	"io/ioutil"
	"os"
)

func main() {
	message := "hello!!"
	tpl := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Document</title>
	</head>
	<body>
		<h1>` + message + `</h1>
	</body>
	</html>`

	os.Create("sample.html")

	ioutil.WriteFile("sample.html", []byte(tpl), os.ModePerm)
}

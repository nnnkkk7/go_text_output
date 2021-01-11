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

	//sampple.htmlファイルを作成
	os.Create("sample.html")

	//ファイルにhtmlを書き込む
	ioutil.WriteFile("sample.html", []byte(tpl), os.ModePerm)
}

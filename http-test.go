package main

import (
	"html/template"
	"net/http"
)

type pp struct {
	Num []int
}

func index(w http.ResponseWriter, r *http.Request) {
	t1, err := template.New("he").ParseFiles("hehe.html")
	if err != nil {
		panic(err)
	}
	var tem pp

	//ss := []int{1, 2, 3, 6}
	tem.Num = []int{1, 2, 3, 6}

	// for ii, kk := range tem.Num {
	// 	fmt.Println(ii, kk)
	// }
	t1.ExecuteTemplate(w, "hehe.html", tem)
	// errr := t1.Execute(w, tem)
	// fmt.Println(errr)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	server.ListenAndServe()
}

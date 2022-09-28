package resume_create

import (
	"fmt"
	"log"
	"net/http"
)

func handlePostForm(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	// 第一种方式
	// username := request.Form["username"][0]
	// password := request.Form["password"][0]

	// 第二种方式
	username := request.Form.Get("username")
	password := request.Form.Get("password")

	fmt.Printf("POST form-urlencoded: username=%s, password=%s\n", username, password)

	fmt.Fprintf(writer, `{"code":0}`)
}

func main() {

	http.HandleFunc("/handlePostForm", handlePostForm)
	log.Println("Running at port 9999 ...")
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}

}

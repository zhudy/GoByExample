//https://blog.csdn.net/yjp19871013/article/details/83444148
package main
 
import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
 
	"golang.org/x/net/websocket"  //可以和httprouter一起配合
	"github.com/julienschmidt/httprouter"	
)
 
func upper(ws *websocket.Conn) {
	var err error
	for {
		var reply string
 
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println(err)
			continue
		}
 
		if err = websocket.Message.Send(ws, strings.ToUpper(reply)); err != nil {
			fmt.Println(err)
			continue
		}
	}
}
 
func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}
 
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}
 
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
    router := httprouter.New()
    router.GET("/hello/:name", Hello)
    //http.HandleFunc("/", index)	//一样不管用了, 需要改成下面这个
    router.HandlerFunc("GET", "/", index)
    //http.Handle("/upper", websocket.Handler(upper))  //ListenAndServe加上router这个后不管用了, 需要改成下面这个
    router.Handler("GET", "/upper", websocket.Handler(upper))
 
	if err := http.ListenAndServe(":9999", router); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


package alive

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const Note = "smoke_test"

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/json")
	resPing := `{code:0,msg:"alive"}`
	_, _ = fmt.Fprintf(w, resPing) //这个写入到w的是输出到客户端的
}

// RunServer
//runSrv := "127.0.0.1:8080" //本地 MAC调试不会弹窗，但只能本机调用
//":8080","0.0.0.0:8080"可以对外
func RunServer(addr string) {
	http.HandleFunc("/api/ping", pingHandler)
	http.HandleFunc("/api/smoke", pingHandler)
	//设置访问的路由

	var port = ""
	if strings.Contains(addr, ":") {
		port = strings.Split(addr, ":")[1]
	}

	log.Println(fmt.Sprintf("http://127.0.0.1:%s/api/ping\n", port))
	err := http.ListenAndServe(addr, nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

package alive

import (
	"fmt"
	"log"
	"net/http"
)

var Settings = struct {
	Port int
}{}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/json")
	resPing := `{code:0,msg:"alive"}`
	_, _ = fmt.Fprintf(w, resPing) //这个写入到w的是输出到客户端的
}

func SetPort(port int) {
	Settings.Port = port
}
func PingPort() int {
	if Settings.Port != 0 {
		return Settings.Port
	}
	return 8080
}

// RunPingServer 服务的响应,探活机制
func RunPingServer() {
	port := PingPort()
	http.HandleFunc("/api/ping", pingHandler)
	http.HandleFunc("/api/smoke", pingHandler)
	//设置访问的路由
	log.Println(fmt.Sprintf("http://127.0.0.1:%d/api/ping\n", port))
	//runSrv := "127.0.0.1:8080" //本地 MAC调试不会弹窗
	err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", port), nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

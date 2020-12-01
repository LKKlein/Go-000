package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"week02/dao"
	"week02/service"
)

func searchUser(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("user_id")
	user, err := service.FindUserById(id)
	if err != nil {
		if errors.Is(err, dao.ErrorNotFound) {
			// 对ErrorNotFound错误进行处理，返回特定消息
			w.Write([]byte(fmt.Sprintf("{\"code\": 404, \"message\": \"user: %v not found\"}", id)))
		} else {
			// 打印错误堆栈信息，并向外抛出内部错误
			log.Printf("%+v\n", err)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		}
	}
	// 没有错误，返回正确结果
	result := fmt.Sprintf("{\"id\": \"%v\", \"name\": \"%v\", \"age\": %v}", id, user.Name, user.Age)
	w.Write([]byte(fmt.Sprintf("{\"code\": 0, \"result\": %v}", result)))
}

func main() {
	http.HandleFunc("/search", searchUser)
	err := http.ListenAndServe("127.0.0.1:18000", nil)
	if err != nil {
		log.Println("server stop.")
	}
}

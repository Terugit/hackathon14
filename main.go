package main

import (
	"back/controller"
	"back/dao"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

//func handler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Access-Control-Allow-Headers", "*")
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
//	switch r.Method {
//	case http.MethodGet:
//		controller.SearchControl(w, r)
//		//controller.FromUserCont(w, r)
//		//controller.ToUserCont(w, r)
//	//case http.MethodPost:
//	//controller.RegisterControl(w, r)
//	case http.MethodOptions:
//		return
//	default:
//		log.Printf("fail: HTTP Method is %s\n", r.Method)
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//}
//
//func handlerCont(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Access-Control-Allow-Headers", "*")
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
//	switch r.Method {
//	case http.MethodGet:
//		controller.ShowAllCont(w, r)
//	case http.MethodPost:
//		controller.RegisterCont(w, r)
//	case http.MethodOptions:
//		return
//	default:
//		log.Printf("fail: HTTP Method is %s\n", r.Method)
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//}
//
//func handlerFromCont(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Access-Control-Allow-Headers", "*")
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
//	switch r.Method {
//	case http.MethodGet:
//		controller.FromUserCont(w, r)
//	//case http.MethodPost:
//	//	controller.RegisterCont(w, r)
//	case http.MethodOptions:
//		return
//	default:
//		log.Printf("fail: HTTP Method is %s\n", r.Method)
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//}
//
//func handlerToCont(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Access-Control-Allow-Headers", "*")
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
//	switch r.Method {
//	case http.MethodGet:
//		controller.ToUserCont(w, r)
//	//case http.MethodPost:
//	//	controller.RegisterCont(w, r)
//	case http.MethodOptions:
//		return
//	default:
//		log.Printf("fail: HTTP Method is %s\n", r.Method)
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//}
//
//func handlerContEdit(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Access-Control-Allow-Headers", "*")
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
//	switch r.Method {
//	//case http.MethodGet:
//	//controller.EditCont(w, r)
//	case http.MethodPost:
//		controller.EditCont(w, r)
//	case http.MethodOptions:
//		return
//	default:
//		log.Printf("fail: HTTP Method is %s\n", r.Method)
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//}
//
//func handlerContDelete(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Access-Control-Allow-Headers", "*")
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
//	switch r.Method {
//	//case http.MethodGet:
//	//controller.EditCont(w, r)
//	case http.MethodPost:
//		controller.DeleteCont(w, r)
//	case http.MethodOptions:
//		return
//	default:
//		log.Printf("fail: HTTP Method is %s\n", r.Method)
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//}
//func handlerRanking(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Access-Control-Allow-Headers", "*")
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
//	switch r.Method {
//	case http.MethodGet:
//		controller.Ranking(w, r)
//	//case http.MethodPost:
//	//	controller.RegisterCont(w, r)
//	case http.MethodOptions:
//		return
//	default:
//		log.Printf("fail: HTTP Method is %s\n", r.Method)
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//}

func handlerEditMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	switch r.Method {
	case http.MethodPost:
		controller.UserEdit(w, r)
	case http.MethodOptions:
		return
	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
func handlerRegisterMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	switch r.Method {
	case http.MethodPost:
		controller.UserAdd(w, r)
	case http.MethodOptions:
		return
	default:
		log.Printf("fail: HTTP Method not allowed: %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func main() {
	log.Println("main start")
	//http.HandleFunc("/user", handler)
	//http.HandleFunc("/home", handlerCont)
	//http.HandleFunc("/fromcont", handlerFromCont)
	//http.HandleFunc("/tocont", handlerToCont)
	//http.HandleFunc("/edit", handlerContEdit)
	//http.HandleFunc("/delete", handlerContDelete)
	//http.HandleFunc("/ranking", handlerRanking)
	http.HandleFunc("/edit/user", handlerEditMember)
	http.HandleFunc("/add/user", handlerRegisterMember)
	closeDBWithSysCall()
	log.Println("Listening...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
func closeDBWithSysCall() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)

		if err := dao.Db.Close(); err != nil {
			log.Fatal(err)
		}
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}

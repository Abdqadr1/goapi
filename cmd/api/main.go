package main

import (
	"fmt"
	"github.com/abdqadr1/goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {

	log.SetReportCaller(true)

	var r *chi.Mux = chi.NewRouter()

	handlers.Handle(r)

	fmt.Println("Starting GO API services...")

	fmt.Println(`
 ______     ______        ______     ______   __    
/\  ___\   /\  __ \      /\  __ \   /\  == \ /\ \   
\ \ \__ \  \ \ \/\ \     \ \  __ \  \ \  _-/ \ \ \  
 \ \_____\  \ \_____\     \ \_\ \_\  \ \_\    \ \_\ 
  \/_____/   \/_____/      \/_/\/_/   \/_/     \/_/ `)

	err := http.ListenAndServe("localhost:8000", r)

	if err != nil {
		log.Error(err)
	}

}

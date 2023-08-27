package routes

import (
	"github.com/Ezinnem/workers/pkg/controllers"
	"github.com/gorilla/mux"
)

var MembersRoutes = func(router *mux.Router) {
	router.HandleFunc("/member", controllers.CreateMember).Methods("POST")
	router.HandleFunc("/members", controllers.GetAllMembers).Methods("GET")
	router.HandleFunc("/member/{memberId}", controllers.GetMemberById).Methods("GET")
	router.HandleFunc("/member/{memberId}", controllers.UpdateMember).Methods("PUT")
	router.HandleFunc("/member/{memberId}", controllers.DeleteMember).Methods("DELETE")
}

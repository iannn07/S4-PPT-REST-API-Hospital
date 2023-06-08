package main
// Caca
import (
<<<<<<< HEAD
	"github.com/gin-gonic/gin"
	"github.com/iann07/S4-PPT-REST-API-Hospital.git/handler/diagnosehandler"
	"github.com/iann07/S4-PPT-REST-API-Hospital.git/handler/doctorhandler"
	"github.com/iann07/S4-PPT-REST-API-Hospital.git/handler/patienthandler"
	"github.com/iann07/S4-PPT-REST-API-Hospital.git/handler/paymenthandler"
	"github.com/iann07/S4-PPT-REST-API-Hospital.git/handler/roomhandler"
	"github.com/iann07/S4-PPT-REST-API-Hospital.git/hospital"
)

func main() {
	router := gin.Default()
	hospital.ConnectDB()

	router.GET("api/hospital", diagnosehandler.SelectAll)
	router.POST("api/hospital", diagnosehandler.Create)
	router.GET("api/hospital/:id", diagnosehandler.Read)
	router.PUT("api/hospital/:id", diagnosehandler.SelectAll)
	router.DELETE("api/hospital/:id", diagnosehandler.Delete)

	router.GET("api/hospital", doctorhandler.SelectAll)
	router.POST("api/hospital", doctorhandler.Create)
	router.GET("api/hospital/:id", doctorhandler.Read)
	router.PUT("api/hospital/:id", doctorhandler.SelectAll)
	router.DELETE("api/hospital/:id", doctorhandler.Delete)

	router.GET("api/hospital", patienthandler.SelectAll)
	router.POST("api/hospital", patienthandler.Create)
	router.GET("api/hospital/:id", patienthandler.Read)
	router.PUT("api/hospital/:id", patienthandler.SelectAll)
	router.DELETE("api/hospital/:id", patienthandler.Delete)

	router.GET("api/hospital", paymenthandler.SelectAll)
	router.POST("api/hospital", paymenthandler.Create)
	router.GET("api/hospital/:id", paymenthandler.Read)
	router.PUT("api/hospital/:id", paymenthandler.SelectAll)
	router.DELETE("api/hospital/:id", paymenthandler.Delete)

	router.GET("api/hospital", roomhandler.SelectAll)
	router.POST("api/hospital", roomhandler.Create)
	router.GET("api/hospital/:id", roomhandler.Read)
	router.PUT("api/hospital/:id", roomhandler.SelectAll)
	router.DELETE("api/hospital/:id", roomhandler.Delete)

	router.Run("localhost:8080")
=======
	"encoding/json"
	"io"
	"net/http"
	us "user"
)

func cekError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func selectAll(w http.ResponseWriter, r *http.Request) {
	usr := us.SelectAll()
	data, err := json.Marshal(usr)
	cekError(err)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(data))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/tampilUserAll", selectAll)

	http.ListenAndServe(":5050", mux)
>>>>>>> 4ddc1a594af2da5174929ba49b82ad8f052bcf82
}

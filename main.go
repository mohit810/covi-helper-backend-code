package main

import (
	"covi-helper-backend-code/routes"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"net"
	"net/http"
	"os"
)

func main() {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			fmt.Println("IPv4: ", ipv4)
		}
	}
	r := httprouter.New()
	uc := routes.NewUserController(getSession())
	r.GET("/location", uc.Location)
	r.GET("/states", uc.GetStates)
	r.GET("/resource/:stateName/:cityName/:cityCode", uc.GetResources)
	r.GET("/verify/:emailId", uc.VerifyResources)
	r.POST("/registerVolunteer", uc.RegisterVolunteer)
	r.POST("/addoxygen", uc.AddOxygen)
	r.POST("/addbed", uc.AddBed)
	r.POST("/addplasma", uc.AddPlasma)
	r.POST("/addmedicine", uc.AddMedicine)
	r.POST("/addambulance", uc.AddAmbulance)
	r.POST("/addhelping", uc.AddHelpingHand)
	r.POST("/verifyoxygen", uc.VerifyOxygenLead)
	r.POST("/verifybed", uc.VerifyBedLead)
	r.POST("/verifyplasma", uc.VerifyPlasmaLead)
	r.POST("/verifymedicine", uc.VerifyMedicineLead)
	r.POST("/verifyambulance", uc.VerifyAmbulanceLead)
	r.POST("/verifyhelping", uc.VerifyHelpingHandLead)
	r.DELETE("/removeoxygen/:id", uc.RemoveOxygenLead)
	r.DELETE("/removebed/:id", uc.RemoveBedLead)
	r.DELETE("/removeplasma/:id", uc.RemovePlasmaLead)
	r.DELETE("/removemedicine/:id", uc.RemoveMedicineLead)
	r.DELETE("/removeambulance/:id", uc.RemoveAmbulanceLead)
	r.DELETE("/removehelping/:id", uc.RemoveHelpingHandLead)
	r.GET("/getVolunteers/:emailId", uc.GetVolunteers)
	r.GET("/getResources/:emailId", uc.GetAllResources)
	r.POST("/verifyVolunteer/:emailId", uc.VerifyVolunteer)
	r.POST("/removeVolunteer/:emailId", uc.RemoveVolunteer)
	handler := cors.AllowAll().Handler(r)
	http.ListenAndServe(":8080", handler)
}

func getSession() *sql.DB {
	//Supporting Lib for supporting .env file
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	// Initialize connection string.
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", os.Getenv("DB_USER"), os.Getenv("PASSWORD"), os.Getenv("HOST"), os.Getenv("DATABASE"))

	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	return db
}

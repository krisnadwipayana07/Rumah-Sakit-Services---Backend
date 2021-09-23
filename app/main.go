package main

import (
	"backend/app/routes"
	_doctorUsecase "backend/business/doctors"
	_doctorControllers "backend/controllers/doctors"
	_doctordb "backend/drivers/databases/doctors"
	_mysqlDriver "backend/drivers/mysql"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("This Service RUN on DEBUG Mode")
	}
}

func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(&_doctordb.Doctors{})
}

func main() {
	configDB := _mysqlDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	Conn := configDB.InitialDB()
	DBMigrate(Conn)

	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	doctorRepository := _doctordb.NewMysqlDoctorRepository(Conn)
	doctorUseCase := _doctorUsecase.NewDoctorUsecase(doctorRepository, timeoutContext)
	DoctorController := _doctorControllers.NewDoctorController(doctorUseCase)

	routesInit := routes.ControllerList{
		DoctorController: *DoctorController,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}

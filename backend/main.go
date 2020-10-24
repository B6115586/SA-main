package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/thanawat/app/controllers"
	"github.com/thanawat/app/ent"
	"github.com/thanawat/app/ent/user"
)

type Users struct {
	User []User
}

type User struct {
	Name  string
	Email string
}

type Rooms struct {
	Room []Room
}

type Room struct {
	User int
}

type Problemtitles struct {
	Problemtitle []Problemtitle
}

type Problemtitle struct {
	Problemtitle string
}

type Problemstatuss struct {
	Problemstatus []Problemstatus
}

type Problemstatus struct {
	Problemstatus string
}

// @title SUT SA Example API Playlist Vidoe
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewUserController(v1, client)
	controllers.NewRoomController(v1, client)
	controllers.NewProblemTitleController(v1, client)
	controllers.NewProblemStatusController(v1, client)
	controllers.NewProblemController(v1, client)

	// Set Users Data
	users := Users{
		User: []User{
			User{"Thanawat Srikaewsiew", "hin.thanawat@gmail.com"},
			User{"Kunsusevit Nomoney", "hardWorker@hotmail.com"},
			User{"Name Surname", "me@example.com"},
		},
	}

	for _, u := range users.User {
		client.User.
			Create().
			SetEmail(u.Email).
			SetName(u.Name).
			Save(context.Background())
	}

	// Set Rooms Data
	rooms := Rooms{
		Room: []Room{
			Room{2},
			Room{3},
		},
	}

	for _, ro := range rooms.Room {

		u, err := client.User.
			Query().
			Where(user.IDEQ(int(ro.User))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.Room.
			Create().
			SetUser(u).
			Save(context.Background())
	}

	// Set Problemtitle Data
	problemtitles := Problemtitles{
		Problemtitle: []Problemtitle{
			Problemtitle{"ปัญหาไฟฟ้าภายในห้อง"},
			Problemtitle{"ปัญหาประปาภายในห้อง"},
			Problemtitle{"ปัญหาอินเทอร์เน็ต"},
			Problemtitle{"ปัญหาอุปกรณ์เเละเครื่องใช้ชำรุด"},
			Problemtitle{"ปัญหาตัวห้องพักชำรุด"},
			Problemtitle{"ปัญหาอื่นๆ"},
		},
	}

	for _, pt := range problemtitles.Problemtitle {

		client.ProblemTitle.
			Create().
			SetProblemtitle(pt.Problemtitle).
			Save(context.Background())
	}

	// Set Problemstatus Data
	problemstatuss := Problemstatuss{
		Problemstatus: []Problemstatus{
			Problemstatus{"ข้อร้องทุกข์ได้ทำการส่งเสร็จสิ้น"},
			Problemstatus{"ข้อร้องทุกข์ไม่อนุมัติ"},
			Problemstatus{"ข้อร้องทุกข์อยู่ระหว่างการดำเนินงาน"},
			Problemstatus{"ข้อร้องทุกข์ดำเนินการเสร็จสิ้น"},
		},
	}

	for _, ps := range problemstatuss.Problemstatus {

		client.ProblemStatus.
			Create().
			SetProblemstatus(ps.Problemstatus).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}

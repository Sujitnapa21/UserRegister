package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/Sujitnapa21/app/controllers"
	"github.com/Sujitnapa21/app/ent"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Titles  defines the struct for the titles
type Titles struct {
	Title []Title
}

// Title  defines the struct for the Title
type Title struct {
	Titlename string
}

// Sexs  defines the struct for the sexs
type Sexs struct {
	Sex []Sex
}

// Sex  defines the struct for the Sex
type Sex struct {
	Sexname string
}

// Roles  defines the struct for the roles
type Roles struct {
	Role []Role
}

// Role  defines the struct for the role
type Role struct {
	Rolename string
}

// @title SUT SA Example API
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

	client, err := ent.Open("sqlite3", "file:Users.db?&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewUserController(v1, client)
	controllers.NewTitleController(v1, client)
	controllers.NewSexController(v1, client)
	controllers.NewRoleController(v1, client)

	// Set Titles Data
	titles := Titles{
		Title: []Title{
			Title{"นาย"},
			Title{"นาง"},
			Title{"นางสาว"},
		},
	}

	for _, t := range titles.Title {
		client.Title.
			Create().
			SetTitlename(t.Titlename).
			Save(context.Background())
	}

	// Set Sexs Data
	sexs := Sexs{
		Sex: []Sex{
			Sex{"ชาย"},
			Sex{"หญิง"},
		},
	}

	for _, s := range sexs.Sex {
		client.Sex.
			Create().
			SetSexname(s.Sexname).
			Save(context.Background())
	}

	// Set Roles Data
	roles := Roles{
		Role: []Role{
			Role{"รายชั่วโมง"},
			Role{"รายวัน"},
		},
	}

	for _, r := range roles.Role {
		client.Role.
			Create().
			SetRolename(r.Rolename).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	_ "github.com/Calmantara/go-kominfo-2024/ch8/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			GO DTS USER API DUCUMENTATION
// @version		2.0
// @description	golong kominfo 006 api documentation
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:3000
// @BasePath		/
// @schemes		http
func main() {
	// jsonEncodingDecoding()
	// urlParsing()
	swaggerExample()
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserPlaceholder struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	} `json:"company"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func jsonEncodingDecoding() {
	// apa itu json?
	// format data / data structure
	// json format data yang biasa digunakan untuk pertukaran data

	// key value pair
	// mobile -> server
	// web -> server
	jsonBody := `
		{
			"id":1,
			"name":"golang",
			"email":"golang@gmail.com"
		}
	`
	// format data untuk pertukaran data
	// selain json
	// - xml
	// - yaml
	// - protobuf

	// bisa kita translate ke struct
	// Json Unmarshal adalah translate dari json string / []bytes
	// ke dalam suatu struct golang
	user := User{}
	err := json.Unmarshal([]byte(jsonBody), &user)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

	// translate dari struct / map
	// ke dalam json string / []byte
	b, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(b, string(b))

	// array of struct
	jsonArr := `
		[
			{
				"id":1,
				"name":"golang1",
				"email":"golang1@gmail.com"
			},
			{
				"id":2,
				"name":"golang2",
				"email":"golang2@gmail.com"
			}
		]
	`
	userArr := []User{}
	err = json.Unmarshal([]byte(jsonArr), &userArr)
	if err != nil {
		panic(err)
	}
	fmt.Println(userArr)

	jsonPlaceholder := `
	{
		"id": 10,
		"name": "Clementina DuBuque",
		"username": "Moriah.Stanton",
		"email": "Rey.Padberg@karina.biz",
		"address": {
		  "street": "Kattie Turnpike",
		  "suite": "Suite 198",
		  "city": "Lebsackbury",
		  "zipcode": "31428-2261",
		  "geo": {
			"lat": "-38.2386",
			"lng": "57.2232"
		  }
		},
		"phone": "024-648-3804",
		"website": "ambrose.net",
		"company": {
		  "name": "Hoeger LLC",
		  "catchPhrase": "Centralized empowering task-force",
		  "bs": "target end-to-end models"
		}
	  }`
	userPlaceholder := map[string]interface{}{}
	err = json.Unmarshal([]byte(jsonPlaceholder), &userPlaceholder)
	if err != nil {
		panic(err)
	}
	fmt.Println(userPlaceholder)
}

func urlParsing() []UserPlaceholder {
	var client = &http.Client{}

	urlAddress := "https://jsonplaceholder.typicode.com/users"
	u, err := url.Parse(urlAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println(u)

	request, err := http.NewRequest("GET", urlAddress, nil)
	if err != nil {
		panic(err)
	}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	users := []UserPlaceholder{}
	err = json.NewDecoder(response.Body).Decode(&users)
	if err != nil {
		panic(err)
	}
	fmt.Println(users)
	return users
}

func swaggerExample() {
	// doc dari api kita
	// yang mengikuti standard OpenAPI
	// go get -u github.com/swaggo/swag/cmd/swaggo
	// go get -u github.com/swaggo/swag/http-swaggergo
	// go get -u github.com/alecthomas/template
	g := gin.Default()

	g.GET("/users", getUsers)
	g.GET("/users/:id", getUsersId)

	// swagger
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g.Run(":3000")
}

// ShowUsers godoc
//
//	@Summary		Show users list
//	@Description	will fetch 3rd party server to get users data
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]UserPlaceholder
//	@Failure		400	{object}	ErrorResponse
//	@Failure		404	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/users [get]
func getUsers(ctx *gin.Context) {
	users := urlParsing()
	if len(users) <= 0 {
		ctx.JSON(http.StatusNotFound, ErrorResponse{Message: "user not found"})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// ShowUsersById godoc
//
//	@Summary		Show users detail
//	@Description	will fetch 3rd party server to get users data to get detail user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	UserPlaceholder
//	@Failure		400	{object}	ErrorResponse
//	@Failure		404	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/users/{id} [get]
func getUsersId(ctx *gin.Context) {
	client := &http.Client{}

	// get id user
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid required param"})
		return
	}
	fmt.Println(id)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: "cannot convert param"})
		return
	}

	urlAddress := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", idInt)
	u, err := url.Parse(urlAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println(u)

	request, err := http.NewRequest("GET", urlAddress, nil)
	if err != nil {
		panic(err)
	}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	user := UserPlaceholder{}
	err = json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, user)
}

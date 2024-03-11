package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// HTTP Request from Golang
type Post struct {
	ID     uint64 `json:"id"`
	UserID uint64 `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

const (
	STATIC_USERNAME = "golang006awesome"
	STATIC_PASSWORD = "mysecretpassword"
)

func getAllPosts() (data []Post, err error) {
	// get response
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Println("error when fetching posts data:", err)
		return
	}

	// check status code
	if res.StatusCode != http.StatusOK {
		log.Println("got not ok response:", res.StatusCode)
		return
	}

	// fetch body from response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("error fetching body from response:", err)
		return
	}
	defer res.Body.Close()

	// unmarshal body
	json.Unmarshal(body, &data)
	return
}

func createPosts(body Post) (err error) {
	// marshall body menjadi json
	bodyJson, err := json.Marshal(body)
	if err != nil {
		log.Println("error marshall body to json")
		return
	}
	// generate client request
	client := &http.Client{}
	res, err := client.Post(
		"https://jsonplaceholder.typicode.com/posts",
		"application/json",
		bytes.NewBuffer(bodyJson))
	if err != nil {
		log.Println("error performing post request")
		return
	}
	// check status code
	if res.StatusCode >= 400 {
		err = errors.New("something went wrong with request")
		log.Println("got error status code", res.StatusCode)
		return
	}

	// fetch body from response
	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("error fetching body from response:", err)
		return
	}
	defer res.Body.Close()
	log.Println(string(respBody))
	return
}

func updatePosts(id uint64, body Post) (err error) {
	// marshall body menjadi json
	bodyJson, err := json.Marshal(body)
	if err != nil {
		log.Println("error marshall body to json")
		return
	}
	// generate client request
	client := &http.Client{}

	req, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%v", id),
		bytes.NewBuffer(bodyJson))
	if err != nil {
		log.Println("error performing post request")
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// do request
	res, err := client.Do(req)
	if err != nil {
		log.Println("error performing post request")
		return
	}
	defer res.Body.Close()

	// check status code
	if res.StatusCode >= 400 {
		err = errors.New("something went wrong with request")
		log.Println("got error status code", res.StatusCode)
		return
	}

	// fetch body from response
	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("error fetching body from response:", err)
		return
	}
	log.Println(string(respBody))
	return
}

func deletePost(id uint64) (err error) {
	// generate client request
	client := &http.Client{}

	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%v", id),
		nil)
	if err != nil {
		log.Println("error performing post request")
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// do request
	res, err := client.Do(req)
	if err != nil {
		log.Println("error performing post request")
		return
	}
	defer res.Body.Close()

	// check status code
	if res.StatusCode >= 400 {
		err = errors.New("something went wrong with request")
		log.Println("got error status code", res.StatusCode)
		return
	}

	// fetch body from response
	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("error fetching body from response:", err)
		return
	}
	log.Println(string(respBody))
	return
}

type Response struct {
	Message string   `json:"message"`
	Data    any      `json:"data"`
	Errors  []string `json:"errors"`
}

func main() {
	g := gin.Default()

	v1 := g.Group("/api/v1", checkAuthBasic)

	posts := v1.Group("/posts")
	{
		posts.GET("", func(ctx *gin.Context) {
			data, err := getAllPosts()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError,
					Response{
						Message: "error",
						Errors:  []string{err.Error()},
					})
				return
			}
			ctx.JSON(http.StatusOK, Response{
				Message: "ok",
				Data:    data,
			})
		})

		posts.PUT("/:id", func(ctx *gin.Context) {
			// get id
			id, err := strconv.Atoi(ctx.Param("id"))
			if err != nil || id <= 0 {
				ctx.JSON(http.StatusBadRequest,
					Response{
						Message: "error",
						Errors:  []string{"invalid id"},
					})
				return
			}
			// bind body
			body := Post{}
			if err := ctx.Bind(&body); err != nil {
				ctx.JSON(http.StatusBadRequest,
					Response{
						Message: "error",
						Errors:  []string{"invalid body request"},
					})
				return
			}

			err = updatePosts(uint64(id), body)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError,
					Response{
						Message: "error",
						Errors:  []string{err.Error()},
					})
				return
			}
			ctx.JSON(http.StatusOK, Response{
				Message: "ok",
				Data:    nil,
			})
		})

		posts.DELETE("/:id", func(ctx *gin.Context) {
			// get id
			id, err := strconv.Atoi(ctx.Param("id"))
			if err != nil || id <= 0 {
				ctx.JSON(http.StatusBadRequest,
					Response{
						Message: "error",
						Errors:  []string{"invalid id"},
					})
				return
			}
			err = deletePost(uint64(id))
			if err != nil {
				ctx.JSON(http.StatusInternalServerError,
					Response{
						Message: "error",
						Errors:  []string{err.Error()},
					})
				return
			}
			ctx.JSON(http.StatusOK, Response{
				Message: "ok",
				Data:    nil,
			})
		})
	}
	g.Run(":3000")
}

func checkAuthBasic(ctx *gin.Context) {
	// check authorization request
	// step1: ambil data auth dari header
	auth := ctx.GetHeader("Authorization")
	// "Basic Z29sYW5nMDA2YXdlc29tZTpteXNlY3JldHBhc3N3b3Jk"
	// step2: dapatkan base64 string
	authArr := strings.Split(auth, " ")
	if len(authArr) < 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, Response{
			Message: "unauthorized",
			Errors:  []string{"invalid token"},
		})
		return
	}
	if authArr[0] != "Basic" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, Response{
			Message: "unauthorized",
			Errors:  []string{"invalid authorization method"},
		})
		return
	}
	// step3: decode base64 string
	token := authArr[1]
	basic, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, Response{
			Message: "unauthorized",
			Errors:  []string{"invalid token", "failed to decode"},
		})
		return
	}
	// step4: compare dengan variable static
	if string(basic) != fmt.Sprintf("%v:%v", STATIC_USERNAME, STATIC_PASSWORD) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, Response{
			Message: "unauthorized",
			Errors:  []string{"invalid username or password"},
		})
		return
	}
	ctx.Next()
}

func perform() {
	// data, err := getAllPosts()
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println(data)

	createPosts(Post{
		UserID: 1,
		Title:  "golang 006 awesome",
		Body:   "you guys special",
	})

	// Challenge:
	// buat http client
	// untuk PUT
	// dan DELETE
	// jsonplaceholder untuk Posts
}

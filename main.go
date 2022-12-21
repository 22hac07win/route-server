package main

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"github.com/22hac07win/route-server.git/handler"
	"github.com/22hac07win/route-server.git/repository"
	"github.com/22hac07win/route-server.git/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"strings"
)

func InitFirebase() *firebase.App {

	var opt option.ClientOption
	ctx := context.Background()

	/*
		config := os.Getenv("FIREBASE_CONFIG")

			if config != "" {
				credentials, err := google.CredentialsFromJSON(ctx, []byte(config))
				if err != nil {
					log.Printf("error credentials from json: %v\n", err)
				}
				opt = option.WithCredentials(credentials)
			} else {
	*/
	opt = option.WithCredentialsFile("routeAccountKey.json")
	// }

	app, err := firebase.NewApp(ctx, nil, opt)

	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	return app
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		app := InitFirebase()

		client, err := app.Auth(c.Request.Context())

		// リクエストヘッダーからIDトークンを取得
		authHeader := c.Request.Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		log.Println(idToken)

		if idToken == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// IDトークンを使ってユーザーを取得
		token, err := client.VerifyIDToken(c.Request.Context(), idToken)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		log.Println(token.UID)

		// ユーザーをコンテキストに保存
		c.Set("userID", token.UID)
	}
}

func main() {
	r := gin.Default()

	// ここからCorsの設定
	r.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://localhost:3001",
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
	}))

	s := repository.NewSupabaseDBClient()
	fbs := service.NewBlockFuncService(s)
	rp := service.NewRouteProvider(s, fbs)
	rh := handler.NewRouteHandler(rp)

	r.GET("/ping", func(c *gin.Context) {
		log.Println("pass")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/init", func(c *gin.Context) {
		c.Set("userID", "test-user1")
		rh.GetInit(c)
	})
	r.POST("/message", func(c *gin.Context) {
		c.Set("userID", "test-user1")
		rh.PostMessage(c)
	})

	api := r.Group("/api")
	api.Use(authMiddleware())
	{
		api.GET("/ping", func(c *gin.Context) {

			message := fmt.Sprintf("Your userID is %s", c.GetString("userID"))
			c.JSON(http.StatusOK, gin.H{
				"message": message,
			})
		})

		api.GET("/init", func(c *gin.Context) { rh.GetInit(c) })
		api.POST("/message", func(c *gin.Context) { rh.PostMessage(c) })
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

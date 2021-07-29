package infra

import (
	"net/http"
	"time"

	"d-meeting.com/d-server/interfaces/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter() *Router {
	router := &Router{
		Engine: gin.Default(),
	}
	handler := NewEnvHandler()
	url, _ := handler.ReadEnv("CORS_URL")
	sqlHandler := NewSqlHandler()

	router.Engine.Use(cors.New(cors.Config{
		// 許可したいHTTPメソッドの一覧
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		// 許可したいHTTPリクエストヘッダの一覧
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		// 許可したいアクセス元の一覧
		AllowOrigins: []string{
			url,
		},
		MaxAge: 24 * time.Hour,
	}))

	//コントローラー単位でルーティング関数を定義
	//コントローラー作成
	taskController := controller.NewTaskController(sqlHandler)

	//ルーティングテーブルの初期化
	router.initTaskRoute(*taskController)

	return router
}

func (router *Router) initTaskRoute(taskController controller.TaskController) {
	taskRoute := router.Engine.Group("/task")
	{
		taskRoute.POST("/create", func(c *gin.Context) {
			taskController.Create(c)
		})

		testTaskRoute := taskRoute.Group("/test")
		{
			testTaskRoute.GET("/create", func(c *gin.Context) {
				taskController.TestCreate(c)
			})
			testTaskRoute.GET("/read", func(c *gin.Context) {
				taskController.Read(c)
			})
		}
	}
}

func (router *Router) initTestRoute() {
	router.Engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
}

func (router *Router) Run() {
	router.Engine.Run()
}

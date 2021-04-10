package main

import (
	"os"
	"github.com/selvakraj/go-redis-cache/controller"
	router "github.com/selvakraj/go-redis-cache/http"
)

var(
	postController controller.PostController = controller.NewPostController()
	httpRouter     router.Router             = router.NewChiRouter()
)

func main()  {
httpRouter.GET("/posts", postController.GetPosts)
httpRouter.GET("/posts/{id}", postController.GetPostByID)
httpRouter.POST("/posts", postController.AddPost)

httpRouter.SERVE(os.Getenv("PORT"))
}
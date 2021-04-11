package main

import (
	"os"
	"github.com/selvakraj/go-redis-cache/repository"
	"github.com/selvakraj/go-redis-cache/service"
	"github.com/selvakraj/go-redis-cache/controller"
	router "github.com/selvakraj/go-redis-cache/http"
)

var(
	postRepository repository.PostRepository = repository.NewSqliteRepository()
	postService    service.PostService       = service.NewPostService(postRepository)

	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewChiRouter()
)

func main()  {
httpRouter.GET("/posts", postController.GetPosts)
httpRouter.GET("/posts/{id}", postController.GetPostByID)
httpRouter.POST("/posts", postController.AddPost)

httpRouter.SERVE(os.Getenv("PORT"))
}
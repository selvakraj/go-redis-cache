package controller

import(
	"encoding/json"
	"net/http"

	"github.com/selvakraj/go-redis-cache/entity"
	"github.com/selvakraj/go-redis-cache/service"
)
type controller struct{}

var (
	postService service.PostService
)


type PostController interface{
	GetPostByID(response http.ResponseWriter, request *http.Request)
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPost(response http.ResponseWriter, request *http.Request)
}

func NewPostController(service service.PostService) PostController{
	postService = service
	return &controller{}
}

func (*controller) GetPostByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

}
func (*controller) GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)

}
func (*controller) AddPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var post entity.Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)	
		return
	}

	err1 := postService.Validate(&post)
	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)		
		return
	}

	result, err2 := postService.Create(&post)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)		
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)


}
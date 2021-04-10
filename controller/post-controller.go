package controller

import(
	"net/http"
)
type controller struct{}

type PostController interface{
	GetPostByID(response http.ResponseWriter, request *http.Request)
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPost(response http.ResponseWriter, request *http.Request)
}

func NewPostController() PostController{
	return &controller{}
}

func (*controller) GetPostByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

}
func (*controller) GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

}
func (*controller) AddPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

}
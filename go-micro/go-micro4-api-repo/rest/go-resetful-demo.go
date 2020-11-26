package main

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
)

type User struct {
	Id   string
	Name string
}

type UserResource struct {
	users map[string]User
}

func (u UserResource) Register(c *restful.Container) {
	// 声明一个关于user的webService实例
	webService := new(restful.WebService)
	webService.Path("/users") // the root URL template path of the WebService
	webService.Consumes(restful.MIME_XML, restful.MIME_JSON)
	webService.Produces(restful.MIME_JSON, restful.MIME_XML)

	// 新建路由
	webService.Route(webService.GET("/{user-id}").To(u.findUser))
	webService.Route(webService.POST("").To(u.updateUser))
	webService.Route(webService.PUT("/{user-id}").To(u.createUser))
	webService.Route(webService.DELETE("/{user-id}").To(u.removeUser))

	// 添加
	// 把webServer放到容器里面，可以理解成一个container里面由很多的webService
	// 每一个webService里面有预先定义Path路径
	// 在这里面是/users，在这个路径下面会有CRUD的方法，通过Route去确定方法匹配和跳转
	c.Add(webService)
}

/*
users的一套CURD
*/

// PUT http://localhost:8080/users/1
func (u UserResource) createUser(req *restful.Request, resp *restful.Response) {
	usr := User{
		Id: req.PathParameter("user-id"),
	}
	err := req.ReadEntity(&usr)
	if err == nil {
		u.users[usr.Id] = usr
		resp.WriteHeaderAndEntity(http.StatusCreated, err.Error())
	} else {
		resp.AddHeader("Content-Type", "text/plain")
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

// DELETE http://localhost:8080/users/1
func (u UserResource) removeUser(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("user-id")
	delete(u.users, id)
}

// POST http://localhost:8080/users
func (u UserResource) updateUser(req *restful.Request, resp *restful.Response) {
	usr := new(User)
	err := req.ReadEntity(&usr)
	if err == nil {
		u.users[usr.Id] = *usr
		resp.WriteEntity(usr)
	} else {
		resp.AddHeader("Content-Type", "text/plain")
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

// GET http://localhost:8080/users/1
func (u UserResource) findUser(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("user-id")
	usr, ok := u.users[id]
	if !ok {
		resp.AddHeader("Content-Type", "text/plain")
		resp.WriteErrorString(http.StatusNotFound, "User could not be found.")
	} else {
		resp.WriteEntity(usr)
	}
}

// func main() {
func main2() {
	// 首先定义一个web container的容器
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	u := UserResource{map[string]User{}}
	u.Register(wsContainer)

	log.Printf("start listening on localhost:8080")
	server := &http.Server{Addr: ":8080", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}

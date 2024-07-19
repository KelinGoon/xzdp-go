// Code generated by hertz generator. DO NOT EDIT.

package xzdp

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	xzdp "xzdp/biz/handler/xzdp"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.GET("/hello", append(_hellomethodMw(), xzdp.HelloMethod)...)
	root.GET("/me", append(_usermethodMw(), xzdp.UserMethod)...)
	{
		_user := root.Group("/user", _userMw()...)
		_user.GET("/:id", append(_userinfoMw(), xzdp.UserInfo)...)
		_user.POST("/login", append(_userloginMw(), xzdp.UserLogin)...)
	}
}

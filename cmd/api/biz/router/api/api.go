// Code generated by hertz generator. DO NOT EDIT.

package api

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	api "github.com/west2-online/fzuhelper-server/cmd/api/biz/handler/api"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		{
			_v1 := _api.Group("/v1", _v1Mw()...)
			{
				_common := _v1.Group("/common", _commonMw()...)
				{
					_classroom := _common.Group("/classroom", _classroomMw()...)
					_classroom.GET("/empty", append(_getemptyclassroomsMw(), api.GetEmptyClassrooms)...)
				}
			}
			{
				_jwch := _v1.Group("/jwch", _jwchMw()...)
				{
					_user := _jwch.Group("/user", _userMw()...)
					_user.GET("/login", append(_getlogindataMw(), api.GetLoginData)...)
				}
			}
		}
	}
}
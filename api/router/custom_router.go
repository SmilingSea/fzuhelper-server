/*
Copyright 2024 The west2-online Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by hertz generator.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"

	"github.com/west2-online/fzuhelper-server/api/handler"
	"github.com/west2-online/fzuhelper-server/api/handler/api"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	// 20241016兼容-2024肖垲
	/*
		旧的福uu客户端请求增加图片点击次数埋点，可以直接对接到重构后的后端
		r.GET("/api/image/point", api.AddImagePointTime)

		旧的福uu客户端通过学号设备请求开屏页，同样可以直接对接到重构后的后端，
		前端发送的结构体定义在飞书文档API.MobileGetImage
		飞书文档:https://west2-online.feishu.cn/wiki/YMtTwhwAOimxkIkeZfAcfAzgnle
		r.GET("/api/screen", api.MobileGetImage)
	*/
	r.GET("/api/image/point", api.AddImagePointTime)
	r.GET("/api/screen", api.MobileGetImage)

	// 历年卷兼容
	r.GET("/api/v1/list", api.ListDirFiles)
	r.GET("/api/v1/downloadUrl", api.GetDownloadUrl)
	// 登录验证码兼容
	r.POST("/api/login/validateCode", api.ValidateCode)
}

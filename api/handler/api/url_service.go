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

package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"

	"github.com/cloudwego/hertz/pkg/protocol/consts"

	api "github.com/west2-online/fzuhelper-server/api/model/api"
	"github.com/west2-online/fzuhelper-server/api/pack"
	"github.com/west2-online/fzuhelper-server/pkg/base"
	"github.com/west2-online/fzuhelper-server/pkg/errno"
	"github.com/west2-online/fzuhelper-server/pkg/logger"
)

var ClientSet *base.ClientSet

// APILogin .
// @router /api/v1/url/login [POST]
func APILogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.APILoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		logger.Errorf("api.APILogin: BindAndValidate error %v", err)
		pack.RespError(c, errno.ParamError.WithError(err))
		return
	}

	url := "http://127.0.0.1:5000/api/login"

	request := new(protocol.Request)
	request.SetMethod(consts.MethodPost)
	request.SetRequestURI(url)
	request.SetFormData(
		map[string]string{
			"password": req.Password,
		},
	)

	res := new(protocol.Response)

	if err = ClientSet.HzClient.Do(ctx, request, res); err != nil {
		pack.RespError(c, err)
		return
	}

	c.String(http.StatusOK, res.BodyBuffer().String())
}

// UploadVersionInfo .
// @router /api/v1/url/api/upload [POST]
func UploadVersionInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UploadVersionInfoRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		logger.Errorf("api.UploadVersionInfo: BindAndValidate error %v", err)
		pack.RespError(c, errno.ParamError.WithError(err))
		return
	}

	url := "http://127.0.0.1:5000/api/upload"

	request := new(protocol.Request)
	request.SetMethod(consts.MethodPost)
	request.SetRequestURI(url)
	request.SetFormData(
		map[string]string{
			"password": req.Password,
			"type":     req.Type,
			"version":  req.Version,
			"code":     req.Code,
			"feature":  req.Feature,
			"url":      req.URL,
		},
	)

	res := new(protocol.Response)

	if err = ClientSet.HzClient.Do(ctx, request, res); err != nil {
		pack.RespError(c, err)
		return
	}

	c.String(consts.StatusOK, res.BodyBuffer().String())
}

// GetUploadParams .
// @router /api/v1/url/api/uploadparams [POST]
func GetUploadParams(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetUploadParamsRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		logger.Errorf("api.GetUploadParams: BindAndValidate error %v", err)
		pack.RespError(c, errno.ParamError.WithError(err))
		return
	}

	url := "http://127.0.0.1:5000/api/uploadparams"
	resp := new(api.GetUploadParamsResponse)

	request := &protocol.Request{}
	request.SetMethod(consts.MethodPost)
	request.SetRequestURI(url)
	request.SetFormData(
		map[string]string{
			"password": req.Password,
		},
	)

	res := &protocol.Response{}

	if err = ClientSet.HzClient.Do(ctx, request, res); err != nil {
		pack.RespError(c, err)
		return
	}

	if res.BodyBuffer().String() == "illegal access" {
		c.String(consts.StatusOK, res.BodyBuffer().String())
		return
	}

	err = sonic.Unmarshal(res.BodyBytes(), resp)
	if err != nil {
		pack.RespError(c, err)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// GetDownloadRelease .
// @router /api/v1/url/release.apk [GET]
func GetDownloadRelease(ctx context.Context, c *app.RequestContext) {
	c.Redirect(http.StatusOK, []byte("http://127.0.0.1:5000/release.apk"))
}

// GetDownloadBeta .
// @router /api/v1/url/beta.apk [GET]
func GetDownloadBeta(ctx context.Context, c *app.RequestContext) {
	c.Redirect(http.StatusOK, []byte("http://127.0.0.1:5000/beta.apk"))
}

// GetReleaseVersion .
// @router /api/v1/url/version.json [GET]
func GetReleaseVersion(ctx context.Context, c *app.RequestContext) {
	url := "http://127.0.0.1:5000/version.json" // 和json无关，仅为一个路径

	request := &protocol.Request{}
	request.SetMethod(consts.MethodGet)
	request.SetRequestURI(url)

	res := new(protocol.Response)
	keyValue := make(map[string]interface{})

	if err := ClientSet.HzClient.Do(ctx, request, res); err != nil {
		pack.RespError(c, err)
		return
	}

	err := sonic.Unmarshal(res.BodyBytes(), &keyValue)
	if err != nil {
		pack.RespError(c, err)
		return
	}

	c.JSON(consts.StatusOK, keyValue)
}

// GetBetaVersion .
// @router /api/v1/url/versionbeta.json [GET]
func GetBetaVersion(ctx context.Context, c *app.RequestContext) {
	url := "http://127.0.0.1:5000/versionbeta.json" // 和json无关，仅为一个路径

	request := &protocol.Request{}
	request.SetMethod(consts.MethodGet)
	request.SetRequestURI(url)

	res := new(protocol.Response)
	keyValue := make(map[string]interface{})

	if err := ClientSet.HzClient.Do(ctx, request, res); err != nil {
		pack.RespError(c, err)
		return
	}

	err := sonic.Unmarshal(res.BodyBytes(), &keyValue)
	if err != nil {
		pack.RespError(c, err)
		return
	}

	c.JSON(consts.StatusOK, keyValue)
}

// GetCloudSetting .
// @router /api/v1/url/settings.php [GET]
func GetCloudSetting(ctx context.Context, c *app.RequestContext) {
	var err error

	account := c.DefaultQuery("account", "")
	verison := c.DefaultQuery("version", "")
	beta := c.DefaultQuery("beta", "false")
	phone := c.DefaultQuery("phone", "")
	isLogin := c.DefaultQuery("isLogin", "false")
	loginType := c.DefaultQuery("loginType", "0")

	url := "http://127.0.0.1:5000/settings.php" // 和php无关，仅为一个路径

	queryParams := strings.Join(
		[]string{
			"?account=" + account,
			"version=" + verison,
			"beta=" + beta,
			"phone=" + phone,
			"isLogin=" + isLogin,
			"loginType=" + loginType,
		}, "&",
	)

	request := protocol.Request{}
	request.SetMethod(consts.MethodGet)
	request.SetRequestURI(url + queryParams)

	res := new(protocol.Response)
	keyValue := make(map[string]interface{})

	if err = ClientSet.HzClient.Do(ctx, &request, res); err != nil {
		pack.RespError(c, err)
		return
	}

	if err = sonic.Unmarshal(res.BodyBytes(), &keyValue); err != nil {
		pack.RespError(c, err)
		return
	}

	c.JSON(consts.StatusOK, keyValue)
}

// GetAllCloudSetting .
// @router /api/v1/url/api/getcloud [GET]
func GetAllCloudSetting(ctx context.Context, c *app.RequestContext) {
	var err error

	url := "http://127.0.0.1:5000/api/getcloud"

	request := &protocol.Request{}
	request.SetMethod(consts.MethodGet)
	request.SetRequestURI(url)

	res := new(protocol.Response)
	keyValue := make(map[string]interface{})

	if err = ClientSet.HzClient.Do(ctx, request, res); err != nil {
		pack.RespError(c, err)
		return
	}

	err = sonic.Unmarshal(res.BodyBytes(), &keyValue)
	if err != nil {
		pack.RespError(c, err)
		return
	}

	c.JSON(consts.StatusOK, keyValue)
}

// SetAllCloudSetting .
// @router /api/v1/url/api/setcloud [POST]
func SetAllCloudSetting(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.SetAllCloudSettingRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		logger.Errorf("api.SetAllCloudSetting: BindAndValidate error %v", err)
		pack.RespError(c, errno.ParamError.WithError(err))
		return
	}

	url := "127.0.0.1:5000/api/setcloud"

	request := &protocol.Request{}
	request.SetMethod(consts.MethodPost)
	request.SetRequestURI(url)

	res := new(protocol.Response)

	if err = ClientSet.HzClient.Do(ctx, request, res); err != nil {
		pack.RespError(c, err)
		return
	}

	c.JSON(consts.StatusOK, res.BodyBuffer().String())
}

// TestSetting .
// @router /api/v1/url/api/test [POST]
func TestSetting(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.TestSettingRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		logger.Errorf("api.TestSetting: BindAndValidate error %v", err)
		pack.RespError(c, errno.ParamError.WithError(err))
		return
	}

	url := "http://127.0.0.1:5000/api/test"

	request := &protocol.Request{}
	request.SetMethod(consts.MethodPost)
	request.SetRequestURI(url)

	res := new(protocol.Response)

	if err = ClientSet.HzClient.Do(ctx, request, res); err != nil {
		pack.RespError(c, err)
		return
	}

	keyValue := make(map[string]interface{})

	err = sonic.Unmarshal(res.BodyBytes(), keyValue)
	if err != nil {
		pack.RespError(c, err)
		return
	}

	c.JSON(consts.StatusOK, keyValue)
}

// DumpVisit .
// @router /api/v1/url/dump [GET]
func DumpVisit(ctx context.Context, c *app.RequestContext) {
	var err error

	url := "http://127.0.0.1:5000/upupdowndownleftleftrightrightbaba_dump_visit"
	request := &protocol.Request{}
	request.SetMethod(consts.MethodGet)
	request.SetRequestURI(url)

	res := new(protocol.Response)
	keyValue := make(map[string]interface{})

	if err = ClientSet.HzClient.Do(ctx, request, res); err != nil {
		pack.RespError(c, err)
		return
	}

	err = sonic.Unmarshal(res.BodyBytes(), &keyValue)
	if err != nil {
		pack.RespError(c, err)
		return
	}

	c.JSON(consts.StatusOK, keyValue)
}

// FZUHelperCSS .
// @router /api/v1/url/onekey/FZUHelper.css [GET]
func FZUHelperCSS(ctx context.Context, c *app.RequestContext) {
	var err error

	url := "http://127.0.0.1:5000/onekey/FZUHelper.css" // 和html无关，仅为一个路径

	request := &protocol.Request{}
	request.SetMethod(consts.MethodGet)
	request.SetRequestURI(url)

	res := new(protocol.Response)
	if err = ClientSet.HzClient.Do(ctx, request, res); err != nil {
		pack.RespError(c, err)
		return
	}

	c.SetContentType("text/css")
	c.Data(consts.StatusOK, "text/css", res.BodyBytes())
}

// FZUHelperHTML .
// @router /api/v1/url/onekey/FZUHelper.html [GET]
func FZUHelperHTML(ctx context.Context, c *app.RequestContext) {
	var err error

	url := "http://127.0.0.1:5000/onekey/FZUHelper.html" // 和html无关，仅为一个路径

	request := &protocol.Request{}
	request.SetMethod(consts.MethodGet)
	request.SetRequestURI(url)

	res := new(protocol.Response)

	if err = ClientSet.HzClient.Do(ctx, request, res); err != nil {
		pack.RespError(c, err)
		return
	}
	c.Data(consts.StatusOK, "text/html", res.BodyBytes())
}

// UserAgreementHTML .
// @router /api/v1/url/onekey/UserAgreement.html [GET]
func UserAgreementHTML(ctx context.Context, c *app.RequestContext) {
	var err error

	url := "http://127.0.0.1:5000/onekey/UserAgreement.html" // 和html无关，仅为一个路径

	request := &protocol.Request{}
	request.SetMethod(consts.MethodGet)
	request.SetRequestURI(url)

	res := new(protocol.Response)

	if err = ClientSet.HzClient.Do(ctx, request, res); err != nil {
		pack.RespError(c, err)
		return
	}

	c.Data(consts.StatusOK, "text/html", res.BodyBytes())
}

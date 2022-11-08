/*
 * SPDX-FileCopyrightText: 2021-2022 Darren <1912544842@qq.com>
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package middleware

import (
	"emqx-user-manager/errno"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

func TokenAuth(r *ghttp.Request) {
	// 首先从头部获取auth鉴权
	auth := r.GetHeader("auth")
	do, err := g.Redis().Do("GET", auth)
	if err != nil {
		glog.Error("redis get error:", err)
		r.Response.WriteJsonExit(g.Map{"code": errno.RedisGetErr.Code, "msg": errno.RedisGetErr.Message})
	}
	// auth无效 超时或者无记录
	if do == nil {
		r.Response.WriteJsonExit(g.Map{"code": errno.AuthErr.Code, "msg": errno.AuthErr.Message})
	}
	r.Middleware.Next()
}

func EmqAuth(r *ghttp.Request) {
	baseAuthHead := r.GetHeader("Authorization")
	authCode, err := emqAuth(baseAuthHead)
	if err != nil {
		glog.Error("emqAuth get err", err)
	}
	if authCode != 200 {
		// 返回值不为200 鉴权失败 返回失败码
		r.Response.WriteJsonExit(g.Map{"code": errno.EMQAuthErr.Code, "msg": errno.EMQAuthErr.Message, "http": authCode})
	}
	r.Middleware.Next()
}

func emqAuth(auth string) (code int, err error) {
	c := g.Client()
	c.SetHeader("Authorization", auth)
	b, err := c.Get(g.Config().GetString("emqx.default") + "/stats")
	return b.StatusCode, err
}

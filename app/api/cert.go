/*
 * SPDX-FileCopyrightText: 2021-2022 Darren <1912544842@qq.com>
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package api

import (
	"emqx-user-manager/errno"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/guid"
	"time"
)

var Cert = certApi{}

type certApi struct{}

func (*certApi) Index(r *ghttp.Request) {
	baseAuthHead := r.GetHeader("Authorization")[6:]
	uid := guid.S([]byte(baseAuthHead))
	do, err := g.Redis().Do("GET", uid)
	if err != nil {
		glog.Error("redis get err:", err)
		r.Response.WriteJsonExit(g.Map{"code": errno.RedisGetErr.Code, "msg": errno.RedisGetErr.Message})
	}
	if do == nil {
		//glog.Debug(uid)
		_, err := g.Redis().Do("SET", uid, baseAuthHead, "EX", 3600) // 1h
		if err != nil {
			glog.Error("redis set err:", err)
			r.Response.WriteJsonExit(g.Map{"code": errno.RedisSetErr.Code, "msg": errno.RedisSetErr.Message})
		}
		r.Response.WriteJsonExit(g.Map{"code": errno.OK.Code, "msg": errno.OK.Message, "token": uid, "date": time.Now()})
	} else {
		//todo uid不唯一的处理
	}
}

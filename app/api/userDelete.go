/*
 * SPDX-FileCopyrightText: 2021-2022 Darren <1912544842@qq.com>
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package api

import (
	"emqx-user-manager/config"
	"emqx-user-manager/errno"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	_ "github.com/lib/pq"
)

var UserDelete = userDeleteApi{}

type userDeleteApi struct{}

func (*userDeleteApi) Index(r *ghttp.Request) {
	body := r.GetBodyString()
	glog.Debug(body)
	decode, err := gjson.DecodeToJson(body)
	if err != nil {
		// json解析出错 返回
		glog.Error("json decode err:", err)
		r.Response.WriteJsonExit(g.Map{"code": errno.UserIOErr.Code, "msg": errno.UserIOErr.Message})
	}
	// 获取要删除的用户名
	username := decode.GetString("username")
	db := g.DB("default").Model(config.EmqTableName)
	re, err := db.Delete("username", username)
	if err != nil {
		glog.Error("SQL del err:", err)
		r.Response.WriteJsonExit(g.Map{"code": errno.SQLDelErr.Code, "msg": errno.SQLDelErr.Message})
	}
	num, _ := re.RowsAffected()
	if num < 1 {
		r.Response.WriteJsonExit(g.Map{"code": errno.SQLDelErr.Code, "msg": errno.SQLDelErr.Message})
	}
	// 正常返回
	r.Response.WriteJsonExit(g.Map{"code": errno.OK.Code, "msg": errno.OK.Message})
}

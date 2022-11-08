/*
 * SPDX-FileCopyrightText: 2021-2022 Darren <1912544842@qq.com>
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package api

import (
	"crypto/sha256"
	"emqx-user-manager/errno"
	"encoding/hex"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	_ "github.com/lib/pq"
)

var UserAdd = userAddApi{}

type userAddApi struct{}

func (*userAddApi) Index(r *ghttp.Request) {
	// auth有效
	body := r.GetBodyString()
	glog.Debug(body)
	decode, err := gjson.DecodeToJson(body)
	if err != nil {
		// json解析出错 返回
		glog.Error("json decode err:", err)
		r.Response.WriteJsonExit(g.Map{"code": errno.UserIOErr.Code, "msg": errno.UserIOErr.Message})
	}
	// 获取三元组信息
	username := decode.GetString("username")
	passwd := decode.GetString("password")
	salt := decode.GetString("salt")
	// 先进行username唯一性校验
	db := g.DB("default").Model(g.Config().GetString("emqx.tableName"))
	all, err := db.Where("username=?", username).One()
	if err != nil {
		glog.Error("SQL get err:", err)
		r.Response.WriteJsonExit(g.Map{"code": errno.SQLGetErr.Code, "msg": errno.SQLGetErr.Message})
	}
	// 用户已经存在了.那就不能添加,直接返回错误
	if all != nil {
		r.Response.WriteJsonExit(g.Map{"code": errno.SQLHaveUser.Code, "msg": errno.SQLHaveUser.Message})
	}
	// 既然用户不存在 那就开始计算并存入数据库
	if salt != "" {
		passwd = salt + passwd
	}
	sum := sha256.Sum256([]byte(passwd))
	re, err := db.Insert(g.Map{"username": username,
		"password_hash": hex.EncodeToString(sum[:]),
		"salt":          salt})
	if err != nil {
		glog.Error("SQL set err:", err)
		r.Response.WriteJsonExit(g.Map{"code": errno.SQLSetErr.Code, "msg": errno.SQLSetErr.Message})
	}
	num, _ := re.RowsAffected()
	if num < 1 {
		r.Response.WriteJsonExit(g.Map{"code": errno.SQLSetErr.Code, "msg": errno.SQLSetErr.Message})
	}
	r.Response.WriteJsonExit(g.Map{"code": errno.OK.Code, "msg": errno.OK.Message})
}

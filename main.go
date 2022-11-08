/*
 * SPDX-FileCopyrightText: 2021-2022 Darren <1912544842@qq.com>
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"emqx-user-manager/app/api"
	"emqx-user-manager/app/middleware"
	_ "emqx-user-manager/init"
	"github.com/gogf/gf/frame/g"
)

func main() {
	s := g.Server()

	group := s.Group("/api")
	certGroup := group.POST("/cert", api.Cert)
	certGroup.Middleware(middleware.EmqAuth)

	userAddGroup := group.POST("/user", api.UserAdd)
	userAddGroup.Middleware(middleware.TokenAuth)

	userDelGroup := group.DELETE("/user", api.UserDelete)
	userDelGroup.Middleware(middleware.TokenAuth)

	s.Run()
}

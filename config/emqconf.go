/*
 * SPDX-FileCopyrightText: 2021-2022 Darren <1912544842@qq.com>
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import "github.com/gogf/gf/frame/g"

var EmqAddress string
var EmqTableName string
var EmqAuthType string

func init() {
	EmqAddress = g.Config().GetString("emqx.address")
	EmqTableName = g.Config().GetString("emqx.tableName")
	EmqAuthType = g.Config().GetString("emqx.authType")
}

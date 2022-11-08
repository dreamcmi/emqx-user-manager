/*
 * SPDX-FileCopyrightText: 2021-2022 Darren <1912544842@qq.com>
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package errno

type Err struct {
	Code    int
	Message string
}

var (
	OK = &Err{Code: 0, Message: "OK"}

	RedisGetErr = &Err{Code: 1001, Message: "REDIS GET ERROR"}
	RedisSetErr = &Err{Code: 1002, Message: "REDIS SET ERROR"}

	SQLGetErr   = &Err{Code: 2001, Message: "SQL GET ERROR"}
	SQLSetErr   = &Err{Code: 2002, Message: "SQL SET ERROR"}
	SQLHaveUser = &Err{Code: 2003, Message: "The User is already in the database"}
	SQLDelErr   = &Err{Code: 2004, Message: "SQL DELETE INFO ERROR"}

	AuthErr = &Err{Code: 3001, Message: "Auth Failed"}

	EMQAuthErr = &Err{Code: 4001, Message: "EMQ Auth Failed"}

	UserIOErr = &Err{Code: 5001, Message: "User input error"}
)

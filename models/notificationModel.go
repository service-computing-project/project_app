/*
 * @Descripttion:
 * @version:
 * @Author: sunylin
 * @Date: 2020-12-16 16:06:27
 * @LastEditors: sunylin
 * @LastEditTime: 2020-12-16 16:12:03
 */
package models

import "gopkg.in/mgo.v2"

//NotifiationDB 通知数据库
type NotifiationDB struct {
	DBN *mgo.Collection //通知数据库

}

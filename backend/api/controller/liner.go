package controller

import (
	"github.com/bragfoo/saman/backend/api/global"
	"github.com/gin-gonic/gin"
	"github.com/bragfoo/saman/backend/api/model/liner"
	"github.com/bragfoo/saman/util/db"
	"github.com/siddontang/go/log"
	"github.com/bragfoo/saman/backend/api/common"
	"github.com/bragfoo/saman/backend/api/model"
	"net/http"
)

func GetLinerPlayAmountByPlat(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		platIds := c.Query("platIds")

		var con []interface{}
		var sql = ""
		if "" != platIds {
			sql = liner.GetPlayAmountByPlatIds
			con = append(con, platIds)
		}
		stm, err := db.Prepare(sql)
		defer stm.Close()
		if nil != err {
			log.Error(err)
			common.StandardError(c)
		} else {
			rows, err := stm.Query(con...)
			defer rows.Close()
			if nil != err {
				log.Error(err)
				common.StandardError(c)
			} else {
				var result []model.PlayAmountLiner
				for rows.Next() {
					var m = model.PlayAmountLiner{}
					err := rows.Scan(&m.Sum, &m.CreateTime, &m.PlatIds)
					if nil == err {
						result = append(result, m)
					}
				}
				c.JSON(http.StatusOK, result)
			}
		}
	}
}

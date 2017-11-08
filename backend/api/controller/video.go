package controller

import (
	"github.com/bragfoo/saman/backend/api/global"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/bragfoo/saman/backend/api/model"
	"github.com/bragfoo/saman/backend/api/model/video"
	"github.com/siddontang/go/log"
	"github.com/bragfoo/saman/util"
	"github.com/bragfoo/saman/backend/api/common"
)

func GetVideoPlayAmount(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {
		stm, err := video.GetVideo()
		if nil != err {
			context.Status(http.StatusInternalServerError)
		} else {
			rows, _ := stm.Query()
			defer rows.Close()
			var result []model.VideoPlayAmount
			for rows.Next() {
				var model = model.VideoPlayAmount{}
				err := rows.Scan(&model.Ids,
					&model.Title,
					&model.Link,
					&model.CreateTime,
					&model.Sum,
					&model.NameChinese,
				)
				if nil == err {
					result = append(result, model)
				}
			}
			context.JSON(http.StatusOK, result)
		}
	}
}

func PostVideo(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		var m = model.Video{}
		common.ReadJSON(c, &m)
		stm, err := video.PostVideo()
		if nil != err {
			c.Status(http.StatusInternalServerError)
		} else {
			_, err := stm.Exec(util.GetObjectId(),
				m.VideoIds,
				m.PlatIds,
				m.Title,
				m.Link)
			if nil != err {
				log.Error(err)
				c.Status(http.StatusInternalServerError)
			} else {
				c.Status(http.StatusOK)
			}
		}
	}
}

func PostVideoPlayAmount(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		var m = model.PlayAmount{}
		common.ReadJSON(c, &m)
		stm, err := video.PostVideoPlayAmount()
		if nil != err {
			c.Status(http.StatusInternalServerError)
		} else {
			_, err := stm.Exec(util.GetObjectId(),
				m.VideoIds,
				m.CreateTime,
				m.Sum)
			if nil != err {
				log.Error(err)
				c.Status(http.StatusInternalServerError)
			} else {
				c.Status(http.StatusOK)
			}
		}
	}
}

func PutVideoPlayAmount(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {
		var m = model.VideoPlayAmount{}
		common.ReadJSON(context, &m)
		stm, err := video.PutVideoPlayAmount()
		if nil != err {
			log.Error(err)
			common.StandardError(context)
		} else {
			_, err := stm.Exec()
			if nil != err {
				log.Error(err)
				common.StandardError(context)
			} else {
				common.StandardSuccess(context)
			}
		}
	}
}

func PutVideo(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {
		var m = model.Video{}
		common.ReadJSON(context, &m)
		stm, err := video.PutVideo()
		if nil != err {
			log.Error(err)
			common.StandardError(context)
		} else {
			_, err := stm.Exec()
			if nil != err {
				log.Error(err)
				common.StandardError(context)
			} else {
				common.StandardSuccess(context)
			}
		}
	}
}

func DelVideo(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		ids := c.Param("ids")
		if "" == ids {
			common.StandardBadRequest(c)
		} else {
			stm, err := video.DelVideo()
			if nil != err {
				log.Error(err)
				common.StandardError(c)
			} else {
				_, err := stm.Exec(ids)
				if nil != err {
					log.Error(err)
					common.StandardError(c)
				} else {
					common.StandardSuccess(c)
				}
			}
		}
	}
}

func DelVideoPlayAmount(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		ids := c.Param("ids")
		if "" == ids {
			common.StandardBadRequest(c)
		} else {
			stm, err := video.DelVideoPlayAmount()
			if nil != err {
				log.Error(err)
				common.StandardError(c)
			} else {
				_, err := stm.Exec(ids)
				if nil != err {
					log.Error(err)
					common.StandardError(c)
				} else {
					common.StandardSuccess(c)
				}
			}
		}
	}
}

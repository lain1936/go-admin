package actions

import (
	"errors"
	"go-admin/service/dto"
	"go-admin/tools/app"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-admin/tools"
	"go-admin/tools/model"
)

// DeleteAction 通用删除动作
func DeleteAction(control dto.Control) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		idb, exist := c.Get("db")
		if !exist {
			err = errors.New("db connect not exist")
			tools.HasError(err, "", 500)
		}
		switch idb.(type) {
		case *gorm.DB:
			//删除操作
			db := idb.(*gorm.DB)
			req := control.Generate()
			err = req.Bind(c)
			tools.HasError(err, "参数验证失败", 422)
			var object model.ActiveRecord
			object, err = req.GenerateM()
			tools.HasError(err, "模型生成失败", 500)
			object.SetUpdateBy(tools.GetUserIdStr(c))
			err = db.WithContext(c).Delete(object).Error
			tools.HasError(err, "更新失败", 500)
			app.OK(c, object.GetId(), "更新成功")
		default:
			err = errors.New("db connect not exist")
			tools.HasError(err, "", 500)
		}
		c.Next()
	}
}
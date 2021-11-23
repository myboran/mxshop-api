package goods

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	// 将grpc的code转换为http的状态码
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"code": codes.NotFound,
					"msg":  e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusOK, gin.H{
					"msg": "用户服务不可用",
				})
			default:
				c.JSON(http.StatusOK, gin.H{
					"msg": "其他错误" + e.Message(),
				})
			}
			return
		}
	}
}

func HandleValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": errs.Error(),
	})
	return
}

func List(c *gin.Context) {

}

package controller

import (
	"blog/api/service"
	"blog/models"
	util "blog/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	service service.PostService
}

func NewPostController(s service.PostService) PostController {
	return PostController{
		service: s,
	}
}

func (p PostController) GetPosts(ctx *gin.Context) {
	var posts models.Post
	keyword := ctx.Query("keyword")
	data, total, err := p.service.FindAll(posts, keyword)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "不能查找到对应的问题")
		return
	}
	//respArr := make([]map[string]interface{}, 0, 0)
	respArr := make([]map[string]interface{}, 0)
	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}
	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		//Code:    "OK",
		Message: "数据获取成功",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

func (p *PostController) AddPost(ctx *gin.Context) {
	var post models.Post
	ctx.ShouldBindJSON(&post)
	util.SuccessJSON(ctx, http.StatusCreated, "成功新增数据")
}

func (p *PostController) GetPost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "需要ID数据")
		return
	}
	var post models.Post
	post.ID = id
	foundPost, err := p.service.Find(post)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "未能找到Post Model")
		return
	}
	response := foundPost.ResponseMap()
	c.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "数据获取成功",
		Data:    &response})
}

func (p *PostController) DeletePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	//err := p.service.Delete(id) //这里会引起"已定义但未使用"的错误，修改为判断
	if err == nil {
		p.service.Delete(id)
	}
	response := &util.Response{
		Success: true,
		Message: "数据删除成功"}
	c.JSON(http.StatusOK, &response)
}

func (p PostController) UpdatePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "需要ID数据")
		return
	}
	var post models.Post
	post.ID = id
	postRecord, err := p.service.Find(post)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "未能找到Post Model")
		return
	}
	c.ShouldBindJSON(&postRecord)
	response := postRecord.ResponseMap()
	c.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "数据更新成功",
		Data:    &response,
	})
}

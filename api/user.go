package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rocket049/gostructcopy"
	"gowork/dao"
	"gowork/public"
	"gowork/result"
	"gowork/validata"
	"log"
	"strconv"
	"time"
)

type Ids struct {
	Ids []string `json:"ids"`
}

type UserVo struct {
	ID    int    `json:"id"`                                  // 可以为空
	Name  string `json:"name" example:"zhangsan"`             // 页码
	Age   int    `json:"age" example:"23"`                    // 每页数据量
	Email string `json:"email" example:"l1277314169@163.com"` // 排序依据
}

func RegisterUser(router *gin.RouterGroup) {
	router.POST("/addUser", AddUser)
	router.POST("/updateUser", UpdateUser)
	router.POST("/del/:id", Del)
	router.GET("/queryAll", AllUser)
	router.GET("/findById/:id", FindById)
	router.POST("/findByIds", FindByInIds)
	router.POST("/queryByUserParam", FindByCondition)
	router.POST("/queryPage", FindByUserPage)
}

// @Summary 新增一个用户
// @Accept application/json
// @Produce application/json
// @Param message body UserVo true "ok"
// @Success 200 {bool} string	"ok"
// @Router /user/addUser [Post]
func AddUser(c *gin.Context) {
	var userVo UserVo
	validata.ShouldBin(c, &userVo)
	u := &dao.UserInfo{Name: userVo.Name, Age: userVo.Age, Email: userVo.Email, CreateTime: time.Now(), UpdateTime: time.Now()}
	user := dao.AddUser(u)
	result.ResponseSuccess(c, user)
}

// @Summary 编辑一个用户
// @Accept application/json
// @Produce application/json
// @Param message body api.UserVo true "UserVo"
// @Success 200 {bool} string	"ok"
// @Router /user/updateUser [Post]
func UpdateUser(c *gin.Context) {
	var userVo UserVo
	validata.ShouldBin(c, &userVo)
	u := &dao.UserInfo{ID: userVo.ID, Name: userVo.Name, Age: userVo.Age, Email: userVo.Email}
	result.ResponseSuccess(c, dao.UpdateUser(u))
}

// @Summary 删除一个用户
// @Param id path int true "ID"
// @Success 200 {bool} string	"ok"
// @Router /user/del/{id} [Post]
func Del(c *gin.Context) {
	id := c.Param("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalf("id parse error = %v", err)
		return
	}
	result.ResponseSuccess(c, dao.Del(atoi))
}

// @Summary 查询所用用户
// @Produce application/json
// @Success 200 {json} json "{"errno":0,"errmsg": "成功","data":userInfo}"
// @Router /user/queryAll [Get]
func AllUser(c *gin.Context) {
	idIn := dao.QueryUserInfoAll()
	var userVo []UserVo
	for _, info := range idIn {
		var user UserVo
		gostructcopy.StructCopy(&info, &user)
		userVo = append(userVo, user)
	}
	result.ResponseSuccess(c, userVo)
}

// @Summary 查询一个用户
// @Param id path int true "ID"
// @Success 200 {object}  dao.UserInfo "ok"
// @Router /user/findById/{id} [Get]
func FindById(c *gin.Context) {
	id := c.Param("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalf("id parse error = %v", err)
		return
	}
	info := dao.FindUserInfoById(atoi)
	var userVo UserVo
	gostructcopy.StructCopy(info, &userVo)
	result.ResponseSuccess(c, userVo)
}

// @Summary 查询多个用户Ids
// @Accept application/json
// @Produce application/json
// @Param message body Ids true "ok"
// @Success 200 {json} json "{"errno":0,"errmsg": "成功","data":userInfo}"
// @Router /user/findByIds [Post]
func FindByInIds(c *gin.Context) {
	var a Ids
	validata.ShouldBin(c, &a)
	var b []int
	for _, id := range a.Ids {
		atoi, err := strconv.Atoi(id)
		if err != nil {
			return
		}
		b = append(b, atoi)
	}
	idIn := dao.FindUserInfoByIdIn(b)
	var userVo []UserVo
	for _, info := range idIn {
		var user UserVo
		gostructcopy.StructCopy(&info, &user)
		userVo = append(userVo, user)
	}
	result.ResponseSuccess(c, userVo)
}

// @Summary 条件查询多个用户
// @Accept application/json
// @Produce application/json
// @Param message body UserVo true "ok"
// @Success 200 {json} json "{"errno":0,"errmsg": "成功","data":userInfo}"
// @Router /user/queryByUserParam [Post]
func FindByCondition(c *gin.Context) {
	var userParam UserVo
	validata.ShouldBin(c, &userParam)
	u := &dao.UserInfo{ID: userParam.ID, Name: userParam.Name, Age: userParam.Age, Email: userParam.Email}
	idIn := dao.FindUserInfoByCondition(u)
	var userVo []UserVo
	for _, info := range idIn {
		var user UserVo
		gostructcopy.StructCopy(&info, &user)
		userVo = append(userVo, user)
	}
	result.ResponseSuccess(c, userVo)
}

// @Summary 查询多个用户
// @Accept application/json
// @Produce application/json
// @Param message body dao.Page true "ok"
// @Success 200 {json} json "{"errno":0,"errmsg": "成功","data":userVo}"
// @Router /user/queryPage [Post]
func FindByUserPage(c *gin.Context) {
	var page public.Page
	validata.ShouldBin(c, &page)
	idIn := dao.FinUserInfoPage(&page)
	var userVo []UserVo
	for _, info := range idIn {
		var user UserVo
		gostructcopy.StructCopy(&info, &user)
		userVo = append(userVo, user)
	}
	result.ResponseSuccess(c, userVo)
}

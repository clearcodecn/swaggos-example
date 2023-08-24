package main

import (
	"github.com/clearcodecn/swaggos"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {

	g := gin.Default()

	{
		// 文档
		doc := swaggos.Default()
		doc.Response(200, BizVo{})

		cpg := doc.Group("/api/pkgs").Tag("充值包管理 - C端")
		cpg.Get("/list").JSON(make([]PackageVo, 0)).Summary("充值包列表")

		cog := doc.Group("/api/order").Tag("用户订单 - C端")
		cog.Post("/create").Body(CreateOrderRequest{}).JSON(CreateOrderResponse{}).Summary("创建订单")
		cog.Post("/check_order_status").Body(CheckOrderStatusRequest{}).JSON(CheckOrderStatusResponse{}).Summary("查询订单状态")
		cog.Post("/list").Body(ListOrdersRequest{}).JSON(UserOrderVo{}).Summary("用户订单列表")

		pg := doc.Group("/api/admin/pkgs").Tag("充值包管理")
		pg.Get("list").QueryObject(ListPackageRequest{}).JSON(make([]PackageVo, 0)).Summary("充值包列表")
		pg.Post("create").Body(AddPackageRequest{}).JSON(BizVo{}).Summary("创建充值包")
		pg.Post("update").Body(PackageVo{}).JSON(BizVo{}).Summary("更新充值包")
		pg.Get("detail").QueryObject(PkgIdRequest{}).JSON(PackageVo{}).Summary("充值包详情")
		pg.Post("delete").Body(PkgIdRequest{}).JSON(BizVo{}).Summary("删除充值包")
		pg.Post("toggle_status").Body(UpdateStatusRequest{}).JSON(BizVo{}).Summary("修改状态")

		og := doc.Group("/api/admin/order").Tag("订单管理")
		og.Get("/list").QueryObject(ListOrdersRequest{}).JSON(make([]OrderVo, 0)).Summary("订单列表")

		// 在线文档
		g.GET("/_swagger/*action", gin.WrapH(swaggos.UI("/_swagger", "/doc.json")))
		g.GET("/doc.json", gin.WrapH(doc))
	}

	g.Run(":1234")
}

// BizVo 业务返回 VO
type BizVo struct {
	Page     int         `json:"page,omitempty"`
	PageSize int         `json:"page_size,omitempty"`
	Total    int         `json:"total,omitempty"`
	Message  string      `json:"message,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}

type PackageVo struct {
	Name        string  `json:"name" description:"名称"`
	Description string  `json:"description" description:"描述"`     // 描述信息
	Token       int     `json:"token" description:"token数量, > 0"` // token 数量
	Times       int     `json:"times" description:"次数, >0"`       // 次数
	Status      int     `json:"status" description:"状态: 1启用，2禁用"` // 状态
	Price       float64 `json:"price" description:"价格, > 0"`      // 价格.
	Sort        int     `json:"sort" description:"排序，越小越靠前"`      // 排序，越小越靠前
}

type CreateOrderRequest struct {
	PkgId    int    `json:"pkg_id"`
	Platform string `json:"platform" description:"平台：wx=在微信浏览器发起直接支付, pc=在web页面,手机浏览器页面生成二维码，扫码支付"`
}

type CheckOrderStatusRequest struct {
	OrderNo string `json:"order_no"`
}

type CheckOrderStatusResponse struct {
	Status int `json:"status" description:"订单状态: 1=已创建 , 2=支付成功, 3=已退款"`
}

type ListOrdersRequest struct {
	Page     int    `json:"page" form:"page" query:"page"`
	PageSize int    `json:"page_size" form:"page_size" query:"page_size"`
	SortBy   string `json:"sort_by" query:"sort_by" form:"sort_by" description:"排序字段：create_time=创建时间"`
	Asc      bool   `json:"asc" query:"asc" form:"asc" description:"排序规则：asc=true 从小到大, false 从大到小"`
}

type CreateOrderResponse struct {
	OrderNo string `json:"order_no" description:"订单号"`
}

type AddPackageRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"` // 描述信息
	Token       int     `json:"token"`       // token 数量
	Times       int     `json:"times"`       // 次数
	Status      int     `json:"status"`      // 状态
	Price       float64 `json:"price"`       // 价格.
}

type ListPackageRequest struct {
	Page     int `json:"page" form:"page" query:"page"`
	PageSize int `json:"page_size" form:"page_size" query:"page_size"`
}

type PkgIdRequest struct {
	Id int `json:"id" form:"id" query:"id"`
}

type UpdateStatusRequest struct {
	Id     int `json:"id"`
	Status int `json:"status"`
}

type UserOrderVo struct {
	Id        uint      `json:"id" description:"id"`
	CreatedAt time.Time `json:"created_at" description:"创建时间"`
	UpdatedAt time.Time `json:"updated_at" description:"更新时间"`
	OrderNo   string    `description:"订单号" json:"order_no" gorm:"index"` // 订单号
	Price     float32   `description:"价格" json:"price"`                  // 支付金额 , 价格
	PkgId     int64     `description:"包id" json:"pkg_id"`                // 包id.
	PkgName   string    `description:"包名" json:"pkg_name"`               // 包名
	Times     int       `description:"次数" json:"times"`                  // 次数
	Token     int       `description:"token数量" json:"token"`             // token的数量
	Status    int       `description:"状态" json:"status"`                 // 订单状态
	Uid       int       `description:"用户id" json:"uid"`                  // 用户id
	PayTime   time.Time `description:"支付时间" json:"pay_time"`             // 支付时间
	WxOrderNo string    `description:"微信订单号" json:"wx_order_no"`         // 微信订单号
}

type OrderVo struct {
	Id        uint      `json:"id" description:"id"`
	CreatedAt time.Time `json:"created_at" description:"创建时间"`
	UpdatedAt time.Time `json:"updated_at" description:"更新时间"`
	OrderNo   string    `description:"订单号" json:"order_no" gorm:"index"` // 订单号
	Price     float32   `description:"价格" json:"price"`                  // 支付金额 , 价格
	PkgId     int64     `description:"包id" json:"pkg_id"`                // 包id.
	PkgName   string    `description:"包名" json:"pkg_name"`               // 包名
	Times     int       `description:"次数" json:"times"`                  // 次数
	Token     int       `description:"token数量" json:"token"`             // token的数量
	Status    int       `description:"状态" json:"status"`                 // 订单状态
	Uid       int       `description:"用户id" json:"uid"`                  // 用户id
	PayTime   time.Time `description:"支付时间" json:"pay_time"`             // 支付时间
	Username  string    `description:"用户名" json:"username"`              // 用户名
	WxOrderNo string    `description:"微信订单号" json:"wx_order_no"`         // 微信订单号
}

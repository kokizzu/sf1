package cProductManager

import (
	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/S"
	"sf1/model"
	"sf1/model/mProduct"
)

type Product_List_Response struct {
	Rows  []model.Product
	Count int64
	Error string `json:",omitempty"`
}

// @Summary get product list
// @Description
// @Accept json
// @Produce json
// @Param offset path int true "page, default: 0"
// @Param limit path int true "page size, default: 10"
// @Success 200 {object} Product_List_Response
// @Router /product_manager/product/list/{limit}/{offset} [get]
func Product_List(ctx *model.Ctx) {
	offset := I.MaxOf(0, S.ToInt(ctx.Param(`offset`)))
	limit := I.MinOf(10, S.ToInt(ctx.Param(`limit`)))
	count, err := mProduct.Total(ctx)
	if ctx.ReturnError(err) {
		return
	}
	rows, err := mProduct.ListAll(ctx, limit, offset)
	if ctx.ReturnError(err) {
		return
	}
	ctx.JSON(200, Product_List_Response{
		Rows:  rows,
		Count: count,
	})
}

type Product_ById_Response struct {
	Row   model.Product
	Error string `json:",omitempty"`
}

// @Summary get product by id
// @Description
// @Accept json
// @Produce json
// @Param productId path int true "product id"
// @Success 200 {object} Product_ById_Response
// @Router /product_manager/product/by-id/{productId} [get]
func Product_ById(ctx *model.Ctx) {
	productId := S.ToInt(ctx.Param(`productId`))
	row, err := mProduct.One(ctx, productId)
	if ctx.ReturnError(err) {
		return
	}
	ctx.JSON(200, Product_ById_Response{
		Row: row,
	})
}

type Product_Upsert_Request struct {
	Product model.Product `json:",omitempty"`
	Action  string        `json:",omitempty"`
}

type Product_Upsert_Response struct {
	Product model.Product `json:",omitempty"`
	Error   string        `json:",omitempty"`
}

// @Summary update product information
// @Description
// @Accept json
// @Produce json
// @Param product body Product_Upsert_Request true "id and fields of product that want to be updated/deleted, inserted=0"
// @Success 200 {object} Product_Upsert_Response
// @Router /product_manager/product/modify [post]
func Product_Upsert(ctx *model.Ctx) {
	p := Product_Upsert_Request{}
	if err := ctx.Bind(&p); ctx.IsError(err, `failed parsing post`) {
		return
	}
	var err error
	switch p.Action {
	case `delete`:
		err = mProduct.Delete(ctx, p.Product.ID)
	case `restore`:
		err = mProduct.Restore(ctx, p.Product.ID)
	default: // upsert
		err = mProduct.Upsert(ctx, &p.Product)
	}
	if ctx.ReturnError(err) {
		return
	}
	ctx.JSON(200, Product_Upsert_Response{Product: p.Product})
}

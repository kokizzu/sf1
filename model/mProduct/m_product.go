package mProduct

import (
	"sf1/model"
)

var This = &model.Product{}

func Total(dep *model.Ctx) (count int64, err error) {
	res := dep.Db.Model(This).Count(&count)
	err = res.Error
	return
}

func ListAll(dep *model.Ctx, limit, offset int) (users []model.Product, err error) {
	res := dep.Db.Limit(limit).Offset(offset).Find(&users)
	err = res.Error
	return
}

func One(dep *model.Ctx, productId int) (user model.Product, err error) {
	res := dep.Db.Model(This).First(&user, `id = ?`, productId)
	err = res.Error
	return
}

func Delete(dep *model.Ctx, productId uint) (err error) {
	res := dep.Db.Delete(This, productId) // TODO: change with update when actor exists
	err = res.Error
	return
}

func Restore(dep *model.Ctx, productId uint) (err error) {
	res := dep.Db.Model(This).Where(`id = ?`, productId).Updates(map[string]interface{}{
		`deleted_at`: nil,
	})
	err = res.Error
	return
}

func Upsert(dep *model.Ctx, product *model.Product) (err error) {
	res := dep.Db.Save(product)
	err = res.Error
	return
}

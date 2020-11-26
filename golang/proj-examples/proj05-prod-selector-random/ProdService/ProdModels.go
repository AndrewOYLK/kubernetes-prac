package ProdService

import "strconv"

// 商品相关的模型
type ProdModel struct {
	ProdID   int
	ProdName string
}

// 返回指针或者返回值也行，这里返回指针，万一需要改，就可以直接去改
func NewProd(id int, pname string) *ProdModel {
	return &ProdModel{
		ProdID:   id,
		ProdName: pname,
	}
}

func NewProdList(n int) []*ProdModel {
	ret := make([]*ProdModel, 0)
	for i := 0; i < n; i++ {
		ret = append(ret, NewProd(100+i, "prodname"+strconv.Itoa(i)))
	}
	return ret
}

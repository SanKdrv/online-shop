package types

type GetIdByBrandRequest struct {
	BrandName string `json:"brand_name"`
}

type GetIdByBrandResponse struct {
	BrandId int64 `json:"brand_id"`
}

type GetBrandByIdRequest struct {
	BrandId int64 `json:"brand_id"`
}

type GetBrandByIdResponse struct {
	BrandName string `json:"brand_name"`
}

type CreateBrandRequest struct {
	BrandName string `json:"brand_name"`
}

type CreateBrandResponse struct {
	BrandId int64 `json:"brand_id"`
}

type DeleteBrandRequest struct {
	BrandId int64 `json:"brand_id"`
}

type DeleteBrandResponse struct {
	Status string `json:"status"`
}

type UpdateBrandRequest struct {
	OldBrandName string `json:"new_brand_name"`
	NewBrandName string `json:"old_brand_name"`
}

type UpdateBrandResponse struct {
	Status string `json:"status"`
}

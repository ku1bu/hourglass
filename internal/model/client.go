package model

type WatcBack func(*Client)

type Client struct {
	Id       string
	Quota    int32
	WatcBack WatcBack
	Weights  float32
	Hunger   bool
}

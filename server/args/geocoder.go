package args

type Geocoder struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}
func (this *Geocoder)QQmapLocation()string {
	return  this.Lat+","+this.Lng
}
//经纬度地址解析
type GeocoderResp struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Result struct{
		AddressComponent struct{
			Province string `json:"province"`
			City string `json:"city"`
			District string `json:"district"`
			Ip    string `json:"ip"`
			Lat   string `json:"lat"`
			Lng   string `json:"lng"`
			Method string `json:"method"`
		} `json:"address_component"`
	} `json:"result"`
}
func (this *GeocoderResp)IsOk()bool{
	return this.Status == 0
}
func (this *GeocoderResp)Error()string{
	return this.Message
}


//经纬度地址解析
type IpcoderResp struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Result struct{
		AdInfo struct{
			Province string `json:"province"`
			City string `json:"city"`
			District string `json:"district"`
			Ip    string `json:"ip"`
			Lat   string `json:"lat"`
			Lng   string `json:"lng"`
		} `json:"ad_info"`
	} `json:"result"`
}
func (this *IpcoderResp)IsOk()bool{
	return this.Status == 0
}
func (this *IpcoderResp)Error()string{
	return this.Message
}
package db

func AddBrand(brand *DeviceBrand)error{
	_,err:=MysqlDB.Table("sxiot_brand").Insert(brand)
	return err
}

func FindBrands()([]DeviceBrand,error){
	var brands []DeviceBrand
	err:=MysqlDB.Table("sxiot_brand").Find(&brands)
	return brands,err
}

func GetBrand(brandID string)(*DeviceBrand,error){
	var brand DeviceBrand
	res,err:=MysqlDB.Table("sxiot_brand").Where("brand_id=?",brandID).Get(&brand)
	if !res{
		return nil,nil
	}
	return &brand,err
}

func GetBrandByB(b string)(*DeviceBrand,error){
	var brand DeviceBrand
	res,err:=MysqlDB.Table("sxiot_brand").Where("brand_type=?",b).Get(&brand)
	if !res{
		return nil,nil
	}
	return &brand,err
}


func UpdateBrand(brand *DeviceBrand)error{
	_,err:=MysqlDB.Table("sxiot_brand").Where("brand_id=?",brand.ID).Update(brand)
	return err
}

func DelBrand(brandID string)error{
	brand:=new(DeviceBrand) 
	_,err:=MysqlDB.Table("sxiot_brand").Where("brand_id=?",brandID).Delete(brand)
	return err
}
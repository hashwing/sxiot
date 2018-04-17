package db

func AddNews(news *News)error{
	_,err:=MysqlDB.Table("sxiot_news").Insert(news)
	return err
}

func FindNews()([]News,error){
	var news []News
	err:=MysqlDB.Table("sxiot_news").Find(&news)
	return news,err
}


func FindTopNews()([]News,error){
	var news []News
	err:=MysqlDB.Table("sxiot_news").Limit(10,0).Desc("created").Find(&news)
	return news,err
}

func DelNews(nid string)error{
	news:=new(News) 
	_,err:=MysqlDB.Table("sxiot_news").Where("news_id=?",nid).Delete(news)
	return err
}
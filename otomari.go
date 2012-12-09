package main

import 
(
	"fmt" 
	// "strconv"
	"io"
	"net/http"
	// "net/url"
	"io/ioutil"
	"encoding/json"
	"strings";	"log"
)

type TravelAPIInfo struct {
	Body BodyT
}
type BodyT struct{
	PagingInfo PagingInfoT
	Hotel []HotelT
	header HeaderT
	StatusMsg string
}
type PagingInfoT struct{
	RecordCount, PageCount, Last, Page, First int
}
type HotelT struct {
	RoomInfo []RoomT
	HotelInfo []HotelInfoT
}
type RoomT struct {
	PlanName string
	DailyCharge []DailyT
}
type HotelInfoT struct {
	HotelBasicInfo HotelBasicInfoT
}
type HotelBasicInfoT struct{
	HotelName string
	HotelSpecial string
}
type HeaderT struct {
	Status int
}
type DailyT struct {
	Total,RakutenCharge,ChargeFlag int
	StayDate string
}



func main() {
	// getUrl
	url := "http://api.rakuten.co.jp/rws/3.0/json?" +
		"developerId=1013796620208528384&" +
		"operation=VacantHotelSearch&" +
		"version=2009-10-20&" +
		"largeClassCode=japan&middleClassCode=kumamoto&" + 
		"smallClassCode=aso&" +
		"checkinDate=2012-12-09&" + 
		"checkoutDate=2012-12-10&" + 
		"adultNum=1" + 
		"&maxCharge=10000" ;
	// getJson from API
	resp, err := http.Get(url);
	// check it
	if err != nil {
		fmt.Println("error now!!")
		log.Fatal(err);

		return
	}

	// extract body from resp
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error now!!")
		fmt.Println(err)
		return
	}

	// cast array to string
	str_body := string(body)
	// fmt.Printf("%T:%v\n",str_body , str_body )

	// json parse
	var tinfo TravelAPIInfo
	dec := json.NewDecoder(strings.NewReader(str_body))
	if err := dec.Decode(&tinfo); err == io.EOF {
		return
	} else if err != nil {
		log.Fatal(err)
	}

	// extractKurokawa
	// fmt.Println(tinfo)
	hotels := tinfo.Body.Hotel
	// size := hotels.length
	
	for i, hotel := range hotels {

		fmt.Println("Hotel : ", i)
		fmt.Println(" Name : " , hotel.HotelInfo[0].HotelBasicInfo.HotelName)
		fmt.Println(" Spec : " , hotel.HotelInfo[0].HotelBasicInfo.HotelSpecial)
		for i, plan := range hotel.RoomInfo{
			fmt.Println("    => Plan[" , i , "] : ", plan.PlanName)
			fmt.Println("    => place     : ", plan.DailyCharge[0].RakutenCharge)
		}
	}
	// output
}
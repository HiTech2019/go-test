package main

import (
	"encoding/json"
	"fmt"
)

type RespBlockChainSrv struct {
	Code string `json:"code,omitempty"`
	//Data Points `json:"data,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg,omitempty"`
}

type TxID struct {
	Txid string `json:"txid"`
}

type Points struct {
	Points string `json:"points"`
	Frozen string `json:"frozen"`
}

type DescTrans struct {
	User              string `json:"user"`
	Points            string `json:"points"`
	PointsType        string `json:"points_type"`
	Source            string `json:"source"`
	Type              string `json:"type"`
	Status            string `json:"status"`
	TimeStamp         string `json:"time_stamp"`
	TimeStampUnixNano string `json:"time_stamp_unix_nano"`
	Txhash            string `json:"txhash"`
	Txid              string `json:"txid"`
}

type Records struct {
	Records []DescTrans `json:"records"`
}

func main() {
	jsonstr := `{"code":"0", "data":{"records": [{"points":"5","points_type":"easystore_points","source":"kj","status":"committed","time_stamp":"2019-03-19 16:52:47","time_stamp_unix_nano":"1552985567416728632","txhash":"275841b3ed39efb7ea13b9aaee2b7e8470eb59f371069f5bda751112af9ed639","txid":"dfde73ca-bfe8-4c14-8545-3ca953513d92","type":"","user":"18500314047"},{"points":"5","points_type":"easystore_points","source":"kj","status":"committed","time_stamp":"2019-03-19 16:43:22","time_stamp_unix_nano":"1552985002110651989","txhash":"8dfc46a58d0678d3932e5fa52864221a5dd335596e8ce17de761289bad7a6615","txid":"b9f6871b-9f40-4ef4-976d-654de02ebba6","type":"","user":"18500314047"}]}}`
	// var resp RespBlockChainSrv
	// resp.Code = "0"
	// resp.Data = Points{Points: "111", Frozen: "100"}

	// bytes, err := json.Marshal(resp)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println(string(bytes))

	var respBcSrv RespBlockChainSrv
	err := json.Unmarshal([]byte(jsonstr), &respBcSrv)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if respBcSrv.Code == "0" {
		//fmt.Println(respBcSrv.Data)
		data1 := (respBcSrv.Data).(interface{})
		//fmt.Println(data1)
		//data1.(map[string]interface{})

		if data, ok := data1.(map[string]interface{}); ok {
			fmt.Println(data["points"].(string))
			//fmt.Println(data["points"].(string))
		}

		// if date, ok := (respBcSrv.Data).(map[string]interface{}); ok {
		// 	fmt.Println(date)
		// 	fmt.Println(date["points"].(string))
		// 	fmt.Println(date["points"].(string))
		// } else {
		// 	fmt.Println("返回非%v类型,异常失败 ", &Points{})
		// 	fmt.Println(date)
		// 	fmt.Println(respBcSrv.Data)
		// }

		// if date, ok := (respBcSrv.Data).(Points); ok {
		// 	fmt.Println(date)
		// } else {
		// 	fmt.Println("返回非%v类型,异常失败 ", &Points{})
		// 	fmt.Println(date)
		// 	fmt.Println(respBcSrv.Data)
		// }
	} else {
		fmt.Println("获取区块链积分余额失败 %s", respBcSrv.Msg)
	}

}

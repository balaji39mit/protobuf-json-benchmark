package proto_files

import (
	"bitbucket.org/grabpay/risk-service/data_layer/client/dto"
	"encoding/json"
)

func populateIdsandMerchantIds(data *AggrData, stats *dto.Stats) {
	data.Ids = make(map[string]string)
	for k, _ := range stats.Obj.Ids {
		data.Ids[k] = ""
	}
	for k, _ := range stats.Obj.MerchantId {
		data.MerchantId[k] = ""
	}
}

func convertStatstoPBStats(statsData *dto.Stats) Stats {
	aggrData := AggrData{
		TotalAmount: float32(statsData.Obj.TotalAmount),
		Count:       int32(statsData.Obj.Count),
	}
	populateIdsandMerchantIds(&aggrData, statsData)
	stats := Stats{
		Year:     int32(statsData.Year),
		Month:    int32(statsData.Month),
		Day:      int32(statsData.Day),
		Hour:     int32(statsData.Hour),
		AggrData: &aggrData,
	}
	return stats
}

func UnMarshalStruct(passInStr string, targetStruct interface{}) {
	if passInStr == "" {
		return
	}

	json.Unmarshal([]byte(passInStr), targetStruct)
}

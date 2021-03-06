package proto_files

import (
	"bitbucket.org/grabpay/risk-service/data_layer/client/dto"
	"fmt"
	"testing"
)

var (
	aggrData   AggrData
	stats      Stats
	statsArray StatsArray
)

func init() {
	statsBytes := "[{\"year\":2018,\"month\":8,\"day\":218,\"hour\":3,\"obj\":{\"t_amt\":128.7,\"cnt\":111,\"ids\":{\"11390265\":{},\"117921155\":{},\"11971561\":{},\"119877915\":{},\"12137936\":{},\"13130768\":{},\"13938385\":{},\"14147098\":{},\"14310147\":{},\"14844203\":{},\"14898979\":{},\"14951554\":{},\"15146058\":{},\"15397125\":{},\"16154156\":{},\"16443971\":{},\"17003208\":{},\"17674769\":{},\"18180689\":{},\"1836364\":{},\"18618799\":{},\"19134128\":{},\"19366466\":{},\"1974526\":{},\"2041979\":{},\"2141869\":{},\"22134067\":{},\"2390077\":{},\"2556128\":{},\"2575287\":{},\"2646069\":{},\"26813337\":{},\"27981017\":{},\"28355437\":{},\"28618752\":{},\"2880101\":{},\"29219897\":{},\"30375777\":{},\"30671059\":{},\"3218534\":{},\"33386529\":{},\"344319\":{},\"3940747\":{},\"398822\":{},\"39892615\":{},\"4004916\":{},\"40555717\":{},\"41564267\":{},\"4202873\":{},\"42795562\":{},\"43540696\":{},\"4381083\":{},\"45194835\":{},\"4524799\":{},\"45856873\":{},\"4609901\":{},\"46150899\":{},\"4809945\":{},\"4928901\":{},\"50320799\":{},\"51225510\":{},\"54583071\":{},\"5616799\":{},\"5636602\":{},\"6067843\":{},\"60747337\":{},\"6146823\":{},\"61730241\":{},\"6186773\":{},\"6494416\":{},\"6554031\":{},\"6563176\":{},\"668915\":{},\"6857030\":{},\"6887799\":{},\"692910\":{},\"707917\":{},\"7276616\":{},\"7673976\":{},\"767494\":{},\"8060859\":{},\"8178807\":{},\"8542484\":{},\"8583628\":{},\"88295469\":{},\"896413\":{},\"9112644\":{},\"9134135\":{},\"937208\":{},\"9421793\":{},\"9446832\":{},\"9571278\":{},\"9970808\":{}},\"merchant_id\":null}},{\"year\":2018,\"month\":8,\"day\":219,\"hour\":4,\"obj\":{\"t_amt\":150.6,\"cnt\":96,\"ids\":{\"11216984\":{},\"1138711\":{},\"1233226\":{},\"13216286\":{},\"13575790\":{},\"14379044\":{},\"14631370\":{},\"14844203\":{},\"14951554\":{},\"15243516\":{},\"16047572\":{},\"16162105\":{},\"16381048\":{},\"16686089\":{},\"1797722\":{},\"18384501\":{},\"18638381\":{},\"1936946\":{},\"1974526\":{},\"20182102\":{},\"2181342\":{},\"21959496\":{},\"2202536\":{},\"22065097\":{},\"2307686\":{},\"2360952\":{},\"2370958\":{},\"23827343\":{},\"24134228\":{},\"24437252\":{},\"24450151\":{},\"24511694\":{},\"24558274\":{},\"24797304\":{},\"24827816\":{},\"25016584\":{},\"25347015\":{},\"2579642\":{},\"26813337\":{},\"27479763\":{},\"28145347\":{},\"2931143\":{},\"29834496\":{},\"3006966\":{},\"30517182\":{},\"30671059\":{},\"31947952\":{},\"32041576\":{},\"3290846\":{},\"3308109\":{},\"33386529\":{},\"33972921\":{},\"3456176\":{},\"36121326\":{},\"37454774\":{},\"3817402\":{},\"39160984\":{},\"42539754\":{},\"43141604\":{},\"4364332\":{},\"4378404\":{},\"4403517\":{},\"4437032\":{},\"5049306\":{},\"5083978\":{},\"52276783\":{},\"5234305\":{},\"5313096\":{},\"5350811\":{},\"5515416\":{},\"5636602\":{},\"59203303\":{},\"5988966\":{},\"6056097\":{},\"6268099\":{},\"66515792\":{},\"678388\":{},\"707917\":{},\"742628\":{},\"742669\":{},\"7603587\":{},\"7863489\":{},\"80924198\":{},\"81295495\":{},\"8242510\":{},\"825292\":{},\"8337896\":{},\"94942796\":{},\"957884\":{},\"9970808\":{}},\"merchant_id\":null}},{\"year\":2018,\"month\":8,\"day\":220,\"hour\":3,\"obj\":{\"t_amt\":106.9,\"cnt\":78,\"ids\":{\"106239594\":{},\"109147206\":{},\"1105097\":{},\"11216984\":{},\"1136540\":{},\"11464625\":{},\"1207295\":{},\"13130768\":{},\"13480561\":{},\"14605738\":{},\"14619261\":{},\"14844203\":{},\"14948901\":{},\"15458809\":{},\"15837719\":{},\"17003208\":{},\"17014279\":{},\"18208840\":{},\"18715464\":{},\"20437259\":{},\"20985301\":{},\"2181342\":{},\"21846573\":{},\"21982889\":{},\"22748358\":{},\"2365497\":{},\"24002336\":{},\"2417621\":{},\"2636787\":{},\"2770096\":{},\"28516649\":{},\"2885932\":{},\"29985999\":{},\"3006966\":{},\"31896474\":{},\"32869183\":{},\"35033364\":{},\"3594363\":{},\"37481509\":{},\"39616341\":{},\"41671609\":{},\"4294006\":{},\"4312855\":{},\"4364332\":{},\"44253760\":{},\"4501878\":{},\"45168361\":{},\"4671781\":{},\"46765906\":{},\"4791089\":{},\"5313096\":{},\"5619826\":{},\"5851726\":{},\"5936667\":{},\"5995333\":{},\"6191153\":{},\"6504398\":{},\"688012\":{},\"7210247\":{},\"7232233\":{},\"72484947\":{},\"7313191\":{},\"75730616\":{},\"7674452\":{},\"77768710\":{},\"849914\":{},\"88295469\":{},\"9134135\":{},\"926149\":{},\"9421793\":{},\"94942796\":{},\"99803840\":{}},\"merchant_id\":null}}]"
	statsLarge := new([]*dto.Stats)
	UnMarshalStruct(string(statsBytes), statsLarge)
	dummy := *statsLarge
	for _, value := range dummy {
		statsData := convertStatstoPBStats(value)
		statsArray.StatsArray = append(statsArray.StatsArray, &statsData)
	}
}

func TestDataAllocations(_ *testing.T) {
	//testAggrData(&aggrData)
	testStatsArray(&statsArray)
}

func BenchmarkProtobufMarshal(b *testing.B) {
	b.ResetTimer()

	//marshalUtilProtobuf(b, "AggrData", &aggrData)
	//marshalUtilProtobuf(b, "Stats", &stats)
	marshalUtilProtobuf(b, "StatsArray", &statsArray)

	fmt.Printf("\n")
}

func BenchmarkJSONMarshal(b *testing.B) {
	b.ResetTimer()

	//marshalUtilJSON(b, "AggrData", &aggrData)
	//marshalUtilJSON(b, "Stats", &stats)
	marshalUtilJSON(b, "StatsArray", &statsArray)

	fmt.Printf("\n")
}

func BenchmarkProtobufUnmarshal(b *testing.B) {
	/*var ad AggrData
	var s Stats*/
	var sA StatsArray
	b.ResetTimer()

	//unMarshalUtilProtobuf(b, "AggrData", &ad, &aggrData)
	//unMarshalUtilProtobuf(b, "Stats", &s, &stats)
	unMarshalUtilProtobuf(b, "StatsArray", &sA, &statsArray)

	fmt.Printf("\n")

}

func BenchmarkJSONUnmarshal(b *testing.B) {
	/*var ad AggrData
	var s Stats*/
	var sA StatsArray

	b.ResetTimer()

	//unMarshalUtilJSON(b, "AggrData", &aggrData, &ad)
	//unMarshalUtilJSON(b, "Stats", &stats, &s)
	unMarshalUtilJSON(b, "StatsArray", &statsArray, &sA)
	fmt.Printf("\n")
}

package proto_files

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"testing"
)

func testStatsArray(statsArray *StatsArray) {
	fmt.Printf("Memory footprint of the final byte array after a Marshal ---------\n")
	js, _ := json.Marshal(statsArray)
	ps, _ := proto.Marshal(statsArray)

	printInfo(js, "json-stats-array")
	printInfo(ps, "protobuf-stats-array")
	fmt.Printf("\n")
}

func testAggrData(aggrData *AggrData) {
	fmt.Printf("Memory footprint of the final byte array for AggrData after a Marshal ---------\n")
	j, _ := json.Marshal(aggrData)
	p, _ := proto.Marshal(aggrData)

	printInfo(j, "json-aggr-data")
	printInfo(p, "protobuf-aggr-data")
	fmt.Printf("\n")

}

func printInfo(d []byte, ser string) {
	used := len(d)
	allocated := cap(d)
	fmt.Printf("Type: %s \t\tData size: %d \t\tTotal Allocated: %d \t\t Used/Allocated: %.2f%%\n", ser, used, allocated, percentUsed(used, allocated)*100)
}

func percentUsed(used, allocated int) float32 {
	return float32(used) / float32(allocated)
}

func marshalUtilProtobuf(b *testing.B, name string, pb proto.Message) {
	b.Run(name, func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := proto.Marshal(pb)
			_ = d
		}
	})

}

func marshalUtilJSON(b *testing.B, name string, v interface{}) {
	b.Run(name, func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := json.Marshal(v)
			_ = d
		}
	})
}

func unMarshalUtilProtobuf(b *testing.B, name string, messageType proto.Message, pb proto.Message) {
	data, _ := proto.Marshal(pb)
	b.Run(name, func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = proto.Unmarshal(data, messageType)
		}
	})

}

func unMarshalUtilJSON(b *testing.B, name string, v interface{}, desType interface{}) {
	data, _ := json.Marshal(v)
	b.Run(name, func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = json.Unmarshal(data, desType)
		}
	})

}

package go_benchmark

import (
	"encoding/json"
	"testing"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/golang/protobuf/proto"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go-benchmark/testobject"
)

func BenchmarkProtobuf_With2Fields(b *testing.B) {
	b.ReportAllocs()
	//obj := PbTestObject{"1","2","3","4","5","6","7","8","9","10"}
	obj := With2Fields{"1", "2"}
	data, _ := proto.Marshal(&obj)
	for i := 0; i < b.N; i++ {
		proto.Unmarshal(data, &obj)
	}
}

func BenchmarkJsoniterMarshal_PbTestObject(b *testing.B) {
	b.ReportAllocs()
	obj := PbTestObject{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	//obj := With2Fields{"1", "2"}
	data, _ := jsoniter.Marshal(&obj)
	//buf := &bytes.Buffer{}
	//stream := jsoniter.NewStream(buf, 4096)
	iter := jsoniter.NewIterator(jsoniter.ConfigDefault)
	for i := 0; i < b.N; i++ {
		iter.ResetBytes(data)
		iter.ReadVal(&obj)
	}
}

func BenchmarkJsoniterMarshal_With2Fields(b *testing.B) {
	b.ReportAllocs()
	obj := With2Fields{"1", "2"}
	data, _ := jsoniter.Marshal(&obj)
	for i := 0; i < b.N; i++ {
		json.Unmarshal(data, &obj)
	}
}

func BenchmarkThrift_TestObject(b *testing.B) {
	b.ReportAllocs()
	obj := testobject.NewThriftTestObject()
	obj.Field1 = "1"
	obj.Field2 = "2"
	obj.Field3 = "3"
	obj.Field4 = "4"
	obj.Field5 = "5"
	obj.Field6 = "6"
	obj.Field7 = "7"
	obj.Field8 = "8"
	obj.Field9 = "9"
	obj.Field10 = "10"
	buf := thrift.NewTMemoryBuffer()
	protocolFactory := &thrift.TCompactProtocolFactory{}
	protocol := protocolFactory.GetProtocol(buf)
	buf.Reset()
	obj.Write(protocol)
	data := buf.Bytes()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		buf.Write(data)
		obj.Read(protocol)
	}
}

package go_benchmark

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"github.com/buger/jsonparser"
	jsoniter "github.com/json-iterator/go"
)

var large = flag.Bool("large", false, "run test with large json file")

func TestJsonParserLargeFile(t *testing.T) {
	if !*large {
		t.Skip("Large test skipped. Use -large to run.")
	}
	file, _ := os.Open("/tmp/large-file.json")
	bytes, _ := ioutil.ReadAll(file)
	file.Close()
	total := 0
	jsonparser.ArrayEach(bytes, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		total++
	})
	if total != 11351 {
		t.Fatal(total)
	}
}

func TestJsoniterLargeFile(t *testing.T) {
	if !*large {
		t.Skip("Large test skipped. Use -large to run.")
	}
	for i := 0; i < 100; i++ {
		file, _ := os.Open("/tmp/large-file.json")
		iter := jsoniter.Parse(jsoniter.ConfigDefault, file, 4096)
		total := 0
		for iter.ReadArray() {
			iter.Skip()
			total++
		}
		file.Close()
		if total != 11351 {
			t.Fatal(total)
		}
	}
}

func BenchmarkJsonParserLargeFile(b *testing.B) {
	if !*large {
		b.Skip("Large benchmark skipped. Use -large to run.")
	}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		file, _ := os.Open("/tmp/large-file.json")
		bytes, _ := ioutil.ReadAll(file)
		file.Close()
		total := 0
		jsonparser.ArrayEach(bytes, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			total++
		})
	}
}

func BenchmarkEncodingJsonFile(b *testing.B) {
	if !*large {
		b.Skip("Large benchmark skipped. Use -large to run.")
	}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		file, _ := os.Open("/tmp/large-file.json")
		result := []struct{}{}
		decoder := json.NewDecoder(file)
		decoder.Decode(&result)
		file.Close()
	}
}

func BenchmarkJsoniterLargeFile(b *testing.B) {
	if !*large {
		b.Skip("Large benchmark skipped. Use -large to run.")
	}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		file, _ := os.Open("/tmp/large-file.json")
		iter := jsoniter.Parse(jsoniter.ConfigDefault, file, 4096)
		for iter.ReadArray() {
			iter.Skip()
		}
		file.Close()
	}
}

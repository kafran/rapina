package parsers

import (
	"reflect"
	"strings"
	"testing"
)

func Test_parseB3(t *testing.T) {

	const file = `012021010412NSLU11      010FII LOURDES CI  ER       R$  000000002840000000000284000000000027700000000002809000000000281900000000028029000000002819000168000000000000001381000000000038793560000000000000009999123100000010000000000000BRNSLUCTF008272
012021010412NVHO11      010FII NOVOHORICI  ER       R$  000000000154000000000015900000000001535000000000153700000000015400000000001536000000000154000092000000000000006200000000000009533490000000000000009999123100000010000000000000BRNVHOCTF003186
012021010412ONEF11      010FII THE ONE CI           R$  000000001478800000000148000000000014717000000001478900000000147360000000014735000000001478700035000000000000002546000000000037652878000000000000009999123100000010000000000000BRONEFCTF003200`

	want := []stockQuote{
		{Stock: "NSLU11", Date: "2021-01-04", Open: 284, High: 284, Low: 277, Close: 281.9, Volume: 387935.6},
		{Stock: "NVHO11", Date: "2021-01-04", Open: 15.4, High: 15.9, Low: 15.35, Close: 15.4, Volume: 95334.9},
		{Stock: "ONEF11", Date: "2021-01-04", Open: 147.88, High: 148, Low: 147.17, Close: 147.36, Volume: 376528.78},
	}

	for i, line := range strings.Split(file, "\n") {
		got, err := parseB3(line)
		if err != nil {
			t.Errorf("parseB3() error = %v", err)
			return
		}

		if err == nil && !reflect.DeepEqual(got, &want[i]) {
			t.Errorf("parseB3() got %+v, want %+v", got, &want[i])
		}

	}

}

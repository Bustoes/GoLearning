package split

import (
	"reflect"
	"testing"
)

// 测试

func TestSplit(t *testing.T) {
	type test struct {
		input string
		sep string
		want []string
	}

	tests := map[string]test {
		"chinese": test{input: "我爱你", sep: "爱", want: []string{"我", "你"}},
		"simple": test{input: "aa:ddd:qwwq", sep: ":", want: []string{"aa", "ddd", "qwwq"}},
		"multi-1": test{input: " how are you", sep: " ", want: []string{"","how", "are", "you"}},
		"multi-2": test{input: "howharehyou", sep: "h", want: []string{"","ow","are","you"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name: %s failed. want: %v, got: %v", name, tc.want, got)
			}
		})
	}
}


func BenchmarkSplit(b *testing.B) {
	// b.N不是固定的数

	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}


}



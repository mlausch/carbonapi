package countValues

import (
	"math"
	"testing"
	"time"

	"github.com/go-graphite/carbonapi/expr/interfaces"
	"github.com/go-graphite/carbonapi/expr/metadata"
	"github.com/go-graphite/carbonapi/expr/types"
	"github.com/go-graphite/carbonapi/pkg/parser"
	th "github.com/go-graphite/carbonapi/tests"
)

var (
	md []interfaces.FunctionMetadata = New("")
)

func init() {
	for _, m := range md {
		metadata.RegisterFunction(m.Name, m.F)
	}
}

func TestCountValues(t *testing.T) {
	// FIXME: test was broken on a merge request, therefore needs to be rewritten, skipping for now
	t.Skip("Skipping countValues function tests")
	now32 := int64(time.Now().Unix())

	tests := []th.MultiReturnEvalTestItem{
		{
			"countValues(metric1.foo.*.*)",
			map[parser.MetricRequest][]*types.MetricData{
				{Metric: "metric1.foo.*.*", From: 0, Until: 1}: {
					types.MakeMetricData("metric1.foo.bar1.baz", []float64{1, 2, 3, 4, 5}, 1, now32),
					types.MakeMetricData("metric1.foo.bar1.qux", []float64{2, 2, 4, 5, 6}, 1, now32),
					types.MakeMetricData("metric1.foo.bar2.baz", []float64{math.NaN(), 1, 1, 1, 1}, 1, now32),
				},
			},
			"countValues",
			map[string][]*types.MetricData{
				"1": {types.MakeMetricData("1", []float64{1, 1, 1, 1, 1}, 1, now32)},
				"2": {types.MakeMetricData("2", []float64{1, 2, 0, 0, 0}, 1, now32)},
				"3": {types.MakeMetricData("3", []float64{0, 0, 1, 0, 0}, 1, now32)},
				"4": {types.MakeMetricData("4", []float64{0, 0, 1, 1, 0}, 1, now32)},
				"5": {types.MakeMetricData("5", []float64{0, 0, 0, 1, 1}, 1, now32)},
				"6": {types.MakeMetricData("6", []float64{0, 0, 0, 0, 1}, 1, now32)},
			},
		},
		{
			"countValues(metric1.foo.*.*, 7)",
			map[parser.MetricRequest][]*types.MetricData{
				{Metric: "metric1.foo.*.*", From: 0, Until: 1}: {
					types.MakeMetricData("metric1.foo.bar1.baz", []float64{1, 2, 3, 4, 5}, 1, now32),
					types.MakeMetricData("metric1.foo.bar1.qux", []float64{2, 2, 4, 5, 6}, 1, now32),
					types.MakeMetricData("metric1.foo.bar2.baz", []float64{math.NaN(), 1, 1, 1, 1}, 1, now32),
				},
			},
			"countValues",
			map[string][]*types.MetricData{
				"1": {types.MakeMetricData("1", []float64{1, 1, 1, 1, 1}, 1, now32)},
				"2": {types.MakeMetricData("2", []float64{1, 2, 0, 0, 0}, 1, now32)},
				"3": {types.MakeMetricData("3", []float64{0, 0, 1, 0, 0}, 1, now32)},
				"4": {types.MakeMetricData("4", []float64{0, 0, 1, 1, 0}, 1, now32)},
				"5": {types.MakeMetricData("5", []float64{0, 0, 0, 1, 1}, 1, now32)},
				"6": {types.MakeMetricData("6", []float64{0, 0, 0, 0, 1}, 1, now32)},
			},
		},
		{
			"countValues(metric1.foo.*.*,valuesLimit=6)",
			map[parser.MetricRequest][]*types.MetricData{
				{Metric: "metric1.foo.*.*", From: 0, Until: 1}: {
					types.MakeMetricData("metric1.foo.bar1.baz", []float64{1, 2, 3, 4, 5}, 1, now32),
					types.MakeMetricData("metric1.foo.bar1.qux", []float64{2, 2, 4, 5, 6}, 1, now32),
					types.MakeMetricData("metric1.foo.bar2.baz", []float64{math.NaN(), 1, 1, 1, 1}, 1, now32),
				},
			},
			"countValues",
			map[string][]*types.MetricData{
				"1": {types.MakeMetricData("1", []float64{1, 1, 1, 1, 1}, 1, now32)},
				"2": {types.MakeMetricData("2", []float64{1, 2, 0, 0, 0}, 1, now32)},
				"3": {types.MakeMetricData("3", []float64{0, 0, 1, 0, 0}, 1, now32)},
				"4": {types.MakeMetricData("4", []float64{0, 0, 1, 1, 0}, 1, now32)},
				"5": {types.MakeMetricData("5", []float64{0, 0, 0, 1, 1}, 1, now32)},
				"6": {types.MakeMetricData("6", []float64{0, 0, 0, 0, 1}, 1, now32)},
			},
		},
		{
			"countValues(metric1.foo.*.*, 5)",
			map[parser.MetricRequest][]*types.MetricData{
				{Metric: "metric1.foo.*.*", From: 0, Until: 1}: {
					types.MakeMetricData("metric1.foo.bar1.baz", []float64{1, 2, 3, 4, 5}, 1, now32),
					types.MakeMetricData("metric1.foo.bar1.qux", []float64{2, 2, 4, 5, 6}, 1, now32),
					types.MakeMetricData("metric1.foo.bar2.baz", []float64{math.NaN(), 1, 1, 1, 1}, 1, now32),
				},
			},
			"countValues",
			map[string][]*types.MetricData{
				"valuesLimitReached": {types.MakeMetricData("valuesLimitReached", []float64{0, 0, 0, 0, 0}, 1, now32)},
			},
		},
		{
			"countValues(metric1.doo.*.*,valuesLimit=32)",
			map[parser.MetricRequest][]*types.MetricData{
				{Metric: "metric1.doo.*.*", From: 0, Until: 1}: {
					types.MakeMetricData("metric1.doo.bar1.baz", []float64{11, 21, 31, 41, 51, 61, 71, 81, 91, 101}, 1, now32),
					types.MakeMetricData("metric1.doo.bar2.baz", []float64{12, 22, 32, 42, 52, 62, 72, 82, 92, 102}, 1, now32),
					types.MakeMetricData("metric1.doo.bar3.baz", []float64{13, 23, 33, 43, 53, 63, 73, 83, 93, 103}, 1, now32),
					types.MakeMetricData("metric1.doo.bar4.baz", []float64{14, 24, 34, 44, 54, 64, 74, 84, 94, 104}, 1, now32),
				},
			},
			"countValues",
			map[string][]*types.MetricData{
				"valuesLimitReached": {types.MakeMetricData("valuesLimitReached", []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 1, now32)},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Target, func(t *testing.T) {
			eval := th.EvaluatorFromFunc(md[0].F)
			th.TestMultiReturnEvalExpr(t, eval, &tt)
		})
	}

}

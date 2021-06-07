package api

import (
	"net/http"
	"strconv"
	"tcp-simulation/pkg/service"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type AimdAPI struct {
	AimdService service.AimdService
}

func NewAimdAPI(p service.AimdService) AimdAPI {
	return AimdAPI{AimdService: p}
}

func (p AimdAPI) GenerateGrafic() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var increase_value int = 1
		var decrease_value int = 2
		increase_value_query := r.URL.Query().Get("increase")
		if increase_value_query != "" {
			result, _ := strconv.Atoi(increase_value_query)
			if result != 0 {
				increase_value = result
			}
		}
		decrease_value_query := r.URL.Query().Get("decrease")
		if decrease_value_query != "" {
			result, _ := strconv.Atoi(decrease_value_query)
			if result != 0 {
				decrease_value = result
			}
		}

		sended_package_per_step := p.AimdService.GenerateGrafic(increase_value, decrease_value)
		line := charts.NewLine()
		line.SetGlobalOptions(
			charts.WithTitleOpts(opts.Title{Title: "basic line example", Subtitle: "This is the subtitle."}),
		)

		chart := make([]opts.LineData, 0)
		chart_axis := make([]string, 0)

		for i, v := range sended_package_per_step {
			chart_axis = append(chart_axis, strconv.Itoa(i))

			chart = append(chart, opts.LineData{Value: v})

		}

		line.SetXAxis(chart_axis).
			AddSeries("Category A", chart)
		line.Render(w)

		RespondWithJSON(w, http.StatusOK, 0)
	}
}

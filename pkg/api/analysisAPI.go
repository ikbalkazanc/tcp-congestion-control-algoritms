package api

import (
	"net/http"
	"strconv"
	"tcp-simulation/pkg/service"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type AnalysisAPI struct {
	AimdService service.AimdService
	SlowStart   service.SlowStartService
}

func NewAnalysisAPI(aimd service.AimdService, slowstart service.SlowStartService) AnalysisAPI {
	return AnalysisAPI{AimdService: aimd, SlowStart: slowstart}
}

func (p AnalysisAPI) GenerateGrafic() http.HandlerFunc {
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

		sended_package_per_step_as_aimd := p.AimdService.GenerateGrafic(increase_value, decrease_value)
		sended_package_per_step_as_slowstart := p.SlowStart.GenerateGrafic(false)
		line := charts.NewLine()
		line.SetGlobalOptions(
			charts.WithTitleOpts(opts.Title{Title: "basic line example", Subtitle: "This is the subtitle."}),
		)

		chart_aimd := make([]opts.LineData, 0)
		chart_slowstart := make([]opts.LineData, 0)
		chart_axis := make([]string, 0)

		for i, v := range sended_package_per_step_as_aimd {
			if len(sended_package_per_step_as_aimd) >= len(sended_package_per_step_as_slowstart) {
				chart_axis = append(chart_axis, strconv.Itoa(i))
			}

			chart_aimd = append(chart_aimd, opts.LineData{Value: v})

		}

		for i, v := range sended_package_per_step_as_slowstart {
			if len(sended_package_per_step_as_aimd) < len(sended_package_per_step_as_slowstart) {
				chart_axis = append(chart_axis, strconv.Itoa(i))
			}
			chart_slowstart = append(chart_slowstart, opts.LineData{Value: v})

		}

		line.SetXAxis(chart_axis).
			AddSeries("aimd", chart_aimd).
			AddSeries("slow", chart_slowstart)
		line.Render(w)

		RespondWithJSON(w, http.StatusOK, 0)
	}
}

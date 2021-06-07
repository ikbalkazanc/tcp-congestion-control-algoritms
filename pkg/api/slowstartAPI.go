package api

import (
	"net/http"
	"strconv"
	"tcp-simulation/pkg/service"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type SlowStartAPI struct {
	SlowStartService service.SlowStartService
}

func NewSlowStartAPI(p service.SlowStartService) SlowStartAPI {
	return SlowStartAPI{SlowStartService: p}
}

func (p SlowStartAPI) GenerateGrafic() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		sended_package_per_step := p.SlowStartService.GenerateGrafic(false)
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

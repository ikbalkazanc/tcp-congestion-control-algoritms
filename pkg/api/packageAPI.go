package api

import (
	"net/http"
	"strconv"
	"tcp-simulation/pkg/service"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type PackageAPI struct {
	PackageService service.PackageService
}

func NewPackageAPI(p service.PackageService) PackageAPI {
	return PackageAPI{PackageService: p}
}

func (p PackageAPI) GetAllPackage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		size, loss := p.PackageService.GetLostPackages()

		line := charts.NewLine()
		line.SetGlobalOptions(
			charts.WithTitleOpts(opts.Title{Title: "basic line example", Subtitle: "This is the subtitle."}),
		)

		chart := make([]opts.LineData, 0)
		chart_axis := make([]string, 0)

		for i := 0; i < size; i++ {
			chart_axis = append(chart_axis, strconv.Itoa(i))
			if find(i, loss) != -1 {
				chart = append(chart, opts.LineData{Value: 1})
			} else {
				chart = append(chart, opts.LineData{Value: 0})
			}
		}

		line.SetXAxis(chart_axis).
			AddSeries("Category A", chart)
		line.Render(w)

		RespondWithJSON(w, http.StatusOK, loss)
	}
}

func (p PackageAPI) GenerateLossPackages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var size_int int = 100
		var rate_float float64 = 0.1
		var incrate_float float64 = 0.0
		size := r.URL.Query().Get("size")
		if size != "" {
			result, _ := strconv.Atoi(size)
			if result != 0 {
				size_int = result
			}
		}
		rate := r.URL.Query().Get("rate")
		if rate != "" {
			result, _ := strconv.ParseFloat(rate, 64)
			if result != 0.0 {
				rate_float = result
			}
		}
		incrate := r.URL.Query().Get("incrate")
		if incrate != "" {
			incrate_float, _ = strconv.ParseFloat(incrate, 64)
		}
		_, loss := p.PackageService.Generate(size_int, rate_float, incrate_float)
		RespondWithJSON(w, http.StatusOK,
			GenerateResponse{Length: size_int, Rate: rate_float, Increment_decrement_rate: incrate_float, Lost_packages: loss})
	}
}

func find(data int, s *[]int) int {
	for _, v := range *s {
		if v == data {
			return data
		}
	}
	return -1
}

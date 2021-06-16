package service

type SlowStartService struct {
}

func NewSlowStartService() SlowStartService {
	return SlowStartService{}
}

func (p *SlowStartService) GenerateGrafic(isActiveReno bool) []int {
	cwnd := 1
	ssthreshold := 9999999 //inf
	a, b := getLossPackages()
	file_length := a
	loss_package_list := *b
	sended_package_per_step := []int{}
	index := 0

	for true {
		sended_package_per_step = append(sended_package_per_step, cwnd)
		unsent_packages, result := sendRange(index, cwnd, file_length, &loss_package_list)

		//all files sended
		if len(unsent_packages) == 0 && !result {
			break
			//failed send process and have unsent packages
		} else if len(unsent_packages) > 0 && !result {
			//set threshold value
			ssthreshold = cwnd / 2
			//set cwnd as reno or tahoe value
			if isActiveReno {
				cwnd = ssthreshold
			} else {
				cwnd = 1
			}
			//index = index - len(unsent_packages)
		} else {
			//incresase linear
			if ssthreshold <= cwnd {
				cwnd++
				//increase multi
			} else {
				cwnd = cwnd * 2
			}
			index = index + cwnd
		}
	}
	return sended_package_per_step
}

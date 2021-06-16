package service

//"fmt"

type AimdService struct {
}

func NewAimdService() AimdService {
	return AimdService{}
}

func (p *AimdService) GenerateGrafic(increase_value int, decrease_value int) []int {
	number_of_packets_wll_send_in_one_step := 1
	//increase_value                    := 1
	//decrease_value                    := 2
	a, b := getLossPackages()
	file_length := a
	loss_package_list := *b
	sended_package_per_step := []int{}
	index := 0

	for true {
		sended_package_per_step = append(sended_package_per_step, number_of_packets_wll_send_in_one_step)
		unsent_packages, result := sendRange(index, number_of_packets_wll_send_in_one_step, file_length, &loss_package_list)
		//increase
		if result {
			number_of_packets_wll_send_in_one_step = number_of_packets_wll_send_in_one_step + increase_value
			index = index + number_of_packets_wll_send_in_one_step
		}
		//fmt.Print("Sending ")
		//fmt.Print(number_of_packets_wll_send_in_one_step)
		//fmt.Println(" piece packet.")

		//all files sended
		if len(unsent_packages) == 0 && !result {
			break
			//failed send process and have unsent packages
		} else if len(unsent_packages) > 0 && !result {
			//require send againt failded to index interval
			index = index - number_of_packets_wll_send_in_one_step
			//decrease
			number_of_packets_wll_send_in_one_step = number_of_packets_wll_send_in_one_step / decrease_value
			//index = index - len(unsent_packages)

		}

	}
	return sended_package_per_step
}

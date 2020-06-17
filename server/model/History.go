package model

type History struct {
	Items []string `json:"items,omitempty"`
}

func Contains(history *History, i string) bool { //Change for hashset or better
	for _,item := range history.Items {
		if item==i {
			return true
		}
	}
	return false
}
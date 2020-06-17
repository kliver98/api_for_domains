package main

type History struct {
	Items []string `json:"items,omitempty"`
}

func Contains(items *[]string, i string) bool { //Change for hashset or better
	for _,item := range items {
		if item==i {
			return true
		}
	}
	return false
}
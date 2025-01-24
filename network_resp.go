package main

type Result_arr struct {
	Name string `json:name`
	Url string `json:url`
}

type Response struct {
	Count int `json:count`
	Next string `json:next`
	Previous string `json:prev`
	Results []Result_arr `json:results`
}

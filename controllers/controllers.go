package controllers

//大伙都能经常用到的玩意都在这

type Basicinfo struct {
	CluserId  string `json:"clusterId"`
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

package MODELS

type BLOCKS struct {
	Id          int    `json:"id"`
	NameBlock   string `json:"nameBlock"`
	Address     string `json:"address"`
	Description string `json:"description"`
	IdOwner     int    `json:"idOwner"`
}

type IDBLOCKS struct {
	BlocksId []int `json:"blocksId"`
}

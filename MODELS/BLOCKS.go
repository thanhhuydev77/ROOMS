package MODELS

type BLOCKS struct {
	Id          int
	NameBlock   string
	Address     string
	Description string
	IdOwner     int
}

type IDBLOCKS struct {
	BlocksId     []int
}

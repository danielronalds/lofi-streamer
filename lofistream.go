package main

var id int

type LofiStream struct {
	Id      int
	Name    string
	VideoId string
}

func newLofiStream(name string, videoid string) LofiStream {
	id += 1
	return LofiStream{
		Id:      id,
		Name:    name,
		VideoId: videoid,
	}
}

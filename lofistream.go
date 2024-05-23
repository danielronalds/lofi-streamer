package main

type LofiStream struct {
	Name    string
	VideoId string
}

func newLofiStream(name string, videoid string) LofiStream {
	return LofiStream{
		Name:    name,
		VideoId: videoid,
	}
}

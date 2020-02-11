package model

// Webページから取得した情報
type EditorialInfo struct {
	Url     string
	Date    string
	MediaId int
	Title   string
	Body    string
}

func (ei EditorialInfo) Count() int {
	editorial := Editorial{}
	editorial.Init()
	return editorial.Count(ei)
}

func (ei EditorialInfo) Insert() {
	editorial := Editorial{}
	editorial.Init()
	editorial.Insert(ei)
}
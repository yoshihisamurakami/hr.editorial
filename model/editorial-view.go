package model

type EditorialView struct {
	Editorial
	FormattedDate string
	FormattedBody []string
	MediaName     string
}

func (ev *EditorialView) Init(e Editorial) {
	ev.ID = e.ID
	ev.Date = e.Date
	ev.MediaId = e.MediaId
	ev.Title = e.Title
	ev.Body = e.Body
	ev.Url = e.Url
	ev.ensureForView()
}

func (e *EditorialView) ensureForView() {
	e.FormattedDate = e.Date[:10]
	e.MediaName = e.getMediaName()
}

func (e EditorialView) getMediaName() string {
	mediaNames := []string{"", "毎日新聞", "東京新聞", "朝日新聞"}
	return mediaNames[e.MediaId]
}

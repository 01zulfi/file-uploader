package controllers

type indexPageFile struct {
	Filename string
	Filelink string
}

type indexPageData struct {
	TitleText string
	Files     []indexPageFile
}

type uploadPageData struct {
	Filename string
}

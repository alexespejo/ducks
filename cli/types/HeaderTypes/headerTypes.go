package headerTypes

type SubSubHeader struct {
	H3 string `json:"h3"`
}

type SubHeader struct {
	H2       string         `json:"h2"`
	Children []SubSubHeader `json:"children"`
}

type Header struct {
	H1       string      `json:"h1"`
	Children []SubHeader `json:"children"`
}

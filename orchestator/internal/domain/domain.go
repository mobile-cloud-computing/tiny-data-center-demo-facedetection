package domain

type ImageOuput struct {
	ImagePath    string  `json:"imagePath"`
	GeneralImage bool    `json:"generalImage"`
	Confidence   float64 `json:"confidence"`
}

package models

type ControllersImage struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Keywords    string `json:"keywords,omitempty"`
	Author      string `json:"author,omitempty"`
	Creator     string `json:"creator,omitempty"`
	CaptureDate string `json:"date,omitempty"`
	Filename    string `json:"filename,omitempty"`
}

type Image struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Keywords    string `json:"keywords,omitempty"`
	Author      string `json:"author,omitempty"`
	Creator     string `json:"creator,omitempty"`
	CaptureDate string `json:"date,omitempty"`
}

type RepositoryImage struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Keywords    string `json:"keywords,omitempty"`
	Author      string `json:"author,omitempty"`
	Creator     string `json:"creator,omitempty"`
	CaptureDate string `json:"creation_date,omitempty"`
	StorageDate string `json:"storage_date,omitempty"`
	Filename    string `json:"filename,omitempty"`
}

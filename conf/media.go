package conf

// Media defines all of the available media types with their properties
type Media struct {
	Movies MediaInfo `yaml:"movies"`
	Music  MediaInfo `yaml:"music"`
	Tv     MediaInfo `yaml:"tv"`
	Books  MediaInfo `yaml:"books"`
}

// MediaInfo defines all the information necessary for for getting the information of a certain media type
type MediaInfo struct {
	Path string `yaml:"path"`
	API  string `yaml:"api"`
	Key  string `yaml:"key"`
	JSON string `yaml:"json"`
	Type string `yaml:"type"`
}

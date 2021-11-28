package config

type Config struct {
	DBType     string `json:"dbtype"`
	DBFile     string `json:"dbfile"`
	Listen     string `json:"listen"`
	APIBase    string `json:"apibase"`
	APIVersion string `json:"apiversion"`
}

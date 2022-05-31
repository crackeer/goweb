package define

// TemplateConfig
type TemplateConfig struct {
	Ext      string `json:"ext"`
	Dir      string `json:"dir"`
	Skeleton string `json:"skeleton"`
}

// Config
type Config struct {
	Port           int64          `json:"port"`
	Database       string         `json:"database"`
	BoltDB         string         `json:"bolt_db"`
	TemplateConfig TemplateConfig `json:"template"`
	Domain         string         `json:"domain"`

	PasswordMD5 string `json:"password_md5"`
	PageSize    int64  `json:"page_size"`
	//Key         string `json:"key"`
}

type PageConf struct {
	Title string `json:"title"`
	TPL   string `json:"tpl"`
	Type  string `json:"type"`
}

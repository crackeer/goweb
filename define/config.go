package define

// Config
type AppConfig struct {
	Port        int64     `yaml:"port"`
	Resource    *Resource `yaml:"resource"`
	PasswordMD5 string    `yaml:"password_md5"`
	PageSize    int64     `yaml:"page_size"`
}

type Resource struct {
	BoltDBFile       string `json:"bolt_db"`
	SqliteDBFile     string `yaml:"sqlite_db_file"`
	PageDir          string `yaml:"page_dir"`
	PageConfDir      string `yaml:"page_conf_dir"`
	APIConfDir       string `yaml:"api_conf_dir"`
	PublicDir        string `yaml:"public_dir"`
	DefaultFrameFile string `yaml:"default_frame_file"`
}

type PageConf struct {
	Title string `json:"title"`
	TPL   string `json:"tpl"`
	Type  string `json:"type"`
}

package define

// Config
type Config struct {
	Port                 int64  `json:"port"`
	Sqlite3DatabaseFile  string `json:"sqlite3_database_file"`
	TemplateFileExt      string `json:"template_file_ext"`
	TemplateFileDir      string `json:"template_file_dir"`
	TemplateSkeletonFile string `json:"template_skeleton_file"`

	PasswordMD5 string `json:"password_md5"`
	Key         string `json:"key"`

	CodeLanguages []map[string]interface{} `json:"code_languages"`

	Page map[string]PageConf `json:"page"`
}

type PageConf struct {
	Title string `json:"title"`
	TPL   string `json:"tpl"`
	Type  string `json:"type"`
}

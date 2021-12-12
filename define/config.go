package define

// Config
type Config struct {
	Port                 int64  `json:"port"`
	Sqlite3DatabaseFile  string `json:"sqlite3_database_file"`
	TemplateFileExt      string `json:"template_file_ext"`
	TemplateFileDir      string `json:"template_file_dir"`
	TemplateSkeletonFile string `json:"template_skeleton_file"`
}

package define

import "github.com/crackeer/gopkg/api"

// AppConfig Config
type AppConfig struct {
	Port                int64             `yaml:"port"`
	Env                 string            `yaml:"env"`
	ResourceDir         string            `yaml:"resource_dir"`
	APIConfDir          string            `yaml:"api_conf_dir"`
	PageConfDir         string            `yaml:"page_conf_dir"`
	PublicFileExtension map[string]string `yaml:"public_file_extension"`
	DefaultFrameFile    string            `yaml:"default_frame_file"`
}

// PageConfig YamlAPIConfig
type PageConfig struct {
	DataAPI       string                 `yaml:"data_api"`
	DataAPIMesh   [][]*api.RequestItem   `yaml:"data_api_mesh"`
	Type          string                 `yaml:"type"`
	Extension     map[string]interface{} `yaml:"extension"`
	DefaultParams map[string]interface{} `yaml:"default_params"`
	Title         string                 `yaml:"title"`
	FrameFile     string                 `yaml:"frame_file"`
	ContentFile   string                 `yaml:"content_file"`
}

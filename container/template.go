package container

import (
	"strings"
)

// GetFullTemplatePath
//  @param name
//  @return string
func GetFullTemplatePath(name string) string {
	return strings.Join([]string{
		config.TemplateConfig.Dir, name, config.TemplateConfig.Ext,
	}, "")
}

// GetFullTemplatePath
//  @param name
//  @return string
func GetSkeletonTemplatePath() string {
	return GetFullTemplatePath(config.TemplateConfig.Skeleton)
}

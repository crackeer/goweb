package container

import (
	"strings"
)

// GetFullTemplatePath
//  @param name
//  @return string
func GetFullTemplatePath(name string) string {
	return strings.Join([]string{
		config.TemplateFileDir, name, config.TemplateFileExt,
	}, "")
}

// GetFullTemplatePath
//  @param name
//  @return string
func GetSkeletionTemplatePath() string {
	return strings.Join([]string{
		config.TemplateFileDir, config.TemplateSkeletonFile, config.TemplateFileExt,
	}, "")
}

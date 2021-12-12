package common

import (
	"flag"
	"sync"
)

var (
	cfgOnce sync.Once
)

var DbFile *string
var Port *string
var ImageDir *string

func GetConfigInstance() {
	cfgOnce.Do(func() {
		DbFile = flag.String("db", "./data.db", "database file")
		ImageDir = flag.String("imagedir", "./image", "image dir")
		Port = flag.String("port", "8888", "port")
		flag.Parse()

	})
	return
}

var ExportMdDir *string
var BaseImageUrl *string

func GetToolConfigInstance() {
	cfgOnce.Do(func() {
		DbFile = flag.String("db", "./data.db", "database file")
		ExportMdDir = flag.String("export-md-dir", "./outputs", "export md dir dir")
		BaseImageUrl = flag.String("base-url", "/", "base image url")
		flag.Parse()

	})
	return
}

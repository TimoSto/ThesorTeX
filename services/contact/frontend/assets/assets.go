package assets

import (
	"embed"
	"io/fs"
)

//go:embed dist/*
var dir embed.FS

var Assets, _ = fs.Sub(dir, "dist")

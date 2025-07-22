package version

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
	"text/template"
)

var (
	AppPlatform = fmt.Sprintf("%s/%s %s", runtime.GOOS, runtime.GOARCH, runtime.Version())
	AppName     string
	AppCommitId string
	AppVersion  string
	AppBuildAt  string
	AppBranch   string
	versionTmpl = `
Application information for {{.appName}}:
    version:	{{.version}}
    branch:		{{.branch}}
    revision:	{{.commitId}}
    buildAt:	{{.buildAt}}
    platform:	{{.platform}}
`
	MetaVersionMap = map[string]string{
		"appName":  AppName,
		"version":  AppVersion,
		"commitId": AppCommitId,
		"branch":   AppBranch,
		"buildAt":  AppBuildAt,
		"platform": AppPlatform,
	}
)

// Version 输出应用版本信息
func Version() string {
	t := template.Must(template.New("version").Parse(versionTmpl))

	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, "version", MetaVersionMap); err != nil {
		panic(err)
	}
	return strings.TrimSpace(buf.String())
}

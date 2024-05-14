package main

import (
	"flag"
	"github.com/Touch/cmd/api/app"
	"github.com/Touch/config"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	confFile = flag.String("f", "conf/config.yaml", "set config file directory")
	mode     = flag.String("mode", "api", "run mode,such as api, cron, queue")
)

func getBinAbPath() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 定义项目根目录
func GetProjectRoot() string {
	dir := getBinAbPath()
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(dir, tmpDir) {
		dir = getCallerAbPath()
	}
	return path.Dir(dir)
}

func getCallerAbPath() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

func main() {
	// 解析配置文件
	flag.Parse()
	config.InitConfig(*confFile)
	app.Run()
}

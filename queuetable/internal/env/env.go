// We put all the environment variables in a single file so that:
// 1. We know which environment variables to add for deployment,
//    without having to search the whoel codebase
// 2. We don't keep reading from all the environment variables

package env

import (
	"os"
	"path"

	"github.com/joho/godotenv"
)

var EntDriver = ""
var EntDatasourceName = ""

func init() {
	err := load()
	if err != nil {
		panic(err)
	}

	EntDriver = os.Getenv("ENT_DRIVER")
	EntDatasourceName = os.Getenv("ENT_DATASOURCE_NAME")
}

func Get(env, defaultvalue string) string {
	value := os.Getenv(env)
	if value == "" {
		return defaultvalue
	}

	return value
}

func load() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	tosearch := []string{}

	if appenv := os.Getenv("APP_ENV"); appenv != "" {
		tosearch = append(tosearch, ".env."+appenv)
	}

	tosearch = append(tosearch, ".env")

	toload := []string{}

	for dir != "/" {
		for _, dotenvfile := range tosearch {

			filename := path.Join(dir, dotenvfile)
			_, err := os.Stat(filename)
			if os.IsNotExist(err) {
				continue
			}

			toload = append(toload, filename)
		}

		dir = path.Dir(dir)
	}

	if len(toload) == 0 {
		return nil
	}

	err = godotenv.Load(toload...)
	return err
}

package b2g

import (
	"fmt"
	"log"
	"os"
)

func ValidateEnv() {
	missing := make([]string, 0)
	envVars := []string{
		"GITHUB_TOKEN",
	}

	for _, v := range envVars {
		_, present := os.LookupEnv(v)
		if !present {
			missing = append(missing, v)
		}
	}
	if len(missing) != 0 {
		fmt.Printf("missing env vars: %v\n", missing)
		os.Exit(1)
	}
}

func CheckIfError(err error) {
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

func LogFatal(f string, err error) {
	log.Fatal(f, " error: ", err)
}

func chdirRepo(r string) {
	wd := "/tmp/" + r
	os.Chdir(wd)
}

func getEnvValue(e string) string {
	return os.Getenv(e)
}

// func createFolder(target string) {
// 	dir := getEnvValue("WORKING_DIR") + "/" + target
// 	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
// 		err := os.Mkdir(dir, os.ModePerm)
// 		CheckIfError(err)
// 	}
// }

// func ListDirectories(target string) []string {
// 	var directories []string
// 	dir := getEnvValue("WORKING_DIR") + "/" + target
// 	dirs, err := ioutil.ReadDir(dir)
// 	CheckIfError(err)
// 	for _, d := range dirs {
// 		if d.IsDir() {
// 			if _, err := os.Stat(dir + "/" + d.Name() + "/.git"); !os.IsNotExist(err) {
// 				directories = append(directories, d.Name())
// 			}
// 		}
// 	}
// 	return directories
// }

package config

import (
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerHost		string `required:"true" split_words:"true"`
	ServerPort		int `required:"true" split_words:"true"`
}

var (
	once sync.Once
	Cfg  Config
)

func Environments() Config {
	once.Do(func() {
		//If you use env file, uncomment this section and the functions:
		//if err := setEnvsFromFile(".env"); err != nil {
		//	log.Panicf("error reading local env file %s", err.Error())
		//	return
		//}

		if err := envconfig.Process("", &Cfg); err != nil {
			log.Panicf("Error parsing environment vars %#v", err)
		}
	})

	return Cfg
}
/*
func setEnvsFromFile(fileName string) error {
	root, err := rootDir()
	if err != nil {
		return err
	}
	file, err := os.Open(fmt.Sprintf("%s/%s", root, fileName))
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		return fmt.Errorf("env file is empty")
	}

	for scanner.Scan() {
		w := strings.Split(scanner.Text(), "=")

		if err = os.Setenv(w[0], w[1]); err != nil {
			return err
		}
	}

	return nil
}

func rootDir() (string, error) {
	dir, err := os.Getwd()
	return dir, err
}
*/

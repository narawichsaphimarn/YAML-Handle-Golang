package goyamlenv

import (
	"strings"

	"github.com/spf13/viper"
)

/**
 * Method defualt
 * *This method is start of method.
 * ?NAME valiable is `application`
 * ?TYPE valiable is `yaml`
 * ?PATH valiable is `$HOME/resourses`
 * ?Format valiable is {`.`, `_`}
 * !If error be have return error.
 */
func Defualt() error {
	viper.SetConfigName(NAME)
	viper.SetConfigType(TYPE)
	viper.AddConfigPath(PATH)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return nil
}

/**
 * Method NewWithBasePath
 * *This method is start of method.
 * @param BASE_PATH valiable is type string.
 * ?NAME valiable is `application`
 * ?TYPE valiable is `yaml`
 * ?PATH valiable is `$HOME/resourses`
 * ?Format valiable is {`.`, `_`}
 * !If error be have return error.
 */
func NewWithBasePath(_basePath string) error {
	viper.SetConfigName(NAME)
	viper.SetConfigType(TYPE)
	viper.AddConfigPath(_basePath)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return nil
}

/**
 * Method new
 * *This method is start of method.
 * @param NAME valiable is type string.
 * @param TYPE valiable is type string.
 * @param PATH valiable is type string.
 * ?Format valiable is {`.`, `_`}
 * !If error be have return error.
 */
func New(_path string, _type string, _name string) error {
	viper.SetConfigName(_name)
	viper.SetConfigType(_type)
	viper.AddConfigPath(_path)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return nil
}

/**
 * Method load configs.
 * *This method is load env from viper to struct.
 * @param VALUE valiable is type pointer interface.
 * !If error be have return error.
 */
func LoadConfig(value any) error {
	if err := viper.Unmarshal(value); err != nil {
		return err
	}
	return nil
}

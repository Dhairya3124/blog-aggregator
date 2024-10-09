package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DbURL string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config,error) {
	homeDir, err := os.UserHomeDir()
	if err!=nil{
		return Config{},err
	}
	 filePath := homeDir + "/" + configFileName
	 fileContent,err := os.ReadFile(filePath)
	 if err != nil {
		return Config{}, err
	 }
	 config:=Config{}
	 if err:=json.Unmarshal(fileContent,&config);err!=nil{
		return Config{},err

	 }
	 return config,err


}
func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName

	homeDir, err := os.UserHomeDir()
	if err != nil {
			return err
	}

	filePath := homeDir + "/" + configFileName
	fileContent, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
			return err
	}

	if err := os.WriteFile(filePath, fileContent, 0644); err != nil {
			return err
	}

	return nil
}
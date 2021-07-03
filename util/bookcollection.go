package util

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
)

type Book struct {
	ID			int `mapstructure:"id"`
	Name    	string `mapstructure:"name"`
	Author 		string `mapstructure:"author"`
	Labels      string `mapstructure:"labels"`
	Quantity 	int `mapstructure:"quantity"`
}

type Collection struct {
	BookList []Book `mapstructure:"booklist"`
}

var BookCollection *Collection

func (i Collection) GetCount() int {
	booklist := i.BookList
	count := 0
	for  _, _ = range booklist {
		count +=1
	}

	return count
}

func (i Collection) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(i)
	return bytes, err
}

func LoadCollection(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("collection")
	viper.SetConfigType("json")

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error in reading config ", err)
		return
	}

	err = viper.Unmarshal(&BookCollection)
	if err != nil {
		fmt.Println("Unable to decode into struct, %v", err)
	}

	return
}




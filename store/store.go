package store

import (
	"encoding/gob"
	"github.com/n1try/telegram-middleman-bot/config"
	"github.com/n1try/telegram-middleman-bot/model"
	"log"
	"os"
)

var store map[string]interface{}

func init() {
	//gob.Register(model.StoreObject{})
	//gob.Register(model.StoreMessageObject{})

	// Backwards compatibility
	gob.RegisterName("main.StoreObject", model.StoreObject{})
	gob.RegisterName("main.StoreMessageObject", model.StoreMessageObject{})

	initEmpty()
}

func initEmpty() {
	store = make(map[string]interface{})
}

func Read(filePath string) {
	log.Println("Loading store.")
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		log.Println("Could not read store from file. Initializing empty one.")
		initEmpty()
		return
	}
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&store)
	if err != nil {
		log.Println(err)
	}
}

func Flush(filePath string) {
	log.Println("Flushing store.")
	file, err := os.Create(filePath)
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(&store)
	if err != nil {
		log.Fatalln(err)
	}
}

func Automigrate() {
	if _, ok := store[config.KeyMessages]; ok {
		Delete(config.KeyMessages)
	}
}

func Get(key string) interface{} {
	return store[key]
}

func Put(key string, value interface{}) {
	store[key] = value
}

func Delete(key string) {
	delete(store, key)
}

func GetMap() map[string]interface{} {
	return store
}

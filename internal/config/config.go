package config

import (
	"os"
)

var ModeDev bool = os.Getenv("MODE_DEV") == "True"
var (
	ChatId         string = os.Getenv("CHAT_ID")
	ApiKey         string = os.Getenv("API_KEY")
	ApiUrl         string = "https://api.telegram.org"
	BotAddress     string = os.Getenv("BOT_ADDRESS")
	BotBindAddress string = os.Getenv("BOT_BIND_ADDRESS")
	BotId          string = os.Getenv("BOT_ID")
	BotKey         string = os.Getenv("BOT_KEY")
	BotPort        string = os.Getenv("BOT_PORT")
	CertPath       string = os.Getenv("BOT_CERT_PATH")
	KeyPath        string = os.Getenv("BOT_KEY_PATH")
	LogPath        string = os.Getenv("LOG_PATH")
	RedisAddress   string = os.Getenv("REDIS_ADDRESS")
	MongoAddress   string = os.Getenv("MONGO_ADDRESS")
	MongoUser      string = os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	MongoPassword  string = os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	MongoDatabase  string = os.Getenv("MONGO_DATABASE")
)

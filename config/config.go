package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	configFile = "./config.yaml"
)

type Server struct {
	Enable   bool
	Port     int
	TLS      bool
	CertPath string
	KeyPath  string
}

type Config struct {
	HttpServer *Server
	GrpcServer *Server
}

type ConfigFlat struct {
	Http_Server_Enable      bool   `yaml:"http_server_enable" env:"http_server_enable"`
	Http_Server_Port        int    `yaml:"http_server_port" env:"http_server_port"`
	Http_Server_ServiceName string `yaml:"http_server_service_name" env:"http_server_service_name"`

	Grpc_Server_Enable      bool   `yaml:"grpc_server_enable" env:"grpc_server_enable"`
	Grpc_Server_Port        int    `yaml:"grpc_server_port" env:"grpc_server_port"`
	Groc_Server_ServiceName string `yaml:"grpc_server_service_name" env:"grpc_server_service_name"`
	Grpc_TLS                bool   `yaml:"grpc_tls" env:"grpc_tls"`
	Grpc_CertPath           string `yaml:"grpc_cert_path" env:"grpc_cert_path"`
	Grpc_KeyPath            string `yaml:"grpc_key_path" env:"grpc_key_path"`
}

func ReadConfig() *Config {
	// read config
	var readConfig ConfigFlat
	log.Print("[INFO] reading config file : ", configFile)
	err := cleanenv.ReadConfig(configFile, &readConfig)
	if err != nil {
		log.Printf("[ERROR] reading config file %s : %s", configFile, err.Error())
		log.Printf("[INFO] reading config from env")
		err = cleanenv.ReadEnv(&readConfig)
		if err != nil {
			log.Fatal("[ERROR] reading config from env : ", err.Error())
		}
	}
	log.Println("[INFO] config file read successfully")

	validate(&readConfig)

	return &Config{
		HttpServer: &Server{
			Enable: readConfig.Http_Server_Enable,
			Port:   readConfig.Http_Server_Port,
		},
		GrpcServer: &Server{
			Enable:   readConfig.Grpc_Server_Enable,
			Port:     readConfig.Grpc_Server_Port,
			TLS:      readConfig.Grpc_TLS,
			CertPath: readConfig.Grpc_CertPath,
			KeyPath:  readConfig.Grpc_KeyPath,
		},
	}
}

func validate(cfg *ConfigFlat) {
	if cfg == nil {
		log.Fatal("[ERROR] config is nil")
	}

	// http server checks
	if cfg.Http_Server_Enable {
		log.Print("[INFO] http server is enabled in config")
		if cfg.Http_Server_Port <= 0 {
			log.Fatal("[ERROR] server_port is invalid")
		}
		if cfg.Http_Server_ServiceName == "" {
			log.Fatal("[ERROR] server_service_name is not set")
		}
	}

	// grpc server checks
	if cfg.Grpc_Server_Enable {
		log.Print("[INFO] grpc server is enabled in config")
		if cfg.Grpc_Server_Port <= 0 {
			log.Fatal("[ERROR] grpc_server_port is invalid")
		}
		if cfg.Groc_Server_ServiceName == "" {
			log.Fatal("[ERROR] grpc_server_service_name is not set")
		}
		if cfg.Grpc_TLS {
			log.Print("[INFO] tls is enabled in grpc server")
			if cfg.Grpc_CertPath == "" {
				log.Fatal("[ERROR] grpc_cert_path is not set")
			}
			if cfg.Grpc_KeyPath == "" {
				log.Fatal("[ERROR] grpc_key_path is not set")
			}
		}
	}
}

package main

import (
	"net/http"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/config")
	viper.SetDefault("port", ":8080")
	viper.SetDefault("mode", "development")
	viper.ReadInConfig()
}
func main() {
	http.Handle("/", &RootHandler{})
	http.Handle("/health", &HealthHandler{})
	http.ListenAndServe(viper.GetString("port"), nil)
}

type RootHandler struct {
}

func (h *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

type HealthHandler struct {
}

func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

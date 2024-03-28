package main

import (
	"context"
	"fmt"
	"os"

	pkgconfig "github.com/yomorun/yomo/pkg/config"
	"github.com/yomorun/yomo/pkg/log"

	"github.com/yomorun/yomo"
	"github.com/yomorun/yomo/core/router"
	"github.com/yomorun/yomo/pkg/bridge/ai"
	"github.com/yomorun/yomo/pkg/bridge/ai/provider/azopenai"
	"github.com/yomorun/yomo/pkg/bridge/ai/provider/cfazure"
	"github.com/yomorun/yomo/pkg/bridge/ai/provider/gemini"
	"github.com/yomorun/yomo/pkg/bridge/ai/provider/openai"
)

func main() {
	config := "../../Desktop/yomo.config.ai.bridge.yaml"

	// config
	conf, err := pkgconfig.ParseConfigFile(config)
	if err != nil {
		log.FailureStatusEvent(os.Stdout, err.Error())
		return
	}
	ctx := context.Background()
	// listening address.
	listenAddr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)

	options := []yomo.ZipperOption{}
	tokenString := ""
	if _, ok := conf.Auth["type"]; ok {
		if tokenString, ok = conf.Auth["token"]; ok {
			options = append(options, yomo.WithAuth("token", tokenString))
		}
	}
	// check llm bridge server config
	// parse the llm bridge config
	bridgeConf := conf.Bridge
	aiConfig, err := ai.ParseConfig(bridgeConf)
	if err != nil {
		if err == ai.ErrConfigNotFound {
			log.InfoStatusEvent(os.Stdout, err.Error())
		} else {
			log.FailureStatusEvent(os.Stdout, err.Error())
			return
		}
	}
	if aiConfig != nil {
		// add AI connection middleware
		options = append(options, yomo.WithZipperConnMiddleware(ai.ConnMiddleware))
	}
	// new zipper
	zipper, err := yomo.NewZipper(
		conf.Name,
		router.Default(),
		nil,
		conf.Mesh,
		options...)
	if err != nil {
		log.FailureStatusEvent(os.Stdout, err.Error())
		return
	}
	zipper.Logger().Info("using config file", "file_path", config)

	// AI Server
	if aiConfig != nil {
		// register the llm provider
		registerAIProvider(aiConfig)
		// start the llm api server
		go func() {
			err := ai.Serve(aiConfig, listenAddr, fmt.Sprintf("token:%s", tokenString))
			if err != nil {
				log.FailureStatusEvent(os.Stdout, err.Error())
				return
			}
		}()
	}

	// start the zipper
	err = zipper.ListenAndServe(ctx, listenAddr)
	if err != nil {
		log.FailureStatusEvent(os.Stdout, err.Error())
		return
	}
}

func registerAIProvider(aiConfig *ai.Config) {
	for name, provider := range aiConfig.Providers {
		switch name {
		case "azopenai":
			ai.RegisterProvider(azopenai.NewProvider(
				provider["api_key"],
				provider["api_endpoint"],
				provider["deployment_id"],
				provider["api_version"],
			))
		case "gemini":
			ai.RegisterProvider(gemini.NewProvider(provider["api_key"]))
		case "openai":
			ai.RegisterProvider(openai.NewProvider(provider["api_key"], provider["model"]))
		case "cloudflare_azure":
			ai.RegisterProvider(cfazure.NewProvider(
				provider["endpoint"],
				provider["api_key"],
				provider["resource"],
				provider["deployment_id"],
				provider["api_version"],
			))
		default:
			log.WarningStatusEvent(os.Stdout, "unknown provider: %s", name)
		}
	}
}

//
// Copyright (c) 2020 Ankur Srivastava
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package app

import (
	"path"

	"github.com/gofiber/embed"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"

	config "github.com/ansrivas/fiber-pongo2-pkger/internal/config"
	"github.com/gofiber/template/django"
	"github.com/markbates/pkger"
)

// App struct consists of our application structure
type App struct {
	Server   *fiber.App
	Config   *config.Config
	ProxyURL string
}

type appConfig struct {
	// Put all your html/jinja templates here, defaults to /templates in pkger
	htmlTemplateDir string

	// Put all your assets like css, js etc.
	staticAssetDir string

	// Proxy url to prefix this app with
	proxyPrefix string

	// configuration
	config *config.Config

	// DB instance
}

// ConfigOption will be a set of configuration options for APP
type ConfigOption struct {
	setup func(ro *appConfig)
}

// WithHTMLTemplateDir will provide the base dir to load templates from pkger
func WithHTMLTemplateDir(htmlTemplateDir string) ConfigOption {
	return ConfigOption{func(ac *appConfig) {
		ac.htmlTemplateDir = htmlTemplateDir
	}}
}

// WithStaticAssetDir will provide the base dir to load static assets from pkger
func WithStaticAssetDir(staticAssetDir string) ConfigOption {
	return ConfigOption{func(ac *appConfig) {
		ac.staticAssetDir = staticAssetDir
	}}
}

// WithProxyURL returns a proxy on wihch we  will register our application
// This could look like "/production/prefix/stuff"
func WithProxyURL(proxyPrefix string) ConfigOption {
	return ConfigOption{func(ac *appConfig) {
		ac.proxyPrefix = proxyPrefix
	}}
}

// WithConfiguration will register a configuration of type *config.Config
func WithConfiguration(config *config.Config) ConfigOption {
	return ConfigOption{func(ac *appConfig) {
		ac.config = config
	}}
}

func prepareRoutes(baseURL, suffix string) string {
	return path.Join(baseURL, suffix)
}

// New creates a new instance of App
func New(configOptions ...ConfigOption) *App {

	ac := &appConfig{}
	for _, option := range configOptions {
		option.setup(ac)
	}

	// Instantiate a fiber application
	app := fiber.New()

	// Register the templates directory which is packaged using pkger
	if ac.htmlTemplateDir != "" {
		engine := django.NewFileSystem(pkger.Dir(ac.htmlTemplateDir), ".html")
		app.Settings.Views = engine
	}

	// Register the static assets like css, js etc
	if ac.staticAssetDir != "" {
		staticAsset := embed.New(embed.Config{
			Root: pkger.Dir(ac.staticAssetDir),
		})
		app.Use(prepareRoutes(ac.proxyPrefix, "/static"), staticAsset)
	}

	// Use default recoverer middleware
	app.Use(middleware.Recover())

	return &App{
		Server:   app,
		Config:   ac.config,
		ProxyURL: ac.proxyPrefix,
	}
}

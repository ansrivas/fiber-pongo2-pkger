package main

import (
	"flag"
	"fmt"
	"os/signal"
	"syscall"

	"os"

	config "github.com/ansrivas/fiber-pongo2-pkger/internal/config"
	figure "github.com/common-nighthawk/go-figure"

	"github.com/ansrivas/fiber-pongo2-pkger/pkg/app"
	"github.com/ansrivas/fiber-pongo2-pkger/pkg/routers"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	// BuildTime gets populated during the build proces
	BuildTime = ""

	//Version gets populated during the build process
	Version = ""
)

// setupConfigOrFatal loads all the variables from the environment variable.
// At this point everything is read as a Key,Value in a map[string]string
func setupConfigOrFatal() config.Config {
	conf, err := config.LoadEnv()
	if err != nil {

		log.Fatal().Msgf("Failed to parse the environment variable. Error %s", err.Error())
	}
	return conf
}

func printBanner() {
	myFigure := figure.NewFigure("fiber-pongo2-pkger", "", true)
	myFigure.Print()
}

// setupLogger will setup the zap json logging interface
// if the --debug flag is passed, level will be debug
func setupLogger(debug bool) {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

}

// printVersionInfo just prints the build time and git commit/tag used
// for this build
func printVersionInfo(version bool) {
	if version {
		fmt.Printf("Version  : %s\nBuildTime: %s\n", Version, BuildTime)
		os.Exit(0)
	}
}

func helloWorld() string {
	return "Hello World"
}

func main() {
	debug := flag.Bool("debug", false, "Set the log level to debug")
	version := flag.Bool("version", false, "Display the BuildTime and Version of this binary")
	flag.Parse()

	printVersionInfo(*version)
	setupLogger(*debug)

	// Register a ctrl-c handler
	errc := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	config := setupConfigOrFatal()
	application := app.New(
		app.WithConfiguration(&config),
		app.WithHTMLTemplateDir("/web/templates"),
		app.WithStaticAssetDir("/web/static"),
	)

	routes := routers.New()
	application.RegisterRoutes(routes)

	go func() {
		errc <- application.Server.Listen(application.Config.Address)
	}()

	fmt.Println("Press ctrl-c to exit")
	log.Info().Msgf("Exiting server. Message: %v", <-errc)
}

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/ryodocx/ical-proxy/pkg/converter"
	"github.com/ryodocx/ical-proxy/pkg/feed/redmine/issue"
	"github.com/ryodocx/ical-proxy/pkg/server"
	"github.com/ryodocx/ical-proxy/pkg/util"
	"github.com/urfave/cli/v2"
)

func main() {

	envPrefix := "ICALPROXY_"
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	app := &cli.App{
		Usage:           "generage iCal from any sources",
		HideHelpCommand: true,
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:  "vcalendar-properties",
				Usage: "Properties for VCALENDAR: https://datatracker.ietf.org/doc/html/rfc5545#section-3.7",
				Value: cli.NewStringSlice(
					"PRODID:https://github.com/ryodocx/ical-proxy",
					"X-PUBLISHED-TTL;VALUE=DURATION:PT30M", // https://docs.microsoft.com/en-us/openspecs/exchange_server_protocols/ms-oxcical/1fc7b244-ecd1-4d28-ac0c-2bb4df855a1f
				),
				EnvVars: []string{envPrefix + "VCALENDAR_PROPERTIES"},
			},
			&cli.StringSliceFlag{
				Category: "rego",
				Name:     "rego-paths",
				Usage:    "paths of *.rego file or dir",
				Value:    cli.NewStringSlice(path.Join(wd, "./configs") + "/"),
				EnvVars:  []string{envPrefix + "REGO_PATHS"},
			},
			&cli.StringFlag{
				Category: "rego",
				Name:     "rego-query",
				Value:    "data.ical.simple",
				EnvVars:  []string{envPrefix + "REGO_QUERY"},
			},
			&cli.StringFlag{
				Category: "server",
				Name:     "listen-addr",
				Value:    "127.0.0.1:8080",
				EnvVars:  []string{envPrefix + "LISTEN_ADDR"},
			},
			&cli.StringFlag{
				Category: "server",
				Name:     "listen-path",
				Value:    "/",
				EnvVars:  []string{envPrefix + "LISTEN_PATH"},
			},
			&cli.StringFlag{
				Category: "server",
				Name:     "listen-query",
				EnvVars:  []string{envPrefix + "LISTEN_QUERY"},
			},
			&cli.DurationFlag{
				Category: "server",
				Name:     "grace-period",
				Value:    time.Second * 3,
				EnvVars:  []string{envPrefix + "GRACE_PERIOD"},
			},
			&cli.StringFlag{
				Category: "redmine",
				Name:     "redmine-url",
				EnvVars:  []string{envPrefix + "REDMINE_URL"},
				Required: true,
			},
			&cli.StringFlag{
				Category: "redmine",
				Name:     "redmine-apikey",
				EnvVars:  []string{envPrefix + "REDMINE_APIKEY"},
				Required: true,
			},
			&cli.StringFlag{
				Category: "redmine",
				Name:     "redmine-query",
				EnvVars:  []string{envPrefix + "REDMINE_QUERY"},
			},
			&cli.IntFlag{
				Category: "redmine",
				Name:     "redmine-maxfetch",
				Usage:    "max fetch number of redmine issues (0 means unlimited)",
				Value:    1000,
				EnvVars:  []string{envPrefix + "REDMINE_MAXFETCH"},
			},
		},
		Action: func(cCtx *cli.Context) error {

			serverConf := &server.Config{
				Addr: cCtx.String("listen-addr"),
				Path: cCtx.String("listen-path"),
			}

			{
				conf := &issue.Config{
					Url:        cCtx.String("redmine-url"),
					ApiKey:     cCtx.String("redmine-apikey"),
					MaxEntries: cCtx.Int("redmine-maxfetch"),
				}
				if q, err := url.ParseQuery(cCtx.String("redmine-query")); err != nil {
					return util.WrapError(err)
				} else {
					conf.Query = q
				}

				// fmt.Printf("%#v\n", conf)
				i, err := issue.New(conf)
				if err != nil {
					return util.WrapError(err)
				}
				serverConf.Feed = i
			}

			{
				conf := &converter.Config{
					RegoPaths:         cCtx.StringSlice("rego-paths"),
					RegoQuery:         cCtx.String("rego-query"),
					CalendarPropaties: cCtx.StringSlice("vcalendar-properties"),
				}
				// fmt.Printf("%#v\n", conf)
				c, err := converter.New(conf)
				if err != nil {
					return util.WrapError(err)
				}
				serverConf.Converter = c
			}

			{
				if q, err := url.ParseQuery(cCtx.String("listen-query")); err != nil {
					return util.WrapError(err)
				} else {
					serverConf.Query = q
				}
				// fmt.Printf("%#v\n", serverConf)
			}

			s, err := server.New(serverConf)
			if err != nil {
				return util.WrapError(err)
			}

			idleConnsClosed := make(chan struct{})
			go func() {
				// signal monitoring
				for {
					sigChan := make(chan os.Signal, 1)
					signal.Notify(sigChan)
					signal.Ignore(syscall.SIGURG) // https://golang.hateblo.jp/entry/golang-signal-urgent-io-condition
					recievedSignal := <-sigChan
					log.Println("signal recieved:", fmt.Sprintf("%d(%s)", recievedSignal, recievedSignal.String()))

					for _, s := range []os.Signal{os.Interrupt, syscall.SIGTERM} {
						if recievedSignal == s {
							goto shutdown
						}
					}
				}

				// graceful shutdown
			shutdown:
				log.Println("shutting down...")

				ctx, cancel := context.WithTimeout(context.Background(), cCtx.Duration("grace-period"))
				defer cancel()

				if err := s.Shutdown(ctx); err != nil {
					log.Printf("HTTP server Shutdown: %v", err)
				}
				close(idleConnsClosed)
			}()

			// start
			log.Printf("start servering at %s\n", serverConf.Addr)
			if err := s.ListenAndServe(); err != http.ErrServerClosed {
				return util.WrapError(err)
			}
			<-idleConnsClosed

			return nil
		},
	}

	// https://github.com/urfave/cli/issues/734#issuecomment-597344796
	globalOptionsTemplate := `{{if .VisibleFlags}}GLOBAL OPTIONS:
   {{range $index, $option := .VisibleFlags}}{{if $index}}
   {{end}}{{$option}}{{end}}
{{end}}
`
	origHelpPrinterCustom := cli.HelpPrinterCustom
	defer func() {
		cli.HelpPrinterCustom = origHelpPrinterCustom
	}()
	cli.HelpPrinterCustom = func(out io.Writer, templ string, data interface{}, customFuncs map[string]interface{}) {
		origHelpPrinterCustom(out, templ, data, customFuncs)
		if data != app {
			origHelpPrinterCustom(app.Writer, globalOptionsTemplate, app, nil)
		}
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

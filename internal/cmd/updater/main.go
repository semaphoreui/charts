package main

import (
	"context"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/semaphoreui/charts/internal/chart"
	"github.com/semaphoreui/charts/internal/chart/oci"
)

func main() {
	if len(os.Args) != 2 {
		_, _ = fmt.Fprintf(os.Stderr, "usage: %s <directory>\n", os.Args)
		os.Exit(2)
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stdout,
			NoColor: false,
		},
	)

	if err := run(
		context.Background(),
		os.Args[1],
	); err != nil {
		os.Exit(1)
	}
}

func run(ctx context.Context, root string) error {
	charts, err := chart.Find(root)

	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to find charts")

		return err
	}

	for _, c := range charts {
		log.Info().
			Str("chart", c.Meta.Name).
			Msg("scanning chart")

		if !c.Source.Disabled {
			log.Debug().
				Str("chart", c.Meta.Name).
				Msg("checking app version")

			latest, err := oci.Latest(
				path.Join(
					c.Source.Registry,
					c.Source.Image,
				),
				c.Source.Ignored,
				c.Source.Prerelease,
			)

			if err != nil {
				log.Error().
					Err(err).
					Str("chart", c.Meta.Name).
					Msg("failed to check latest app version")

				continue
			}

			if c.Meta.AppVersion != latest {
				log.Info().
					Str("chart", c.Meta.Name).
					Str("old", c.Meta.AppVersion).
					Str("new", latest).
					Msg("found newer app version")

				if err := chart.UpdateApp(
					c,
					latest,
				); err != nil {
					log.Error().
						Err(err).
						Str("chart", c.Meta.Name).
						Str("old", c.Meta.AppVersion).
						Str("new", latest).
						Msg("failed to update app version")
				}
			} else {
				log.Info().
					Str("chart", c.Meta.Name).
					Str("version", c.Meta.AppVersion).
					Msg("no newer app version found")
			}
		}

		for _, d := range c.Meta.Dependencies {
			log.Debug().
				Str("chart", c.Meta.Name).
				Str("dependency", d.Name).
				Msg("checking dependency version")

			switch {
			case strings.HasPrefix(d.Repository, "oci://"):
				latest, err := oci.Latest(
					path.Join(
						strings.TrimPrefix(
							d.Repository,
							"oci://",
						),
						d.Name,
					),
					[]string{},
					false,
				)

				if err != nil {
					log.Error().
						Err(err).
						Str("chart", c.Meta.Name).
						Str("dependency", d.Name).
						Msg("failed to check latest dependency version")

					continue
				}

				if d.Version != latest {
					log.Info().
						Str("chart", c.Meta.Name).
						Str("dependency", d.Name).
						Str("old", d.Version).
						Str("new", latest).
						Msg("found newer dependency version")

					if err := chart.UpdateDependency(
						c,
						d,
						latest,
					); err != nil {
						log.Error().
							Err(err).
							Str("chart", c.Meta.Name).
							Str("dependency", d.Name).
							Str("old", d.Version).
							Str("new", latest).
							Msg("failed to update dependency version")
					}
				} else {
					log.Info().
						Str("chart", c.Meta.Name).
						Str("dependency", d.Name).
						Str("version", d.Version).
						Msg("no newer dependency version found")
				}
			default:
				log.Error().
					Str("chart", c.Meta.Name).
					Str("dependency", d.Name).
					Str("repo", d.Repository).
					Msg("unknown repository type")
			}
		}

		if err := chart.UpdateVersion(
			c,
		); err != nil {
			log.Error().
				Err(err).
				Str("chart", c.Meta.Name).
				Msg("failed to bump chart version")
		}

		if err := chart.CommitChanges(
			c,
		); err != nil {
			log.Error().
				Err(err).
				Str("chart", c.Meta.Name).
				Msg("failed to commit changed files")
		}
	}

	return nil
}

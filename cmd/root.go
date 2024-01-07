package cmd

import (
	"fmt"
	"github.com/hov1417/conserve-clean/conserve"
	"github.com/hov1417/conserve-clean/retention"
	"github.com/spf13/cobra"
	"os"
)

// This format is stolen from
// https://duplicati.readthedocs.io/en/latest/06-advanced-options/#retention-policy
const longHelp = `
conserve-clean, version 0.1.0
author: hov1417 <hovhannes1417@gmail.com>

Duplicati style retention policy for sourcefrog's Conserve backups.
By default this tool will print the names of backups that would be deleted.
To actually delete them, use the --delete flag.

FORMAT OF RETENTION POLICY:
The retention policy is a comma separated list of time frames, each of which is a
timeframe:interval duple. The time frame is the duration of time to keep backups for,
and the interval is the time between backups to keep.
Both time frame and interval are specified as a number followed by a letter.
The number is the number of time units, and the letter is the unit of time.

Valid letters for time: 's', 'm', 'h', 'D', 'W', 'M', 'Y' (case sensitive)
corresponding to seconds, minutes, hours, days, weeks, months and years.
Year and Month are approximations, and assumed to be 365 days per year and 30 days per month.

When overlapping time frames are specified, the smallest time frame takes priority,
thus the effective duration of longer time frames becomes shorter.

For example the value '7D:0s,3M:1D,10Y:2M' means "during the next 7 day keep all backups,
during the next 3 months from now keep a daily backup and for 10 years from now keep one
backup every 2nd month."`

// rootCmd represents the base command when called without any subcommands
var (
	path          string
	executable    string
	deleteBackups bool
	rootCmd       = &cobra.Command{
		Use:   "conserve-clean",
		Short: "Clean Conserve Backups by a Given Filter",
		Long:  longHelp,
		Run:   Run,
		Args:  cobra.ExactArgs(1),
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func Run(cmd *cobra.Command, args []string) {
	if err := execute(args[0]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func execute(patter string) error {
	policy, err := retention.Parse(patter)
	if err != nil {
		return err
	}

	backups, err := conserve.Versions(executable, path)
	if err != nil {
		return err
	}

	_, remove, err := retention.SplitByPolicy(backups, policy)
	if err != nil {
		return err
	}

	if deleteBackups {
		for _, backup := range remove {
			if err := conserve.Delete(executable, path, backup.Name()); err != nil {
				return err
			}
		}
	} else {
		for _, backup := range remove {
			fmt.Println(backup.Name())
		}
	}
	return nil
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&path, "path", "p", ".", "directory to search for backups")
	rootCmd.PersistentFlags().StringVarP(&executable, "executable", "e", "conserve", "executable to use for listing backups")
	rootCmd.PersistentFlags().BoolVarP(&deleteBackups, "delete", "d", false, "delete filtered backups")
}

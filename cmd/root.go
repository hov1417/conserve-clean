package cmd

import (
    "conserve-clean/conserve"
    "conserve-clean/retention"
    "fmt"
    "github.com/spf13/cobra"
)

// This format is stolen from
// https://duplicati.readthedocs.io/en/latest/06-advanced-options/#retention-policy
const longHelp = `Clean backups from 'sourcefrog/conserve' tool, using syntax like
Specify one or more timeframe:interval duples, such as '7D:0s'
Valid letters for time: 's', 'm', 'h', 'D', 'W', 'M', 'Y'

Year and Month are approximations, and assume 365 days per year and 30 days per month.

Multiple duples shall be comma separated, and time frames shall be increasing.
For example the value '7D:0s,3M:1D,10Y:2M' means "during the next 7 day keep all backups,
during the next 3 months from now keep a daily backup and for 10 years from now keep one
backup every 2nd month. 0s stands for an interval of zero length, allowing unlimited versions to be kept,
which can be also noted as 'U' To avoid gaps, time frames all start at "now" and overlap, with smaller time
frames taking priority, thus the effective duration of longer time frames becomes shorter.`

// rootCmd represents the base command when called without any subcommands
var (
    dir           string
    executable    string
    deleteBackups bool
    rootCmd       = &cobra.Command{
        Use:   "conserve-clean",
        Short: "Clean Conserve RawBackup by a Given Filter",
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

    }
}

func execute(patter string) error {
    policy, err := retention.Parse(patter)
    if err != nil {
        return err
    }

    backups, err := conserve.Versions(executable, dir)
    if err != nil {
        return err
    }

    _, remove, err := retention.SplitByPolicy(backups, policy)
    if err != nil {
        return err
    }

    if deleteBackups {
        for _, backup := range remove {
            if err := conserve.Delete(executable, dir, backup.Name()); err != nil {
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
    rootCmd.PersistentFlags().StringVarP(&dir, "dir", "d", ".", "directory to search for backups")
    rootCmd.PersistentFlags().StringVarP(&executable, "executable", "e", "conserve", "executable to use for listing backups")
    rootCmd.PersistentFlags().BoolVarP(&deleteBackups, "delete", "d", false, "delete filtered backups")
}

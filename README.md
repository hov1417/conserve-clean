# Conserve Clean
[Duplicati](https://duplicati.readthedocs.io/en/latest/06-advanced-options/#retention-policy) style retention policy for sourcefrog's Conserve backups.
By default, this tool will print the names of backups that would be deleted.
To actually delete them, use the --delete flag.

## Format of the retention policy
The retention policy is a comma separated list of time frames, each of which is a
`timeframe:interval` duple. The time frame is the duration of time to keep backups for,
and the interval is the time between backups to keep.
Both time frame and interval are specified as a number followed by a letter.
The number is the number of time units, and the letter is the unit of time.

Valid letters for time: `'s', 'm', 'h', 'D', 'W', 'M', 'Y'` (case sensitive)   
corresponding to seconds, minutes, hours, days, weeks, months and years.
Year and Month are approximations, and assumed to be 365 days per year and 30 days per month.

When overlapping time frames are specified, the smallest time frame takes priority,
thus the effective duration of longer time frames becomes shorter.

For example the value '7D:0s,3M:1D,10Y:2M' means "during the next 7 days keep all backups,
during the next 3 months from 7 days keep a daily backup and for 10 years from 3 months keep one
backup every 2nd month".

## Usage

```bash
conserve-clean 1W:1D,4W:1W,10Y:2M --path /path/to/backup --delete --executable /path/to/conserve
```
`--executable` is optional, and defaults to 'conserve'  
`--path` is also optional, and defaults to the current directory  
without `--delete` by default, the tool will only print the names of backups that would be deleted.

## Installation

```bash
go install github.com/hov1417/conserve-clean@latest
```
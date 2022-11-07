# Verse CLI

## Getting Started

### Installation

```
go install github.com/ansidev/verse-cli@latest
```

### Usage

```
NAME:
   verse - Verse CLI

USAGE:
   Get verses by month and day. Example command: verse --month=12 --day=12 --format=md

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --month value, -m value   Month to fetch verse for. Value range: [1-12] (default: current local month)
   --day value, -d value     Day to fetch verse for. Value range: [1-31] (default: current local day)
   --format value, -f value  Verse address format. dm: {book} {month}:{day}, md: {book} {day}:{month}. (default: "dm")
   --help, -h                show help (default: false)
```

## Contact

Le Minh Tri [@ansidev](https://ansidev.xyz/about).

## License

This source code is available under the [AGPL-3.0 LICENSE](/LICENSE).
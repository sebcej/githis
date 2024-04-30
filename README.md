# githis

Simple commits aggregator for your projects folders

## Why

* Get all your projects commits in a single place
* Filter all projects by date, author or list of authors
* Support for multiple projects folders

## Installation

### Mac

```bash
xattr -d com.apple.quarantine githis
chmod +x githis
sudo mv githis /usr/local/bin
```

## Get started

Set your projects folder, the cli will scan this folder and will get all commits

`githis sources add projects /home/sebcej/Projects`

Set yourself as the commit author filter if you want only your logs

`githis config set author sebcej`

## Use cases and examples

See your yesterday commits in all your local projects

`githis logs -o -1`

... or filter by date range

`githis logs --from 2024-04-10 --to 2024-04-13`

... or filter by single day

`githis logs -d 2024-04-10`

Date autocomplete is supported for single day param

`githis logs -d 10` equals to `githis logs -d 2024-04-10`

`githis logs -d 04-10` equals to `githis logs -d 2024-04-10`

Filter by multiple authors

`githis logs -a sebcej -a anotherdev`

To enable auto-pull use the `-p` param

`githis logs -o -1 -p`

Use `githis help` to see all functionalities
# githis

Simple commits aggregator for your projects folder

With this cli you can track all contributions from one or multiple developers in all company projects at once.

## Functionalities

* Get all your projects commits in a single place
* Filter by date, author or list of authors
* Support for multiple projects folders

## Examples

Set your projects folder

`githis sources add projects /home/sebcej/Projects`

Set yourself as the commit author filter

`githis config set author sebcej`

See your yesterday commits in all your local projects

`githis logs -o -1`

... or filter by date range

`githis logs --fromDate 2024-04-10 --toDate 2024-04-13`

Filter by group of authors

`githis logs -a sebcej -a anotherdev`

Use `githis help` to see all functionalities
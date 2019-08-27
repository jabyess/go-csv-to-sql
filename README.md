# Go CSV to SQL

Converts CSV files to SQL scripts. Pet project to learn go, built at recurse during fall 1 2019.

## Usage

`go run main.go file.csv`

## Notes

Put the file in the same directory as the script, just because.

Currently tested with postgresql types only.

Creates a table with the same name as the input file.

Coverts header values " " and "-" to "_".

Parses and writes:
* 126MB, 735,000 line file in ~4 seconds.
* 1008MB, 5,884,018 line file in ~32 seconds.

## Supported Types

* Date (`YYYY-MM-DD`)
* Timestamp (`YYYY-MM-DD HH:MM:SS`)
* Integer (`NNNNN`)
* Text (`[A-Za-z0-9]`)

## Todo

* sanitize filename for sql standards (no `-` or special chars)

* Add support for more datatypes in psql
  * different dates
  * bool
  * enumerated types?
  * floats vs integers


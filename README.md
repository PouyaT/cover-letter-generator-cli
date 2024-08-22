# cover-letter-generator-cli
A cli to generate cover letters for companies

## Installation
1. clone the repo
2. `go install`


## Usage:
To use the cli first install the package. As of now the only command that exists is `generate` 

`cover-letter-generator generate {company} {position}`

This will output a pdf file to the `C:/Users/Admin/OneDrive/Documents/Resume and Cover Letter/Cover Letters/` dir with a file name `{company}_Coverletter.pdf`

Ex: `cover-letter-generator generate Google "Senior Engineer"` (Note if the position is two words pass it in with quotes)

## TODO
* Add a flag such as `-l` when using the `generate` command to set a new location to download the pdf file
* Add a config file to set a new default location to download the pdf file
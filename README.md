# go-automate-db - A database automation tool
***

A database automation tool to help automate 
common operation tasks

## License
***
Licensed under `GNU General Public License v3.0`
This project aims to be usable and extensible by anyone interested.

## Use
### With Linux
```bash
# Download latest release. Check the link below for the latest release
# https://github.com/alacrity-sg/go-automate-db/releases

OS_ARCHI="linux_amd64"
VERSION="1.0.0"
wget https://github.com/alacrity-sg/go-automate-db/releases/download/v${VERSION}/go-automate-db_${VERSION}_${OS_ARCHI}.tar.gz
tar -xvf go-automate-db_${VERSION}_${OS_ARCHI}.tar.gz

# Run your DB operations
./go-automate-db \
  -t "postgres" -h "localhost" \
  -P "5432" -u "root" \
  -p "some_password" \
  ...other arguments
# Refer to Usage for documentation on arguments
```
## Usage
go-automate-db is a tool that aims to automate some of the common database operations work.
This tool accepts inputs in 3 modes:
- .yaml configuration file
- environment variables
- direct flag arguments

### Feature Mode - Postgres DB creation

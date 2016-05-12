# log-crayon
A simple devops utility to emphasize / colorize certain parts of a log file. Used with grep and tail unix utilities.

# Installation
```
go get github.com/kvlar/log-crayon
```
Customize `config.yml.example` and put it in your home directory as `.log_crayon.yml`
Alternatively you can provide a path to the config file using `-c` flag.

# Usage
```
tail -f log_file.log | log-crayon
```

# Colors
Please see https://github.com/mgutz/ansi for available colors.

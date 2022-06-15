# wkcli

Search wikimedia sites from the command line

## Usage
```
wkcli [ -lx ] [ -w wiki ]
```

 - `-l` List page references as well
 - `-x` Return exact match
 - `-w wiki` Select which wikimedia site to search (Default en.wikipedia.org)

## Bugs

Issues with how the json parsing is done leads to fatal errors when bad results come in. This will be addressed in the future as I have time

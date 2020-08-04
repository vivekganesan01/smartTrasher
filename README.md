# Smart Auto Trasher

Smart Auto Trasher helps to monitor and track file system within a specific directory and trashes the files if modified time is not within a given timeframe.
Below file's are being monitored ....

  - ".png", ".jpeg", ".jpg", ".svn"
  - ".doc", ".xlsx", ".xls", ".pdf", ".log", ".csv", ".log", ".txt", ".jmx", ".docx", ".html", ".xml", ".json"
  - ".zip", ".dmg", ".gz"

# Features!

  - Tracks last modification time of the files and trash the unused files based on current date
  - Move the unused files to auto trash directory

### Installation

requires [golang].

```sh
$ ./smartTrash
```

For development...
```sh
$ go build smartTrash.go
```

###### Todo:
 - logger
 - control over the file extensions
 - recursive dir reading

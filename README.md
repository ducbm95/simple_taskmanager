# Overview
This is a simple task manager application that running on console.

# Usage
To start the application:
```bash
$ go run main.go
```
Basic usage example:
```bash
Please enter command: ADD|REMOVE|START|STOP|IMPORT|EXPORT|EXIT
> add
> Task name: task1
> Cron expression: * * * * *
> Task content: Do Math exercise
task1 added successfully
Please enter command: ADD|REMOVE|START|STOP|IMPORT|EXPORT|EXIT
> start
> Task name: task1
task1 started successfully
```

The tasks can be exported/imported to a dump file. Usage:
```bash
Please enter command: ADD|REMOVE|START|STOP|IMPORT|EXPORT|EXIT
> export
> Task name: task1
task1 exported successfully
Please enter command: ADD|REMOVE|START|STOP|IMPORT|EXPORT|EXIT
> remove
> Task name: task1
task1 removed successfully
Please enter command: ADD|REMOVE|START|STOP|IMPORT|EXPORT|EXIT
> import
> Task name: task1
task1 imported successfully
Please enter command: ADD|REMOVE|START|STOP|IMPORT|EXPORT|EXIT
> start   
> Task name: task1
task1 started successfully
```


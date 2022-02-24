
                            ▄▄▄▄▄▄▄ ▄▄▄▄▄▄   ▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄ 
                            █       █   ▄  █ █       █       █       █
                            █   ▄▄▄▄█  █ █ █ █   ▄   █   ▄   █▄     ▄█
                            █  █  ▄▄█   █▄▄█▄█  █ █  █  █ █  █ █   █  
                            █  █ █  █    ▄▄  █  █▄█  █  █▄█  █ █   █  
                            █  █▄▄█ █   █  █ █       █       █ █   █  
                            █▄▄▄▄▄▄▄█▄▄▄█  █▄█▄▄▄▄▄▄▄█▄▄▄▄▄▄▄█ █▄▄▄█  


Groot is the Go logging library for Arcana. It wraps up libraries/services like zap, sentry and nicely expose the standard logging functions and features that's needed by most projects. – like levelled logging, context logging, HTTP request tracing, log rotation and backups.

You mostly don't have to worry about format of logs, which library we use behind the scenes, which service we are using for monitoring. Groot standardizes the way we log in our code so that devops can easily pickup, collect and process logs.

# Quick start
- Go through the `example/main.go` to find out how to use groot in your project

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


# Usage
- Create a logger instance and pass it around to packages. This is a better pattern as it avoids globals. Though this might not be possible if the code is not structured to inject dependencies 
and heavily rely on globals, for this purpose groot provides a singleton logger

# Todo
- [ ] Implement log rotation and retention
- [ ] Implement log backups to cold storage
- [ ] Implement context based logging
- [ ] Implement HTTP request tracing
- [ ] Implement sentry logs
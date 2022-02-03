# logger
It's a simple logger, as an extension from log of standard golang library, only need register 3 environment variables in the operative System.

``` console
LOG_FILE_IS_ACTIVE = true
LOG_FILE_NAME = miloggerfile.log || .txt
LOG_LEVEL = warning
```

**LOG_FILE_IS_ACTIVE (bool)**: It is the flag that indicates if they save the logs in a file, by default is false.

**LOG_FILE_NAME**: It is the name of the file where the logs are saved, by default is logger.log.

---
**LOG_LEVEL**: It is the output level of the messages both in the file and in the console, by default is information (info).

The writing level can be worked in the following way:
``` diff
+ debug || debugger
+ info  || information
- warn  || warning
- err   || error
```
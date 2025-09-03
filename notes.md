# main
- The main function creates a new Config type by reading the 
json file **configFileName**.
- **SetUser()** is a method of the Config type which allows us to
define username for the json file, using a helper function called
**write**, which takes a Config type as argument and creates or 
truncates with it the json file given in configFileName.
- Then reads again the file.


# JSON
The gatorconfig.json file keeps track of:
1. who is currently logged
2. connection credential for the PostgreSQL database    
`{"db_url":"postgres://example","current_user_name":"lalobg"}`

# internal/config
config package contains the following:
- constants
    - configFileName
- types
    - Config struct
- functions
    - Read
    - SetUser
    - write
    - getConfigFilePath
    

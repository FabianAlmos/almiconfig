# Almi config
Almi config is a lightweight configuration loader module for Go based on struct tags.

## Currently supported types are: 
- **int**
- **int8**
- **int16**
- **int32**
- **int64**
- **uint**
- **uint8**
- **uint16**
- **uint32**
- **uint64**
- **uintptr**
- **float32**
- **float64**
- **byte**
- **rune**
- **string**

Almi config reads in the values from the environment,
which means that you only have to load the values to the environment,
for example with **joho's godotenv** you can load all the environment variables
to the environment and the module will do the rest with the
**almi.ValidateConfig\[T any\]\(cfg T\) T** function.

For the config validating function to work the config struct fields must use
the **almi** struct tag.

## Currently supported struct tags:
- **required**:
  - specifies whether the field of the config must be set in the environment or not.
- **env**:
  - **env** must be specified for all fields even if they are not used
  - usage:
    ```go
    package main
    
    // env: SECRET=secret
    
    type Config struct {
        Secret string `almi:"env=SECRET"`
    }
    ```
- **type**:
  - The **type** constraint is used to convert the read in environment variable's
    type to the desired type in the config, **by default** environment variables
    are read in as **string**.
  - usage:
    ```go
    package main

    // env: SECRET_LIFETIME=1440
    
    type Config struct {
        SecretLifetime int `almi:"env=SECRET_LIFETIME,type=int"`
    }
    ```
  - The **type** constraint can also be used to read in slices from the environment,
    with slices you must specify the separator character in the square brackets
    of the type.
  - usage:
    ```go
    package main
    
    // env: BROKERS=broker1,broker2,broker3
    
    type Config struct {
        Brokers []string `almi:"env=BROKERS,type=[,]string"`
    }
    ```
- **default**:
  - The **default** constraint can be used to set a default value for environment variables in-case they are not set in the environment.
  - If the **required** constraint is set on a config field and the **default** constraint is also set,
    the **required** constraint won't raise an error during the validation of the config.
  - If the field is a slice type set by the **type** constraint,
    the **default** constraint's value must be set with square brackets around the default value.
    Like in the second example shown below, also make sure that the separator character matches the separator set in the **type** constraint.
  - usage:
    ```go
    package main
    
    // regular field example
    // env: empty
    
    type Config struct {
        SecretLifetime int `almi:"required,env=SECRET_LIFETIME,type=int,default=1440"`
    }
    ```
    ```go
    package main
    
    // slice type field example
    // env: empty
    
    type Config struct {
        Brokers []string `almi:"required,env=BROKERS,type=[,]string,default=[broker1,broker2,broker3]"`
    }
    ```
## Usage example:
**.env**:
```
ACCESS_SECRET=access_secret
REFRESH_SECRET=refresh_secret
ACCESS_LIFETIME=2
REFRESH_LIFETIME=30

POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_HOST=postgres
PGPORT=5432
POSTGRES_DB=postgres

KAFKA_BROKERS=broker1,broker2,broker3
```

**main.go**:
```go
package main

import (
  "almi"
  "fmt"
  "github.com/joho/godotenv"
)

type Config struct {
	AccessSecret    string `almi:"required,env=ACCESS_SECRET"`
	RefreshSecret   string `almi:"required,env=REFRESH_SECRET"`
	AccessLifetime  int    `almi:"required,env=ACCESS_LIFETIME,type=int"`
	RefreshLifetime int    `almi:"required,env=REFRESH_LIFETIME,type=int"`
	
	PostgresRootUser     string `almi:"env=POSTGRES_ROOT_USER"`
	PostgresRootPassword string `almi:"env=POSTGRES_ROOT_PASSWORD"`
	PostgresUser         string `almi:"required,env=POSTGRES_USER"`
	PostgresPassword     string `almi:"required,env=POSTGRES_PASSWORD"`
	PostgresHost         string `almi:"required,env=POSTGRES_HOST"`
	PostgresPort         int    `almi:"required,env=PGPORT,type=int"`
	PostgresDatabase     string `almi:"required,env=POSTGRES_DB"`
	
	KafkaBrokers []string `almi:"required,env=KAFKA_BROKERS,type=[,]string"`
}

func main() {
  if err := godotenv.Load(); err != nil {
    panic(err)
  }

  cfg, err := almi.ValidateConfig(Config{})
  if err != nil {
    panic(err)
  }

  fmt.Println(cfg)
}
```
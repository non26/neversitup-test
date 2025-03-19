```
project.name
├── cmd
│   ├── app
│   │   ├── init.go                          # init database, redis here
│   │   └── rate-ui.css
│   └── main.go                              # main entry
├── configs
│   ├── config.yml                           # app config
├── internal
│   ├── core
│   │    ├── domain                          # domain (model) for usecases
│   │    ├── ports
│   │    │    ├── usecases.go                # usecase interface
│   │    │    └── repositories.go            # repository interface ex. DB, Redis, Third party
│   │    └── usecases                        # usecase implementation
│   ├── handlers
│   │     ├── router.go                      # http api
│   │     └── handles.go                     # handlers 
│   │
│   └── repositories                         # repository implementation
│
└── README.md
│
└── DockerFile
```
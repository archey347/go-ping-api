## goping-api

`goping-api` is a basic http api that serves pong when you request `/GET`. This project serves no specific purpose, other than as a learning exercise for myself on how to structure go projects.

Uses zap logging, viper config management, and chi.

### Building

`make build`

### Running

`make run`

This will run using the config file in `./resources/configs/goping-api.yaml`
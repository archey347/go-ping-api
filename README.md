## goping

`goping` is a basic http api that serves pong (just the word, not a nasty smell!) when you request `/GET`. This project serves no specific purpose, other than as a learning exercise for myself on how to structure go projects.

Uses zap logging, viper config management, and chi.

### Building

`make build`

### Running

`make run`

This will run using the config file in `./resources/configs/goping-api.yaml`

### Packaging

`make package`

Note that this package does not include the appropriate files to make it a system service (yet).

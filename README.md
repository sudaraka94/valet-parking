#Valet Parking

## Design Decisions
- While loading the configuration I decided to crash the application
in case of any err because, if there is any issue with the config, there is no point in continuing the execution
- You can find all the utility methods inside `util.go` file
- I tried to minimize the usage of third party libraries and do most of the 
implementations myself

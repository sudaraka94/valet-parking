# Valet Parking
[![Go](https://github.com/sudaraka94/valet-parking/actions/workflows/go.yml/badge.svg)](https://github.com/sudaraka94/valet-parking/actions/workflows/go.yml)
## Running the program on a bare Ubuntu 16.04 VM
1. Install Go
```shell
sudo apt-get update
sudo apt-get -y upgrade
wget https://dl.google.com/go/go1.16.4.linux-amd64.tar.gz
sudo tar -xvf go1.16.4.linux-amd64.tar.gz
sudo mv go /usr/local
```
2. Setup Go Env
```shell
export GOROOT=/usr/local/go
export PATH=$GOROOT/bin:$PATH
```
3. Clone Project
```shell
git clone https://github.com/sudaraka94/valet-parking.git
cd valet-parking
```
4. Run Tests
```shell
go test ./...
```
5. Run the program
```shell
go build -o .
./main
```
7. Run the program with a custom data file
```shell
./main -data=<path to datafile>
```
## Running with docker
1. Clone the repository
```shell
git clone https://github.com/sudaraka94/valet-parking.git
cd valet-parking
```
2. OPTIONAL : Replace the `datafile` inside project directory with the desired datafile 
3. Build docker image
```shell
docker build --tag valet-parking .
```
4. Run on docker
```shell
docker run -it valet-parking
```
## Design Decisions
- You can find all the utility methods inside `util.go` file
- I tried to minimize the usage of third party libraries and do most of the 
implementations myself
  
## Algorithm
Here is the basic algorithm used in this program. We can separate out the algorithm into
four sections,
#### 1. Adding a new vehicle to the park - `Overall time complexity : O(n)`
- Look for an available slot with the least slot number - `Time Complexity: O(n)`
- Mark the slot as occupied -  `Time Complexity: O(1)`
- Keep reference of the slot in a hashmap with the vehicle registration number -  `Time Complexity: O(1)`

#### 2. Removing a vehicle from the park - `Overall time complexity : O(1)`
- Lookup the vehicle in the hashmap using the vehicle registration number - `Time Complexity: O(1)`
- Remove the vehicle from the hashmap and mark the slot available - `Time Complexity: O(1)`
- Calculate the fare - `Time Complexity: O(1)`

## Logger implementation
Throughout the whole project, Logger is exposed through an interface. Currently logger interface is 
implemented by cliLogger, which directly writes to the std out. But in case we need to writeout the results to a file,
we can easily do so by introducing a Logger implementation which writes to a file. Logger type can be made configurable
via config.yml 
```yaml
logger_config:
  logger_type: "cli" #allowed logger types: cli
```

## Configurability
In this application, supported types of vehicles are configurable. On the application start,
it fetches configurations from `confi.yml` and adjusts accordingly. For an example, if we need to
add a third vehicle type we just have to define the vehicle type in the `config.yml`.

eg:
```yaml
vehicle_types:
  - name: car
    price_per_hour: 2 #pricePerHour in USD
  - name: motorcycle
    price_per_hour: 1 #pricePerHour in USD
  - name: van
    price_per_hour: 3 #pricePerHour in USD
```

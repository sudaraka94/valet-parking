# Valet Parking
## Running the program on a bare ubuntu 16.04 vm
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
2. OPTIONAL : Replace the datafile with the desired datafile 
3. Build docker image
```shell
docker build --tag valet-parking .
```
4. Run on docker
```shell
docker run -it valet-parking
```
## Design Decisions
- While loading the configuration I decided to crash the application
in case of any error because, if there is any issue with the config, there is no point in continuing the execution
- You can find all the utility methods inside `util.go` file
- I tried to minimize the usage of third party libraries and do most of the 
implementations myself

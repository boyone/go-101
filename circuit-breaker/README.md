# Demo Go app with Circuit Breaker

## Move circuit breaker to Gopath

```sh
cd ..
cp -r circuit-breaker $GOPATH/src/
cd $GOPATH/src/circuit-breaker
```

## Build and run front

```sh
cd $GOPATH/src/circuit-breaker/cmd/front
go build
./front
```

## Build and run back

```sh
cd $GOPATH/src/circuit-breaker/cmd/back
go build
./back
```

## Build Hystrix-Dashboard

```sh
cd $GOPATH/src/circuit-breaker/hystrixDashboard
mvn package
java -jar target/hystrixDashboard-0.0.1-SNAPSHOT.jar
```

## Open Hystrix-Dashboard

* Go To: http://localhost:8080/hystrix
* Add Single Hystrix App: http://localhost:8881  
* Click Monitor Stream

## Call Service

* New Tab
* Go To: http://localhost:3000/account 
* Should see acount: 1234
* Should see green color on hystrix-dashboard

## Test Circuit Breaker

* Kill back
* Connect to http://localhost:3000/account at lease 30 times
* Should see red color on hystrix-dashboard
* Run ./back
* Connect to http://localhost:3000/account
* Hystrix-dashboard should back to green

## Option Connect http://localhost:3000/account with load test tools

```sh
go get -u github.com/rakyll/hey
hey -z 10s http://localhost:3000/account
```
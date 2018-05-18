# Create Hystrix-Dashboard with springboot
![logo](./hystrix.png)
* go to: [start.spring.io](start.spring.io)  
* create project  
* add dependecies
  * web
  * hystrix-dashboard

## Add this lines to main class

```java
import org.springframework.cloud.netflix.hystrix.dashboard.EnableHystrixDashboard;

@EnableHystrixDashboard
```

## Build

```java
$mvn package
```

# Run
```java
$cd target
$java -jar XXXX.jar
```

# URL : http://localhost:8080/hystrix
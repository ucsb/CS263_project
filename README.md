# CS 263 Final Project, Fall 2021

## Project Description and Goals

For this project we will be surveying and analyzing the ways in which developers write microservices today. We hope to come to an understanding of what frameworks, languages, libraries, and platforms are popular among microservices developers, and gain experience on how these tools work. How do they compare to each other and how easy are they to use from an application developers perspective? Is there a lot of community support and documentation for these tools? What types of applications are best implemented as microservices and what are the pros and cons of developing an application as a microservices? Ultimately we aim to learn more about how applications designed using a microservice architecture are written, and how the technologies involved in microservice implementation perform in terms of ease of use. 

## Mircroservice Applications Implemented

### Doughnut Shop
We implemented serveral applications using a microservice architecture. One of the applications we developed was a mock doughnut shop. In this application we first began by creating three different microservices using Spring Boot with Spring Cloud: a eureka server, a frontend service, and an inventory service. The Eureka server is responsible for acting as a service registry and registers all microservices within the application with details such as their name and location. The Frontend displays doughnut inventory and communicates with the Inventory Service in order to get the inventory of avaliable doughnuts. The Inventory Service initally contains a hardcoded list of doughnuts that includes information about doughnuts such as their name, price, and quantity. The Inventory Service is responsible for sending this hardcoded list of doughnuts to the Frontend when requested. This application is contained in the doughnut_shop directory.

The application contained in the doughnut_shop_with_database directory builds off the application in the doughnut_shop direcory. Instead of using a hardcoded list of doughnuts however, data is now stored and retrevied from MongoDB Altas. The Frontend service has also been modified in this version of the application to display doughnuts in a different format and in order of price.

Finally, the application contained in the doughnut_shop_using_flask_and_spring_cloud directory builds off the previous application implemented using Spring Cloud and Spring MongoDB. In this application we've added a fourth microservice written in Python using the Flask framework and py-eureka-client. This microservice acts as a checkout service. Users make a post request to the checkout service which includes the doughnut name and the quantity of doughnut that the user wishes to buy. The checkout service then communicates with the inventory service, which decreases the doughnut inventory in the database, and returns a message to the checkout service on success. The Checkout Service keeps a list of these doughnut orders. The Frontend then communicates to the Checkout Service and requests the doughnut orders, which it displays at the orders route.

In order to run this final microservice application, the services must be started in a specific order. First the Eureka server must begin running, then the Inventory Service, then the Frontend Service, and finally the Checkout Service. Below we walk through an example illustrating how the doughnut shop application works. Because of the limited functionality of the Frontend, certain requests are made using Python that would otherwise be made using the Frontend service.

The services are configured to start at the following locations:

Eureka Server: localhost:8761

Inventory Service: localhost:8085

Frontend Service: localhost:8081

Checkout Service: localhost:8084

### Instructions to Build and Deploy
                    
First start each mircoservice in the order described above by running the following commands.

For the Eureka Server, Inventory Service, and Frontend Service:

Make sure you have Maven installed and configured correctly. Then change into the correct directories and run:

```./mvnw spring-boot:run```

For the Checkout Service:

First run

```cd doughnut_shop_using_flask_and_spring_cloud/checkout-service```

```python3 -m venv flask-microservice```

Then run either

on Mac/Linux: ```source flask-microservice/bin/activate``` 

on Windows: ```flask-microservice/Scripts/activate.bat``` 

Then run:

```pip3 install py-eureka-client flask requests```

```python3 flask_app.py```

You should now be able to access http://localhost:8084/orders successfully. Additionally you should see all services registered with the Eureka server after refreshing http://localhost:8761/. 

You can view the doughnuts currently in the database at http://localhost:8085/doughnuts/. In order to create, update, or delete a doughnut in the database, edit client.py accordingly and run:

```python3 client.py```

You should see the corresponding changes displayed at http://localhost:8085/doughnuts/ upon refreshing the page. 

The Frontend service which requests this list of doughnuts from the Inventory Service will display them at http://localhost:8081/doughnuts/ in a different format and in ascending order by price. When a doughnut is added, updated, or deleted the Inventory Service will communicate this information to the Frontend Service, which should reflect this change. 


You can view the doughnut orders from the Checkout Service at http://localhost:8084/orders. Because of the limited functionality of the Frontend Service, in order to purchase a doughnut and hence add it to the orders list, edit client.py accordingly and run:

```python3 client.py```

After making a POST request to the Checkout Service in order to purchase a doughnut, you can refresh http://localhost:8084/orders and see that the doughnut order has been added. Upon refreshing the Frontend Service you can see that the doughnut inventory has decreased. You can see these same changes reflected in the Inventory Service as well. This is because the Checkout Service communicates with the Inventory Service, which communicates with the Frontend service to reflect this change in inventory. Finally, you can see on the Frontend by visitng http://localhost:8081/doughnuts/orders that the Checkout Service sends the list of orders to the Frontend upon request.

This application demonstrates how microservices communicate and exchange information, and how important it is for communication between microservices to be quick. It also illustrates how applications are divided into seperate services and have their own functionality. 

### SleepTracker

Another sample microservice we wrote was SleepTracker, a simple program for calculating the average number of hours you slept. SleepTracker was written in Go and uses gRPC as the method of communication between the server and client, as well as Protocol Buffers for exchanging data. 

To run SleepTracker locally on the commandline:

1. Install the latest version of Go (see [here](https://go.dev/doc/install) for installation instructions).
2. Open your terminal and navigate to your home directory by typing ```cd``` into your terminal.
3. Git clone the SleepTracker repo (found [here](https://github.com/aprilsanchez/SleepTracker)).
4. cd into the cloned repository.
  ```cd SleepTracker```
5. You may need to add ```$GOPATH/bin``` to your ```$PATH``` so that executables installed via ```go get``` are available on your $PATH.
  For example, on MacOS, you can type the following into the terminal:
  ```export PATH=$PATH:$(go env GOPATH)/bin```
7. Run the server by typing the following command into your terminal:
  ```go run sleep_tracker/server/server.go```
8. Open another terminal window and navigate to the SleepTracker directory again.
9. Run the client by typing the following command into your terminal:
  ```go run sleep_tracker/client/client.go```
10. Follow the commands prompted by the client to interact with the SleepTracker.

Note: SleepTracker was implented to gain familiarity with Go, gRPC, and Protocol Buffers. It was not designed to be fully \
functional and as such, may contain bugs.

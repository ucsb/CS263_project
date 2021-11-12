## What is a Microservice?

A microservice architecture designs applications as a collection of small autonomous services, each with their own unique functionality. Each of these small modularized services are referred to as a **microservice**. Microservices communicate with each other and work together to create a fully functional multi-feature application.

## What are the Pros and Cons of Microservices?

The independent nature of microservices allow applications to be more efficiently and reliably implemented, and lends well to the agile development model for software teams that do utilize this methodology. Because each microservice performs a specialized task independent from the other, microservices can be added, removed, and updated without interfering with the functionality of other microservices that contribute to the application at large. This means that issues in one microservice may not necessarily cause a failure in the rest of the application (improved fault isolation), and that parts of the application can be upgraded easily. The modularity of the application often leads to increased programmer productivity as well, as code is not only more readable, but developers are also able to pinpoint and resolve issues quickly when testing in smaller units. Developers can also focus on scaling individual components in the application that need to be scaled as opposed to modifying the entirety of the application, and deployment is arguably simpler and faster for these smaller services.

Despite their many advantages, applications designed using microservices can also be difficult to implement because of the need for these microservices to seamlessly communicate with each other. Communication becomes more complicated when considering the asynchronous nature of network communication and experiences with latency. Adding on to that, more services generally means the use of more resources. For example, instead of sharing a database with multiple tables, separate microservices may need to utilize their own database. Contrary to what was mentioned earlier, testing and deploying with microservices can also be more difficult. This is because each service has its own set of logs that require shifting through, and global testing is almost impossible without each microservice being individually tested first. While deployment of a single microservice may be easier, orchastering and coordination among multiple services when deploying is complicated. Because designing a microservice can be difficult, and time consuming some applications are better suited to being designed as microservices than others.

## What Applications are Best Suited to be Written as Microservices?

TODO

## What are Some Tools Used to Develop Microservices? 
### Spring Boot with Spring Cloud
### *Why use Spring Boot with Spring Cloud to Develop Microservices?*

Spring Boot includes libraries with user-friendly APIs that make it easy to deal with many of the difficulties that come with writing Microservices. For example, dealing with partial failures, an issue that arises in many distributed systems of which microservices are not exempt, is made easy with Spring Boot’s fault tolerance library Resilience4J. Resilience4J provides various options to handle partial failures including circuit breakers, retry, timeout, rate limiters, and fallback handlers. Spring Cloud also adopts Eureka, an open source library created by Netflix, which handles service discovery for developers. Eureka uses client-side-discovery, and additionally implements load balancing. The daunting task of microservice communication is also made easy in Spring Boot with libraries such as RestTemplate and WebClient. Some added benefits of Spring Cloud include cloud native support (easy to deploy on Docker and Kuberenetes), as well as libraries to make load balancing and clustering easy. In the microservice we wrote using Spring Cloud, we focused primarily on experimenting with the RestTemplate and Eureka libraries, with a small side investigation into WebClient.

## What are Some Tools Used to Deploy Microservices? 

(Discuss the learning curve of these tools, availability of documentation/community help/support, ease of development and performance for example microservice benchmarks and applications.)

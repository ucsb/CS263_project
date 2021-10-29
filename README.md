# CS 263 Final Project, Fall 2021

## What is Ambience?

Ambience is an operating system designed to support distributed applications structured as microservices. Ambience is designed to overcome two major challenges: the inability to run a common OS on multiple and various hardware platforms, and the negative impact general commodity operating systems can have on performance.

## Project Description and Project Goals

For this project, we will be implementing applications (e.g. microservices) to run on Ambience and identifying areas that are lacking from an application developer’s perspective. For instance, if an application developer wishes to run their microservice on Ambience, what are the tools needed to streamline this process that are not already present in Ambience? What additional abstractions could be added to Ambience to facilitate the application deployment process?

We will begin by creating a sample microservice and running it using Kubernetes to get a feel for what is involved in deploying a service on a system like Kubernetes. Then, we will take the same service, try to run it on Ambience, and compare and contrast the experience of running the microservice on Ambience as opposed to Kubernetes. Ultimately, we aim to identify ways in which Ambience can be improved. 

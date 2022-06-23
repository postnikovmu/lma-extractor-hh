# Project name
***
Go microservice for getting data from API hh.ru
## General Info
***
This is a microservice, which can be used to grab information
about useful skills for certain profession.

Example of request (local usage):
http://localhost:8080/hh4?area=<Example_City>&text=<Example_developer>

Responce will the data in JSON format.

Deploy on SAP BTP CloudFoundry platform: the settings is in manifest.yml
## Technologies
***
Go (Golang)
Web API https://github.com/hhru/api
JSON
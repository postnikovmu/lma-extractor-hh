# lma-extractor-hh
Go microservice for getting data with usage of  WebAPI hh.ru

![](https://img.shields.io/github/languages/code-size/postnikovmu/lma-extractor-hh)
![](https://img.shields.io/github/directory-file-count/postnikovmu/lma-extractor-hh)
![](https://img.shields.io/github/languages/count/postnikovmu/lma-extractor-hh)
![](https://img.shields.io/github/languages/top/postnikovmu/lma-extractor-hh)

## General Info

This is a microservice, which can be used to grab information
about useful skills for certain professions.

Example of request (local usage):
http://localhost:8080/hh4?area=<Example_City>&text=<Example_developer>

Responce is the data in JSON format.

## Install

### Prerequisites:
There are no prerequisites.

### Deploy:
Deploy is on SAP BTP CloudFoundry platform: the settings are in manifest.yml. (or another CloudFoundry)
## Technologies

Go (Golang)

Web API https://github.com/hhru/api

JSON

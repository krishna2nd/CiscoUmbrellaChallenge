# README #

Application is for url lookup service. Which provides the input url is safe to proceed or not.
Server listen to 9090 by default.

End points
#### GET urlinfo/1/
	Provides information to client that the url is safe to access or not
	url to check has to be in urlencoded format if it consists query parameters

eg:

	192.168.64.3:9090/urlinfo/1/gogle.com%2Fq%3Dtest
	which provide following json result for safe url

	```{
		"status": "safe"
		"message": ""
	}```

	192.168.64.3:9090/urlinfo/1/goglee.com%2Ftest%3Fhawk`
	
	which provide following json result for unsafe url
	
	```{
		"status": "unsafe"
		"message": ""
	}```

#### POST /urls
	If you have any records of some unsafe urls and parameters, you can uload to lookup service.
	POST raw body with new line serperated formated will upload to lookup service

eg: 

	http://192.168.64.3:9090/urls
	
	body
		```goglee.com/test?test=10000
		goglee.com/test?hawk```

	
	For MAC users: docker-machine ip <machine-name> will give you the docker machien ip


### What is this repository for? ###

* Repository contains the solution code for the problem "url lookup service".
* Implementation is in golang.


### How do I get set up? ###

* Summary of set up
	make build - To run the production containers and access website via 9090
	make test - To test code in test containers
	make run-test - perform the unitest on go modules

* Dependencies
	golang 1.8
	Docker
	Docker-compose

* Database configuration
	Redis is the database used for this implemenation. not configured for persistance (It can be done).
	Default database, no password.

* How to run tests
	make test will be executing the unit tests in containers
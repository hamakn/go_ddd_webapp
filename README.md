# my Golang DDD webapp example

## Features

* GAE/datastore
* Multiple unique constaints on datastore
* Layered Architecture
* No framework library
* Use [negroni](https://github.com/urfave/negroni) to handle HTTP
* Use [goon](https://github.com/mjibson/goon) for autocaching

## Boot
````
% dev_appserver.py config/app.yaml
````

## Request examples
````
% curl -X POST -d '{"email": "foo@foo.test", "screen_name": "foo", "age": 17}' http://localhost:8080/users/
{"ID":5629499534213120,"Email":"foo@foo.test","ScreenName":"foo","Age":17,"CreatedAt":"2018-04-08T10:21:07.617449Z","UpdatedAt":"2018-04-08T10:21:07.617449Z"}

% curl http://localhost:8080/users/
[{"ID":5629499534213120,"Email":"foo@foo.test","ScreenName":"foo","Age":17,"CreatedAt":"2018-04-08T10:21:07.617449Z","UpdatedAt":"2018-04-08T10:21:07.617449Z"}]

% curl http://localhost:8080/users/5629499534213120
{"ID":5629499534213120,"Email":"foo@foo.test","ScreenName":"foo","Age":17,"CreatedAt":"2018-04-08T10:21:07.617449Z","UpdatedAt":"2018-04-08T10:21:07.617449Z"}

% curl -X POST -d '{"email": "foo@foo.test", "screen_name": "new", "age": 17}' http://localhost:8080/users/
{"error":"Unprocessable Entity"}

% curl -X POST -d '{"email": "new@foo.test", "screen_name": "foo", "age": 17}' http://localhost:8080/users/
{"error":"Unprocessable Entity"}

% curl -X PUT -d '{"screen_name": "new"}' http://localhost:8080/users/5629499534213120
{"ID":5629499534213120,"Email":"foo@foo.test","ScreenName":"new","Age":17,"CreatedAt":"2018-04-08T10:21:07.617449Z","UpdatedAt":"2018-04-08T10:22:59.279485Z"}

% curl -X DELETE http://localhost:8080/users/5629499534213120

% curl http://localhost:8080/users/
[]
````

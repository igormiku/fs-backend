## About fs-backend
Backend for Feature Toggle Management. Includes following API endpoints
1. /api/v1/features - GET all features toggles
2. /api/v1/features/2 - GET one feature toggle by id
3. /api/v1/features/add - POST add new feature toggle to the list
4. /api/v1/features - PUT update one feature toggle by id
5. /api/v1/features - DELETE delete one feature toggle by id
6. /api/v1/features - POST endpoint allows to query what features are ON or OFF given a customer and features of interest

## Getting Started fs-backend
This project was build on go version 1.17.3

## Configuration
Configuration resides in config.json file.

### Storage 
Two persistence layers are available BuntDB (file based  storage) or MockDB(in memory go arrays). MockDB persistence layer was tested, BuntDB not yet.

### Use config.json to configure 
```
    "DBType": "MockDB",			//MockDB in memory, or FileDB for BuntDB (https://github.com/tidwall/buntdb)
    "DBFile": "/db/data.db",	//file location for BuntDB
    
    "listen": "0.0.0.0:10000",	//server IP/port configuration
    "APIBase": "api",			//API base
    "APIVersion": "v1"			//API version
```


## Build and Run instructions
Open Terminal and cd root of project
```
make
build/bin/ft_backend-1.0
```

## Some curl for backend checks
```
QUERY ALL 
curl -X GET -H 'content-type: application/json' http://localhost:10000/api/v1/features

QUERY BY ID 
curl -X GET -H 'content-type: application/json' http://localhost:10000/api/v1/features/2

ADD
curl -X POST -H 'content-type: application/json' --data '{"technicalname": "test", "displayname":"Display Test", "expiresOn":1637958974, "customerIds":["1","2"]}' http://localhost:10000/api/v1/features/add

UPDATE 
curl -X PUT -H 'content-type: application/json' --data '{"id": 2, "technicalname": "github", "displayname":"GitHub", "description": "Usage of GitHub feature", "inverted":false, "expiresOn":"123", "customerIds":["3","5"]}' http://localhost:10000/api/v1/features

DELETE
curl -X DELETE -H 'content-type: application/json' --data '{"id": 1}' http://localhost:10000/api/v1/features

QUERY by Customer  (Example API in requirement document)
curl -X POST -H 'content-type: application/json' --data '{"customerId": "1", "features":["github","bitbucket"]}' http://localhost:10000/api/v1/features
```

## ASSUMPTIONS & Remarks
1. No security requirements implemented(e.g. authentication, authorization)
2. No Unit test for this project
3. Inverted default is false, which means feature is available for a customers, if true it is not.
4. FeatureToggle entity has an id, which is supposed to be unique key, but unique checks are not implemented. For new entries random id is provided. 
5. Archiving means deletion
6. Mind Chrome CORS rules, don't allowing cross origin, bypassed for this project. Made OPTIONS method to reply HTTP 204 to compile with Chrome rules
7. API example in requirement document is implemented, and available only in backend call.(refer to above in this readme for curl to check it)
8. Datetime for FeatureToggle is stored as string in unix datetime format (shell sample commands: date +%s, date -r 1638043542)


## Potential TODOs:
1. Dedicated Server package to load from config and run Listener 
2. Test real DB CRUD used in backend. Tested only MockDB as persitence layer

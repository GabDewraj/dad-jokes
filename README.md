# Dad-jokes

## A fun little API for your dad jokes !

## This API allows a user to do the following

- Store a new dad joke.
- Get a list of dad jokes.
- Get a random dad joke.

## How to run the is project locally 

- Navigate to the root folder of this project.
- and paste this command into your terminal : ```sh server/build/run.sh```
- This command achieves the following:
    - Composes all containers to be used.
    - Initialises the database and seeds it with some dummy data.

## API'S Present and their descriptions.

### Health-Check

- Description: This endpoint allows for health check probes from a load balancer, target group and container to ensure service is running.
- Method: `` Any method is fine``
- Endpoint: ``/health-check``
- Port: ``8080``
- Reponses:
    - Status OK (200) body format: empty.

### Generate an API key

- Description: This endpoint allows a user to retrieve an api key.
- Method: ``GET``
- Endpoint: ``/apikey``
- Port: ``8080``
- Reponse: ``Here is your key tgZ6bLqzjwoZvoD2iThbPS``

### Create a new dad joke.

- Description: This endpoint allows an authorized user to save his or her own dad jokes.
- Method: ``Post``
- Endpoint: ``/new/jokes``
- Port: ``8080``
- Authorization:   API key must be present in the header                                     ``X-Dad-Jokes-Access-Token``

- Request Body Data Type: ``application/json``
- Request Body format: 
    ``
        "body": string,
        "author": string,
    ``

- Response Body data type: ``application/json``
- Responses
    - Status OK(200) Body format: ``"message":"Joke with id 10 saved successfully"``
    - Bad Request (400) Body format: 
        - `` json: cannot unmarshal number into Go struct field Joke.author of type string ``
        - `` json: cannot unmarshal bool into Go struct field Joke.author of type string ``
    - Unauthorized: `` secret not found for key ``

### Get a list of dad jokes.
- Description: This unauthenticated endpoint allows a user to retrieve a paginated list 
            of dad jokes.
- Method: ``GET``
- Endpoint: ``/jokes``
- Port: ``8080``
- Pagination Type: ``offset``
- Query Params: 


        limit: string *number of results to receive.

        offset: string *number of entries you want to skip.


- Response Body data type: ``application/json``
- Responses
    - Status OK (200) Response Body format: 


            [
                {
                    "id": 3,
                    "body": "I broke my arm in two places. <>My doctor told me to stop going to those places.",
                    "author": "github.com",
                    "createdAt": "2022-09-07T05:29:00.807276Z",
                    "updatedAt": "2022-09-07T05:29:00.807276Z"
                },
                {
                    "id": 4,
                    "body": "I quit my job at the coffee shop the other day. <>It was just the same old grind over and over.",
                    "author": "github.com",
                    "createdAt": "2022-09-07T05:29:00.807276Z",
                    "updatedAt": "2022-09-07T05:29:00.807276Z"
                },

                ...
            ]
    - Bad request (400) Body format : ``strconv.Atoi: parsing "": invalid syntax``


### Get a Random Dad Joke.

- Description: This unauthenticated endpoint allows a user to retrieve a random dad jokes.
- Method: ``GET``
- Endpoint: ``/random/jokes``
- Port: ``8080``
- Response Body data type: ``application/json``
- Status OK (200) Response Body format: 

        {
            "id": 4,
            "body": "I quit my job at the coffee shop the other day. <>It was just the same old grind over and over.",
            "author": "github.com",
            "createdAt": "2022-09-07T05:29:00.807276Z",
            "updatedAt": "2022-09-07T05:29:00.807276Z"
        }
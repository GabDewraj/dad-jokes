# Dad-jokes

## A fun little API for your dad jokes !

## This API allows a user to do the following

- Store a new dad joke.
- Get a list of dad jokes.
- Get a random dad joke.

## How to run the is project locally 

- This small project includes a microservice labelled as "server" (just for context) and 
    some code to initialise the database structure with seeding ("db").
- Navigate to the root folder of this project (dad-jokes).
- Paste this command into your terminal : ```sh server/build/run.sh```
- This command achieves the following:
    - Composes all containers to be used.
    - Initialises the database and seeds it with some dummy data.

## API's present and their descriptions.

### Health-Check

- Protocal: HTTPS
- Description: This endpoint allows for health check probes from a load balancer, target group and container to ensure service is running.
- Method: `` Any method is fine``
- Endpoint: ``https://localhost:8080/health-check``
- Responses:
    - Status OK (200) body format: empty.

### Generate an API key

- Protocal: HTTPS
- Description: This endpoint allows a user to generate and retrieve a personalized api key.
- Method: ``GET``
- Endpoint: ``https://localhost:8080/apikey``
- Response: ``Here is your key tgZ6bLqzjwoZvoD2iThbPS`` (*this is just an example)

### Create a new dad joke.

- Protocal: HTTPS
- Description: This endpoint allows an authorized user to save his or her own dad jokes.
- Method: ``Post``
- Endpoint: ``https://localhost:8080/new/jokes``
- Authorization:   API key must be present in the header                                     ``X-Dad-Jokes-Access-Token``

- Request Body Data Type: ``application/json``
- Request Body format: 
    ``
       {
         "body": string,
        "author": string,
        }
    ``

- Response Body data type: ``application/json``
- Responses
    - Status OK(200) Body format: ``Joke with id 11 and author Gabriel saved successfully``
    - Bad Request (400) Body format: 
        - `` json: cannot unmarshal number into Go struct field Joke.author of type string ``
        - `` json: cannot unmarshal bool into Go struct field Joke.author of type string ``
    - Unauthorized(401): `` secret not found for key ``

### Get a list of dad jokes.

- Protocal: HTTPS
- Description: This unauthenticated endpoint allows a user to retrieve a paginated list 
            of dad jokes.
- Method: ``GET``
- Endpoint: ``https://localhost:8080/jokes?limit=8&offset=10`` (*this is just an example, you can change the limit and offset)
- Pagination Type: ``offset``
- Query Params: 


        limit: number of results to receive.

        offset: number of entries you want to skip.


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

- Protocal: HTTPS
- Description: This unauthenticated endpoint allows a user to retrieve a random dad jokes.
- Method: ``GET``
- Endpoint: ``https://localhost:8080/random/jokes``
- Response Body data type: ``application/json``
- Status OK (200) Response Body format: 

        {
            "id": 4,
            "body": "I quit my job at the coffee shop the other day. <>It was just the same old grind over and over.",
            "author": "github.com",
            "createdAt": "2022-09-07T05:29:00.807276Z",
            "updatedAt": "2022-09-07T05:29:00.807276Z"
        }

## Approach

- Understand deliverables carefully to produce the desired API behavior as described.
- Create a quick an easy project that can be deployed (infrastructure code was ommitted due to time) and run locally.
- Ensure data is protected from sql injection by using an ORM that actively prevents against this.
- GORM uses the database/sql???s argument placeholders to construct the SQL statement,
which will automatically escape arguments to avoid SQL injection
- Use a microservice and domain driven approach that is go idiomatic for ease of reading.

## Assumptions

- A cache (Redis) would be used in production to store the API keys of respective 3rd party APIs between being stored in an RDS after the expired ttl.
- Use an offset type pagination to retrieve a list of dad jokes.
- This API may serve both 3rd party API users and Frontend Browser clients, that is why CORS is included as apart of the middleware stack.


## Improvements given more time

- Include gh workflows to include unit testing and code linting to ensure good code coverage
and clean code.
- Increased unit testing of code (Positive and negative cases).
- Better error handling with more specific and cleaner error descriptions.
- Inclusion of a more appropriate API documentation tool such as  OPENAPI Swagger.
- Inclusion of infrastructure code for build and deployment (buildspec.yml,appspec.yml and taskDef.json), this is if using aws managed services (CODEBUILD,CODEDEPLOY)
- Durable Storage implementation to rotate API keys between a Redis Cache for performance and long term storage in PostgreSQL.
- Spend more time creating a more robust ERD that can be extensible to future usecases.
- Take more time to think about better RESTful endpoint naming.






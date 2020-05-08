# Go-PostgreSQL
Restful API to perform CRUD operations.
## To run :
- Run the below command in terminal
```
 go run main.go
```
## CRUD
- ### Create
URL format \
`POST /todo`

Send body containing, 

     { 
       "title":"Todo title, 
       "body" : "what to do",  
       "completed" : "status"   
     }
- ### Get All todos
URL format  \
`GET /todos`

- ### Get a todo
URL format \
`GET /todo/{id}`

- ### Edit a todo
URL format   \
`PUT /todo/{id}`

Send body containing,


     { 
       "title":"Todo title, 
       "body" : "what to do",  
       "completed" : "status"   
     }
`

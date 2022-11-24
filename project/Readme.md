# ToDo App

## User Story:
### Part One:
    -   As a User i want to be able to add and delete my todo's to an list
    -   The Input takes place via an Form field.
    -   After submitting the the todo entry, the todo list is updated automaticly

### Part Two:
    - As a User i want to be able to login into an existing account
    - I want further more be able to save my todo list

### Part Three:
    - As a User I want to be able to create diffrent todo lists and save them.
    - As a User I want to be able to create sub todolists for an todo



## The Plan: 
    - first creating the front end
        - create the right folder structur
        - init go project
        - create html templates
            - base, header, footer, page
            - CSS and Javascript extra


    - second the server and routes
        -   server, handler, routes
    - third save it in and DB


## Usefull stuff: 

## Commands: 
 - go mod init front-end
 - protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative dbs.protoc

 ## ToDo

 ### DB-Service
    - Check if Repostitory is needed

 ### Rename API URLs
    - change urls from addTodo -> add-todo
    - change from checkTodo -> check-todo

 ### Delete Todo:
    - add an Icon to the front end and see if its rendered corectly
    - add the eventListener -> Post to delete-todo
    - create the gRPC function DeleteTodo
    - add DeleteTodo function to DB-service

    - add Route in distribution service
    - add Handerl in Distribution service
    


### Form Handleing

### Account
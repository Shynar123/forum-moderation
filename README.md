# Forum

## Description
   
  This project consists in creating a web forum that allows :

  * communication between users.
  * associating categories to posts.
  * liking and disliking posts and comments.
  * filtering posts.

  SQLite should be used with at least one SELECT, one CREATE and one INSERT queries

  Instructions for user registration:

  * Must ask for email
      * When the email is already taken return an error response.
    * Must ask for username
    * Must ask for password
        * The password must be encrypted when stored (this is a Bonus task)

Project should have Likies, Dislikes, Comments and filter


## Authors
@dokshiye, @dtyuligu nd @szhigero

## Usage
```
$ git clone git@git.01.alem.school:szhigero/forum.git
```
To run using docker:
```make run```

To stop:
```make stop```

To run without docker:

``` go run ./cmd .```

Go to ```http://localhost:8080```
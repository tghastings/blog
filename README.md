# Hastings Blog
[![Build Status](https://travis-ci.org/tghastings/blog.svg?branch=master)](https://travis-ci.org/tghastings/blog) [![Go Report Card](https://goreportcard.com/badge/github.com/tghastings/blog)](https://goreportcard.com/report/github.com/tghastings/blog) 


`Hastings Blog` is a lightweight blogging platform that provides a RESTful API which serves JSON and utilizes CRUD functions to manipulate data in SQL databases for posts and users. The platform leverages JSON Web Tokens for authentication and bcrypt to protect sensitive data for persistent data storage.

# Features
* [x] Postgres, MySQL, SQLite and Foundation database support
* [x] CRUD framework for users and posts
* [x] User authentication using JWTs and bcrypt with salt for secure password storage

## Motivation
This is my first project built with golang. I have only been playing with it for a week and building a blogging platform is usually the first project I make with a new language.

## Road Map
* [x] API which allows CRUD for user management, posts, and user authentication
* [ ] Automated API testing
* [ ] Swagger API documentation
* [ ] Javascript front-end for readers to read posts
* [ ] Javascript front-end for authors to manage users and posts

# Usage
Below is a quick snippet on how to use the platform.

## Default Username and Password
username: `root` 
Password: `12345` 

## Routes
There are a number of routes defined for viewing and manipulating data for users and posts. The routes are distinguished into two groups, unauthenticated and authenticated. Unauthenticated routes do NOT require authentication. Authenticated routes DO require authentication using a JWT.

### Unauthenticated Routes
The routes below do not require authentication are accessible by everyone.

`GET http://localhost:8090/` returns a list of all posts
`GET http://localhost:8090/post/{id}` returns one post with the specified ID


### Authenticated Routes
`PUT http://localhost:8090/auth` returns a cookie with `Token`. Use this JSON Web Token's value in a header for API requests.

```
{
  "username" : "root",
  "password" : "12345"
}
```

# Licence

This project is released under the MIT licence. See [LICENCE](LICENCE) for more details.

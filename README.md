# Hastings Blog
[![Build Status](https://travis-ci.org/tghastings/blog.svg?branch=master)](https://travis-ci.org/tghastings/blog) [![Go Report Card](https://goreportcard.com/badge/github.com/tghastings/blog)](https://goreportcard.com/report/github.com/tghastings/blog) 


`Hastings Blog` is a lightweight blogging platform that provides a RESTful API which serves JSON and utalizes CRUD functions to manipulate data for posts and users. The platform leverages JSON Web Tokens for authentication and bcrypt to protect sensative data for persistent data.

# Features
* [x] Postgres, MySQL, SQLite and Foundation database support
* [x] CRUD framework for users and posts
* [x] User authentication using JWTs and bcrypt with salt for secure password storage


## Motivation
This is my first project built with golang. I have only been playing with it for a week and building a blogging platform is usually the first project I make with a new language.

## Road Map
* [x] API which allows CRUD for user management, posts, and user authentication
* [ ] Automated API testing
* [ ] Javascript front-end for readers to read posts
* [ ] Javascript front-end for authors to manage users and posts
* [ ] Swagger API Documentation

# Licence

This project is released under the MIT licence. See [LICENCE](LICENCE) for more details.

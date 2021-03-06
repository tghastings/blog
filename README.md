# Hastings Blog
[![Build Status](https://travis-ci.org/tghastings/blog.svg?branch=master)](https://travis-ci.org/tghastings/blog) [![Go Report Card](https://goreportcard.com/badge/github.com/tghastings/blog)](https://goreportcard.com/report/github.com/tghastings/blog) 


`Hastings Blog` is a lightweight blogging platform that provides a RESTful API which serves JSON and utilizes CRUD functions to manipulate data in SQL databases for posts and users. The platform leverages JSON Web Tokens for authentication and bcrypt to protect sensitive data for persistent data storage.

# Features
* [x] Postgres, MySQL, SQLite and Foundation database support
* [x] CRUD functions for users and posts
* [x] User authentication using JWTs and bcrypt with salt for secure password storage

## Motivation
This is my first project built with golang. I've worked on this project for a little over 3 days and building a blogging platform is usually the first project I make with a new language.

## Road Map
* [x] API which allows CRUD for user management, posts, and user authentication
* [ ] Automated API testing
* [x] Swagger API documentation
* [ ] Javascript front-end for readers to read posts
* [ ] Javascript front-end for authors to manage users and posts

# Usage
There are two ways to start the blog platform. The first is to download the binary and the second is to clone the repository, build, and run. 

## From Binary
Download the binary [Download](https://res.cloudinary.com/innopar/raw/upload/v1546122423/blog-0.0.1.tar_b9s505.gz)

Run `./blog`

## From Source
Clone the repo `git clone git@github.com:tghastings/blog.git`

`cd blog`

`go build`

`./blog`

## Default Username and Password
username: `root` 
Password: `12345` 

# [View the API Documentation](https://tghastings.github.io/blog/)
The API is documented using Swagger. [View the API Documentation](https://tghastings.github.io/blog/).


# Licence
This project is released under the MIT licence. See [LICENCE](LICENCE) for more details.

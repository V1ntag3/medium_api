
# NewMedium - A Medium Clone API


This project is a clone of the https://medium.com article site where users can post articles on a wide range of subjects


## API Reference

Documentation API: https://marcos-team.postman.co/workspace/4927b179-fda5-47d3-b94c-14c71518e289/documentation/17463653-e3d2ac4f-a213-4357-8ddc-a0bb9dc8035f

## Front End

https://github.com/V1ntag3/medium_clone_front

## Lessons Learned

I put into practice the development of REST APIs in addition to learning more about the functioning of the Golang language and its GORM and Fiber libraries to build a safe and fast application



## Features

- Auth User
  - Login
  - Register
  - Logout
- User
  - Updade user data
  - Delete user
  - Follow user
  - UnFollow user
  - List Followings
  - List Followers
- Upload Image Profile and Banner
- Articles
  - List Articles
  - Create Articles
  - Delete Articles


## Installation

Install Golang and run this command in terminal 

```bash
  go run main.go
```

### Alert

Change the secret key than generate the user token

. \
├── controllers \
│&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;└── authControllers.go
 
```javascript
const SecretKey = "newSecretKey"
```


## Tech Stack

**Server:** Golang, SQLite3, GORM, Fiber

## License

MIT License

Copyright (c) 2023 Marcos Vinícius Ribeiro Alencar

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
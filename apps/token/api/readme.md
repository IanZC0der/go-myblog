## RESTful API for Login and Logout

- http method: POST/GET/PUT/DELETE/PATCH...
- http path: /users/list, /users/get

RESTful: representational state transfer

params:

- header
- URL
  - path param: /users/02
  - url param: /users/01?key1=...
- HTTP body

for token apis:

- POST /tokens/: create tokens, params in http body
- DELETE /tokens/: delete tokens, params in header

## API

exposed @ localhost:8000

### Run the API

run the command,

```bash
air
```

### Routes

All routes are prefixed with `/api/v1` for the version one API routes.

1. **POST** `/check?question_id=<int>&test_id=<int>`
   It takes parameters like question_id and test_id.
   The code block is accepted as a JSON object. Code block is found inside the `content` field.
   eg:

   ```json
   {
     "content": "a, b = 23, 34\nprint(a + b)"
   }
   ```

2. **POST** `/run`
   It takes the python code as `content` in the body of the request, and throws out metrics and run of the code.

   sample url,

   ```json
   {
     "content": "def add(a, b):\n    return a + b\nprint(add(2, 5))"
   }
   ```

   eg:

   ```json
   {
     "error": "",
     "ok": true,
     "output": "7\r\n",
     "time_taken": 8993200,
     "time_units": "ns"
   }
   ```

3. **GET** `/generate/:<id>`
   It takes the id of the question and generates the template for the code.
   It returns the template as a plain-text.

   sample url,

   ```
   http://localhost:8000/api/v1/generate/d5da605a-0a79-474b-bd29-682e41f6f2c6
   ```

   eg:

   ```py
   from typing import List

   type IMG = List[List[int]]
   type D = List[int]
   type Z = int

   def change_contrast(image: IMG, contrast: Z) -> IMG:
     pass
   ```

4. **POST** `/auth/login`
   It takes email and password of the user and returns generated token upon success.

   sample body,

   ```json
   {
     "email": "anuragsrivastav0027@gmail.com",
     "password": "123456"
   }
   ```

   response,

   ```json
   {
      "error": "Incase of error",
      "message"?:"In case of success",
      "token"?:"In case of success"
   }
   ```

5. **POST** `/auth/register`
   It takes email and password of the user creates and returns new user upon success.

   sample body,

   ```json
   {
     "email": "anuragsrivastav0027@gmail.com",
     "password": "123456"
   }
   ```

   response,

   ```json
   {
      "error"?:"Incase of error",
      "id"?:"User Id",
      "email"?:"User's email"
      "fullName"?:"",
      "userName"?:""
   }
   ```

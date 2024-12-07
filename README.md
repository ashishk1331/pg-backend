## API

exposed @ localhost:8000

### Run the API

run the command,

```bash
go run main.go
```

### Routes

1. **POST** `/check?question_id=<int>&test_id=<int>`
   It takes parameters like question_id and test_id.
   The code block is accepted as a JSON object. Code block is found inside the `content` field.
   eg:
    ```json
    {
    	"content": "a, b = 23, 34\nprint(a + b)"
    }
    ```

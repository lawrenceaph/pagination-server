# API Documentation

This API provides paginated data using a Go serverless function deployed on Vercel. The API returns a list of objects with pagination support, allowing clients to specify the page number and the number of objects per page.

## Endpoint

`GET /api`

## Query Parameters

* `page` (optional): The page number to start fetching data from. Defaults to `1` if not provided or invalid.
* `perPage` (optional): The number of objects to return per page. Defaults to `10` if not provided or invalid. Maximum is `1000`.
* `longContent` (optional): If set to `true`, the content of each object will contain 50 words of "Lorem Ipsum" text.

## Response

The API returns a JSON object containing the requested objects and the next page number (if available).

### Success Response

```json
{
  "objects": [
    {
      "name": "random name 1",
      "content": "random content 1",
      "image": "https://placehold.co/600x400"
    },
    {
      "name": "random name 2",
      "content": "random content 2",
      "image": "https://placehold.co/600x400"
    }
    ...
  ],
  "nextPage": 2
}
```

### Error Response

```json
{
  "error": "error message"
}
```

## Example Requests

Fetch Default Page:

```bash
curl -X GET "https://your-vercel-project/api"
```

Fetch Specific Page:

```bash
curl -X GET "https://your-vercel-project/api?page=2"
```

Fetch Specific Page with Custom Objects per Page:

```bash
curl -X GET "https://your-vercel-project/api?page=2&perPage=20"
```

Fetch with Long Content:

```bash
curl -X GET "https://your-vercel-project/api?page=1&perPage=10&longContent=true"
```
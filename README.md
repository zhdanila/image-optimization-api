# Image Optimization Service API

This API provides endpoints for uploading, listing, and retrieving images, as well as retrieving images with different
compression quality levels.

---

## Technologies Used

- **Go 1.23.4**: The programming language used to build the service, known for its performance and simplicity.
- **AWS SDK for Go**: Used for interacting with Amazon Web Services (AWS), particularly for storing images in S3.
- **Bimg**: A Go library for fast image processing, used for resizing and optimizing images.
- **Echo**: A lightweight and fast web framework for building the RESTful API.
- **EasyJSON**: A JSON serialization package that provides high-performance and zero-allocation encoding/decoding for efficient API communication.
- **Validator**: A library for validating incoming data to ensure that requests meet the required structure and integrity.
- **Viper**: A configuration management library used to handle environment variables and configuration files.
- **AMQP (RabbitMQ)**: A message broker used to manage image optimization tasks in a queue, ensuring asynchronous processing.

---

## How to Run the Service

### 1. Create `.env` File

Before running the service, make sure to create a `.env` file in the root directory of the project. You can use the
example file `env.default` to set the necessary environment variables. Copy the contents of `env.default` into a new
`.env` file, and update the variables as needed for your environment.

--- 

### 2. Start the Service Locally

To run the service locally, use the following command:

```bash
make up
```

### 3. Build Docker Image

To build the Docker image for the service, use the following command:

```bash
docker build -t image-optimization-api .
```

### 4. Start the Service Using Docker Compose

To build and start the service using Docker Compose, run:

```bash
docker-compose up --build
```

---

## Endpoints

### POST /api/image

Uploads an image.

#### Request Body:

- `images` (file, required): The image file to be uploaded.

#### Response:

- **200 OK**: A successful response with no content.

```json
{
  "status": "success",
  "message": "Image uploaded successfully"
}
```

---

### GET /api/image/list

Retrieves a list of images.

#### Response:

- **200 OK**: A list of images, each containing its key and URL.

```json
{
  "images": [
    {
      "key": "image_key_1",
      "url": "https://example.com/images/image_key_1"
    },
    {
      "key": "image_key_2",
      "url": "https://example.com/images/image_key_2"
    }
  ]
}
```

---

### GET /api/image/origin

Retrieves a list of original images (excluding those with compression quality suffixes).

#### Response:

- **200 OK**: A list of keys for original images.

```json
{
  "keys": [
    "image_key_1",
    "image_key_2"
  ]
}
```

---

### GET /api/image/{image_id}

Retrieves an image by ID. Optionally, you can specify a compression quality level.

#### Path Parameters:

- `image_id` (string, required): The ID of the image to retrieve.

#### Query Parameters:

- `quality` (integer, optional): Compression quality level. Possible values: `100`, `75`, `50`, `25`.

#### Response:

- **200 OK**: The image with the requested ID and compression quality.

```json
{
  "image": {
    "key": "image_key_1",
    "url": "https://example.com/images/image_key_1?quality=75"
  }
}
```


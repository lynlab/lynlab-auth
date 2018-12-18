# API Specifications

## List

  - GET /v1/me
  - POST /v1/recovery/reset
  - POST /v1/recovery/send_email
  - POST /v1/register
  - POST /v1/token/generate
  - POST /v1/token/refresh

## GET /v1/me

Authorization required.

### Response Body

```json
{
  "uuid": "string",
  "username": "string",
  "email": "string"
}
```

## POST /v1/recovery/reset

Reset password.

### Request Body

```json
{
  "recovery_token": "string",
  "password": "string"
}
```

## POST /v1/recovery/send_email

Send a password recovery email.

### Request Body

```json
{
  "email": "string"
}
```

## POST /v1/register

### Request Body

Register a new account.

```json
{
  "username": "string",
  "email": "string",
  "password": "string",
}
```

## POST /v1/token/generate

Sign in and generate access token.

### Request Body

```json
{
  "email": "string",
  "password": "string"
}
```

### Response Body

```json
{
  "access_token": "string",
  "refresh_token": "string"
}
```

## POST /v1/token/refresh

Generate new access token using refresh token.

### Request Body

```json
{
  "refresh_token": "string"
}
```

### Response Body

```json
{
  "access_token": "string",
  "refresh_token": "string"
}
```

## Requirements

- Implement a Rest API with CRUD functionality. 
- Database: MySQL or PostgreSQL.
- Unit test as much as you can.
- Set up service with docker compse.

## Installation

- `Install docker`
- `git clone https://github.com/moneyforwardvietnam/rory-exercise`

## Usage
- `docker-compose up`

## Detail of endpoints

#### 1. Get transactions of an user with paginate

- URL path: `/api/v2/users/<user_id>/transactions`
- HTTP method: `GET`
- Request:
    - Parameters:
        |Name|Required|Data type|Description|
        | --- | --- | --- | --- |
        |`user_id`|Yes|Integer|User's ID|
        |`account_id`|No|Integer|Account's ID|
        |`page`|No|Integer|Page|
        |`size`|No|Integer|Size of a Page|
- Response:
    - Content type: `application/json` 
    - HTTP status: `200 OK`
    - Body:

        |Name|Data type|Description|
        | --- | --- | --- |
        | `data` |Array| Array of user's transactions |
        | `total` |Integer| Transaction's Total |
        | `page` |Integer| Page |

#### 2. Get transactions of an user

- URL path: `/api/users/<user_id>/transactions`
- HTTP method: `GET`
- Request:
    - Parameters:
        |Name|Required|Data type|Description|
        | --- | --- | --- | --- |
        |`user_id`|Yes|Integer|User's ID|
        |`account_id`|No|Integer|Account's ID|
    - Note: When `account_id` is not specified, return all transactions of the user.
    - Please have validations for required fields

- Response:
    - Content type: `application/json` 
    - HTTP status: `200 OK`
    - Body: Array of user's transactions, each of which has the following fields:

        |Name|Data type|Description|
        | --- | --- | --- |
        | `id` |Integer| Transaction's ID |
        | `account_id` |Integer| Account's id |
        | `amount` |Decimal| Amount of money |
        | `bank` |String| Bank's name |
        | `transaction_type` |String| Type of transaction |
        | `created_at` |String| Created date of transaction |

- Example:  GET `/api/users/1/transactions?account_id=1`
  - Response:
    ```json
    [{
      "id": 1,
      "account_id": 1,
      "amount": 100000.00,
      "bank": "VCB",
      "transaction_type": "deposit",
      "created_at": "2020-02-10 20:00:00 +0700"
    }, { ... }]
    ```

#### 3. Create a transaction for an user
- URL path: `/api/users/<user_id>/transactions`
- HTTP method: `POST`
- Request:
    - Parameters:

        |Name|Required|Data type|Description|
        | --- | --- | --- | --- |
        |`user_id`|Yes|Integer|User's ID|

    - Body:

        |Name|Required|Data type|Description|
        | --- | --- | --- | --- |
        |`account_id`|Yes|Integer|Account's ID|
        | `amount`|Yes|Decimal| Amount of money |
        | `transaction_type`|Yes |String| Type of transaction |
    - Please have validations for required fields

- Response:
    - Content type: `application/json` 
    - HTTP status: `201 Created`
    - Body: Details of the created transaction with the following fields:

        |Name|Data type|Description|
        | --- | --- | --- |
        | `id` |Integer| Transaction's ID |
        | `account_id` |Integer| Account's id |
        | `amount` |Decimal| Amount of transaction |
        | `bank` |String| Bank's name |
        | `transaction_type` |String| Type of transaction |
        | `created_at` |String| Created date of transaction |

- Example: POST `/api/users/1/transactions`
  - Request body:
    ```json
    {
      "account_id": 2,
      "amount": 100000.00,
      "transaction_type": "deposit"
    }
     ```  
  - Response
    ```json
    {
      "id": 10,
      "account_id": 2,
      "amount": 100000.00,
      "bank": "VCB",
      "transaction_type": "deposit",
      "created_at": "2020-02-10 20:10:00 +0700"
    }
    ```

# Go E-Commerce API & GraphQL Service

A Go-based backend providing both REST and GraphQL APIs for an e-commerce platform with authentication, products, categories, shopping cart, orders, asynchronous event publishing (AWS SQS), and email notifications.

---

## Quick Start with Docker

### 1. Build and Run Containers
```bash
# Build images with latest changes
make docker-build

# Start all services
make docker-up
```

### Services & Endpoints
| Service | Endpoint | Description |
|---|---|---|
| **GraphQL Playground** | `http://localhost:8081/playground` | Interactive GraphQL UI |
| **GraphQL API** | `http://localhost:8081/graphql` | GraphQL POST Endpoint |
| **Swagger UI** | `http://localhost:8081/swagger` | Interactive REST Docs |
| **RapiDoc** | `http://localhost:8081/docs` | REST API Documentation |
| **MailHog Web UI** | `http://localhost:8025` | Inspect outgoing emails |
| **LocalStack** | `http://localhost:4566` | Mocked AWS S3 & SQS |
| **PostgreSQL** | `localhost:5439` | Database (`postgres`/`postgres`) |

---

## GraphQL Testing Guide

Open GraphQL Playground at: **`http://localhost:8081/playground`**

---

### 1. Public Operations (No Authorization Header Required)

#### 1.1 Register User
```graphql
mutation Register {
  register(input: {
    email: "john.doe@example.com"
    password: "password123"
    firstName: "John"
    lastName: "Doe"
    phone: "9876543210"
  }) {
    accessToken
    refreshToken
    user {
      id
      email
      firstName
      lastName
      phone
      role
      isActive
    }
  }
}
```

#### 1.2 Login (Copy the `accessToken`)
```graphql
mutation Login {
  login(input: {
    email: "john.doe@example.com"
    password: "password123"
  }) {
    accessToken
    refreshToken
    user {
      id
      email
      firstName
      role
    }
  }
}
```

#### 1.3 Refresh Token
```graphql
mutation RefreshToken {
  refreshToken(input: {
    refreshToken: "PASTE_REFRESH_TOKEN_HERE"
  }) {
    accessToken
    refreshToken
  }
}
```

#### 1.4 List Categories (Public)
```graphql
query GetCategories {
  categories {
    id
    name
    description
    isActive
  }
}
```

#### 1.5 List Products (Public with Pagination)
```graphql
query GetProducts {
  products(page: 1, limit: 10) {
    id
    name
    sku
    price
    stock
    description
    category {
      id
      name
    }
    images {
      id
      imageUrl
    }
  }
}
```

#### 1.6 Get Product by ID (Public)
```graphql
query GetProduct {
  product(id: "1") {
    id
    name
    price
    stock
    description
    category {
      name
    }
  }
}
```

---

### 2. Setting Up Authentication for Protected Operations

In GraphQL Playground, click on the **`HTTP HEADERS`** tab at the bottom-left and add your JWT token:

```json
{
  "Authorization": "Bearer YOUR_ACCESS_TOKEN_HERE"
}
```

---

### 3. Protected Operations (Requires `Authorization` Header)

#### 3.1 Get Current User (`me`)
```graphql
query Me {
  me {
    id
    email
    firstName
    lastName
    phone
    role
    isActive
  }
}
```

#### 3.2 Update Profile
```graphql
mutation UpdateProfile {
  updateProfile(input: {
    firstName: "Johnny"
    lastName: "Doe"
    phone: "9123456780"
  }) {
    id
    firstName
    lastName
    phone
  }
}
```

#### 3.3 Create Category (Admin / Protected)
```graphql
mutation CreateCategory {
  createCategory(input: {
    name: "Electronics"
    description: "Gadgets, smartphones, and computers"
    isActive: true
  }) {
    id
    name
    description
    isActive
  }
}
```

#### 3.4 Create Product (Protected)
```graphql
mutation CreateProduct {
  createProduct(input: {
    categoryId: 1
    name: "Wireless Headphones"
    description: "Noise cancelling bluetooth headphones"
    price: 99.99
    stock: 50
    sku: "HEADPHONE-WL-01"
  }) {
    id
    name
    sku
    price
    stock
    category {
      id
      name
    }
  }
}
```

#### 3.5 Add to Cart
```graphql
mutation AddToCart {
  addToCart(input: {
    productId: 1
    quantity: 2
  }) {
    id
    userId
    total
    cartItem {
      id
      quantity
      subTotal
      product {
        id
        name
        price
      }
    }
  }
}
```

#### 3.6 Get Cart
```graphql
query GetCart {
  cart {
    id
    userId
    total
    cartItem {
      id
      quantity
      subTotal
      product {
        id
        name
        price
        stock
      }
    }
  }
}
```

#### 3.7 Update Cart Item Quantity
```graphql
mutation UpdateCart {
  updateCart(input: {
    productId: 1
    quantity: 3
  }) {
    id
    total
    cartItem {
      quantity
      subTotal
    }
  }
}
```

#### 3.8 Remove Item from Cart
```graphql
mutation RemoveFromCart {
  removeFromCart(input: {
    productId: 1
  }) {
    id
    total
    cartItem {
      id
    }
  }
}
```

#### 3.9 Create Order (Converts Cart into Order + Publishes SQS Event)
```graphql
mutation CreateOrder {
  createOrder {
    id
    userId
    status
    totalAmount
    orderItems {
      id
      quantity
      price
      product {
        id
        name
      }
    }
  }
}
```

#### 3.10 List Orders
```graphql
query GetOrders {
  orders {
    id
    status
    totalAmount
    orderItems {
      id
      quantity
      price
      product {
        name
      }
    }
  }
}
```

#### 3.11 Logout
```graphql
mutation Logout {
  logout(refreshToken: "PASTE_REFRESH_TOKEN_HERE")
}
```

---

## Code Generation

If you modify any schema files in `graph/schema/*.graphql`, regenerate code with:
```bash
make graph-gen
```

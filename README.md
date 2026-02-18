# Mini-Event-Management-System
Mini Event Management System in Go

# Day 1 Progress - Formatted Folder Structure

Today we created folders of the system Such as Frontend,Backend,Services and Files
 
---

## Day 2 Progress — User Registration Feature

Today we implemented the first real functionality of the system: **User Registration**.

### Implemented Endpoint

**POST /users/register**

### Request Body

```json
{
  "name": "string",
  "email": "string",
  "password": "string"
}
```

### Functionality

* Accepts user details in JSON format
* Stores user information in MongoDB Atlas cloud database
* Returns confirmation response after successful insertion

### Result

The system now supports real user account creation.
Users are successfully stored inside the `eventdb.users` collection in MongoDB Atlas.

### Tools Used

* Golang (Gin Framework)
* MongoDB Atlas (Cloud Database)
* Postman Desktop API Testing
* GitHub Version Control

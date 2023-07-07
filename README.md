# Health Monitoring Application

## Table of Contents

- [Introduction](#introduction)
- [Web Application](#web-application)
- [Mobile Application](#mobile-application)
  - [Features](#features)
  - [API Endpoints](#api-endpoints)
  - [Features](#features)
- [Technology Stack](#technology-stack)

## Introduction

This project is a Health Monitoring Application developed using the Quasar Framework. It consists of a web application and a mobile application that allow users to track and monitor their health data in real-time. The application utilizes WebSocket technology for real-time updates and PostgreSQL for data storage.

## Web Application

The web application is built using the Quasar Framework, a high-performance Vue.js-based framework for building responsive Single Page Applications (SPAs). The application provides a user-friendly interface for managing user accounts and accessing health data.

## Mobile Application
![ScreenShot1](https://github.com/AYGA2K/Health-Monitor/tree/main/Screenshots/Screenshot_2023-07-07-20-21-24_1920x1080.png)
The mobile application is also developed using the Quasar Framework. 

### Features

- User registration and login functionality.
- User profile management.
- Real-time updates of heartbeat, sleep, and step data using WebSocket connections.
- Visualization of health data through interactive charts and graphs.
- Integration with the backend API for data retrieval and storage.

### API Endpoints

- `POST /api/signup`: Endpoint for user registration.
- `POST /api/login`: Endpoint for user login.
- `GET /api/users`: Endpoint to retrieve a list of users.
- `GET /api/user/:id`: Endpoint to retrieve user information.
- `DELETE /api/user/:id`: Endpoint to delete a user.
- `GET /ws/heartbeat/:id`: WebSocket endpoint for real-time heartbeat data.
- `GET /ws/sleep/:id`: WebSocket endpoint for real-time sleep data.
- `GET /ws/step/:id`: WebSocket endpoint for real-time step data.
- `POST /api/heartbeat`: Endpoint for creating a new heartbeat entry.

### Features

- User authentication and authorization.
- Real-time updates of heartbeat, sleep, and step data using WebSocket connections.
- Interactive and intuitive user interface for viewing and analyzing health data.


## Technology Stack

The project utilizes the following technologies and tools:

- Quasar Framework: A Vue.js-based framework for building web and mobile applications.
- Vue.js: A progressive JavaScript framework for building user interfaces.
- WebSocket: A communication protocol providing full-duplex communication channels over a single TCP connection.
- PostgreSQL: A powerful open-source relational database management system for data storage.
- Fiber: A fast and lightweight web framework for building APIs in Go.
- JWT (JSON Web Tokens): A method for securely transmitting information between parties as a JSON object.
- Git: A distributed version control system for tracking changes in the project's source code.



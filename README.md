# Project Tracker

Project Tracker is a web application designed to facilitate communication between lecturers and students during their final project period at the university. Students can submit their progress in the form of a document file, which lecturers can review and provide feedback on. The application keeps track of the communication history between lecturers and students.

## Table of Contents

- [Features](#features)
- [Technologies](#technologies)
- [Setup](#setup)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Features

- **User Authentication**: Secure login and registration for students and lecturers.
- **File Upload**: Students can upload their project documents.
- **Feedback System**: Lecturers can review and provide feedback on student submissions.
- **Communication History**: Tracks the communication between students and lecturers.
- **Dashboard**: Provides an overview of project status and submissions.
- **Profile Management**: Users can update their profile information and upload profile pictures.

## Technologies

- **Frontend**: Next.js, React, Tailwind CSS
- **Backend**: Golang, Echo framework
- **Database**: PostgreSQL
- **Authentication**: JWT
- **State Management**: Zustand
- **Other Tools**: Docker, Firebase for storage, GitHub for version control

## Setup

### Prerequisites

- Node.js and npm/pnpm
- Golang
- Docker
- PostgreSQL
- Firebase account

### Installation

1. **Set up the backend:**

   - Navigate to the `backend` directory:

     ```sh
     cd backend
     ```

   - Create a `.env` file and add your environment variables:

     ```env
     DATABASE_URL=your_database_url
     JWT_SECRET=your_jwt_secret
     ```

   - Build and run the backend:

     ```sh
     go build
     ./backend
     ```

2. **Set up the frontend:**

   - Navigate to the `frontend` directory:

     ```sh
     cd ../frontend
     ```

   - Install dependencies:

     ```sh
     pnpm install
     ```

   - Create a `.env.local` file and add your environment variables:

     ```env
     NEXT_PUBLIC_API_URL=your_api_url
     NEXT_PUBLIC_FIREBASE_API_KEY=your_firebase_api_key
     NEXT_PUBLIC_FIREBASE_AUTH_DOMAIN=your_firebase_auth_domain
     NEXT_PUBLIC_FIREBASE_PROJECT_ID=your_firebase_project_id
     ```

   - Run the frontend:

     ```sh
     pnpm run dev
     ```

3. **Run with Docker:**

   - Ensure Docker is installed and running.
   - Use Docker Compose to build and run the application:

     ```sh
     docker-compose up --build
     ```

## Usage

- **Student Actions:**
  - Register and log in to the application.
  - Upload project documents and view submission status.
  - View feedback from lecturers.

- **Lecturer Actions:**
  - Log in to the application.
  - Review student submissions and provide feedback.
  - Track project progress and communication history.

## Contributing

We welcome contributions! Please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature/your-feature`).
6. Open a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
```

Feel free to further customize and expand upon this template to better fit the specifics of your project and any additional details you want to include.
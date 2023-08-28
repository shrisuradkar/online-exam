# Online Examination Project

This project aims to provide an online examination platform where users can take exams, submit their answers, and view their results. It is built using Go programming language, Gin web framework, and MongoDB for data storage.

## Features

- User authentication and authorization
- Exam creation and management
- Exam question creation and management
- Exam submission and result retrieval

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/shrisuradkar/online-exam-go-gin-monogdb.git
   ```

2. Change to the project directory:

   ```bash
   cd online-exam-go-gin-monogdb
   ```

3. Install the dependencies:

   ```bash
   go mod download
   ```

4. Set up the MongoDB database:

   - Install MongoDB on your system
   - Create a new MongoDB database
   - Update the database connection details in `.env` if required.

5. Run the application:

   ```bash
   go run main.go
   ```

6. Access the application in your browser at `http://localhost:9000`

## Usage

1. Register a new user or log in with existing credentials.
2. Create exams from the admin dashboard.
3. Add questions to the exams.
4. Share the exam link with users.
5. Users can take the exams, submit their answers, and view their results.

## Contributing

Contributions are welcome! Please follow these guidelines when contributing to the project:

1. Fork the repository.
2. Create a new branch.
3. Make your changes and test them thoroughly.
4. Submit a pull request explaining the changes you've made.

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).

## Acknowledgements

- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework for Go
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver) - Official MongoDB driver for Go

## Support

If you have any questions or issues, please [raise an issue](https://github.com/shrisuradkar/online-exam-go-gin-monogdb/issues) on the project repository.

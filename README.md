# Order App

## Backend Installation

### Prerequisites

- **Docker:** Ensure that Docker is installed on your machine. Follow the official Docker installation guide: [Get Docker](https://docs.docker.com/get-docker/).

### Installation Steps

1. Open a terminal and navigate to the root folder of the backend project.
2. Run the following command to start the Docker containers and build the backend project:

    ```bash
    docker-compose up --build
    ```

   This command uses the `docker-compose.yml` configuration file in your backend project to set up the required services.

## Frontend Installation

### Prerequisites

- **Node Version Manager (NVM):** Install NVM to manage Node.js versions. Follow the instructions on the official NVM GitHub repository: [NVM GitHub](https://github.com/nvm-sh/nvm).
- **Node.js version 14.20:** Once NVM is installed, run the following commands in the terminal:

    ```bash
    nvm install 14.20
    nvm use 14.20
    ```

- **Yarn:** Install Yarn for package management. Follow the instructions on the official Yarn installation guide: [Yarn Installation](https://yarnpkg.com/getting-started/install).

### Installation Steps

1. Open a terminal and navigate to the root folder of the frontend project.
2. Run the following commands to install dependencies and serve the frontend application:

    ```bash
    yarn install
    yarn serve
    ```

   The first command installs the project dependencies specified in `package.json`, and the second command starts a development server to serve the frontend application.

Now, your backend and frontend projects should be up and running. If there are any specific configuration steps or environment variables needed for your projects, make sure to include those in your instructions.

Feel free to update this README file based on your project's specific details and requirements. Happy coding!

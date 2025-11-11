# Tutorial on Go

The code in this repository is part of the Programming Languages and Paradigms (PLP) Seminar at the University of Zurich.

In the Tutorial session, we will implement a simple web crawler that fetches the titles of a list of websites both synchronously and concurrently using goroutines.

## Preparation
### Install Go
1. Download and install Go from the official website: https://golang.org/dl/
2. Follow the installation instructions for your operating system.
3. Verify the installation by opening a terminal and running:
   ```bash
   go version
   ```  
    You should see the installed Go version.

### Set Up the Project
1. Clone this repository to your local machine:
    ```bash
    git clone https://github.com/vichcruz/programming-languages-tutorial.git
    ```
2. Navigate to the project directory:
    ```bash
    cd programming-languages-tutorial
    ```

## Running the Code
1. Open a terminal and navigate to the project directory.
2. The **exercise code** is located in the `tutorial-starter` folder. You can run it using:
   ```bash
   cd tutorial-starter
   go run .
   ```
3. The **solution code** is located in the `solution` folder. Please don't look at it yet. If you want to see the expected output of the crawler, you can run it using:
   ```bash
   cd solution
   go run .
   ``` 

## Project Structure
- `main.go`: The main entry point of the application.
- `Fetcher.go`: Contains the function to fetch the title of a webpage.
- `ConcurrentCrawler.go`: Implements the concurrent web crawler using goroutines.
- `go.mod`: Go module file specifying dependencies.



# Finance Tracker

## Summary
This project is a command-line finance tracker built with Go. It provides an interactive interface for users to manage their expenses, allowing them to add new expenses, list existing ones, view a comprehensive manual, and track their financial activities efficiently.

## Key Features
*   **Expense Management**: Add new expenses with details like name, category, and amount.
*   **Expense Listing**: View a complete list of all recorded expenses.
*   **Interactive CLI**: Engage with the application through an intuitive command-line interface.
*   **User Manual**: Access an in-app manual detailing all available commands and their usage.
*   **Input Validation**: Basic checks for valid input for expense details.
*   **Cross-platform Clarity**: Clears the terminal screen for a clean user experience across different operating systems.

## Tech stack
*   Go

## Installation
To get this project up and running on your local machine, follow these steps:

1.  **Clone the repository**:
    ```bash
    git clone https://github.com/Astra-max/finance.git
    cd finance
    ```

2.  **Run the application**:
    ```bash
    go run main.go
    ```
    This will start the interactive command-line interface.

## Folder Structure

```
finance/                        # Root directory for the finance tracker application
├── go.mod                      # Go module definitions and dependencies
├── lib/                        # Core logic and utility functions
│   ├── finance.go              # Defines the Finance struct and its core methods (add, list, etc.)
│   ├── userInput.go            # Handles user input operations
│   └── util.go                 # Contains utility functions like clearing the screen
└── main.go                     # Main application entry point, initializes and runs the controller
```

## API Documentation
This section outlines the key functions and their descriptions, serving as the functional "API" for interacting with the finance tracker's logic.

*   `NewFinance()`
    *   **DocString**: Initializes a new Finance struct.

*   `(f *Finance) List()`
    *   **DocString**: Lists all recorded expenses. If no expenses are available, it prints a message indicating that.

*   `(f *Finance) AddExpense(name, category string, amount int)`
    *   **DocString**: Adds a new expense to the Finance tracker.
    *   **Args**:
        *   `name` (string): The name of the expense.
        *   `category` (string): The category of the expense.
        *   `amount` (int): The amount of the expense. Must be a non-zero value.

*   `(f *Finance) IsEmpty() bool`
    *   **DocString**: Checks if the list of expenses is empty.
    *   **Returns**: `bool`: `True` if the list of expenses is empty, `false` otherwise.

*   `(f *Finance) Manual()`
    *   **DocString**: Displays the manual for the Finance tracker, outlining available commands and their usage.

*   `(f *Finance) Controller()`
    *   **DocString**: Provides an interactive command-line interface for the Finance tracker. It continuously prompts the user for commands and executes the corresponding actions such as adding expenses, listing expenses, viewing total expenses, and exiting the application.

*   `Message(str string)`
    *   **DocString**: Displays a message to the console.
    *   **Args**:
        *   `str`: The string message to be displayed.

*   `UserInput(message string)`
    *   **DocString**: Reads a line of input from the user via the console.
    *   **Args**:
        *   `message`: The prompt message to display to the user before awaiting input.
    *   **Returns**: The trimmed string input provided by the user.

*   `ProcessInputs(name string)`
    *   **DocString**: Checks if the input string `name` is empty. It returns `true` if the string is empty or contains only an empty string, otherwise it returns `false`.

*   `Clear()`
    *   **DocString**: Clears the terminal screen. It detects the operating system and executes the appropriate command (`cls` for Windows, `clear` for others) to clear the console output.

*   `main()`
    *   **DocString**: This is the main function of the application. It initializes a new Finance object and then calls its Controller method to start the application's main logic.

## Contributing
Contributions are welcome! If you have suggestions for improvements, new features, or bug fixes, please feel free to:
1.  Fork the repository.
2.  Create a new branch (`git checkout -b feature/AmazingFeature`).
3.  Make your changes.
4.  Commit your changes (`git commit -m 'Add some AmazingFeature'`).
5.  Push to the branch (`git push origin feature/AmazingFeature`).
6.  Open a Pull Request.

## License
This project is open-source and available under the [MIT License](LICENSE).
```
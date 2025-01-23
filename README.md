# Expense Tracker

[https://roadmap.sh/projects/expense-tracker](https://roadmap.sh/projects/expense-tracker)

Usage Examples:

- Help:
  ```
  expense-tracker --help
  expense-tracler help
  ```
- Version

  ```
  expense-tracker --version
  expense-tracker -V
  ```

- Add a category:

  ```
  $ expense-tracker add-category <category>
  # Category set: <category_name>
  ```

- Set a budget for a month:

  ```
  $ expense-tracker set-budget --month <month e.g 1-12> --amount <amount>
  # Budget set for month <month> of <amount>
  ```

- Add an expense:

  ```
  $ expense-tracker add --description "Lunch" --amount 20
  # Expense added successfully (ID: 1)
  ```

- Add an expense for a specific month:

  ```
  $ expense-tracker add --description "Dinner" --amount 10 --month 3
  # Expense added successfully (ID: 2)
  ```

- List all expenses:

  ```
  $ expense-tracker list
  # ID Date       Description Amount
  # 1  2024-08-06 Lunch       $20
  # 2  2024-08-06 Dinner      $10
  ```

- View a summary of all expenses:

  ```
  $ expense-tracker summary
  # Total expenses: $30
  ```

- View a summary by category:

  ```
  $ expense-tracker summary --category <category_name>
  # Total expenses for category_name: $20
  ```

- View a summary for a specific month:

  ```
  $ expense-tracker summary --month <monthValue 1-12>
  # Total expenses for August: $20
  ```

- Delete an expense by ID:

  ```
  $ expense-tracker delete --id <expenseId>
  # Expense deleted successfully
  ```

- Export expenses to a CSV file:
  ```
  $ expense-tracker export <filename.csv>
  # Expenses exported to specified file name
  ```

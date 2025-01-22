# Expense Tracker

Command: `expense-tracker --help`

[Project Roadmap](https://roadmap.sh/projects/expense-tracker)

Features:

- Add an expense with a description and amount. (Completed)
- Update an expense. (Completed)
- Delete an expense. (Completed)
- View all expenses.
- View a summary of all expenses.
- View a summary of expenses for a specific month (of the current year).

Upcoming Features:

- Add expense categories and filter expenses by category.
- Set a budget for each month and display a warning when the budget is exceeded.
- Export expenses to a CSV file.

Usage Examples:

- Add a category:

  ```
  $ expense-tracker add-category <category>
  # Category set: <category_name>
  ```

- Set a budget for a month:

  ```
  $ expense-tracker set-budget --month <amount>
  # Budget set for month <month> of <amount>
  ```

- Add an expense (with budget warning):

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
  $ expense-tracker summary --month 8
  # Total expenses for August: $20
  ```

- Delete an expense by ID:

  ```
  $ expense-tracker delete --id 2
  # Expense deleted successfully
  ```

- Export expenses to a CSV file:
  ```
  $ expense-tracker export
  # Expenses exported to expenses.csv
  ```

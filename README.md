# Project url: https://roadmap.sh/projects/task-tracker

# Adding a new task
go run *.go add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
go run *.go update 1 "Buy groceries and cook dinner" 
# Output: Task updated successfully
go run *.go delete 1
# Output: Task deleted

# Marking a task as in progress or done
go run *.go mark-in-progress 1
go run *.go mark-done 1

# Listing all tasks
go run *.go list

# Listing tasks by status
go run *.go list done
go run *.go list todo
go run *.go list in-progress


interface Task {
    id: string;
    title: string;
    status: string;
    created_at: string;
}

const API_URL = 'http://localhost:8081/tasks';

document.addEventListener('DOMContentLoaded', () => {
    fetchTasks();

    const addTaskBtn = document.getElementById('add-task-btn') as HTMLButtonElement;
    const newTaskInput = document.getElementById('new-task-title') as HTMLInputElement;

    addTaskBtn.addEventListener('click', () => {
        const title = newTaskInput.value.trim();
        if (title) {
            createTask(title);
            newTaskInput.value = '';
        }
    });

    newTaskInput.addEventListener('keypress', (e) => {
        if (e.key === 'Enter') {
            const title = newTaskInput.value.trim();
            if (title) {
                createTask(title);
                newTaskInput.value = '';
            }
        }
    });
});

async function fetchTasks() {
    try {
        const response = await fetch(API_URL);
        if (!response.ok) {
            throw new Error('Failed to fetch tasks');
        }
        const tasks: Task[] = await response.json();
        renderTasks(tasks);
    } catch (error) {
        console.error('Error fetching tasks:', error);
    }
}

async function createTask(title: string) {
    try {
        const response = await fetch(API_URL, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ title }),
        });

        if (!response.ok) {
            throw new Error('Failed to create task');
        }

        fetchTasks(); // Refresh list
    } catch (error) {
        console.error('Error creating task:', error);
    }
}

async function updateTaskStatus(id: string, status: string) {
    try {
        const response = await fetch(`${API_URL}/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ status }),
        });

        if (!response.ok) {
            throw new Error('Failed to update task status');
        }

        fetchTasks(); // Refresh list
    } catch (error) {
        console.error('Error updating task status:', error);
    }
}

async function updateTaskTitle(id: string, title: string) {
    try {
        const response = await fetch(`${API_URL}/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ title }),
        });

        if (!response.ok) {
            throw new Error('Failed to update task title');
        }

        fetchTasks(); // Refresh list
    } catch (error) {
        console.error('Error updating task title:', error);
    }
}

async function deleteTask(id: string) {
    try {
        const response = await fetch(`${API_URL}/${id}`, {
            method: 'DELETE',
        });

        if (!response.ok) {
            throw new Error('Failed to delete task');
        }

        fetchTasks(); // Refresh list
    } catch (error) {
        console.error('Error deleting task:', error);
    }
}

function renderTasks(tasks: Task[]) {
    const taskList = document.getElementById('task-list') as HTMLUListElement;
    taskList.innerHTML = '';

    tasks.forEach(task => {
        const li = document.createElement('li');
        
        const taskContent = document.createElement('div');
        taskContent.className = 'task-content';

        const checkbox = document.createElement('input');
        checkbox.type = 'checkbox';
        checkbox.checked = task.status === 'done';
        checkbox.addEventListener('change', () => {
            const newStatus = checkbox.checked ? 'done' : 'todo';
            updateTaskStatus(task.id, newStatus);
        });

        const titleSpan = document.createElement('span');
        titleSpan.className = 'task-title';
        if (task.status === 'done') {
            titleSpan.classList.add('completed');
        }
        titleSpan.textContent = task.title;
        
        // Double click to edit
        titleSpan.addEventListener('dblclick', () => {
            const input = document.createElement('input');
            input.type = 'text';
            input.value = task.title;
            input.addEventListener('blur', () => {
                if (input.value.trim() !== task.title) {
                    updateTaskTitle(task.id, input.value.trim());
                } else {
                    renderTasks(tasks); // Re-render to restore span
                }
            });
            input.addEventListener('keypress', (e) => {
                if (e.key === 'Enter') {
                    if (input.value.trim() !== task.title) {
                        updateTaskTitle(task.id, input.value.trim());
                    } else {
                        renderTasks(tasks); 
                    }
                }
            });
            taskContent.replaceChild(input, titleSpan);
            input.focus();
        });

        const deleteBtn = document.createElement('button');
        deleteBtn.className = 'delete-btn';
        deleteBtn.innerHTML = '&times;'; // X symbol
        deleteBtn.title = 'Delete Task';
        deleteBtn.addEventListener('click', () => {
            deleteTask(task.id);
        });

        taskContent.appendChild(checkbox);
        taskContent.appendChild(titleSpan);

        li.appendChild(taskContent);
        li.appendChild(deleteBtn);

        taskList.appendChild(li);
    });
}

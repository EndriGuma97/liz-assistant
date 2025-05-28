package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Task represents a single task
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Type        string    `json:"type"`
	Owner       string    `json:"owner"`
	Priority    string    `json:"priority"`
	Completed   bool      `json:"completed"`
	Notes       string    `json:"notes"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

// TaskManager holds all tasks
type TaskManager struct {
	Tasks  []Task `json:"tasks"`
	NextID int    `json:"next_id"`
}

var taskManager = TaskManager{
	Tasks:  []Task{},
	NextID: 1,
}

// Initialize with predefined tasks
func initializeTasks() {
	tasks := []Task{
		{
			ID:        1,
			Title:     "Schedule follow-up meeting - Tomorrow at 7:30 AM",
			Type:      "Immediate Tasks (24-48 hours)",
			Owner:     "Tariro & Endri",
			Priority:  "High",
			Completed: false,
			Notes:     "Endri to create recurring weekly meeting",
			CreatedAt: time.Now(),
		},
		{
			ID:        2,
			Title:     "Audit current Calendly setup",
			Type:      "Immediate Tasks (24-48 hours)",
			Owner:     "Tariro & Endri",
			Priority:  "High",
			Completed: false,
			Notes:     "Review all existing schedules, clones, and configurations",
			CreatedAt: time.Now(),
		},
		{
			ID:        3,
			Title:     "Complete resident list cleanup",
			Type:      "Immediate Tasks (24-48 hours)",
			Owner:     "Michael & Tariro",
			Priority:  "High",
			Completed: false,
			Notes:     "Verify all current vs. former residents in Zoho",
			CreatedAt: time.Now(),
		},
		{
			ID:        4,
			Title:     "Implement SMS capability",
			Type:      "Immediate Tasks (24-48 hours)",
			Owner:     "Michael & Technical Team",
			Priority:  "Medium",
			Completed: false,
			Notes:     "Deploy text messaging API integration in Zoho within one week",
			CreatedAt: time.Now(),
		},
		{
			ID:        5,
			Title:     "Establish Calendly change procedures",
			Type:      "Process Improvement Tasks (1-2 weeks)",
			Owner:     "Liz, Tariro, Endri",
			Priority:  "Medium",
			Completed: false,
			Notes:     "Define who updates what and when",
			CreatedAt: time.Now(),
		},
		{
			ID:        6,
			Title:     "Fix recurring meeting setup",
			Type:      "Process Improvement Tasks (1-2 weeks)",
			Owner:     "Endri",
			Priority:  "Medium",
			Completed: false,
			Notes:     "Ensure Wednesday 7:30 AM meetings auto-schedule properly",
			CreatedAt: time.Now(),
		},
		{
			ID:        7,
			Title:     "Track residency completion dates",
			Type:      "Ongoing Management Tasks",
			Owner:     "Liz",
			Priority:  "Low",
			Completed: false,
			Notes:     "Monitor approaching end dates for proper offboarding",
			CreatedAt: time.Now(),
		},
		{
			ID:        8,
			Title:     "Establish clear handoff procedures",
			Type:      "Communication & Coordination",
			Owner:     "All Team",
			Priority:  "Medium",
			Completed: false,
			Notes:     "Between Liz and Tariro for scheduling",
			CreatedAt: time.Now(),
		},
	}

	taskManager.Tasks = tasks
	taskManager.NextID = 9
}

func main() {
	// Initialize tasks
	initializeTasks()

	// Serve static files (CSS, JS, images)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/tasks", tasksHandler)
	http.HandleFunc("/api/tasks/", taskHandler)

	fmt.Println("🚀 AMSKU Task Management Server starting on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	
	// Write the HTML directly to avoid template parsing issues
	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>AMSKU Task Management</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            padding: 20px;
        }

        .container {
            max-width: 1200px;
            margin: 0 auto;
            background: white;
            border-radius: 20px;
            box-shadow: 0 20px 40px rgba(0,0,0,0.1);
            overflow: hidden;
        }

        .header {
            background: linear-gradient(135deg, #4f46e5 0%, #7c3aed 100%);
            color: white;
            padding: 30px;
            text-align: center;
        }

        .header h1 {
            font-size: 2.5rem;
            margin-bottom: 10px;
            font-weight: 700;
        }

        .header p {
            font-size: 1.1rem;
            opacity: 0.9;
        }

        .controls {
            padding: 20px 30px;
            background: #f8fafc;
            border-bottom: 1px solid #e2e8f0;
            display: flex;
            gap: 15px;
            flex-wrap: wrap;
            align-items: center;
        }

        .filter-group {
            display: flex;
            align-items: center;
            gap: 10px;
        }

        .filter-group label {
            font-weight: 600;
            color: #374151;
        }

        select, input[type="text"] {
            padding: 8px 12px;
            border: 2px solid #e5e7eb;
            border-radius: 8px;
            font-size: 14px;
            transition: all 0.2s;
        }

        select:focus, input[type="text"]:focus {
            outline: none;
            border-color: #4f46e5;
            box-shadow: 0 0 0 3px rgba(79, 70, 229, 0.1);
        }

        .btn {
            padding: 10px 20px;
            border: none;
            border-radius: 8px;
            font-weight: 600;
            cursor: pointer;
            transition: all 0.2s;
            text-decoration: none;
            display: inline-block;
            text-align: center;
        }

        .btn-primary {
            background: #4f46e5;
            color: white;
        }

        .btn-primary:hover {
            background: #4338ca;
            transform: translateY(-1px);
        }

        .btn-success {
            background: #10b981;
            color: white;
            padding: 6px 12px;
            font-size: 12px;
        }

        .btn-success:hover {
            background: #059669;
        }

        .btn-secondary {
            background: #6b7280;
            color: white;
            padding: 6px 12px;
            font-size: 12px;
        }

        .btn-secondary:hover {
            background: #4b5563;
        }

        .stats {
            padding: 20px 30px;
            background: #f1f5f9;
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
            gap: 20px;
        }

        .stat-card {
            background: white;
            padding: 20px;
            border-radius: 12px;
            text-align: center;
            box-shadow: 0 4px 6px rgba(0,0,0,0.05);
        }

        .stat-number {
            font-size: 2rem;
            font-weight: 700;
            color: #1e293b;
        }

        .stat-label {
            color: #64748b;
            font-size: 0.9rem;
            margin-top: 5px;
        }

        .task-types {
            padding: 30px;
        }

        .task-type {
            margin-bottom: 30px;
        }

        .task-type-header {
            background: #f8fafc;
            padding: 15px 20px;
            border-radius: 10px;
            border-left: 4px solid #4f46e5;
            margin-bottom: 15px;
        }

        .task-type-title {
            font-size: 1.3rem;
            font-weight: 700;
            color: #1e293b;
        }

        .task-grid {
            display: grid;
            grid-template-columns: repeat(auto-fill, minmax(350px, 1fr));
            gap: 20px;
        }

        .task-card {
            background: white;
            border: 2px solid #e5e7eb;
            border-radius: 12px;
            padding: 20px;
            transition: all 0.3s;
            position: relative;
        }

        .task-card:hover {
            box-shadow: 0 8px 25px rgba(0,0,0,0.1);
            transform: translateY(-2px);
        }

        .task-card.completed {
            border-color: #10b981;
            background: #f0fdf4;
        }

        .task-title {
            font-size: 1.1rem;
            font-weight: 600;
            color: #1e293b;
            margin-bottom: 10px;
            line-height: 1.4;
        }

        .task-meta {
            display: flex;
            gap: 10px;
            margin-bottom: 15px;
            flex-wrap: wrap;
        }

        .badge {
            padding: 4px 8px;
            border-radius: 6px;
            font-size: 12px;
            font-weight: 600;
        }

        .badge-high { background: #fef2f2; color: #dc2626; }
        .badge-medium { background: #fffbeb; color: #d97706; }
        .badge-low { background: #f0f9ff; color: #0284c7; }

        .task-owner {
            color: #6b7280;
            font-size: 0.9rem;
            margin-bottom: 10px;
        }

        .task-notes {
            background: #f8fafc;
            padding: 12px;
            border-radius: 8px;
            font-size: 0.9rem;
            color: #4b5563;
            line-height: 1.4;
            margin-bottom: 15px;
            border-left: 3px solid #e5e7eb;
        }

        .task-actions {
            display: flex;
            gap: 10px;
            justify-content: space-between;
            align-items: center;
        }

        .task-date {
            font-size: 0.8rem;
            color: #9ca3af;
        }

        .completion-badge {
            position: absolute;
            top: 15px;
            right: 15px;
            background: #10b981;
            color: white;
            padding: 4px 8px;
            border-radius: 12px;
            font-size: 11px;
            font-weight: 600;
        }

        .modal {
            display: none;
            position: fixed;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
            background: rgba(0,0,0,0.5);
            z-index: 1000;
        }

        .modal-content {
            background: white;
            margin: 5% auto;
            padding: 30px;
            border-radius: 20px;
            width: 90%;
            max-width: 500px;
            position: relative;
        }

        .modal-header {
            margin-bottom: 20px;
        }

        .modal-title {
            font-size: 1.5rem;
            font-weight: 700;
            color: #1e293b;
        }

        .form-group {
            margin-bottom: 20px;
        }

        .form-group label {
            display: block;
            margin-bottom: 5px;
            font-weight: 600;
            color: #374151;
        }

        .form-group input,
        .form-group select,
        .form-group textarea {
            width: 100%;
            padding: 12px;
            border: 2px solid #e5e7eb;
            border-radius: 8px;
            font-size: 14px;
        }

        .form-group textarea {
            resize: vertical;
            min-height: 100px;
        }

        .close {
            position: absolute;
            top: 15px;
            right: 20px;
            font-size: 28px;
            font-weight: bold;
            cursor: pointer;
            color: #9ca3af;
        }

        .close:hover {
            color: #374151;
        }

        @media (max-width: 768px) {
            .container {
                margin: 10px;
                border-radius: 15px;
            }
            
            .header {
                padding: 20px;
            }
            
            .header h1 {
                font-size: 2rem;
            }
            
            .controls {
                flex-direction: column;
                align-items: stretch;
            }
            
            .task-grid {
                grid-template-columns: 1fr;
            }
            
            .stats {
                grid-template-columns: repeat(2, 1fr);
                gap: 15px;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>🎯 AMSKU Task Management</h1>
            <p>Track progress and coordinate team tasks efficiently</p>
        </div>

        <div class="controls">
            <div class="filter-group">
                <label for="typeFilter">Filter by Type:</label>
                <select id="typeFilter" onchange="filterTasks()">
                    <option value="">All Types</option>
                    <option value="Immediate Tasks (24-48 hours)">Immediate Tasks (24-48 hours)</option>
                    <option value="Process Improvement Tasks (1-2 weeks)">Process Improvement Tasks (1-2 weeks)</option>
                    <option value="Ongoing Management Tasks">Ongoing Management Tasks</option>
                    <option value="Communication & Coordination">Communication & Coordination</option>
                </select>
            </div>

            <div class="filter-group">
                <label for="statusFilter">Filter by Status:</label>
                <select id="statusFilter" onchange="filterTasks()">
                    <option value="">All Tasks</option>
                    <option value="pending">Pending</option>
                    <option value="completed">Completed</option>
                </select>
            </div>

            <div class="filter-group">
                <label for="ownerFilter">Filter by Owner:</label>
                <select id="ownerFilter" onchange="filterTasks()">
                    <option value="">All Owners</option>
                </select>
            </div>

            <button class="btn btn-primary" onclick="openAddTaskModal()">+ Add New Task</button>
        </div>

        <div class="stats">
            <div class="stat-card">
                <div class="stat-number" id="totalTasks">0</div>
                <div class="stat-label">Total Tasks</div>
            </div>
            <div class="stat-card">
                <div class="stat-number" id="completedTasks">0</div>
                <div class="stat-label">Completed</div>
            </div>
            <div class="stat-card">
                <div class="stat-number" id="pendingTasks">0</div>
                <div class="stat-label">Pending</div>
            </div>
            <div class="stat-card">
                <div class="stat-number" id="completionRate">0%</div>
                <div class="stat-label">Completion Rate</div>
            </div>
        </div>

        <div class="task-types" id="taskContainer">
            <!-- Tasks will be loaded here -->
        </div>
    </div>

    <!-- Add/Edit Task Modal -->
    <div id="taskModal" class="modal">
        <div class="modal-content">
            <span class="close" onclick="closeTaskModal()">&times;</span>
            <div class="modal-header">
                <h2 class="modal-title" id="modalTitle">Add New Task</h2>
            </div>
            <form id="taskForm">
                <input type="hidden" id="taskId" value="">
                <div class="form-group">
                    <label for="taskTitle">Task Title</label>
                    <input type="text" id="taskTitle" required>
                </div>
                <div class="form-group">
                    <label for="taskType">Task Type</label>
                    <select id="taskType" required>
                        <option value="">Select Type</option>
                        <option value="Immediate Tasks (24-48 hours)">Immediate Tasks (24-48 hours)</option>
                        <option value="Process Improvement Tasks (1-2 weeks)">Process Improvement Tasks (1-2 weeks)</option>
                        <option value="Ongoing Management Tasks">Ongoing Management Tasks</option>
                        <option value="Communication & Coordination">Communication & Coordination</option>
                    </select>
                </div>
                <div class="form-group">
                    <label for="taskOwner">Owner</label>
                    <input type="text" id="taskOwner" required>
                </div>
                <div class="form-group">
                    <label for="taskPriority">Priority</label>
                    <select id="taskPriority" required>
                        <option value="">Select Priority</option>
                        <option value="High">High</option>
                        <option value="Medium">Medium</option>
                        <option value="Low">Low</option>
                    </select>
                </div>
                <div class="form-group">
                    <label for="taskNotes">Notes</label>
                    <textarea id="taskNotes" placeholder="Add any notes or progress updates..."></textarea>
                </div>
                <div class="form-group">
                    <label>
                        <input type="checkbox" id="taskCompleted"> Mark as completed
                    </label>
                </div>
                <button type="submit" class="btn btn-primary">Save Task</button>
            </form>
        </div>
    </div>

    <script>
        let tasks = [];

        // Load tasks on page load
        document.addEventListener('DOMContentLoaded', function() {
            loadTasks();
        });

        async function loadTasks() {
            try {
                const response = await fetch('/api/tasks');
                tasks = await response.json();
                renderTasks();
                updateStats();
                populateOwnerFilter();
            } catch (error) {
                console.error('Error loading tasks:', error);
            }
        }

        function renderTasks() {
            const container = document.getElementById('taskContainer');
            const typeFilter = document.getElementById('typeFilter').value;
            const statusFilter = document.getElementById('statusFilter').value;
            const ownerFilter = document.getElementById('ownerFilter').value;

            // Filter tasks
            let filteredTasks = tasks.filter(task => {
                if (typeFilter && task.type !== typeFilter) return false;
                if (statusFilter === 'completed' && !task.completed) return false;
                if (statusFilter === 'pending' && task.completed) return false;
                if (ownerFilter && task.owner !== ownerFilter) return false;
                return true;
            });

            // Group tasks by type
            const tasksByType = {};
            filteredTasks.forEach(task => {
                if (!tasksByType[task.type]) {
                    tasksByType[task.type] = [];
                }
                tasksByType[task.type].push(task);
            });

            let html = '';
            Object.keys(tasksByType).forEach(type => {
                html += '<div class="task-type">';
                html += '<div class="task-type-header">';
                html += '<div class="task-type-title">' + type + '</div>';
                html += '</div>';
                html += '<div class="task-grid">';
                
                tasksByType[type].forEach(task => {
                    html += '<div class="task-card ' + (task.completed ? 'completed' : '') + '">';
                    if (task.completed) {
                        html += '<div class="completion-badge">✓ Completed</div>';
                    }
                    html += '<div class="task-title">' + task.title + '</div>';
                    html += '<div class="task-meta">';
                    html += '<span class="badge badge-' + task.priority.toLowerCase() + '">' + task.priority + '</span>';
                    html += '</div>';
                    html += '<div class="task-owner">👤 ' + task.owner + '</div>';
                    if (task.notes) {
                        html += '<div class="task-notes">' + task.notes + '</div>';
                    }
                    html += '<div class="task-actions">';
                    html += '<div class="task-date">Created: ' + new Date(task.created_at).toLocaleDateString() + '</div>';
                    html += '<div>';
                    if (!task.completed) {
                        html += '<button class="btn btn-success" onclick="toggleTaskCompletion(' + task.id + ')">Mark Complete</button>';
                    } else {
                        html += '<button class="btn btn-secondary" onclick="toggleTaskCompletion(' + task.id + ')">Mark Pending</button>';
                    }
                    html += '<button class="btn btn-secondary" onclick="editTask(' + task.id + ')">Edit</button>';
                    html += '</div>';
                    html += '</div>';
                    html += '</div>';
                });
                
                html += '</div>';
                html += '</div>';
            });

            container.innerHTML = html;
        }

        function updateStats() {
            const total = tasks.length;
            const completed = tasks.filter(t => t.completed).length;
            const pending = total - completed;
            const completionRate = total > 0 ? Math.round((completed / total) * 100) : 0;

            document.getElementById('totalTasks').textContent = total;
            document.getElementById('completedTasks').textContent = completed;
            document.getElementById('pendingTasks').textContent = pending;
            document.getElementById('completionRate').textContent = completionRate + '%';
        }

        function populateOwnerFilter() {
            const ownerFilter = document.getElementById('ownerFilter');
            const owners = [...new Set(tasks.map(t => t.owner))].sort();
            
            ownerFilter.innerHTML = '<option value="">All Owners</option>';
            owners.forEach(owner => {
                ownerFilter.innerHTML += '<option value="' + owner + '">' + owner + '</option>';
            });
        }

        function filterTasks() {
            renderTasks();
        }

        async function toggleTaskCompletion(taskId) {
            try {
                const response = await fetch('/api/tasks/' + taskId + '/toggle', {
                    method: 'POST',
                });
                if (response.ok) {
                    loadTasks();
                }
            } catch (error) {
                console.error('Error toggling task completion:', error);
            }
        }

        function openAddTaskModal() {
            document.getElementById('modalTitle').textContent = 'Add New Task';
            document.getElementById('taskForm').reset();
            document.getElementById('taskId').value = '';
            document.getElementById('taskModal').style.display = 'block';
        }

        function editTask(taskId) {
            const task = tasks.find(t => t.id === taskId);
            if (task) {
                document.getElementById('modalTitle').textContent = 'Edit Task';
                document.getElementById('taskId').value = task.id;
                document.getElementById('taskTitle').value = task.title;
                document.getElementById('taskType').value = task.type;
                document.getElementById('taskOwner').value = task.owner;
                document.getElementById('taskPriority').value = task.priority;
                document.getElementById('taskNotes').value = task.notes || '';
                document.getElementById('taskCompleted').checked = task.completed;
                document.getElementById('taskModal').style.display = 'block';
            }
        }

        function closeTaskModal() {
            document.getElementById('taskModal').style.display = 'none';
        }

        document.getElementById('taskForm').addEventListener('submit', async function(e) {
            e.preventDefault();
            
            const taskData = {
                title: document.getElementById('taskTitle').value,
                type: document.getElementById('taskType').value,
                owner: document.getElementById('taskOwner').value,
                priority: document.getElementById('taskPriority').value,
                notes: document.getElementById('taskNotes').value,
                completed: document.getElementById('taskCompleted').checked
            };

            const taskId = document.getElementById('taskId').value;
            const url = taskId ? '/api/tasks/' + taskId : '/api/tasks';
            const method = taskId ? 'PUT' : 'POST';

            try {
                const response = await fetch(url, {
                    method: method,
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(taskData)
                });

                if (response.ok) {
                    closeTaskModal();
                    loadTasks();
                }
            } catch (error) {
                console.error('Error saving task:', error);
            }
        });

        // Close modal when clicking outside
        window.onclick = function(event) {
            const modal = document.getElementById('taskModal');
            if (event.target === modal) {
                closeTaskModal();
            }
        }
    </script>
</body>
</html>`

	fmt.Fprint(w, html)
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(taskManager.Tasks)

	case "POST":
		var task Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		task.ID = taskManager.NextID
		task.CreatedAt = time.Now()
		taskManager.NextID++
		taskManager.Tasks = append(taskManager.Tasks, task)

		json.NewEncoder(w).Encode(task)
	}
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract task ID from URL
	path := r.URL.Path[len("/api/tasks/"):]
	
	if path == "" {
		http.Error(w, "Task ID required", http.StatusBadRequest)
		return
	}

	// Handle toggle endpoint
	if len(path) > 7 && path[len(path)-7:] == "/toggle" {
		taskIDStr := path[:len(path)-7]
		taskID, err := strconv.Atoi(taskIDStr)
		if err != nil {
			http.Error(w, "Invalid task ID", http.StatusBadRequest)
			return
		}

		if r.Method == "POST" {
			for i, task := range taskManager.Tasks {
				if task.ID == taskID {
					taskManager.Tasks[i].Completed = !task.Completed
					if taskManager.Tasks[i].Completed {
						now := time.Now()
						taskManager.Tasks[i].CompletedAt = &now
					} else {
						taskManager.Tasks[i].CompletedAt = nil
					}
					json.NewEncoder(w).Encode(taskManager.Tasks[i])
					return
				}
			}
			http.Error(w, "Task not found", http.StatusNotFound)
		}
		return
	}

	// Handle regular task operations
	taskID, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "PUT":
		var updatedTask Task
		if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for i, task := range taskManager.Tasks {
			if task.ID == taskID {
				updatedTask.ID = taskID
				updatedTask.CreatedAt = task.CreatedAt
				if updatedTask.Completed && !task.Completed {
					now := time.Now()
					updatedTask.CompletedAt = &now
				} else if !updatedTask.Completed {
					updatedTask.CompletedAt = nil
				} else {
					updatedTask.CompletedAt = task.CompletedAt
				}
				taskManager.Tasks[i] = updatedTask
				json.NewEncoder(w).Encode(updatedTask)
				return
			}
		}
		http.Error(w, "Task not found", http.StatusNotFound)

	case "DELETE":
		for i, task := range taskManager.Tasks {
			if task.ID == taskID {
				taskManager.Tasks = append(taskManager.Tasks[:i], taskManager.Tasks[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		http.Error(w, "Task not found", http.StatusNotFound)
	}
}
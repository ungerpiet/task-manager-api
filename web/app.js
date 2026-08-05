function loadTasks() {
    fetch('http://localhost:8080/tasks')
        .then(response => response.json())
        .then(tasks => {
            console.log("Empfangene Tasks:", tasks);
            const list = document.getElementById('task-list');
            list.innerHTML = '';
            
            if (tasks) {
                tasks.forEach(task => {
                const date = new Date(task.created_at).toLocaleDateString('de-DE');
                const status = task.completed ? 'Erledigt' : 'Offen'

                const li = document.createElement('li');

                const headerDiv = document.createElement('div')
                headerDiv.textContent = task.title;

                const checkbox = document.createElement('input')
                checkbox.type = 'checkbox'
                checkbox.checked = task.completed

                const detailDiv = document.createElement('div')
                detailDiv.innerHTML = `
                    <p>Beschreibung: ${task.description}</p>
                    <p>Status: ${status}</p>
                    <p>Erstellt am ${date}</p>
                `;
                detailDiv.style.display = 'none'

                headerDiv.addEventListener('click', function() {
                    if (detailDiv.style.display === 'none') {
                        detailDiv.style.display = 'block';
                    } else {
                        detailDiv.style.display = 'none';
                    }
                });

                checkbox.onchange = function() {
                    toggleTaskStatus(task.id, task.title, task.description, checkbox.checked);
                }

                const deleteBtn = document.createElement('button');
                deleteBtn.textContent = 'Löschen';
                deleteBtn.onclick = function() {
                    deleteTask(task.id);
                };
                

                list.appendChild(li);
                li.append(checkbox);
                li.appendChild(headerDiv);
                li.appendChild(detailDiv);
                li.appendChild(deleteBtn);
                });
            }
        });
}

function toggleTaskStatus (id, title, description, isCompleted) {
    fetch('http://localhost:8080/tasks/${id}', {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ 
            title: title,
            description: description,
            completed: isCompleted 
        })
    })
    .then(() => {
        loadTasks();
    });
}

function deleteTask(id) {
    fetch(`http://localhost:8080/tasks/${id}`, {
        method: 'DELETE'
    })
    .then(() => {
        loadTasks();
    });
}

function createTask() {
    const input = document.getElementById('task-input');
    const meinTitel = input.value.trim();

    if (!meinTitel) return;

    fetch('http://localhost:8080/tasks', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ title: meinTitel })
    })
    .then(() => {
        input.value = '';
        loadTasks();
    });
}

window.onload = function(){
    loadTasks();

    document.getElementById('task-form').addEventListener('submit', function(event) {
        event.preventDefault();
        createTask();
    });
}
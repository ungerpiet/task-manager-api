function loadTasks() {
    fetch('http://localhost:8080/tasks')
        .then(response => response.json())
        .then(tasks => {
            const list = document.getElementById('task-list');
            list.innerHTML = '';
            
            if (tasks) {
                tasks.forEach(task => {
                const li = document.createElement('li');
                li.textContent = task.title;

                const deleteBtn = document.createElement('button');
                deleteBtn.textContent = 'Löschen';
                deleteBtn.onclick = function() {
                    deleteTask(task.id);
                };
                

                list.appendChild(li);
                li.appendChild(deleteBtn);
                });
            }
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
    const meinTitel = input.value;

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
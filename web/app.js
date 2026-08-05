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
                 list.appendChild(li);
                });
            }
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
        loadTasks();
    });
}

window.onload = loadTasks;
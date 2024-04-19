fetch('http://localhost:3000/api/login', {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
    },
    body: JSON.stringify({
        email: 'joao@gmail.com',
        password: '1234'
    })
})
.then(response => response.json())
.then(data => {
    console.log(data),
    console.log(data.message);
})
.catch(error => console.log('Error', error));


// Fazer logout
fetch('http://localhost:3000/api/logout', {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
    }
})
.then(response => response.json())
.then(data => {
    console.log(data),
    console.log(data.message);
})
.catch(error => console.log('Error', error));


// POST de tarefas
fetch('http://localhost:3000/api/cadastro/tarefa', {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
    },
    body: JSON.stringify({
        title: "Aprender Javascript",
        description: "Descrição",
        hora:"23:59"
      })
})
.then(response => response.json())
.then(data => {
    console.log(data),
    console.log(data.message);
})
.catch(error => console.log('Error', error));

// GET de tarefas
fetch('http://localhost:3000/api/tarefas')
.then(response => response.json())
.then(data => console.log(data))
.catch(error => console.log('Error', error));

// PUT de tarefas
fetch('http://localhost:3000/api/tarefa/9', {
    method: 'PUT',
    headers: {
        'Content-Type': 'application/json',
    },
    body: JSON.stringify({
        title: "Javascript",
        description: "OK",
        hora:"13:59"
      })
})
.then(response => response.json())
.then(data => {
    console.log(data),
    console.log(data.message);
})
.catch(error => console.log('Error', error));

// DELETE de Tarefas
fetch('http://localhost:3000/api/tarefa/9', {
    method: 'DELETE', 
    headers: {
        'Content-Type': 'application/json',
    }
}
)
.then(response => response.json())
.then(data => {
    console.log(data),
    console.log(data.message);
})
.catch(error => console.log('Error', error));
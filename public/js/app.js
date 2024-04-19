// logout
fetch('http://localhost:3000/api/logout', {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
    },
    body: JSON.stringify({}),
}).then(response => {
    if (response.ok) {
        alert('Logout bem sucedido!');
        window.location.href = 'Tela_login.html';
    }
}
    )
    .catch(error => {
        console.error(error);
    });
import { backendAPI } from './script.js';

document.getElementById('form-cadastro').addEventListener('submit', async function(event) {
    event.preventDefault();
    var name = document.getElementById('name').value;
    var email = document.getElementById('email').value;
    var password = document.getElementById('password').value;

    try {
        const response = await backendAPI.post('api/cadastro/usuario', { name, email, password });
        if (response.error) {
            throw new Error(response.error);
        } else {
            alert('Cadastro bem sucedido!');
            console.log(response);
            window.location.href = 'Tela_login.html';
        }
    } catch (error) {
        alert('Erro ao cadastrar usuário: ' + error.message);
        console.error(error);
    }
});
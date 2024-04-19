import { backendAPI } from "./script.js";

document.getElementById('login-form').addEventListener('submit', async function(event) {
    event.preventDefault();
    var email = document.getElementById('email').value;
    var password = document.getElementById('password').value;

    try {
        const response = await backendAPI.post('api/login', { email, password });
        if (response.error) {
            throw new Error(response.error);
        } else {
            alert('Login bem sucedido!');
            console.log(response);
            window.location.href = 'Tela_lista.html';
        }
    } catch (error) {
        alert('Erro ao fazer login: ' + error.message);
        console.error(error);
    }
});
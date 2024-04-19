import { backendAPI } from './script.js';

document.getElementById('adicionarTarefa').addEventListener('click', async function() {
    const nomeTarefa = document.getElementById('nomeTarefa').value;
    const horaTarefa = document.getElementById('horarioTarefa').value;
    const descricaoTarefa = document.getElementById('descricaoTarefa').value;

    try {
        const response = await backendAPI.post('api/cadastro/tarefa', {
            title: nomeTarefa,
            hora: horaTarefa,
            description: descricaoTarefa
        });
        if (response.error) {
            throw new Error(response.error);
        } else {
            alert('Tarefa cadastrada com sucesso!');
            window.location.href = 'Tela_lista.html';
        }
    } catch (error) {
        alert('Erro ao cadastrar tarefa: ' + error.message);
        console.error(error);
    }
}
);
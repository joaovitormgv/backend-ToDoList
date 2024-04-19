import { backendAPI } from "./script.js";

function carregarDetalhesTarefa(idTarefa) {
    backendAPI.get(`api/tarefa/${idTarefa}`)
        .then(data => {
            document.getElementById('titulo').innerText = data.title;            

            // verificar se há descrição
            if (data.description) {
                document.getElementById('nota').value = data.description;
            } else {
                document.getElementById('nota').placeholder = 'Digite uma descrição';
            }
        })
        .catch(error => { console.error(error); });
}

// Obter id da tarefa
const urlParams = new URLSearchParams(window.location.search);
const idTarefa = urlParams.get('id');

if (idTarefa) {
    carregarDetalhesTarefa(idTarefa);
} else {
    alert('ID da tarefa não informado na URL.');
}

// Evento de clique no botão salvar
document.getElementById('salvar').addEventListener('click', function() {
    var nota = document.getElementById('nota').value;
    var titulo = document.getElementById('titulo').innerText;

    salvarNota(idTarefa, titulo, nota);
});

// Função para salvar nota
function salvarNota(idTarefa, titulo, nota) {
    if (nota) {
        backendAPI.put(`api/tarefa/${idTarefa}`, { title: titulo, description: nota })
            .then(response => {
                if (response.error) {
                    throw new Error(response.error);
                } else {
                    alert('Nota salva com sucesso!');
                }
            })
            .catch(error => {
                alert('Erro ao salvar nota: ' + error.message);
                console.error(error);
            });
    } else {
        alert('Digite uma nota para salvar.');
    }
}
import { backendAPI } from './script.js';

document.addEventListener('DOMContentLoaded', async function () {
    // Função para carregar a lista de tearefas
    function carregarTarefas() {
        backendAPI.get('api/tarefas')
            .then(data => {
                const lista = document.getElementById('taskList');
                lista.innerHTML = '';

                data.forEach(tarefa => {
                    const item = document.createElement('li');
                    
                    const nomeSpan = document.createElement('span');
                    nomeSpan.textContent = tarefa.title;
                    const horarioSpan = document.createElement('span');
                    horarioSpan.textContent = tarefa.hora;
                    // checkbox para marcar a tarefa como concluída com base em completed no backend
                    const checkbox = document.createElement('input');
                    checkbox.type = 'checkbox';
                    checkbox.checked = tarefa.completed;
                    checkbox.addEventListener('change', async function () {
                        await backendAPI.put(`api/tarefa/${tarefa.id}`, {
                            completed: checkbox.checked
                        });
                        carregarTarefas();
                    });

                    item.appendChild(checkbox);
                    item.appendChild(nomeSpan);
                    item.appendChild(horarioSpan);

                    lista.appendChild(item);
                });
            })
            .catch(error => console.error('Erro ao carregar tarefas:', error));
        }

    // Carregar tarefas
    carregarTarefas();
    });

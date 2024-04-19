import { backendAPI } from './script.js';

document.addEventListener('DOMContentLoaded', async function () {
    // Função para carregar a lista de tearefas
    function carregarTarefas() {
        backendAPI.get('api/tarefas')
            .then(data => {
                const lista = document.getElementById('Lista_de_tarefas');
                lista.innerHTML = '';

                data.forEach(tarefa => {
                    const div = document.createElement('div');
                    div.classList.add('tarefa');
                    div.innerHTML = `
                        <input type="checkbox" id="tarefa-${tarefa.id}" ${tarefa.completed ? 'checked' : ''}>
                        <p >${tarefa.title}</p>
                        <p>${tarefa.hora}</p>
                    `;
                    
                    // Adicionar evento para marcar tarefa como concluída
                    div.querySelector('input').addEventListener('change', async function () {
                        tarefa.completed = this.checked;
                        await backendAPI.put(`api/tarefa/${tarefa.id}`, tarefa);
                        carregarTarefas();
                    });
                    
                    if (tarefa.completed) {
                        div.style.textDecoration = 'line-through';
                    }

                    lista.appendChild(div);
                });
            })
            .catch(error => console.error('Erro ao carregar tarefas:', error));
        }

    // Carregar tarefas
    carregarTarefas();
    });

    // Adicionar evento para ir para a página de adicionar tarefa
    document.getElementById('Ir_para_a_pagina_de_adicionar_tarefa').addEventListener('click', function() {
        window.location.href = 'Tela_adicionar.html'; 
    });


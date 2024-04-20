import { backendAPI } from './script.js';

// Função para carregar tarefas
export async function carregarTarefas(search = '') {
    let endpoint = 'api/tarefas';
    if (search) {
        endpoint += `?search=${encodeURIComponent(search)}`;
    }
    await backendAPI.get(endpoint)
        .then(data => {
            const lista = document.getElementById('Lista_de_tarefas');
            lista.innerHTML = '';

            data.forEach(tarefa => {
                const div = document.createElement('div');
                div.classList.add('Tarefa');
                div.innerHTML = `
                <div class="dados">
                    <input type="checkbox" id="tarefa-${tarefa.id}" ${tarefa.completed ? 'checked' : ''}>
                    <p>${tarefa.title}</p>
                </div>
                <div class="dados">
                    <p>${tarefa.hora}</p>
                    <button id="deletar_tarefa">X</button>
                </div>
                `;
                
                // Adicionar evento para marcar tarefa como concluída
                div.querySelector('input').addEventListener('change', async function () {
                    tarefa.completed = this.checked;
                    await backendAPI.put(`api/tarefa/${tarefa.id}`, tarefa);
                    carregarTarefas(search);
                });
                
                if (tarefa.completed) {
                    div.style.textDecoration = 'line-through';
                }

                lista.appendChild(div);
            });
        })
        .catch(error => console.error('Erro ao carregar tarefas:', error));
        }

document.addEventListener('DOMContentLoaded', async function () {
    // Carregar tarefas
    carregarTarefas();
    });

    // Adicionar evento para pesquisar tarefas
    document.getElementById("Botão_de_pesquisar").addEventListener('click', async function() {
        const pesquisa = document.getElementById("Botão_de_pesquisar_input").value;
        carregarTarefas(pesquisa);
    });

    // Adicionar evento para ir para a página de adicionar tarefa
    document.getElementById('Ir_para_a_pagina_de_adicionar_tarefa').addEventListener('click', function() {
        window.location.href = 'Tela_adicionar.html'; 
    });


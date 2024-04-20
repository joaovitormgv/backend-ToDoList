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
                // Verificar se hora é nula
                console.log(tarefa.hora);
                if (tarefa.hora == null) {
                    div.innerHTML = `
                    <div class="dados">
                        <input type="checkbox" id="tarefa-${tarefa.id}" ${tarefa.completed ? 'checked' : ''}>
                        <p class="tarefa-titulo">${tarefa.title}</p>
                    </div>
                    <div class="dados">
                        <p class="tarefa-hora">Dia todo</p>
                            <a class="deletar-tarefa""><img src="../assets/trash-icon.svg"></img</a>
                            <a href="Tela_nota.html?id=${tarefa.id}"><img src="../assets/edit-icon.svg"></a>
                    </div>
                    `;
                } else {
                    div.innerHTML = `
                    <div class="dados">
                        <input type="checkbox" id="tarefa-${tarefa.id}" ${tarefa.completed ? 'checked' : ''}>
                        <p class="tarefa-titulo">${tarefa.title}</p>
                    </div>
                    <div class="dados">
                        <p class="tarefa-hora">${tarefa.hora}</p>
                            <a class="deletar-tarefa""><img src="../assets/trash-icon.svg"></img</a>
                            <a href="Tela_nota.html?id=${tarefa.id}"><img src="../assets/edit-icon.svg"></a>
                    </div>
                    `;
                }

            // Mudar titulo
                const titleElement = div.querySelector('.tarefa-titulo');
                titleElement.addEventListener('click', async function() {
                    const editInput = document.createElement('input');
                    editInput.type = 'text';
                    editInput.value = tarefa.title;
                    editInput.classList.add('input_editado');
                    titleElement.replaceWith(editInput);
                    editInput.focus();

                    editInput.addEventListener('keydown', async function(event) {
                        if (event.key === 'Enter') {
                            tarefa.title = editInput.value;
                            // Atualizar tarefa
                            await backendAPI.put(`api/tarefa/${tarefa.id}`, {title: tarefa.title, completed: tarefa.completed});
                            carregarTarefas(search);
                        }
                    });

                    editInput.addEventListener('blur', async function() {
                        tarefa.title = editInput.value;
                        // Atualizar tarefa
                        await backendAPI.put(`api/tarefa/${tarefa.id}`, {title: tarefa.title, completed: tarefa.completed});
                        carregarTarefas(search);
                    });
                });
            
            // Mudar horário
                const timeElement = div.querySelector('.tarefa-hora');
                timeElement.addEventListener('click', function() {
                    const editInput = document.createElement('input');
                    editInput.type = 'text';
                    editInput.value = tarefa.hora;
                    editInput.classList.add('input_editado'); 
                    timeElement.replaceWith(editInput); 

                    editInput.addEventListener('keydown', async function(event) {
                        if (event.key === 'Enter') {
                            tarefa.hora = editInput.value;
                            // Atualizar tarefa
                            await backendAPI.put(`api/tarefa/${tarefa.id}`, {hora: tarefa.hora, tarefa: tarefa.completed});
                            carregarTarefas(search);
                        }
                    });

                    editInput.addEventListener('blur', async function() {
                        tarefa.hora = editInput.value;
                        // Atualizar tarefa
                        await backendAPI.put(`api/tarefa/${tarefa.id}`, {hora: tarefa.hora, tarefa: tarefa.completed});
                        carregarTarefas(search);
                    });
            
                    editInput.focus();
                });



                // Adicionar evento para marcar tarefa como concluída
                div.querySelector('input').addEventListener('change', async function () {
                    tarefa.completed = this.checked;
                    await backendAPI.put(`api/tarefa/${tarefa.id}`, tarefa);
                    carregarTarefas(search);
                });
                // Adicionar evento para deletar tarefa
                div.querySelector('.deletar-tarefa').addEventListener('click', async function () {
                    const confirmacao = confirm('Tem certeza que deseja deletar a tarefa?');
                    if (confirmacao) {
                        await backendAPI.delete(`api/tarefa/${tarefa.id}`);
                        carregarTarefas(search);
                    }
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


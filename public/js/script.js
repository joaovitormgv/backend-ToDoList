const API_URL = 'http://localhost:3000';

export const backendAPI = {
    async get(endpoint) {
        const response = await fetch(`${API_URL}/${endpoint}`);
        const data = await response.json();
        return data;
    },
    
    async post(endpoint, body) {
        const response = await fetch(`${API_URL}/${endpoint}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(body),
        });
        const data = await response.json();
        return data;
    },
    
    async put(endpoint, body) {
        const response = await fetch(`${API_URL}/${endpoint}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(body),
        });
        const data = await response.json();
        return data;
    },

    async delete(endpoint) {
        const response = await fetch(`${API_URL}/${endpoint}`, {
            method: 'DELETE',
        });
        const data = await response.json();
        return data;
    }, 

    async getWithBody(endpoint, body) {
        const response = await fetch(`${API_URL}/${endpoint}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(body),
        });
        const data = await response.json();
        return data;
    }
}
const API_URL_DEBETS = '/debets';

async function fetchJson(url) {
    const response = await fetch(url);
    if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
    }
    return response.json();
}

export async function fetchDebetData() {
    try {
        return await fetchJson(API_URL_DEBETS);
    } catch (error) {
        console.error('Failed to fetch data:', error);
        throw error;
    }
}

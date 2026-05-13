const API_URL_DEBET = '/debet';
const API_URL_RESPONSE = '/response';

async function fetchJson(url) {
    const response = await fetch(url);
    if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
    }
    return response.json();
}

export async function fetchDebetData() {
    try {
        return await fetchJson(API_URL_DEBET);
    } catch (error) {
        console.error('Failed to fetch data:', error);
        throw error;
    }
}

export async function fetchResponseData() {
    try {
        return await fetchJson(API_URL_RESPONSE);
    } catch (error) {
        console.error('Failed to fetch data:', error);
        throw error;
    }
}


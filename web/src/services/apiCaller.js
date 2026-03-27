// const API_URL_RESPONSE = 'http://localhost:8181/debet';
const API_URL_DEBET = '/debet'
const API_URL_RESPONSE = '/response'


export async function fetchDebetData() {
    try {
        const response = await fetch(API_URL_DEBET);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Failed to fetch data:', error);
        throw error;
    }
}

/*export async function fetchDebetDataWithMIP() {
    try {
        const response = await fetch(API_URL_DEBET_WITH_MIP);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Failed to fetch data:', error);
        throw error;
    }
}*/

export async function fetchResponseData() {
    try {
        const response = await fetch(API_URL_RESPONSE);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Failed to fetch data:', error);
        throw error;
    }
}

/*
export async function fetchResponseDataWithMIP() {
    try {
        const response = await fetch(API_URL_RESPONSE_WITH_MIP);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Failed to fetch data:', error);
        throw error;
    }
}*/

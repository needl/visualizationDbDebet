const API_URL_RESPONSE = '/response';

export async function fetchResponseData() {
    const response = await fetch(API_URL_RESPONSE);
    if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
    }
    return response.json();
}

const API_URL_RESPONSES = '/responses/with-mip';

export async function fetchResponseData() {
    const response = await fetch(API_URL_RESPONSES);
    if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
    }
    return response.json();
}

export async function fetchObjectData(objectName, counterpartyName) {
    if (!counterpartyName) {
        throw new Error('counterpartyName is required');
    }

    const params = new URLSearchParams({ objectName, counterpartyName });
    const url = `/objects/search?${params.toString()}`;
    const res = await fetch(url);
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    return res.json();
}

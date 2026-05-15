export async function fetchObjectData(objectName) {
    const url = `/objects/search?objectName=${encodeURIComponent(objectName)}`;
    const res = await fetch(url);
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    return res.json();
}

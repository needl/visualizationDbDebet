export async function fetchObjectList(orgName) {
    const url = `/objects/${encodeURIComponent(orgName)}`;
    const res = await fetch(url);
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    return await res.json ();
}

export async function fetchObjectData(orgName, objectName) {
    const url = `/objects/search?orgName=${encodeURIComponent(orgName)}&objectName=${encodeURIComponent(objectName)}`;
    const res = await fetch(url);
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    return await res.json ();
}
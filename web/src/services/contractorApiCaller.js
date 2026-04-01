export async function fetchContractorsByFactor(orgName, columnName) {
    const url = `/contractor/${encodeURIComponent(orgName)}?columnName=${encodeURIComponent(columnName)}`;
    const response = await fetch(url);
    if (!response.ok) {
        throw new Error(`HTTP ${response.status}`);
    }
    return response.json();
}
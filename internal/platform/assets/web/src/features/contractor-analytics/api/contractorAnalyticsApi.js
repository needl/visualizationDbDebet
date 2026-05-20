async function fetchJson(url) {
    const res = await fetch(url);
    if (!res.ok) {
        const message = (await res.text()).trim();
        if (message) {
            throw new Error(`HTTP ${res.status}: ${message}`);
        }
        throw new Error(`HTTP ${res.status}`);
    }
    return res.json();
}

export async function fetchContractorNames() {
    return fetchJson('/contractor-analytics');
}

export async function fetchContractorAnalytics(contractorName) {
    const url = `/contractor-analytics/${encodeURIComponent(contractorName)}`;
    return fetchJson(url);
}

export async function fetchContractorObjectDetails(contractorName, customerName, objectName) {
    const params = new URLSearchParams({
        customerName,
        objectName
    });

    const url = `/contractor-analytics/${encodeURIComponent(contractorName)}/object-details?${params.toString()}`;
    return fetchJson(url);
}

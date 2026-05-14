export async function fetchContractorsByFactor(orgName, columnName) {
    const url = `/contractor/${encodeURIComponent(orgName)}?columnName=${encodeURIComponent(columnName)}`;
    const response = await fetch(url);
    if (!response.ok) {
        throw new Error(`HTTP ${response.status}`);
    }
    return response.json();
}

export async function fetchContractorDebt(orgName, counterpartyName) {
    const url = `/contractor/${encodeURIComponent(orgName)}/debt?counterpartyName=${encodeURIComponent(counterpartyName)}`;
    const res = await fetch(url);
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    return res.json();
}

export async function fetchContractorOverdue(orgName, counterpartyName) {
    const url = `/contractor/${encodeURIComponent(orgName)}/overdue?counterpartyName=${encodeURIComponent(counterpartyName)}`;
    const res = await fetch(url);
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    return res.json();
}

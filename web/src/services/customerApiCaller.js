export async function fetchCustomers() {
    const res = await fetch('/customer');
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    return res.json();
}

export async function fetchCustomerSummary(orgName) {
    const url = `/customer/summary/${encodeURIComponent(orgName)}`;
    const res = await fetch(url);
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    return res.json();
}

export async function fetchCustomerTopDebtors(orgName) {
    const url = `/customer/top-debtors/${encodeURIComponent(orgName)}`;
    const res = await fetch(url);
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    return res.json();
}

export async function fetchCustomerTopOverdue(orgName) {
    const url = `/customer/top-debtors-overdue/${encodeURIComponent(orgName)}`;
    const res = await fetch(url);
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    return res.json();
}

export async function fetchCustomerBlockFactors(orgName) {
    const url = `/customer/blockFactors/${encodeURIComponent(orgName)}`;
    const res = await fetch(url);
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    return res.json();
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

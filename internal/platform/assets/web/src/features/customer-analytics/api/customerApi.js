export async function fetchCustomers() {
    const res = await fetch('/customers');
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    return res.json();
}

export async function fetchCustomerSummary(orgName) {
    const url = `/customers/summary/${encodeURIComponent(orgName)}`;
    const res = await fetch(url);
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    return res.json();
}

export async function fetchCustomerTopDebtors(orgName) {
    const url = `/customers/top-debtors/${encodeURIComponent(orgName)}`;
    const res = await fetch(url);
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    return res.json();
}

export async function fetchCustomerTopOverdue(orgName) {
    const url = `/customers/top-debtors-overdue/${encodeURIComponent(orgName)}`;
    const res = await fetch(url);
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    return res.json();
}

export async function fetchCustomerBlockFactors(orgName) {
    const url = `/customers/block-factors/${encodeURIComponent(orgName)}`;
    const res = await fetch(url);
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    return res.json();
}


export async function fetchCustomers() {
    const res = await fetch('/customer');
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    return res.json();
}

export async function fetchCustomerSummary(orgName) {
    const url = `/customer/${encodeURIComponent(orgName)}/summary`;
    const res = await fetch(url);
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    return res.json();
}

export async function fetchCustomerTopDebtors(orgName) {
    const url = `/customer/${encodeURIComponent(orgName)}/top-debtors`;
    const res = await fetch(url);
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    return res.json();
}

export async function fetchCustomerTopOverdue(orgName) {
    const url = `/customer/${encodeURIComponent(orgName)}/top-debtors-overdue`;
    const res = await fetch(url);
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    return res.json();
}

export async function fetchCustomerBlockFactors(orgName) {
    const url = `/customer/${encodeURIComponent(orgName)}/blockFactors`;
    const res = await fetch(url);
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    return res.json();
}
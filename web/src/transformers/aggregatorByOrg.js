export function aggregateByOrg(items) {
    const aggregated = {};

    items.forEach(item => {
        const orgName = item.source_org_name;
        if (!orgName) return;

        if (!aggregated[orgName]) {
            aggregated[orgName] = {
                contractAmount: 0,
                debetTotal: 0,
                debetOverdose: 0
            };
        }

        if (typeof item.contract_amount === 'number') {
            aggregated[orgName].contractAmount += item.contract_amount;
        }
        if (typeof item.debt_2025_12_31_total === 'number') {
            aggregated[orgName].debetTotal += item.debt_2025_12_31_total;
        }
        if (typeof item.debt_2025_12_31_overdue === 'number') {
            aggregated[orgName].debetOverdose += item.debt_2025_12_31_overdue;
        }
    });
        Math.round(aggregated)

    return aggregated;
}

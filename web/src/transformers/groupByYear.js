export function aggregateByYear(items) {
    const map = new Map(); // key: source_org_name, value: { year2024, year2025 }

    items.forEach(item => {
        const orgName = item.source_org_name;
        if (!orgName) return;

        if (!map.has(orgName)) {
            map.set(orgName, { year2024: 0, year2025: 0 });
        }
        const entry = map.get(orgName);
        if (typeof item.debt_2024_12_31_total === 'number') {
            entry.year2024 += item.debt_2024_12_31_total;
        }
        if (typeof item.debt_2025_12_31_total === 'number') {
            entry.year2025 += item.debt_2025_12_31_total;
        }
    });

    const names = [];
    const series2024 = [];
    const series2025 = [];

    for (const [org, data] of map.entries()) {
        names.push(org);
        series2024.push(data.year2024);
        series2025.push(data.year2025);
    }

    return {
        names: names,
        series: [
            { name: '2024', data: series2024 },
            { name: '2025', data: series2025 }
        ]
    };
}
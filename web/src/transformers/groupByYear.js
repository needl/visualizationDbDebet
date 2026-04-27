export function aggregateByYearStacked(items) {
    const map = new Map();

    items.forEach(item => {
        const orgName = item.source_org_name;
        if (!orgName) return;

        if (!map.has(orgName)) {
            map.set(orgName, {
                total2024: 0,
                overdue2024: 0,
                total2025: 0,
                overdue2025: 0
            });
        }
        const entry = map.get(orgName);

        if (typeof item.debt_2024_12_31_total === 'number') {
            entry.total2024 += item.debt_2024_12_31_total;
        }
        if (typeof item.debt_2024_12_31_overdue === 'number') {
            entry.overdue2024 += item.debt_2024_12_31_overdue;
        }

        if (typeof item.debt_2025_12_31_total === 'number') {
            entry.total2025 += item.debt_2025_12_31_total;
        }
        if (typeof item.debt_2025_12_31_overdue === 'number') {
            entry.overdue2025 += item.debt_2025_12_31_overdue;
        }
    });

    const names = [];
    const current2024 = [];
    const overdue2024 = [];
    const current2025 = [];
    const overdue2025 = [];

    for (const [org, data] of map.entries()) {
        names.push(org);
        // Текущая = Общая − Просроченная (если данные корректны)
        const cur24 = Math.max(0, data.total2024 - data.overdue2024);
        const cur25 = Math.max(0, data.total2025 - data.overdue2025);
        current2024.push(cur24);
        overdue2024.push(data.overdue2024);
        current2025.push(cur25);
        overdue2025.push(data.overdue2025);
    }

    return {
        names,
        series: [
            { name: '31.12.2024 Текущая дебиторская задолженность',   stack: '2024', data: current2024 },
            { name: '31.12.2024 Просроченная дебиторская задолженность', stack: '2024', data: overdue2024 },
            { name: '31.12.2025 Текущая дебиторская задолженность',   stack: '2025', data: current2025 },
            { name: '31.12.2025 Просроченная дебиторская задолженность', stack: '2025', data: overdue2025 }
        ]
    };
}

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
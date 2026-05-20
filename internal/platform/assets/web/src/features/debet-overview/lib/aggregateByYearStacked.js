export function aggregateByYearStacked(items) {
    const map = new Map();

    items.forEach((item) => {
        const orgName = item.source_org_name;
        if (!orgName) return;

        if (!map.has(orgName)) {
            map.set(orgName, {
                total20241231: 0,
                overdue20241231: 0,
                total20250331: 0,
                overdue20250331: 0,
                total20251231: 0,
                overdue20251231: 0,
                total20260331: 0,
                overdue20260331: 0
            });
        }

        const entry = map.get(orgName);

        if (typeof item.debt_2024_12_31_total === 'number') {
            entry.total20241231 += item.debt_2024_12_31_total;
        }
        if (typeof item.debt_2024_12_31_overdue === 'number') {
            entry.overdue20241231 += item.debt_2024_12_31_overdue;
        }

        if (typeof item.debt_2025_03_31_total === 'number') {
            entry.total20250331 += item.debt_2025_03_31_total;
        }
        if (typeof item.debt_2025_03_31_overdue === 'number') {
            entry.overdue20250331 += item.debt_2025_03_31_overdue;
        }

        if (typeof item.debt_2025_12_31_total === 'number') {
            entry.total20251231 += item.debt_2025_12_31_total;
        }
        if (typeof item.debt_2025_12_31_overdue === 'number') {
            entry.overdue20251231 += item.debt_2025_12_31_overdue;
        }

        if (typeof item.debt_2026_03_31_total === 'number') {
            entry.total20260331 += item.debt_2026_03_31_total;
        }
        if (typeof item.debt_2026_03_31_overdue === 'number') {
            entry.overdue20260331 += item.debt_2026_03_31_overdue;
        }
    });

    const names = [];
    const current20241231 = [];
    const overdue20241231 = [];
    const current20250331 = [];
    const overdue20250331 = [];
    const current20251231 = [];
    const overdue20251231 = [];
    const current20260331 = [];
    const overdue20260331 = [];

    for (const [org, data] of map.entries()) {
        names.push(org);

        current20241231.push(Math.max(0, data.total20241231 - data.overdue20241231));
        overdue20241231.push(data.overdue20241231);

        current20250331.push(Math.max(0, data.total20250331 - data.overdue20250331));
        overdue20250331.push(data.overdue20250331);

        current20251231.push(Math.max(0, data.total20251231 - data.overdue20251231));
        overdue20251231.push(data.overdue20251231);

        current20260331.push(Math.max(0, data.total20260331 - data.overdue20260331));
        overdue20260331.push(data.overdue20260331);
    }

    return {
        names,
        series: [
            { name: '31.12.2024 Текущая дебиторская задолженность', stack: '2024-12-31', data: current20241231 },
            { name: '31.12.2024 Просроченная дебиторская задолженность', stack: '2024-12-31', data: overdue20241231 },
            { name: '31.03.2025 Текущая дебиторская задолженность', stack: '2025-03-31', data: current20250331 },
            { name: '31.03.2025 Просроченная дебиторская задолженность', stack: '2025-03-31', data: overdue20250331 },
            { name: '31.12.2025 Текущая дебиторская задолженность', stack: '2025-12-31', data: current20251231 },
            { name: '31.12.2025 Просроченная дебиторская задолженность', stack: '2025-12-31', data: overdue20251231 },
            { name: '31.03.2026 Текущая дебиторская задолженность', stack: '2026-03-31', data: current20260331 },
            { name: '31.03.2026 Просроченная дебиторская задолженность', stack: '2026-03-31', data: overdue20260331 }
        ]
    };
}

export function aggregateByYear(items) {
    return aggregateByYearStacked(items);
}

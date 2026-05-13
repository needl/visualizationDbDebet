

function findMaxContractObject(objects) {
    let maxObj = null;
    let maxAmount = -Infinity;
    objects.forEach(obj => {
        const amount = obj.contract_amount || 0;
        if (amount > maxAmount) {
            maxAmount = amount;
            maxObj = obj;
        }
    });
    return maxObj;
}


function formatDate(dateStr) {
    if (!dateStr) return '—';
    const d = new Date(dateStr);
    if (isNaN(d.getTime())) return '—';
    return d.toLocaleDateString('ru-RU', { day: '2-digit', month: '2-digit', year: 'numeric' });
}


function getEarliestStartDate(objects) {
    let earliest = null;
    objects.forEach(obj => {
        if (obj.work_start_date) {
            const d = new Date(obj.work_start_date);
            if (!isNaN(d.getTime())) {
                if (!earliest || d < earliest) earliest = d;
            }
        }
    });
    return earliest ? formatDate(earliest.toISOString()) : '—';
}


function getLatestEndDate(objects) {
    let latest = null;
    objects.forEach(obj => {
        if (obj.work_end_date) {
            const d = new Date(obj.work_end_date);
            if (!isNaN(d.getTime())) {
                if (!latest || d > latest) latest = d;
            }
        }
    });
    return latest ? formatDate(latest.toISOString()) : '—';
}


export function aggregateObjectMetrics(rawData) {
    if (!rawData || rawData.length === 0) {
        return null;
    }

    const maxObj = findMaxContractObject(rawData);

    const sums = {
        hardContractPrice: 0,
        contractAmount: 0,
        paidAmount: 0,
        acceptedAmount: 0,
    };

    rawData.forEach(item => {
        sums.hardContractPrice += item.hard_contract_price || 0;
        sums.contractAmount += item.contract_amount || 0;
        sums.paidAmount += item.paid_amount || 0;
        sums.acceptedAmount += item.accepted_amount || 0;
    });

    const contractorName = maxObj?.counterparty_name || '—';
    const buildReady = maxObj?.build_ready_percent ?? null;
    const permission = maxObj?.permission_to_enter ?? null;
    const conclusion = maxObj?.conclusion ?? null;

    const startDate = getEarliestStartDate(rawData);
    const endDate = getLatestEndDate(rawData);

    return {
        contractorName,
        workStartDate: startDate,
        workEndDate: endDate,
        buildReadyPercent: buildReady,
        permissionToEnter: permission,
        conclusionMke: conclusion,
        hardContractPrice: sums.hardContractPrice,
        contractAmount: sums.contractAmount,
        paidAmount: sums.paidAmount,
        acceptedAmount: sums.acceptedAmount,
    };
}


export function prepareChartData(rawData) {
    let total2024 = 0, overdue2024 = 0;
    let total2025 = 0, overdue2025 = 0;

    rawData.forEach(item => {
        total2024 += item.debt_2024_12_31_total || 0;
        overdue2024 += item.debt_2024_12_31_overdue || 0;
        total2025 += item.debt_2025_12_31_total || 0;
        overdue2025 += item.debt_2025_12_31_overdue || 0;
    });

    return {
        categories: ['31.12.2024', '31.12.2025'],
        series: [
            {
                name: 'Дебиторская задолженность',
                type: 'line',
                data: [total2024, total2025],
                itemStyle: { color: '#3b82f6' }
            },
            {
                name: 'Просроченная задолженность',
                type: 'line',
                data: [overdue2024, overdue2025],
                itemStyle: { color: '#ef4444' }
            }
        ]
    };
}

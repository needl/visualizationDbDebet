export const dashboardConfig = [
    {
        title: 'Общая статистика',
        type: 'stats',
        metrics: [
            { key: 'count_source_org', label: 'Количество заказчиков', format: 'number' },
            { key: 'count_contracts', label: 'Количество контрактов', format: 'number' },
            { key: 'sum_contract_amount', label: 'Сумма контрактов', format: 'currency' },
            { key: 'sum_debet_total', label: 'Сумма дебиторской задолженности', format: 'currency' },
            { key: 'sum_debet_overdue', label: 'Сумма просроченной задолженности', format: 'currency' }
        ],
        table: true,
        chart: {
            metric: 'debetByYear',
            title: 'Динамика задолженности по годам (2024 vs 2025)'
        }
    },
    {
        title: 'Дебиторка по заказчикам',
        type: 'charts',
        charts: [
            { metric: 'contractAmount', title: 'Сумма контрактов' },
            { metric: 'debetTotal', title: 'Дебиторская задолженность' },
            { metric: 'debetOverdose', title: 'Просроченная дебиторская задолженность' }
        ]
    },
    {
        title: 'Подробная аналитика заказчика',
        type: 'customer-analytics'
    }

];
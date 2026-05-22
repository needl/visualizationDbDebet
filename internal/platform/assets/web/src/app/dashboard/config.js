export const dashboardConfig = [
    {
        title: 'Общие сведения по дебиторской задолженности по состоянию на 31.03.2026',
        type: 'stats',
        metrics: [
            { key: 'count_source_org', label: 'Количество заказчиков', format: 'number' },
            { key: 'count_contracts', label: 'Количество контрактов', format: 'number' },
            { key: 'sum_contract_amount', label: 'Цена контрактов', format: 'currency' },
            { key: 'sum_debet_total', label: 'Дебиторская задолженность', format: 'currency' },
            { key: 'sum_debet_overdue', label: 'Просроченная задолженность', format: 'currency' }
        ],
        table: true,
        chart: {
            metric: 'debetByYear',
            title: 'Динамика дебиторской задолженности за 2024 – 2026 г.г.'
        }
    },
    {
        title: 'Дебиторская задолженность в разрезе заказчиков по состоянию на 31.03.2026',
        type: 'charts',
        charts: [
            { metric: 'contractAmount', title: 'Цена контрактов' },
            { metric: 'debetTotal', title: 'Дебиторская задолженность' },
            { metric: 'debetOverdose', title: 'Просроченная дебиторская задолженность' }
        ]
    },
    {
        title: 'Анализ дебиторской задолженности заказчика',
        type: 'customer-analytics'
    },
    {
        title: 'Анализ подрядчиков, выполняющих работы для нескольких заказчиков',
        type: 'contractor-analytics'
    }
];

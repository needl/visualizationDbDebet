export const TOP_MODAL_TITLES = {
    overdue: 'Топ-10 дебиторов (просроченная дебиторская задолженность)',
    debt: 'Топ-10 дебиторов (текущая дебиторская задолженность)'
};

export const CONTRACTOR_TABLE_HEADERS = {
    object: 'Объект',
    contract_date: 'Дата заключения контракта',
    work_end_date: 'Дата окончания работ',
    number: 'Номер контракта',
    amount: 'Цена контракта, млн ₽',
    debet_total: 'Дебиторская задолженность, млн ₽',
    debet_overdue: 'Просроченная задолженность, млн ₽'
};

export const OBJECT_METRIC_DEFS = [
    { label: 'Подрядчик', key: 'contractorName', format: 'string' },
    { label: 'Дата начала работ', key: 'workStartDate', format: 'date' },
    { label: 'Дата окончания работ', key: 'workEndDate', format: 'date' },
    { label: 'Строительная готовность, %', key: 'buildReadyPercent', format: 'string' },
    { label: 'Разрешение на ввод', key: 'permissionToEnter', format: 'boolean' },
    { label: 'Заключение МКЭ', key: 'conclusionMke', format: 'boolean' },
    { label: 'Твёрдая договорная цена', key: 'hardContractPrice', format: 'money' },
    { label: 'Цена контракта', key: 'contractAmount', format: 'money' },
    { label: 'Кассовые расходы', key: 'paidAmount', format: 'money' },
    { label: 'Объём принятых работ', key: 'acceptedAmount', format: 'money' }
];

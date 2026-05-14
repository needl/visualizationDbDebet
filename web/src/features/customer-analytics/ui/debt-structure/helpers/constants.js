export const TOP_MODAL_TITLES = {
    overdue: 'Топ-10 подрядчиков по просроченной дебиторской задолженности',
    debt: 'Топ-10 подрядчиков по текущей дебиторской задолженности'
};

export const CONTRACTOR_TABLE_HEADERS = {
    object: 'Объект',
    contract_date: 'Дата заключения контракта',
    work_end_date: 'Дата окончания работ',
    number: 'Номер контракта',
    amount: 'Сумма контракта, млн ₽',
    debet_total: 'Общая задолженность, млн ₽',
    debet_overdue: 'Просроченная задолженность, млн ₽'
};

export const OBJECT_METRIC_DEFS = [
    { label: 'Подрядчик', key: 'contractorName', format: 'string' },
    { label: 'Дата начала работ', key: 'workStartDate', format: 'date' },
    { label: 'Дата окончания работ', key: 'workEndDate', format: 'date' },
    { label: 'Строительная готовность', key: 'buildReadyPercent', format: 'boolean' },
    { label: 'Разрешение на ввод', key: 'permissionToEnter', format: 'boolean' },
    { label: 'Заключение МКЭ', key: 'conclusionMke', format: 'boolean' },
    { label: 'Твёрдая договорная цена', key: 'hardContractPrice', format: 'money' },
    { label: 'Сумма договора', key: 'contractAmount', format: 'money' },
    { label: 'Перечислено', key: 'paidAmount', format: 'money' },
    { label: 'Принято', key: 'acceptedAmount', format: 'money' }
];

import { appState } from '../../state/appState.js';
import { loadContractorsByFactor } from '../../state/actions/actionForContractor.js';

export class BlockFactorFilter {
    constructor(container) {
        this.container = container;
        this._unsubscribe = null;
        this._lastCustomer = null;

        this.factorMap = {
            bankrot_count: 'Банкротство',
            likvidatsiya_count: 'Ликвидация',
            nedostovernost_count: 'Недостоверность данных',
            isklyuchenie_count: 'Исключение из ЕГРЮЛ',
            inostrannye_count: 'Реестр иностранных агентов',
            eks_ter_count: 'Экстремизм/Терроризм',
            nedobrosovestn_count: 'Реестр недобросовестных поставщиков',
            admin_otvet_count: 'Административные правонарушения (19.28)',
            nam_bankrot_count: 'Намерение о банкротстве',
            blokirovka_count: 'Блокировка счетов',
            chisl_count: 'Численность менее 1'
        };

        this.columnMap = {
            bankrot_count: 'priznanie_bankrotom',
            likvidatsiya_count: 'likvidatsiya',
            nedostovernost_count: 'nedostovernost_egryul',
            isklyuchenie_count: 'isklyuchenie_egryul',
            inostrannye_count: 'inostrannye_agenty',
            eks_ter_count: 'ekstremizm_terrorizm',
            nedobrosovestn_count: 'reestr_nedobrosovestnyh_postavshchikov',
            admin_otvet_count: 'administrativnaya_otvetstvennost_19_28',
            nam_bankrot_count: 'namerenie_bankrotstvo',
            blokirovka_count: 'blokirovka_schetov',
            chisl_count: 'srednespisochnaya_chislennost_le_1'
        };

        this.render();
        this.subscribe();
    }

    subscribe() {
        this._unsubscribe = appState.subscribe(state => {
            const selectWrapper = this.container.querySelector('.select-wrapper');
            const emptyMessage = this.container.querySelector('.empty-message');

            if (state.selectedCustomer !== this._lastCustomer) {
                this._lastCustomer = state.selectedCustomer;
                const resultTable = this.container.querySelector('#block-factor-result-table');
                if (resultTable) resultTable.innerHTML = '';
            }

            if (state.customerBlockFactors) {
                if (selectWrapper) selectWrapper.style.display = '';
                if (emptyMessage) emptyMessage.style.display = 'none';
                this.updateOptions(state.customerBlockFactors);
            } else {
                if (selectWrapper) selectWrapper.style.display = 'none';
                if (emptyMessage) emptyMessage.style.display = '';
                const resultTable = this.container.querySelector('#block-factor-result-table');
                if (resultTable) resultTable.innerHTML = '';
            }
        });
    }
    render() {
        this.container.innerHTML = `
            <div class="block-factor-filter">
                <div class="select-wrapper">
                    <label for="block-factor-select">Выберите фактор:</label>
                    <select id="block-factor-select">
                        <option value="">-- Выберите --</option>
                    </select>
                </div>
                <div class="empty-message" style="display:none;">Нет данных по блок-факторам</div>
                <div id="block-factor-result-table" class="result-table"></div>
            </div>
        `;

        const selectElement = this.container.querySelector('#block-factor-select');
        selectElement.addEventListener('change', (e) => this.onFactorSelected(e.target.value));
    }

    updateOptions(blockFactors) {
        const selectElement = this.container.querySelector('#block-factor-select');
        if (!selectElement) return;

        const options = Object.entries(blockFactors)
            .filter(([_, value]) => (value ?? 0) > 0)
            .sort((a, b) => b[1] - a[1]);

        let html = '<option value="">-- Выберите --</option>';
        options.forEach(([key, value]) => {
            const displayName = this.factorMap[key] || key;
            html += `<option value="${key}">${displayName} (${value})</option>`;
        });
        selectElement.innerHTML = html;
        selectElement.value = '';
    }

    async onFactorSelected(factorKey) {
        const resultTable = this.container.querySelector('#block-factor-result-table');
        if (!factorKey) {
            if (resultTable) resultTable.innerHTML = '';
            return;
        }
        const selectedOrg = appState._getState().selectedCustomer;
        if (!selectedOrg) return;

        const columnName = this.columnMap[factorKey] || factorKey;

        if (resultTable) resultTable.innerHTML = 'Загрузка...';
        try {
            const contractors = await loadContractorsByFactor(selectedOrg, columnName);
            this.renderContractorsTable(contractors, resultTable);
        } catch (err) {
            if (resultTable) resultTable.innerHTML = `<div class="error">Ошибка: ${err.message}</div>`;
        }
    }

    renderContractorsTable(data, container) {
        if (!data || data.length === 0) {
            container.innerHTML = 'Нет данных';
            return;
        }

        const formatDate = (dateStr) => {
            if (!dateStr) return '—';
            const date = new Date(dateStr);
            if (isNaN(date.getTime())) return dateStr;
            const year = date.getFullYear();
            const month = String(date.getMonth() + 1).padStart(2, '0');
            const day = String(date.getDate()).padStart(2, '0');
            return `${year}-${month}-${day}`;
        };
        const formatNumber = (value) => {
            if (value === null || value === undefined) return '—';
            let num = typeof value === 'number' ? value : parseFloat(value);
            if (isNaN(num)) return value;
            const mln = num / 1_000_000;
            return mln.toLocaleString('ru-RU', { minimumFractionDigits: 2, maximumFractionDigits: 2 });
        };
        const headerMap = {
            name: 'Наименование подрядчика',
            object: 'Объект',
            contract_date: 'Дата заключения контрактов',
            work_end_date: 'Дата окончания работ',
            number: 'Номер контракта',
            amount: 'Сумма контракта, млн ₽',
            debet_total: 'Общая задолженность, млн ₽',
            debet_overdue: 'Просроченная задолженность, млн ₽',
            status: 'Статус'
        };

        let html = '<table class="contractor-table"><thead>';
        const first = data[0];
        Object.keys(first).forEach(key => {
            const header = headerMap[key] || key;
            html += `<th>${header}</th>`;
        });
        html += '</thead><tbody>';
        data.forEach(item => {
            html += '<tr>';
            Object.entries(item).forEach(([key, val]) => {
                let displayValue = val != null ? val : '—';
                if (key === 'contract_date' || key === 'work_end_date') {
                    displayValue = formatDate(val);
                }
                if (key === 'amount' || key === 'debet_total' || key === 'debet_overdue') {
                    displayValue = formatNumber(val);
                }
                html += `<td style="text-align: center;">${displayValue}</td>`;
            });
            html += '</tr>';
        });
        html += '</tbody></table>';
        container.innerHTML = html;
    }
}

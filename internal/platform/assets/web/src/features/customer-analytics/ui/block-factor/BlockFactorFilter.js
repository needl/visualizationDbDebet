import { appState } from '../../../../shared/state/appState.js';
import { getUserFriendlyError } from '../../../../shared/lib/userFriendlyError.js';
import { fetchContractorTable } from '../../api/contractorApi.js';
import { fetchObjectData } from '../../api/objectApi.js';
import { aggregateObjectMetrics, prepareChartData } from '../../lib/objectAggregator.js';
import { loadContractorsByFactor } from '../../model/actions/loadContractorsByFactor.js';
import { renderObjectMetricCards } from '../debt-structure/helpers/renderers.js';

export class BlockFactorFilter {
    constructor(container) {
        this.container = container;
        this._unsubscribe = null;
        this._lastCustomer = null;
        this._currentContractorName = '';

        this.contractorTableModal = null;
        this.objectModal = null;
        this.objectModalChart = null;

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
            chisl_count: 'Численность менее 1',
            all_risks: 'Все риски'
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
            chisl_count: 'srednespisochnaya_chislennost_le_1',
            all_risks: 'Все риски'
        };

        this.render();
        this.subscribe();
    }

    subscribe() {
        this._unsubscribe = appState.subscribe((state) => {
            const selectWrapper = this.container.querySelector('.select-wrapper');
            const emptyMessage = this.container.querySelector('.empty-message');

            if (state.selectedCustomer !== this._lastCustomer) {
                this._lastCustomer = state.selectedCustomer;
                const resultTable = this.container.querySelector('#block-factor-result-table');
                if (resultTable) resultTable.innerHTML = '';
                this.closeContractorTableModal();
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
                    <label for="block-factor-select">Выберите блок-фактор:</label>
                    <select id="block-factor-select">
                        <option value="">-- Блок фактор --</option>
                    </select>
                </div>
                <div class="empty-message" style="display:none;">Нет данных по блок-факторам</div>
                <div id="block-factor-result-table" class="result-table"></div>
            </div>
        `;

        const selectElement = this.container.querySelector('#block-factor-select');
        selectElement.addEventListener('change', (event) => this.onFactorSelected(event.target.value));
    }

    updateOptions(blockFactors) {
        const selectElement = this.container.querySelector('#block-factor-select');
        if (!selectElement) return;

        const options = Object.entries(blockFactors)
            .filter(([_, value]) => (value ?? 0) > 0)
            .sort((a, b) => b[1] - a[1]);

        let html = '<option value="">-- Блок фактор --</option>';
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
            const message = getUserFriendlyError(err, 'Не удалось загрузить подрядчиков по выбранному блок-фактору');
            if (resultTable) resultTable.innerHTML = `<div class="error">${this.escapeHtml(message)}</div>`;
        }
    }

    escapeHtml(value) {
        return String(value)
            .replaceAll('&', '&amp;')
            .replaceAll('<', '&lt;')
            .replaceAll('>', '&gt;')
            .replaceAll('"', '&quot;')
            .replaceAll("'", '&#39;');
    }

    renderContractorsTable(data, container) {
        if (!data || data.length === 0) {
            container.innerHTML = 'Нет данных';
            return;
        }

        const formatNumber = (value) => {
            if (value === null || value === undefined) return '—';
            const num = typeof value === 'number' ? value : parseFloat(value);
            if (Number.isNaN(num)) return value;
            const mln = num / 1_000_000;
            return mln.toLocaleString('ru-RU', { minimumFractionDigits: 2, maximumFractionDigits: 2 });
        };
        const headerMap = {
            name: 'Наименование подрядчика',
            amount: 'Цена контракта, млн ₽',
            debet_total: 'Дебиторская задолженность, млн ₽',
            debet_overdue: 'Просроченная задолженность, млн ₽'
        };

        let html = '<table class="contractor-table"><thead><tr>';
        const first = data[0];
        Object.keys(first).forEach((key) => {
            const header = headerMap[key] || key;
            html += `<th>${this.escapeHtml(header)}</th>`;
        });
        html += '</tr></thead><tbody>';

        data.forEach((item) => {
            html += '<tr>';
            Object.entries(item).forEach(([key, val]) => {
                let displayValue = val != null ? val : '—';
                if (key === 'amount' || key === 'debet_total' || key === 'debet_overdue') {
                    displayValue = formatNumber(val);
                }

                if (key === 'name' && val) {
                    const contractorName = this.escapeHtml(val);
                    html += `<td style="text-align: center;"><button type="button" class="contractor-link-btn" data-contractor="${contractorName}">${contractorName}</button></td>`;
                } else {
                    html += `<td style="text-align: center;">${this.escapeHtml(displayValue)}</td>`;
                }
            });
            html += '</tr>';
        });

        html += '</tbody></table>';
        container.innerHTML = html;

        const contractorButtons = container.querySelectorAll('.contractor-link-btn');
        contractorButtons.forEach((button) => {
            button.addEventListener('click', () => {
                const contractorName = button.dataset.contractor;
                if (!contractorName) return;
                this.showContractorTableModal(contractorName);
            });
        });
    }

    closeContractorTableModal() {
        this.closeObjectModal();
        this._currentContractorName = '';
        if (!this.contractorTableModal) return;
        document.body.removeChild(this.contractorTableModal);
        this.contractorTableModal = null;
    }

    async showContractorTableModal(contractorName) {
        this.closeContractorTableModal();
        this._currentContractorName = contractorName;

        const overlay = document.createElement('div');
        overlay.className = 'modal-overlay';
        overlay.addEventListener('click', (event) => {
            if (event.target === overlay) this.closeContractorTableModal();
        });

        const modal = document.createElement('div');
        modal.className = 'modal-content';
        modal.innerHTML = `
            <div class="modal-header">
                <h3>Контракты подрядчика: ${this.escapeHtml(contractorName)}</h3>
                <button type="button" class="modal-close" aria-label="Закрыть">&times;</button>
            </div>
            <div class="modal-body">Загрузка...</div>
        `;

        overlay.appendChild(modal);
        document.body.appendChild(overlay);
        this.contractorTableModal = overlay;

        const closeButton = modal.querySelector('.modal-close');
        closeButton?.addEventListener('click', () => this.closeContractorTableModal());

        try {
            const data = await fetchContractorTable(contractorName);
            const body = modal.querySelector('.modal-body');
            if (!body) return;
            this.renderContractorDetailsTable(body, data);
        } catch (err) {
            const message = getUserFriendlyError(err, 'Не удалось загрузить таблицу по подрядчику');
            const body = modal.querySelector('.modal-body');
            if (body) {
                body.innerHTML = `<div class="error">${this.escapeHtml(message)}</div>`;
            }
        }
    }

    renderContractorDetailsTable(container, data) {
        if (!data || data.length === 0) {
            container.innerHTML = 'Нет данных';
            return;
        }

        const columns = ['object', 'org_name', 'number', 'work_start_date', 'work_end_date', 'amount', 'debet_total', 'debet_overdue'];
        const headers = {
            object: 'Объект',
            org_name: 'Наименование организации',
            number: 'Номер контракта',
            work_start_date: 'Дата начала работ',
            work_end_date: 'Дата окончания работ',
            amount: 'Цена контракта, млн ₽',
            debet_total: 'Дебиторская задолженность, млн ₽',
            debet_overdue: 'Просроченная задолженность, млн ₽'
        };

        const formatDate = (dateStr) => {
            if (!dateStr) return '—';
            const date = new Date(dateStr);
            if (Number.isNaN(date.getTime())) return dateStr;
            const year = date.getFullYear();
            const month = String(date.getMonth() + 1).padStart(2, '0');
            const day = String(date.getDate()).padStart(2, '0');
            return `${year}-${month}-${day}`;
        };
        const formatNumber = (value) => {
            if (value === null || value === undefined) return '—';
            const num = typeof value === 'number' ? value : parseFloat(value);
            if (Number.isNaN(num)) return value;
            const mln = num / 1_000_000;
            return mln.toLocaleString('ru-RU', { minimumFractionDigits: 2, maximumFractionDigits: 2 });
        };

        let html = '<table class="contractor-table contractor-table-details"><thead><tr>';
        columns.forEach((column) => {
            const dateClass = (column === 'work_start_date' || column === 'work_end_date') ? ' class="date-col"' : '';
            html += `<th${dateClass}>${this.escapeHtml(headers[column] || column)}</th>`;
        });
        html += '</tr></thead><tbody>';

        data.forEach((row) => {
            html += '<tr>';
            columns.forEach((column) => {
                const value = row[column];
                let displayValue = value ?? '—';

                if (column === 'work_start_date' || column === 'work_end_date') {
                    displayValue = formatDate(value);
                }
                if (column === 'amount' || column === 'debet_total' || column === 'debet_overdue') {
                    displayValue = formatNumber(value);
                }

                const dateClass = (column === 'work_start_date' || column === 'work_end_date') ? ' class="date-col"' : '';
                if (column === 'object' && value) {
                    const safeObject = this.escapeHtml(value);
                    html += `<td${dateClass}><button type="button" class="object-link-btn" data-object="${safeObject}">${safeObject}</button></td>`;
                } else {
                    html += `<td${dateClass}>${this.escapeHtml(displayValue)}</td>`;
                }
            });
            html += '</tr>';
        });

        html += '</tbody></table>';
        container.innerHTML = html;

        const objectButtons = container.querySelectorAll('.object-link-btn');
        objectButtons.forEach((button) => {
            button.addEventListener('click', () => {
                const objectName = button.dataset.object;
                if (!objectName) return;
                this.showObjectModal(objectName, this._currentContractorName || '');
            });
        });
    }

    closeObjectModal() {
        if (this.objectModalChart) {
            this.objectModalChart.dispose();
            this.objectModalChart = null;
        }
        if (this.objectModal) {
            document.body.removeChild(this.objectModal);
            this.objectModal = null;
        }
    }

    async showObjectModal(objectName, counterpartyName = '') {
        this.closeObjectModal();

        const overlay = document.createElement('div');
        overlay.className = 'modal-overlay';
        overlay.addEventListener('click', (event) => {
            if (event.target === overlay) this.closeObjectModal();
        });

        const modal = document.createElement('div');
        modal.className = 'modal-content';

        const header = document.createElement('div');
        header.className = 'modal-header';
        const title = document.createElement('h3');
        title.textContent = `Объект: ${objectName}`;

        const closeButton = document.createElement('button');
        closeButton.type = 'button';
        closeButton.className = 'modal-close';
        closeButton.innerHTML = '&times;';
        closeButton.addEventListener('click', () => this.closeObjectModal());

        header.appendChild(title);
        header.appendChild(closeButton);

        const body = document.createElement('div');
        body.className = 'modal-body object-modal-body';
        body.textContent = 'Загрузка...';

        modal.appendChild(header);
        modal.appendChild(body);
        overlay.appendChild(modal);
        document.body.appendChild(overlay);
        this.objectModal = overlay;

        try {
            const rawData = await fetchObjectData(objectName, counterpartyName);
            this.renderObjectAnalytics(body, rawData || []);
        } catch (err) {
            const message = getUserFriendlyError(err, 'Не удалось загрузить аналитику по объекту');
            body.innerHTML = `<div class="error">${this.escapeHtml(message)}</div>`;
        }
    }

    renderObjectAnalytics(container, rawData) {
        const metrics = aggregateObjectMetrics(rawData);
        const chartData = prepareChartData(rawData);
        if (!metrics) {
            container.innerHTML = '<div class="empty-message">Нет данных по объекту</div>';
            return;
        }

        container.innerHTML = '';

        const root = document.createElement('div');
        root.className = 'object-analytics';

        const cardsGrid = document.createElement('div');
        cardsGrid.className = 'object-metrics-grid';
        root.appendChild(cardsGrid);

        const chartContainer = document.createElement('div');
        chartContainer.className = 'analytics-chart object-modal-chart';
        root.appendChild(chartContainer);

        container.appendChild(root);
        renderObjectMetricCards(cardsGrid, metrics);

        if (this.objectModalChart) {
            this.objectModalChart.dispose();
        }
        this.objectModalChart = echarts.init(chartContainer);

        const option = {
            title: { text: 'График задолженности по периодам', left: 'center' },
            tooltip: { trigger: 'axis' },
            legend: {
                data: chartData.series.map((series) => series.name),
                bottom: 0
            },
            xAxis: { type: 'category', data: chartData.categories },
            yAxis: {
                type: 'value',
                axisLabel: {
                    formatter: (value) => `${(value / 1_000_000).toFixed(1)} млн`
                }
            },
            series: chartData.series
        };

        this.objectModalChart.setOption(option, true);
        this.objectModalChart.resize();
    }
}

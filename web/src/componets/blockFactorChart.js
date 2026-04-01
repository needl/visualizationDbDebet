// src/components/blockFactorChart.js
import { appState } from '../state/appState.js';
import {loadContractorsByFactor} from '../state/actions/actionForContractor.js'

export class BlockFactorsChart {
    constructor(container, title = 'Блок-факторы по количеству подрядчиков') {
        this.container = container;
        this.title = title;
        this.chart = null;
        this.modal = null;
        this.createModal();
    }

    createModal() {
        if (document.getElementById('blockFactorModal')) return;
        const modal = document.createElement('div');
        modal.id = 'blockFactorModal';
        modal.className = 'modal';
        modal.style.display = 'none';
        modal.innerHTML = `
            <div class="modal-content">
                <span class="close">&times;</span>
                <h3>Список объектов</h3>
                <div class="modal-body">Загрузка...</div>
            </div>
        `;
        document.body.appendChild(modal);
        this.modal = modal;

        const closeBtn = modal.querySelector('.close');
        closeBtn.addEventListener('click', () => this.hideModal());

        window.addEventListener('click', (e) => {
            if (e.target === modal) this.hideModal();
        });
    }

    showModal(content) {
        if (this.modal) {
            const body = this.modal.querySelector('.modal-body');
            body.innerHTML = content;
            this.modal.style.display = 'block';
        }
    }

    hideModal() {
        if (this.modal) this.modal.style.display = 'none';
    }

    render(blockFactors) {
        if (!blockFactors) {
            this.clear();
            this.container.innerHTML = '<div class="empty-message">Нет данных по блок-факторам</div>';
            return;
        }

        const data = Object.entries(blockFactors)
            .map(([name, value]) => ({ name, value: value ?? 0 }))
            .filter(item => item.value > 0);

        if (data.length === 0) {
            this.clear();
            this.container.innerHTML = '<div class="empty-message">Нет данных по блок-факторам</div>';
            return;
        }

        // Сортировка по убыванию
        data.sort((a, b) => b.value - a.value);

        if (!this.chart) {
            this.chart = echarts.init(this.container);
        }

        const names = data.map(item => this.formatName(item.name));
        const values = data.map(item => item.value);
        const originalKeys = data.map(item => item.name);

        const option = {
            title: { text: this.title, left: 'center' },
            tooltip: {
                trigger: 'axis',
                axisPointer: { type: 'shadow' }
            },
            grid: {
                containLabel: true,
                left: '18%',
                right: '5%',
                top: '15%',
                bottom: '5%'
            },
            xAxis: {
                type: 'value',
                name: 'Количество',
                nameLocation: 'middle',
                nameGap: 30,
                minInterval: 1,
                axisLabel: {
                    formatter: (value) => Math.floor(value) === value ? value : Math.floor(value),
                    fontSize: 10
                }
            },
            yAxis: {
                type: 'category',
                data: names,
                axisLabel: {
                    fontSize: 10,
                    width: 150,
                    overflow: 'break'
                }
            },
            series: [{
                name: 'Количество',
                type: 'bar',
                data: values,
                itemStyle: { color: '#d11b1b' },
                label: {
                    show: true,
                    position: 'right',
                    formatter: (p) => p.value
                }
            }]
        };

        this.chart.setOption(option, true);

        // Удаляем старые слушатели, чтобы не накапливались
        this.chart.off('click');
        this.chart.on('click', async (params) => {
            if (params.componentType === 'series' && params.dataIndex !== undefined) {
                const index = params.dataIndex;
                const originalKey = originalKeys[index];
                const selectedOrg = appState._getState().selectedCustomer;
                if (!selectedOrg) return;

                // Маппинг ключей в имена полей для бэкенда
                const columnMap = {
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
                const columnName = columnMap[originalKey] || originalKey;

                this.showModal('Загрузка...');
                try {
                    const contractors = await loadContractorsByFactor(selectedOrg, columnName);
                    this.renderContractorsTable(contractors);
                } catch (err) {
                    this.showModal(`Ошибка: ${err.message}`);
                }
            }
        });

        this.chart.resize();
    }

    renderContractorsTable(data) {
        if (!data || data.length === 0) {
            this.showModal('Нет данных');
            return;
        }

        const formatDate = (dateStr) => {
            if (!dateStr) return '—';
            const date = new Date(dateStr);
            if (isNaN(date.getTime())) return dateStr; // если не дата – оставляем как есть
            const year = date.getFullYear();
            const month = String(date.getMonth() + 1).padStart(2, '0');
            const day = String(date.getDate()).padStart(2, '0');
            return `${year}-${month}-${day}`;
        };

        const formatNumber = (value) => {
            if (value === null || value === undefined) return '—';
            let num = typeof value === 'number' ? value : parseFloat(value);
            if (isNaN(num)) return value;
            // Форматируем с пробелами (например, 1234567.89 -> 1 234 567.89)
            return num.toLocaleString('ru-RU');
        };

        const headerMap = {
            name: 'Наименование подрядчика',
            object: 'Объект',
            contract_date: 'Дата заключения контракта',
            work_end_date: 'Дата окончания работ',
            number: 'Номер контракта',
            amount: 'Сумма контракта',
            debet_total: 'Общая задолженность',
            debet_overdue: 'Просроченная задолженность',
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
                let displayValue = val !== null && val !== undefined ? val : '—';
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
        this.showModal(html);
    }

    clear() {
        if (this.chart) {
            this.chart.dispose();
            this.chart = null;
        }
        this.container.innerHTML = '';
    }

    formatName(key) {
        const map = {
            bankrot_count: 'Банкротство',
            likvidatsiya_count: 'Ликвидация',
            nedostovernost_count: 'Недостоверность данных',
            isklyuchenie_count: 'Исключение из ЕГРЮЛ',
            inostrannye_count: 'Реестр иностранных агентов',
            eks_ter_count: 'Экстремизм Терроризм',
            nedobrosovestn_count: 'Реестр недобросовестных поставщиков',
            admin_otvet_count: 'Административные правонарушения',
            nam_bankrot_count: 'Намерение о банкротство',
            blokirovka_count: 'Блокировка счётов',
            chisl_count: 'Численность меньше 1'
        };
        return map[key] || key;
    }
}
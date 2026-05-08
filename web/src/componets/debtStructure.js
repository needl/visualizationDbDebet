// src/components/debtStructure.js
import { appState } from '../state/appState.js';
import { HorizontalBarChart } from './chart/horizonBarChart.js';
import { fetchContractorDebt, fetchContractorOverdue } from '../services/customerApiCaller.js';

export class DebtStructure {
    constructor(container) {
        this.container = container;
        this.chart = null;

        this.activeModal = null;
        this.modalChart = null;
        this.modalUnsub = null;
        this.contractorModal = null; // для нового окна

        this.lastState = null;
        this.stateUnsub = appState.subscribe((state) => {
            this.lastState = state;
        });
    }

    render(summary) {
        if (!summary || summary.total_debet === undefined) {
            if (this.chart) {
                this.chart.dispose();
                this.chart = null;
            }
            this.container.innerHTML = '<div class="empty-message">Нет данных по задолженности</div>';
            return;
        }

        const totalDebt = summary.total_debet;
        const overdueDebt = summary.total_debet_overdue || 0;
        const currentDebt = totalDebt - overdueDebt;

        if (totalDebt === 0 && overdueDebt === 0) {
            if (this.chart) {
                this.chart.dispose();
                this.chart = null;
            }
            this.container.innerHTML = '<div class="empty-message">Задолженность отсутствует</div>';
            return;
        }

        if (!this.chart) {
            this.chart = echarts.init(this.container);
        }

        const data = [
            { name: 'Текущая', value: currentDebt, itemStyle: { color: '#10b981' } },
            { name: 'Просроченная', value: overdueDebt, itemStyle: { color: '#ef4444' } },
        ].filter(item => item.value > 0);

        if (data.length === 0) {
            if (this.chart) {
                this.chart.dispose();
                this.chart = null;
            }
            this.container.innerHTML = '<div class="empty-message">Задолженность отсутствует</div>';
            return;
        }

        const option = {
            title: {
                text: 'Структура дебиторской задолженности',
                left: 'center',
                top: 0
            },
            tooltip: {
                trigger: 'item',
                formatter: (params) => {
                    const value = params.value;
                    const formatted = new Intl.NumberFormat('ru-RU', {
                        style: 'currency',
                        currency: 'RUB',
                        minimumFractionDigits: 0,
                        maximumFractionDigits: 0
                    }).format(value);
                    return `${params.name}: ${formatted} (${params.percent}%)`;
                }
            },
            series: [
                {
                    name: 'Дебиторская задолженность',
                    type: 'pie',
                    radius: '80%',
                    avoidLabelOverlap: false,
                    label: {
                        show: true,
                        position: 'outside',
                        formatter: (params) => {
                            const percent = params.percent.toFixed(1);
                            const nameLines = params.name.split(' ').join('\n');
                            return `${nameLines}\n${percent}%`;
                        }
                    },
                    labelLayout: (params) => {
                        if (params.text.includes('Просроченная')) {
                            return {
                                x: params.labelRect.x,
                                y: params.labelRect.y + 30,
                                align: 'centre'
                            };
                        } else {
                            return {
                                x: params.labelRect.x,
                                y: params.labelRect.y,
                                align: 'centre'
                            };
                        }
                    },
                    labelLine: {
                        length: 15,
                        length2: 10,
                        smooth: true
                    },
                    emphasis: {
                        scale: true
                    },
                    data: data
                }
            ]
        };

        this.chart.setOption(option, true);
        this.chart.resize();

        this.chart.off('click');
        this.chart.on('click', (params) => {
            const name = params.name;
            if (name === 'Просроченная') {
                this.showTopModal('overdue');
            } else if (name === 'Текущая') {
                this.showTopModal('debt');
            }
        });
    }

    showTopModal(type) {
        console.log('lastState.customerTopDebtors:', this.lastState.customerTopDebtors);
        console.log('lastState.customerTopOverdue:', this.lastState.customerTopOverdue);


        if (!this.lastState) {
            console.warn('Нет данных для открытия топа подрядчиков');
            return;
        }
        this.closeTopModal();

        const overlay = document.createElement('div');
        overlay.className = 'modal-overlay';
        overlay.addEventListener('click', (e) => {
            if (e.target === overlay) this.closeTopModal();
        });

        const modal = document.createElement('div');
        modal.className = 'modal-content';

        const header = document.createElement('div');
        header.className = 'modal-header';
        const title = document.createElement('h3');
        title.textContent = type === 'overdue'
            ? 'Топ-10 подрядчиков по просроченной дебиторской задолженности'
            : 'Топ-10 подрядчиков по текущей дебиторской задолженности';
        const closeBtn = document.createElement('button');
        closeBtn.className = 'modal-close';
        closeBtn.innerHTML = '&times;';
        closeBtn.addEventListener('click', () => this.closeTopModal());
        header.appendChild(title);
        header.appendChild(closeBtn);

        const chartContainer = document.createElement('div');
        chartContainer.className = 'modal-chart-container';
        chartContainer.style.height = '450px';

        modal.appendChild(header);
        modal.appendChild(chartContainer);
        overlay.appendChild(modal);
        document.body.appendChild(overlay);

        // Коллбэк, вызываемый при клике на столбец
        const onBarClick = (contractorName, value) => {
            const orgName = this.lastState.selectedCustomer;
            if (!orgName) return;
            this.showContractorModal(contractorName, type, orgName);
        };

        const barChart = new HorizontalBarChart(
            chartContainer,
            title.textContent,
            (v) => (v / 1e9).toFixed(2) + ' млрд ₽',
            onBarClick
        );

        const updateChart = (state) => {
            const data = type === 'overdue'
                ? state.customerTopOverdue
                : state.customerTopDebtors;
            barChart.render(data || []);
        };

        updateChart(this.lastState);
        const unsub = appState.subscribe(updateChart);

        this.activeModal = overlay;
        this.modalChart = barChart;
        this.modalUnsub = unsub;
    }

    closeTopModal() {
        if (this.activeModal) {
            document.body.removeChild(this.activeModal);
            this.activeModal = null;
        }
        if (this.modalUnsub) {
            this.modalUnsub();
            this.modalUnsub = null;
        }
        if (this.modalChart) {
            this.modalChart = null;
        }
    }

    showContractorModal(contractorName, type, orgName) {
        console.log('showContractorModal called with:', contractorName, type, orgName);
        this.closeContractorModal();

        const overlay = document.createElement('div');
        overlay.className = 'modal-overlay';
        overlay.addEventListener('click', (e) => {
            if (e.target === overlay) this.closeContractorModal();
        });

        const modal = document.createElement('div');
        modal.className = 'modal-content';
        modal.innerHTML = `
            <span class="close">&times;</span>
            <h3>Контракты подрядчика: ${contractorName}</h3>
            <div class="modal-body">Загрузка...</div>
        `;

        overlay.appendChild(modal);
        document.body.appendChild(overlay);
        this.contractorModal = overlay;

        const closeBtn = modal.querySelector('.close');
        closeBtn.addEventListener('click', () => this.closeContractorModal());

        const fetchFn = type === 'overdue' ? fetchContractorOverdue : fetchContractorDebt;
        fetchFn(orgName, contractorName)
            .then(data => this.renderContractorTable(data, modal.querySelector('.modal-body')))
            .catch(err => {
                modal.querySelector('.modal-body').innerHTML = `<div class="error">Ошибка: ${err.message}</div>`;
            });
    }

    closeContractorModal() {
        if (this.contractorModal) {
            document.body.removeChild(this.contractorModal);
            this.contractorModal = null;
        }
    }

    renderContractorTable(data, container) {
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

        const formatMoney = (value) => {
            if (value === null || value === undefined) return '—';
            const num = Number(value);
            if (isNaN(num)) return '—';
            const mln = num / 1_000_000;
            return mln.toLocaleString('ru-RU', { minimumFractionDigits: 2, maximumFractionDigits: 2 });
        };

        const headers = {
            object: 'Объект',
            contract_date: 'Дата заключения контракта',
            work_end_date: 'Дата окончания работ',
            number: 'Номер контракта',
            amount: 'Сумма контракта, млн ₽',
            debet_total: 'Общая задолженность, млн ₽',
            debet_overdue: 'Просроченная задолженность, млн ₽'
        };

        let html = '<table class="contractor-table"><thead><tr>';
        Object.values(headers).forEach(h => html += `<th>${h}</th>`);
        html += '</tr></thead><tbody>';

        data.forEach(item => {
            html += '<tr>';
            html += `<td>${item.object || '—'}</td>`;
            html += `<td>${formatDate(item.contract_date)}</td>`;
            html += `<td>${formatDate(item.work_end_date)}</td>`;
            html += `<td>${item.number || '—'}</td>`;
            html += `<td>${formatMoney(item.amount)}</td>`;
            html += `<td>${formatMoney(item.debet_total)}</td>`;
            html += `<td>${formatMoney(item.debet_overdue)}</td>`;
            html += '</tr>';
        });

        html += '</tbody></table>';
        container.innerHTML = html;
    }

    clear() {
        this.closeTopModal();
        this.closeContractorModal();
        if (this.stateUnsub) {
            this.stateUnsub();
            this.stateUnsub = null;
        }
        if (this.chart) {
            this.chart.dispose();
            this.chart = null;
        }
        this.container.innerHTML = '';
    }
}
// src/components/debtStructure.js
import { appState } from '../state/appState.js';
import { HorizontalBarChart } from './horizonBarChart.js';

export class DebtStructure {
    constructor(container) {
        this.container = container;
        this.chart = null;

        // Модальное окно
        this.activeModal = null;
        this.modalChart = null;
        this.modalUnsub = null;

        // Сохраняем последнее состояние, чтобы избежать вызова несуществующего getState()
        this.lastState = null;
        this.stateUnsub = appState.subscribe((state) => {
            this.lastState = state;
        });
    }

    render(summary) {
        if (!summary || summary.total_debet === undefined) {
            this.container.innerHTML = '<div class="empty-message">Нет данных по задолженности</div>';
            return;
        }

        const totalDebt = summary.total_debet;
        const overdueDebt = summary.total_debet_overdue || 0;
        const currentDebt = totalDebt - overdueDebt;

        if (totalDebt === 0 && overdueDebt === 0) {
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

        const barChart = new HorizontalBarChart(chartContainer, title.textContent);

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

    clear() {
        this.closeTopModal();
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
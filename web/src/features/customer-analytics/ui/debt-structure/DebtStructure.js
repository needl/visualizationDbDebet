import { appState } from '../../../../shared/state/appState.js';
import { HorizontalBarChart } from './HorizontalBarChart.js';
import { fetchContractorDebt, fetchContractorOverdue } from '../../api/contractorApi.js';
import { fetchObjectData } from '../../api/objectApi.js';
import { getUserFriendlyError } from '../../../../shared/lib/userFriendlyError.js';
import { aggregateObjectMetrics, prepareChartData } from '../../lib/objectAggregator.js';
import { TOP_MODAL_TITLES } from './helpers/constants.js';
import { escapeHtml } from './helpers/formatters.js';
import { renderContractorTable, renderObjectMetricCards } from './helpers/renderers.js';

export class DebtStructure {
    constructor(container) {
        this.container = container;
        this.chart = null;

        this.activeModal = null;
        this.modalChart = null;
        this.modalUnsub = null;
        this.contractorModal = null;
        this.objectModal = null;
        this.objectModalChart = null;

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
            { name: 'Просроченная', value: overdueDebt, itemStyle: { color: '#ef4444' } }
        ].filter((item) => item.value > 0);

        if (data.length === 0) {
            if (this.chart) {
                this.chart.dispose();
                this.chart = null;
            }
            this.container.innerHTML = '<div class="empty-message">Задолженность отсутствует</div>';
            return;
        }

        const option = {
            title: { show: false },
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
                        }
                        return {
                            x: params.labelRect.x,
                            y: params.labelRect.y,
                            align: 'centre'
                        };
                    },
                    labelLine: {
                        length: 15,
                        length2: 10,
                        smooth: true
                    },
                    emphasis: { scale: true },
                    data
                }
            ]
        };

        this.chart.setOption(option, true);
        this.chart.resize();

        this.chart.off('click');
        this.chart.on('click', (params) => {
            if (params.name === 'Просроченная') {
                this.showTopModal('overdue');
            } else if (params.name === 'Текущая') {
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
        title.textContent = TOP_MODAL_TITLES[type] || TOP_MODAL_TITLES.debt;

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

        const onBarClick = (contractorName) => {
            const orgName = this.lastState.selectedCustomer;
            if (!orgName) return;
            this.showContractorModal(contractorName, type, orgName);
        };

        const barChart = new HorizontalBarChart(
            chartContainer,
            title.textContent,
            (v) => `${(v / 1e9).toFixed(2)} млрд ₽`,
            onBarClick
        );

        const updateChart = (state) => {
            const data = type === 'overdue' ? state.customerTopOverdue : state.customerTopDebtors;
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
            .then((data) => {
                const body = modal.querySelector('.modal-body');
                renderContractorTable(body, data, (objectName) => this.showObjectModal(objectName));
            })
            .catch((err) => {
                const message = getUserFriendlyError(err, 'Не удалось загрузить контракты подрядчика');
                modal.querySelector('.modal-body').innerHTML = `<div class="error">${escapeHtml(message)}</div>`;
            });
    }

    closeContractorModal() {
        this.closeObjectModal();
        if (this.contractorModal) {
            document.body.removeChild(this.contractorModal);
            this.contractorModal = null;
        }
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

    async showObjectModal(objectName) {
        const orgName = this.lastState?.selectedCustomer;
        if (!orgName) return;

        this.closeObjectModal();

        const overlay = document.createElement('div');
        overlay.className = 'modal-overlay';
        overlay.addEventListener('click', (e) => {
            if (e.target === overlay) this.closeObjectModal();
        });

        const modal = document.createElement('div');
        modal.className = 'modal-content';

        const header = document.createElement('div');
        header.className = 'modal-header';
        const title = document.createElement('h3');
        title.textContent = `Аналитика объекта: ${objectName}`;

        const closeBtn = document.createElement('button');
        closeBtn.className = 'modal-close';
        closeBtn.innerHTML = '&times;';
        closeBtn.addEventListener('click', () => this.closeObjectModal());

        header.appendChild(title);
        header.appendChild(closeBtn);

        const body = document.createElement('div');
        body.className = 'modal-body object-modal-body';
        body.textContent = 'Загрузка...';

        modal.appendChild(header);
        modal.appendChild(body);
        overlay.appendChild(modal);
        document.body.appendChild(overlay);
        this.objectModal = overlay;

        try {
            const rawData = await fetchObjectData(orgName, objectName);
            this.renderObjectAnalytics(body, rawData || []);
        } catch (err) {
            const message = getUserFriendlyError(err, 'Не удалось загрузить аналитику по объекту');
            body.innerHTML = `<div class="error">${escapeHtml(message)}</div>`;
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
            title: { text: 'График задолженности по годам', left: 'center' },
            tooltip: { trigger: 'axis' },
            legend: { data: chartData.series.map((series) => series.name), bottom: 0 },
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

    clear() {
        this.closeTopModal();
        this.closeContractorModal();
        this.closeObjectModal();
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

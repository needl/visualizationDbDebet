import { ChartComponent } from '../../features/debet-overview/ui/ChartComponent.js';
import { PieChartComponent } from '../../features/debet-overview/ui/PieChartComponent.js';
import { MetricCard } from '../../features/debet-overview/ui/MetricCard.js';
import { StatsTable } from '../../features/debet-overview/ui/StatsTable.js';
import { appState } from '../../shared/state/appState.js';
import { CustomerFilter } from '../../features/customer-analytics/ui/CustomerFilter.js';
import { CustomerCard } from '../../features/customer-analytics/ui/CustomerCard.js';
import { BlockFactorFilter } from '../../features/customer-analytics/ui/block-factor/BlockFactorFilter.js';
import { DebtStructure } from '../../features/customer-analytics/ui/debt-structure/DebtStructure.js';
import { ContractorAnalytics } from '../../features/contractor-analytics/ui/ContractorAnalytics.js';

export class DashboardRenderer {
    constructor(config) {
        this.config = config;
        this.components = [];
        this.statsSubscriptions = [];
    }

    build(container) {
        if (!container) return;
        container.innerHTML = '';

        this.statsSubscriptions.forEach((unsub) => unsub());
        this.statsSubscriptions = [];
        this.components = [];

        this.config.forEach((block) => {
            const blockDiv = document.createElement('div');
            blockDiv.className = 'dashboard-block';
            if (block.type === 'contractor-analytics') {
                blockDiv.classList.add('contractor-analytics-block');
            }

            const titleElem = document.createElement('h2');
            titleElem.textContent = block.title;
            blockDiv.appendChild(titleElem);

            if (block.type === 'charts') {
                const chartsContainer = document.createElement('div');
                chartsContainer.className = 'charts-grid';

                block.charts.forEach((chartConfig) => {
                    const wrapperDiv = document.createElement('div');
                    wrapperDiv.className = 'chart-wrapper';

                    const pieHeader = document.createElement('h3');
                    pieHeader.textContent = chartConfig.title;
                    pieHeader.className = 'pie-chart-title';
                    wrapperDiv.appendChild(pieHeader);

                    const pieTotal = document.createElement('div');
                    pieTotal.className = 'pie-chart-total';
                    pieTotal.textContent = 'Сумма всего: —';
                    wrapperDiv.appendChild(pieTotal);

                    const pieWrapper = document.createElement('div');
                    pieWrapper.className = 'pie-wrapper';

                    const pieContainer = document.createElement('div');
                    pieContainer.style.width = '100%';
                    pieContainer.style.height = '350px';
                    pieWrapper.appendChild(pieContainer);
                    wrapperDiv.appendChild(pieWrapper);

                    const pieMetricKey = chartConfig.metric + 'Pie';
                    const pieChart = new PieChartComponent(pieContainer, pieMetricKey, chartConfig.title);
                    this.components.push(pieChart);

                    const unsubscribePieTotal = appState.subscribe((state) => {
                        const pieData = state.chartData[pieMetricKey]?.data;
                        if (!pieData || !pieData.length) {
                            pieTotal.textContent = 'Сумма всего: —';
                            return;
                        }

                        const total = pieData.reduce((sum, item) => sum + (Number(item.value) || 0), 0);
                        const totalInBillions = total / 1_000_000_000;
                        const formatted = totalInBillions.toLocaleString('ru-RU', {
                            minimumFractionDigits: 2,
                            maximumFractionDigits: 2
                        });
                        pieTotal.textContent = `Сумма всего: ${formatted} млрд ₽`;
                    });
                    this.statsSubscriptions.push(unsubscribePieTotal);

                    chartsContainer.appendChild(wrapperDiv);
                });

                blockDiv.appendChild(chartsContainer);
            } else if (block.type === 'customer-analytics') {
                const filterContainer = document.createElement('div');
                filterContainer.className = 'customer-filter-container';
                blockDiv.appendChild(filterContainer);
                const filter = new CustomerFilter(filterContainer);
                this.components.push(filter);

                const metricsContainer = document.createElement('div');
                metricsContainer.className = 'customer-metrics';
                blockDiv.appendChild(metricsContainer);
                const metrics = new CustomerCard(metricsContainer);
                metrics.mount();
                this.components.push(metrics);

                const row1 = document.createElement('div');
                row1.className = 'stats-row';

                const debtStructCol = document.createElement('div');
                debtStructCol.className = 'stats-col debt-col';

                const debtStructTitle = document.createElement('h3');
                debtStructTitle.textContent = 'Структура дебиторской задолженности на 31.03.2026';
                debtStructTitle.className = 'pie-chart-title';
                debtStructCol.appendChild(debtStructTitle);

                const debtStructureContainer = document.createElement('div');
                debtStructureContainer.className = 'debt-structure-container';
                debtStructCol.appendChild(debtStructureContainer);
                row1.appendChild(debtStructCol);

                const blockFactorsCol = document.createElement('div');
                blockFactorsCol.className = 'stats-col factors-col';

                const blockFactorsTitle = document.createElement('h3');
                blockFactorsTitle.textContent = 'Оценка состояния благонадёжности подрядчика';
                blockFactorsTitle.className = 'pie-chart-title';
                blockFactorsCol.appendChild(blockFactorsTitle);

                const blockFactorsContainer = document.createElement('div');
                blockFactorsContainer.className = 'block-factors-chart-container';
                blockFactorsCol.appendChild(blockFactorsContainer);
                row1.appendChild(blockFactorsCol);
                blockDiv.appendChild(row1);

                const debtStructure = new DebtStructure(debtStructureContainer);
                const blockFactorFilter = new BlockFactorFilter(blockFactorsContainer);
                this.components.push(debtStructure, blockFactorFilter);

                const debtUnsub = appState.subscribe((state) => {
                    debtStructure.render(state.customerSummary);
                });
                this.statsSubscriptions.push(debtUnsub);
            } else if (block.type === 'contractor-analytics') {
                const contractorContainer = document.createElement('div');
                contractorContainer.className = 'contractor-analytics-container';
                blockDiv.appendChild(contractorContainer);

                const contractorAnalytics = new ContractorAnalytics(contractorContainer);
                this.components.push(contractorAnalytics);
            } else if (block.type === 'stats') {
                const statsContainer = document.createElement('div');
                statsContainer.className = 'stats-grid';
                const cards = [];

                block.metrics.forEach((metric) => {
                    const cardContainer = document.createElement('div');
                    cardContainer.className = 'metric-card-wrapper';
                    statsContainer.appendChild(cardContainer);
                    const card = new MetricCard(cardContainer, metric.label, metric.key, metric.format);
                    cards.push(card);
                    this.components.push(card);
                });
                blockDiv.appendChild(statsContainer);

                let tableContainer = null;
                let chartContainer = null;

                if (block.table) {
                    tableContainer = document.createElement('div');
                    tableContainer.className = 'stats-table-container';
                    const tableComponent = new StatsTable(tableContainer);
                    this.components.push(tableComponent);
                    const unsubscribe = appState.subscribe((state) => {
                        if (state.tableData) {
                            tableComponent.render(state.tableData);
                        }
                    });
                    this.statsSubscriptions.push(unsubscribe);
                }

                if (block.chart) {
                    chartContainer = document.createElement('div');
                    chartContainer.className = 'chart-wrapper-full';
                    chartContainer.style.width = '100%';
                    chartContainer.style.height = '700px';
                    const chart = new ChartComponent(chartContainer, block.chart.metric, block.chart.title);
                    this.components.push(chart);
                }

                if (tableContainer) blockDiv.appendChild(tableContainer);
                if (chartContainer) blockDiv.appendChild(chartContainer);

                const unsubscribeStats = appState.subscribe((state) => {
                    if (state.stats) {
                        cards.forEach((card) => card.update(state.stats));
                    }
                });
                this.statsSubscriptions.push(unsubscribeStats);
            }

            container.appendChild(blockDiv);
        });
    }
}

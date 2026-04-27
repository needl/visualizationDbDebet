// src/dashboardRenderer.js
import {ChartComponent} from "./componets/chartForDeb.js";
import {PieChartComponent} from "./componets/pieChartForDeb.js";
import {MetricCard} from "./componets/metricCard.js";
import {StatsTable} from "./componets/statsTableForDebet.js";
import {appState} from "./state/appState.js";
import {CustomerFilter} from "./componets/customerFilter.js";
import {CustomerCard} from "./componets/customerCard.js";
import {BlockFactorsChart} from "./componets/blockFactorChart.js";
import {DebtStructure} from "./componets/debtStructure.js";

export class DashboardRenderer {
    constructor(config) {
        this.config = config;
        this.components = [];
        this.statsSubscriptions = [];
    }

    build(container) {
        if (!container) return;
        container.innerHTML = '';

        // Очищаем старые подписки (если пересобираем)
        this.statsSubscriptions.forEach(unsub => unsub());
        this.statsSubscriptions = [];
        this.components = [];

        this.config.forEach(block => {
            const blockDiv = document.createElement('div');
            blockDiv.className = 'dashboard-block';

            const titleElem = document.createElement('h2');
            titleElem.textContent = block.title;
            blockDiv.appendChild(titleElem);

            if (block.type === 'charts') {
                const chartsContainer = document.createElement('div');
                chartsContainer.className = 'charts-grid';

                block.charts.forEach(chartConfig => {
                    const wrapperDiv = document.createElement('div');
                    wrapperDiv.className = 'chart-wrapper';

                    const pieHeader = document.createElement('h3');
                    pieHeader.textContent = chartConfig.title;
                    pieHeader.className = 'pie-chart-title';
                    wrapperDiv.appendChild(pieHeader);

                    // Контейнер для круговой диаграммы с фоном
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

                    chartsContainer.appendChild(wrapperDiv);
                });

                blockDiv.appendChild(chartsContainer);
            }
            else if (block.type === 'customer-analytics') {
                // Фильтр
                const filterContainer = document.createElement('div');
                filterContainer.className = 'customer-filter-container';
                blockDiv.appendChild(filterContainer);
                const filter = new CustomerFilter(filterContainer);
                this.components.push(filter);

                // Карточки
                const metricsContainer = document.createElement('div');
                metricsContainer.className = 'customer-metrics';
                blockDiv.appendChild(metricsContainer);
                const metrics = new CustomerCard(metricsContainer);
                metrics.mount();
                this.components.push(metrics);

                // Ряд 1: структура ДЗ и блок-факторы (две колонки)
                const row1 = document.createElement('div');
                row1.className = 'stats-row';

                // Левая колонка: структура ДЗ (полукруг) – 30%
                const debtStructCol = document.createElement('div');
                debtStructCol.className = 'stats-col debt-col';
                const debtStructureContainer = document.createElement('div');
                debtStructureContainer.className = 'debt-structure-container';
                debtStructCol.appendChild(debtStructureContainer);
                row1.appendChild(debtStructCol);

                // Правая колонка: блок-факторы – 70%
                const blockFactorsCol = document.createElement('div');
                blockFactorsCol.className = 'stats-col factors-col';
                const blockFactorsContainer = document.createElement('div');
                blockFactorsContainer.className = 'block-factors-chart-container';
                blockFactorsCol.appendChild(blockFactorsContainer);
                row1.appendChild(blockFactorsCol);

                blockDiv.appendChild(row1);

                // Инициализация компонентов внутри колонок
                const debtStructure = new DebtStructure(debtStructureContainer);
                const blockFactorsChart = new BlockFactorsChart(blockFactorsContainer, 'Блок-факторы по количеству подрядчиков');
                this.components.push(debtStructure, blockFactorsChart);

                // Подписки на обновление
                const debtUnsub = appState.subscribe(state => {
                    debtStructure.render(state.customerSummary);
                });
                const blockUnsub = appState.subscribe(state => {
                    blockFactorsChart.render(state.customerBlockFactors);
                });
                this.statsSubscriptions.push(debtUnsub, blockUnsub);
            }
            else if (block.type === 'stats') {
                // Карточки статистики
                const statsContainer = document.createElement('div');
                statsContainer.className = 'stats-grid';
                const cards = [];

                block.metrics.forEach(metric => {
                    const cardContainer = document.createElement('div');
                    cardContainer.className = 'metric-card-wrapper';
                    statsContainer.appendChild(cardContainer);
                    const card = new MetricCard(cardContainer, metric.label, metric.key, metric.format);
                    cards.push(card);
                    this.components.push(card);
                });
                blockDiv.appendChild(statsContainer);

                // Контейнер для таблицы и графика (если оба есть)
                let tableContainer = null;
                let chartContainer = null;

                // Создаём таблицу, если нужно
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

                // Создаём график, если нужно
                if (block.chart) {
                    chartContainer = document.createElement('div');
                    chartContainer.className = 'chart-wrapper-full';
                    chartContainer.style.width = '100%';
                    chartContainer.style.height = '700px';
                    const chart = new ChartComponent(chartContainer, block.chart.metric, block.chart.title);
                    this.components.push(chart);
                }

                if (tableContainer) {
                    blockDiv.appendChild(tableContainer);
                }
                if (chartContainer) {
                    blockDiv.appendChild(chartContainer);
                }

                // Подписка на обновление карточек
                const unsubscribeStats = appState.subscribe((state) => {
                    if (state.stats) {
                        cards.forEach(card => card.update(state.stats));
                    }
                });
                this.statsSubscriptions.push(unsubscribeStats);
            }

            container.appendChild(blockDiv);
        });
    }
}
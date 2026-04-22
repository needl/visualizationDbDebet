// src/dashboardRenderer.js – исправленная версия с добавленными классами для контейнеров

import {ChartComponent} from "./componets/chartForDeb.js";
import {PieChartComponent} from "./componets/pieChartForDeb.js";
import {MetricCard} from "./componets/metricCard.js";
import {StatsTable} from "./componets/statsTableForDebet.js";
import {appState} from "./state/appState.js";
import {CustomerFilter} from "./componets/customerFilter.js";
import {CustomerCard} from "./componets/customerCard.js";
import {HorizontalBarChart} from "./componets/horizonBarChart.js";
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

                    // Контейнер для столбчатого графика
                    /*const barContainer = document.createElement('div');
                    barContainer.style.width = '100%';
                    barContainer.style.height = '600px';
                    wrapperDiv.appendChild(barContainer);

                    const barChart = new ChartComponent(barContainer, chartConfig.metric, chartConfig.title);
                    this.components.push(barChart);*/

                    //const pieHeaderText = chartConfig.pieTitle || `${chartConfig.title}`;

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
                    //const pieTitle = `Распределение ${chartConfig.title.toLowerCase()}`;
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
                debtStructureContainer.className = 'debt-structure-container'; // добавлен класс
                debtStructCol.appendChild(debtStructureContainer);
                row1.appendChild(debtStructCol);

                // Правая колонка: блок-факторы – 70%
                const blockFactorsCol = document.createElement('div');
                blockFactorsCol.className = 'stats-col factors-col';
                const blockFactorsContainer = document.createElement('div');
                blockFactorsContainer.className = 'block-factors-chart-container'; // добавлен класс
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

                // Ряд 2: топ-10 подрядчиков (две колонки)
                const row2 = document.createElement('div');
                row2.className = 'stats-row';
                const topDebtorsCol = document.createElement('div');
                topDebtorsCol.className = 'stats-col';
                const topDebtorsContainer = document.createElement('div');
                topDebtorsContainer.className = 'chart-wrapper-full'; // ← добавлен класс
                topDebtorsContainer.style.height = '400px';
                topDebtorsCol.appendChild(topDebtorsContainer);
                row2.appendChild(topDebtorsCol);

                const topOverdueCol = document.createElement('div');
                topOverdueCol.className = 'stats-col';
                const topOverdueContainer = document.createElement('div');
                topOverdueContainer.className = 'chart-wrapper-full'; // ← добавлен класс
                topOverdueContainer.style.height = '400px';
                topOverdueCol.appendChild(topOverdueContainer);
                row2.appendChild(topOverdueCol);
                blockDiv.appendChild(row2);

                const topDebtorsChart = new HorizontalBarChart(topDebtorsContainer, 'Топ-10 подрядчиков по ДЗ');
                const topOverdueChart = new HorizontalBarChart(topOverdueContainer, 'Топ-10 подрядчиков по ПДЗ');
                this.components.push(topDebtorsChart, topOverdueChart);

                const topDebtorsUnsub = appState.subscribe(state => {
                    topDebtorsChart.render(state.customerTopDebtors);
                });
                const topOverdueUnsub = appState.subscribe(state => {
                    topOverdueChart.render(state.customerTopOverdue);
                });
                this.statsSubscriptions.push(topDebtorsUnsub, topOverdueUnsub);
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
                let rowContainer = null;
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
                    chartContainer.style.height = '470px';
                    const chart = new ChartComponent(chartContainer, block.chart.metric, block.chart.title);
                    this.components.push(chart);
                }

                // Если есть и таблица, и график — помещаем их в общий ряд
                if (tableContainer && chartContainer) {
                    rowContainer = document.createElement('div');
                    rowContainer.className = 'stats-row';
                    tableContainer.classList.add('stats-col');
                    chartContainer.classList.add('stats-col');
                    rowContainer.appendChild(tableContainer);
                    rowContainer.appendChild(chartContainer);
                    blockDiv.appendChild(rowContainer);
                } else if (tableContainer) {
                    blockDiv.appendChild(tableContainer);
                } else if (chartContainer) {
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